package checkpoint_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/maticnetwork/heimdall/app"
	"github.com/maticnetwork/heimdall/checkpoint"
	hmTypes "github.com/maticnetwork/heimdall/types"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	app *app.HeimdallApp
	ctx sdk.Context
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app, suite.ctx, _ = createTestApp(false)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestAddCheckpoint() {
	t, app, ctx := suite.T(), suite.app, suite.ctx
	keeper := app.CheckpointKeeper

	headerBlockNumber := uint64(2000)
	startBlock := uint64(0)
	endBlock := uint64(256)
	rootHash := hmTypes.HexToHeimdallHash("123")
	accountRootHash := hmTypes.HexToHeimdallHash("456")
	proposerAddress := hmTypes.HexToHeimdallAddress("123")
	timestamp := uint64(time.Now().Unix())

	checkpointBlockHeader := hmTypes.CreateBlock(
		startBlock,
		endBlock,
		rootHash,
		accountRootHash,
		proposerAddress,
		timestamp,
	)
	err := keeper.AddCheckpoint(ctx, headerBlockNumber, checkpointBlockHeader)
	require.NoError(t, err)

	result, err := keeper.GetCheckpointByIndex(ctx, headerBlockNumber)
	require.NoError(t, err)
	require.Equal(t, startBlock, result.StartBlock)
	require.Equal(t, endBlock, result.EndBlock)
	require.Equal(t, rootHash, result.RootHash)
	require.Equal(t, accountRootHash, result.AccountRootHash)
	require.Equal(t, proposerAddress, result.Proposer)
	require.Equal(t, timestamp, result.TimeStamp)
}

func (suite *KeeperTestSuite) TestGetCheckpointList() {
	t, app, ctx := suite.T(), suite.app, suite.ctx
	keeper := app.CheckpointKeeper

	count := 5
	childBlockInterval := uint64(10000)

	startBlock := uint64(0)
	endBlock := uint64(0)

	for i := 0; i < count; i++ {
		headerBlockNumber := childBlockInterval * (uint64(i) + 1)

		startBlock = startBlock + endBlock
		endBlock = endBlock + uint64(255)
		rootHash := hmTypes.HexToHeimdallHash("123")
		accountRootHash := hmTypes.HexToHeimdallHash("123")
		proposerAddress := hmTypes.HexToHeimdallAddress("123")
		timestamp := uint64(time.Now().Unix()) + uint64(i)

		checkpointBlockHeader := hmTypes.CreateBlock(
			startBlock,
			endBlock,
			rootHash,
			accountRootHash,
			proposerAddress,
			timestamp,
		)

		keeper.AddCheckpoint(ctx, headerBlockNumber, checkpointBlockHeader)
		keeper.UpdateACKCount(ctx)
	}

	result, err := keeper.GetCheckpointList(ctx, uint64(1), uint64(20))
	require.NoError(t, err)
	require.LessOrEqual(t, count, len(result))

	// GetLastCheckpoint
	lastCheckpoint, _ := keeper.GetLastCheckpoint(ctx)
	// TODO: find a way to mock helper.GetConfig().ChildBlockInterval
	// currently, ackCount is not getting updated as above value is not set
	require.Equal(t, result[0], lastCheckpoint)
}

func (suite *KeeperTestSuite) TestHasStoreValue() {
	t, app, ctx := suite.T(), suite.app, suite.ctx
	keeper := app.CheckpointKeeper
	key := checkpoint.ACKCountKey
	result := keeper.HasStoreValue(ctx, key)
	require.True(t, result)
}

func (suite *KeeperTestSuite) TestFlushCheckpointBuffer() {
	t, app, ctx := suite.T(), suite.app, suite.ctx
	keeper := app.CheckpointKeeper
	key := checkpoint.BufferCheckpointKey
	keeper.FlushCheckpointBuffer(ctx)
	result := keeper.HasStoreValue(ctx, key)
	require.False(t, result)
}
