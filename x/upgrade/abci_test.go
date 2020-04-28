package upgrade_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
)

var s TestSuite

type TestSuite struct {
	module  module.AppModule
	querier sdk.Querier
	handler gov.Handler
	ctx     sdk.Context

	app *simapp.SimApp
	db  dbm.DB
}

func setupTest(t *testing.T, height int64, skip map[int64]bool) TestSuite {
	db := dbm.NewMemDB()
	app := simapp.NewSimApp(
		log.NewNopLogger(), db, nil, true, skip, simapp.DefaultNodeHome, 0,
		baseapp.SetPruning(store.NewPruningOptionsFromString(store.PruningStrategySyncable)),
	)
	genesisState := simapp.NewDefaultGenesisState()

	stateBytes, err := codec.MarshalJSONIndent(app.Codec(), genesisState)
	require.NoError(t, err)

	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)

	s.app = app
	s.db = db
	s.ctx = app.BaseApp.NewContext(false, abci.Header{Height: height, Time: time.Now()})

	s.module = upgrade.NewAppModule(app.UpgradeKeeper)
	s.querier = s.module.NewQuerierHandler()
	s.handler = upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)

	return s
}

func TestUpgradeNoPersistHeight(t *testing.T) {
	s := setupTest(t, 1, map[int64]bool{})
	ctx := s.ctx

	// begin and commit first block
	s.app.BeginBlock(abci.RequestBeginBlock{Header: abci.Header{Height: ctx.BlockHeight()}})
	s.app.Commit()
	require.Equal(t, int64(1), s.app.LastBlockHeight())

	// begin second block where we create an upgrade plan
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	s.app.BeginBlock(abci.RequestBeginBlock{Header: abci.Header{Height: ctx.BlockHeight()}})

	upgradeHeight := int64(3)
	proposal := &upgrade.SoftwareUpgradeProposal{
		Title: "prop",
		Plan:  upgrade.Plan{Name: "test", Height: upgradeHeight},
	}

	// set upgrade plan and commit second block
	require.NoError(t, s.handler(ctx, proposal))
	s.app.Commit()
	require.Equal(t, int64(2), s.app.LastBlockHeight())

	// begin third block where the plan should be triggered and we panic
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	s.app.BeginBlock(abci.RequestBeginBlock{Header: abci.Header{Height: ctx.BlockHeight()}})

	t.Log("verify the upgrade proposal plan is triggered (not skipped and panics)")
	VerifyDoUpgradeWithCtx(t, ctx, proposal.Plan.Name)

	// latest committed block should be 2
	require.Equal(t, int64(2), s.app.LastBlockHeight())

	// create second app with the same database where the app doesn't persist the
	// last height to disk due to the pruning strategy.
	appState, _, _, err := s.app.ExportAppStateAndValidators(false, nil)
	require.NoError(t, err)

	app2 := simapp.NewSimApp(
		log.NewNopLogger(), s.db, nil, true, map[int64]bool{}, simapp.DefaultNodeHome, 0,
		baseapp.SetPruning(store.NewPruningOptionsFromString(store.PruningStrategySyncable)),
	)
	app2.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: appState,
		},
	)

	// Since the plan triggered a panic/halt of the old binary, ensure the last
	// height was persisted to disk.
	require.Equal(t, int64(2), app2.LastBlockHeight())
}

func TestRequireName(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})

	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{}})
	require.NotNil(t, err)
	require.True(t, errors.Is(sdkerrors.ErrInvalidRequest, err), err)
}

func TestRequireFutureTime(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Time: s.ctx.BlockHeader().Time}})
	require.NotNil(t, err)
	require.True(t, errors.Is(sdkerrors.ErrInvalidRequest, err), err)
}

func TestRequireFutureBlock(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: s.ctx.BlockHeight()}})
	require.NotNil(t, err)
	require.True(t, errors.Is(sdkerrors.ErrInvalidRequest, err), err)
}

func TestCantSetBothTimeAndHeight(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Time: time.Now(), Height: s.ctx.BlockHeight() + 1}})
	require.NotNil(t, err)
	require.True(t, errors.Is(sdkerrors.ErrInvalidRequest, err), err)
}

func TestDoTimeUpgrade(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})
	t.Log("Verify can schedule an upgrade")
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Time: time.Now()}})
	require.Nil(t, err)

	VerifyDoUpgrade(t)
}

func TestDoHeightUpgrade(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})
	t.Log("Verify can schedule an upgrade")
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: s.ctx.BlockHeight() + 1}})
	require.Nil(t, err)

	VerifyDoUpgrade(t)
}

func TestCanOverwriteScheduleUpgrade(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})
	t.Log("Can overwrite plan")
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "bad_test", Height: s.ctx.BlockHeight() + 10}})
	require.Nil(t, err)
	err = s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: s.ctx.BlockHeight() + 1}})
	require.Nil(t, err)

	VerifyDoUpgrade(t)
}

func VerifyDoUpgrade(t *testing.T) {
	t.Log("Verify that a panic happens at the upgrade time/height")
	newCtx := s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1).WithBlockTime(time.Now())

	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	require.Panics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})

	t.Log("Verify that the upgrade can be successfully applied with a handler")
	s.app.UpgradeKeeper.SetUpgradeHandler("test", func(ctx sdk.Context, plan upgrade.Plan) {})
	require.NotPanics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})

	VerifyCleared(t, newCtx)
}

func VerifyDoUpgradeWithCtx(t *testing.T, newCtx sdk.Context, proposalName string) {
	t.Log("Verify that a panic happens at the upgrade time/height")
	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	require.Panics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})

	t.Log("Verify that the upgrade can be successfully applied with a handler")
	s.app.UpgradeKeeper.SetUpgradeHandler(proposalName, func(ctx sdk.Context, plan upgrade.Plan) {})
	require.NotPanics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})

	VerifyCleared(t, newCtx)
}

func TestHaltIfTooNew(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})
	t.Log("Verify that we don't panic with registered plan not in database at all")
	var called int
	s.app.UpgradeKeeper.SetUpgradeHandler("future", func(ctx sdk.Context, plan upgrade.Plan) { called++ })

	newCtx := s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1).WithBlockTime(time.Now())
	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	require.NotPanics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})
	require.Equal(t, 0, called)

	t.Log("Verify we panic if we have a registered handler ahead of time")
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "future", Height: s.ctx.BlockHeight() + 3}})
	require.NoError(t, err)
	require.Panics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})
	require.Equal(t, 0, called)

	t.Log("Verify we no longer panic if the plan is on time")

	futCtx := s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 3).WithBlockTime(time.Now())
	req = abci.RequestBeginBlock{Header: futCtx.BlockHeader()}
	require.NotPanics(t, func() {
		s.module.BeginBlock(futCtx, req)
	})
	require.Equal(t, 1, called)

	VerifyCleared(t, futCtx)
}

func VerifyCleared(t *testing.T, newCtx sdk.Context) {
	t.Log("Verify that the upgrade plan has been cleared")
	bz, err := s.querier(newCtx, []string{upgrade.QueryCurrent}, abci.RequestQuery{})
	require.NoError(t, err)
	require.Nil(t, bz)
}

func TestCanClear(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})
	t.Log("Verify upgrade is scheduled")
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Time: time.Now()}})
	require.Nil(t, err)

	err = s.handler(s.ctx, &upgrade.CancelSoftwareUpgradeProposal{Title: "cancel"})
	require.Nil(t, err)

	VerifyCleared(t, s.ctx)
}

func TestCantApplySameUpgradeTwice(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Time: time.Now()}})
	require.Nil(t, err)
	VerifyDoUpgrade(t)
	t.Log("Verify an executed upgrade \"test\" can't be rescheduled")
	err = s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Time: time.Now()}})
	require.NotNil(t, err)
	require.True(t, errors.Is(sdkerrors.ErrInvalidRequest, err), err)
}

func TestNoSpuriousUpgrades(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})
	t.Log("Verify that no upgrade panic is triggered in the BeginBlocker when we haven't scheduled an upgrade")
	req := abci.RequestBeginBlock{Header: s.ctx.BlockHeader()}
	require.NotPanics(t, func() {
		s.module.BeginBlock(s.ctx, req)
	})
}

func TestPlanStringer(t *testing.T) {
	ti, err := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	require.Nil(t, err)
	require.Equal(t, `Upgrade Plan
  Name: test
  Time: 2020-01-01T00:00:00Z
  Info: `, upgrade.Plan{Name: "test", Time: ti}.String())
	require.Equal(t, `Upgrade Plan
  Name: test
  Height: 100
  Info: `, upgrade.Plan{Name: "test", Height: 100}.String())
}

func VerifyNotDone(t *testing.T, newCtx sdk.Context, name string) {
	t.Log("Verify that upgrade was not done")
	height := s.app.UpgradeKeeper.GetDoneHeight(newCtx, name)
	require.Zero(t, height)
}

func VerifyDone(t *testing.T, newCtx sdk.Context, name string) {
	t.Log("Verify that the upgrade plan has been executed")
	height := s.app.UpgradeKeeper.GetDoneHeight(newCtx, name)
	require.NotZero(t, height)
}

func VerifySet(t *testing.T, skipUpgradeHeights map[int64]bool) {
	t.Log("Verify if the skip upgrade has been set")

	for k := range skipUpgradeHeights {
		require.True(t, s.app.UpgradeKeeper.IsSkipHeight(k))
	}
}

func TestContains(t *testing.T) {
	var (
		skipOne int64 = 11
	)
	s := setupTest(t, 10, map[int64]bool{skipOne: true})

	VerifySet(t, map[int64]bool{skipOne: true})
	t.Log("case where array contains the element")
	require.True(t, s.app.UpgradeKeeper.IsSkipHeight(11))

	t.Log("case where array doesn't contain the element")
	require.False(t, s.app.UpgradeKeeper.IsSkipHeight(4))
}

func TestSkipUpgradeSkippingAll(t *testing.T) {
	var (
		skipOne int64 = 11
		skipTwo int64 = 20
	)
	s := setupTest(t, 10, map[int64]bool{skipOne: true, skipTwo: true})

	newCtx := s.ctx

	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: skipOne}})
	require.NoError(t, err)

	t.Log("Verify if skip upgrade flag clears upgrade plan in both cases")
	VerifySet(t, map[int64]bool{skipOne: true, skipTwo: true})

	newCtx = newCtx.WithBlockHeight(skipOne)
	require.NotPanics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})

	t.Log("Verify a second proposal also is being cleared")
	err = s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop2", Plan: upgrade.Plan{Name: "test2", Height: skipTwo}})
	require.NoError(t, err)

	newCtx = newCtx.WithBlockHeight(skipTwo)
	require.NotPanics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})

	// To ensure verification is being done only after both upgrades are cleared
	t.Log("Verify if both proposals are cleared")
	VerifyCleared(t, s.ctx)
	VerifyNotDone(t, s.ctx, "test")
	VerifyNotDone(t, s.ctx, "test2")
}

func TestUpgradeSkippingOne(t *testing.T) {
	var (
		skipOne int64 = 11
		skipTwo int64 = 20
	)
	s := setupTest(t, 10, map[int64]bool{skipOne: true})

	newCtx := s.ctx

	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: skipOne}})
	require.Nil(t, err)

	t.Log("Verify if skip upgrade flag clears upgrade plan in one case and does upgrade on another")
	VerifySet(t, map[int64]bool{skipOne: true})

	// Setting block height of proposal test
	newCtx = newCtx.WithBlockHeight(skipOne)
	require.NotPanics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})

	t.Log("Verify the second proposal is not skipped")
	err = s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop2", Plan: upgrade.Plan{Name: "test2", Height: skipTwo}})
	require.Nil(t, err)
	// Setting block height of proposal test2
	newCtx = newCtx.WithBlockHeight(skipTwo)
	VerifyDoUpgradeWithCtx(t, newCtx, "test2")

	t.Log("Verify first proposal is cleared and second is done")
	VerifyNotDone(t, s.ctx, "test")
	VerifyDone(t, s.ctx, "test2")
}

func TestUpgradeSkippingOnlyTwo(t *testing.T) {
	var (
		skipOne   int64 = 11
		skipTwo   int64 = 20
		skipThree int64 = 25
	)
	s := setupTest(t, 10, map[int64]bool{skipOne: true, skipTwo: true})

	newCtx := s.ctx

	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: skipOne}})
	require.Nil(t, err)

	t.Log("Verify if skip upgrade flag clears upgrade plan in both cases and does third upgrade")
	VerifySet(t, map[int64]bool{skipOne: true, skipTwo: true})

	// Setting block height of proposal test
	newCtx = newCtx.WithBlockHeight(skipOne)
	require.NotPanics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})

	// A new proposal with height in skipUpgradeHeights
	err = s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop2", Plan: upgrade.Plan{Name: "test2", Height: skipTwo}})
	require.Nil(t, err)
	// Setting block height of proposal test2
	newCtx = newCtx.WithBlockHeight(skipTwo)
	require.NotPanics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})

	t.Log("Verify a new proposal is not skipped")
	err = s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop3", Plan: upgrade.Plan{Name: "test3", Height: skipThree}})
	require.Nil(t, err)
	newCtx = newCtx.WithBlockHeight(skipThree)
	VerifyDoUpgradeWithCtx(t, newCtx, "test3")

	t.Log("Verify two proposals are cleared and third is done")
	VerifyNotDone(t, s.ctx, "test")
	VerifyNotDone(t, s.ctx, "test2")
	VerifyDone(t, s.ctx, "test3")
}

func TestUpgradeWithoutSkip(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})
	newCtx := s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1).WithBlockTime(time.Now())
	req := abci.RequestBeginBlock{Header: newCtx.BlockHeader()}
	err := s.handler(s.ctx, &upgrade.SoftwareUpgradeProposal{Title: "prop", Plan: upgrade.Plan{Name: "test", Height: s.ctx.BlockHeight() + 1}})
	require.Nil(t, err)
	t.Log("Verify if upgrade happens without skip upgrade")
	require.Panics(t, func() {
		s.module.BeginBlock(newCtx, req)
	})

	VerifyDoUpgrade(t)
	VerifyDone(t, s.ctx, "test")
}

func TestDumpUpgradeInfoToFile(t *testing.T) {
	s := setupTest(t, 10, map[int64]bool{})

	planHeight := s.ctx.BlockHeight() + 1
	name := "test"
	t.Log("verify if upgrade height is dumped to file")
	err := s.app.UpgradeKeeper.DumpUpgradeInfoToDisk(planHeight, name)
	require.Nil(t, err)

	upgradeInfoFilePath, err := s.app.UpgradeKeeper.GetUpgradeInfoPath()
	require.Nil(t, err)

	data, err := ioutil.ReadFile(upgradeInfoFilePath)
	require.NoError(t, err)

	var upgradeInfo storetypes.UpgradeInfo
	err = json.Unmarshal(data, &upgradeInfo)
	require.Nil(t, err)

	t.Log("Verify upgrade height from file matches ")
	require.Equal(t, upgradeInfo.Height, planHeight)

	// clear the test file
	err = os.Remove(upgradeInfoFilePath)
	require.Nil(t, err)
}