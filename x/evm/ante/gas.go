package ante

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	evmkeeper "github.com/sei-protocol/sei-chain/x/evm/keeper"
	evmtypes "github.com/sei-protocol/sei-chain/x/evm/types"
)

type GasLimitDecorator struct {
	evmKeeper *evmkeeper.Keeper
}

func NewGasLimitDecorator(evmKeeper *evmkeeper.Keeper) *GasLimitDecorator {
	return &GasLimitDecorator{evmKeeper: evmKeeper}
}

// Called at the end of the ante chain to set gas limit properly
func (gl GasLimitDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	msg := evmtypes.MustGetEVMTransactionMessage(tx)
	txData, err := evmtypes.UnpackTxData(msg.Data)
	if err != nil {
		return ctx, err
	}

	adjustedGasLimit := gl.evmKeeper.GetPriorityNormalizer(ctx).MulInt64(int64(txData.GetGas()))
	gasWanted := txData.GetGas()

	if cp := ctx.ConsensusParams(); cp != nil && cp.Block != nil {
		// If there exists a maximum block gas limit, we must ensure that the tx
		// does not exceed it.
		if gasWanted > 0 {
			fmt.Printf("[Debug] GasLimitDecorator Gas wanted is %d and block gas limit is %d at height %d\n", gasWanted, cp.Block.MaxGas, ctx.BlockHeight())
		}
		if cp.Block.MaxGas > 0 && gasWanted > uint64(cp.Block.MaxGas) {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrInvalidGasLimit, "tx gas limit %d exceeds block max gas %d", gasWanted, cp.Block.MaxGas)
		}
	}
	ctx = ctx.WithGasMeter(sdk.NewGasMeterWithMultiplier(ctx, adjustedGasLimit.TruncateInt().Uint64()))
	return next(ctx, tx, simulate)
}
