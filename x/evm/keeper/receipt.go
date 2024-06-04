package keeper

import (
	"bytes"
	"errors"
	"fmt"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/x/evm/types"
	"google.golang.org/protobuf/proto"
)

// Global buffer pool
var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
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
	r := &types.Receipt{}
	if err := proto.Unmarshal(bz, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (k *Keeper) SetReceipt(ctx sdk.Context, txHash common.Hash, receipt *types.Receipt) error {
	store := ctx.KVStore(k.storeKey)

	// Get a buffer from the pool
	buf := bufPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()      // Reset the buffer for reuse
		bufPool.Put(buf) // Return the buffer to the pool
	}()

	// Clear the buffer and ensure it is empty
	buf.Reset()

	// Marshal the receipt
	bz, err := proto.MarshalOptions{Deterministic: true}.Marshal(receipt)
	if err != nil {
		return err
	}

	// Ensure the buffer is large enough
	if cap(buf.Bytes()) < len(bz) {
		buf.Grow(len(bz))
	}

	// Write marshalled data into the buffer
	buf.Write(bz)
	buf.Truncate(len(bz))

	fmt.Printf("[Debug] tx=%s, receiptLen=%d\n", txHash.Hex(), buf.Len())

	store.Set(types.ReceiptKey(txHash), buf.Bytes())
	return nil
}
