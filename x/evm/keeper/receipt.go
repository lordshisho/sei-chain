package keeper

import (
	"errors"
	"fmt"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 4096)
	},
}

// Receipt is a data structure that stores EVM-specific transaction metadata.
// Many EVM applications (e.g., MetaMask) rely on being able to query receipt
// by EVM transaction hash (not Sei transaction hash) to function properly.
func (k *Keeper) GetReceipt(ctx sdk.Context, txHash common.Hash) (*types.Receipt, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ReceiptKey(txHash))
	if bz == nil {
		return nil, errors.New("not found")
	}
	r := types.Receipt{}
	if err := r.Unmarshal(bz); err != nil {
		return nil, err
	}
	return &r, nil
}

func (k *Keeper) SetReceipt(ctx sdk.Context, txHash common.Hash, receipt *types.Receipt) error {
	store := ctx.KVStore(k.storeKey)

	// Get a buffer from the pool and ensure it's clear
	buf := bufPool.Get().([]byte)[:0]
	defer func() {
		bufPool.Put(buf) // Return the buffer to the pool
	}()

	// resize the buffer if necessary
	size := receipt.Size()
	if len(buf) < size {
		buf = make([]byte, size)
	} else if size < len(buf) {
		buf = buf[0:size]
	}

	_, err := receipt.MarshalTo(buf)
	if err != nil {
		ctx.Logger().Error("error marshalling receipt", "err", err)
		return err
	}

	ctx.Logger().Info("[Debug] saving receipt", "tx", txHash.Hex(), "receipt", fmt.Sprintf("%X", buf))

	store.Set(types.ReceiptKey(txHash), buf)
	return nil
}
