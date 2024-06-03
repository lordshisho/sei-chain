package keeper

import (
	"errors"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/x/evm/types"
	"google.golang.org/protobuf/proto"
)

// Create a pool for byte slices
var byteSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 1024) // 1KB initial capacity, adjust as needed
	},
}

// getByteSlice retrieves a byte slice from the pool
func getByteSlice() []byte {
	return byteSlicePool.Get().([]byte)
}

// putByteSlice returns a byte slice to the pool
func putByteSlice(b []byte) {
	if cap(b) <= 1024 { // Only pool slices of 1KB or less, adjust as needed
		byteSlicePool.Put(b[:0])
	}
}

// Receipt is a data structure that stores EVM specific transaction metadata.
// Many EVM applications (e.g. MetaMask) rely on being able to query receipt
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

	// Use the byte slice pool for marshaling
	bz := getByteSlice()
	defer putByteSlice(bz)

	// Marshal the receipt
	data, err := proto.Marshal(receipt)
	if err != nil {
		return err
	}

	// Set the data in the store
	store.Set(types.ReceiptKey(txHash), data)
	return nil
}
