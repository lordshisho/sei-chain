package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/app"
	testkeeper "github.com/sei-protocol/sei-chain/testutil/keeper"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	"github.com/sei-protocol/sei-chain/x/evm/types"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
	"github.com/stretchr/testify/require"
)

func MockEVMKeeper() (*evmkeeper.Keeper, sdk.Context) {
	testApp := app.Setup(false, true)
	ctx := testApp.GetContextForDeliverTx([]byte{}).WithBlockHeight(8).WithBlockTime(time.Now())
	k := testApp.EvmKeeper
	k.InitGenesis(ctx, *evmtypes.DefaultGenesis())

	// mint some coins to a sei address
	seiAddr, err := sdk.AccAddressFromHex(common.Bytes2Hex([]byte("seiAddr")))
	if err != nil {
		panic(err)
	}
	err = testApp.BankKeeper.MintCoins(ctx, "evm", sdk.NewCoins(sdk.NewCoin("usei", sdk.NewInt(10))))
	if err != nil {
		panic(err)
	}
	err = testApp.BankKeeper.SendCoinsFromModuleToAccount(ctx, "evm", seiAddr, sdk.NewCoins(sdk.NewCoin("usei", sdk.NewInt(10))))
	if err != nil {
		panic(err)
	}
	return &k, ctx
}

func TestParams(t *testing.T) {
	k := &testkeeper.EVMTestApp.EvmKeeper
	// k, _ := MockEVMKeeper()
	ctx := testkeeper.EVMTestApp.GetContextForDeliverTx([]byte{}).WithBlockTime(time.Now())
	ctx, _ = ctx.CacheContext()
	require.Equal(t, "usei", k.GetBaseDenom(ctx))
	require.Equal(t, types.DefaultPriorityNormalizer, k.GetPriorityNormalizer(ctx))
	require.Equal(t, types.DefaultMinFeePerGas, k.GetDynamicBaseFeePerGas(ctx))
	require.Equal(t, types.DefaultBaseFeePerGas, k.GetBaseFeePerGas(ctx))
	require.Equal(t, types.DefaultMinFeePerGas, k.GetMinimumFeePerGas(ctx))
	require.Equal(t, types.DefaultDeliverTxHookWasmGasLimit, k.GetDeliverTxHookWasmGasLimit(ctx))

	require.Nil(t, k.GetParams(ctx).Validate())
}

func TestGetParamsIfExists(t *testing.T) {
	k := &testkeeper.EVMTestApp.EvmKeeper
	ctx := testkeeper.EVMTestApp.GetContextForDeliverTx([]byte{}).WithBlockTime(time.Now())

	// Define the expected parameters
	expectedParams := types.Params{
		PriorityNormalizer: sdk.NewDec(1),
		BaseFeePerGas:      sdk.NewDec(1),
	}

	// Set only a subset of the parameters in the keeper
	k.Paramstore.Set(ctx, types.KeyPriorityNormalizer, expectedParams.PriorityNormalizer)
	k.Paramstore.Set(ctx, types.KeyBaseFeePerGas, expectedParams.BaseFeePerGas)

	// Retrieve the parameters using GetParamsIfExists
	params := k.GetParamsIfExists(ctx)

	// Assert that the retrieved parameters match the expected parameters
	require.Equal(t, expectedParams.PriorityNormalizer, params.PriorityNormalizer)
	require.Equal(t, expectedParams.BaseFeePerGas, params.BaseFeePerGas)

	// Assert that the missing parameter has its default value
	require.Equal(t, types.DefaultParams().DeliverTxHookWasmGasLimit, params.DeliverTxHookWasmGasLimit)
}
