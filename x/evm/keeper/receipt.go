package keeper

import (
	"errors"
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/x/evm/types"
	"google.golang.org/protobuf/proto"
)

// Global buffer pool
var bufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024)
	},
}

var poolMutex = &sync.Mutex{}

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

	// Get a buffer from the pool
	poolMutex.Lock()
	buf := bufPool.Get().([]byte)
	poolMutex.Unlock()
	defer func() {
		poolMutex.Lock()
		bufPool.Put(buf[:0]) // Reset buffer length but keep capacity
		poolMutex.Unlock()
	}()

	// Marshal the receipt
	bz, err := proto.Marshal(receipt)
	if err != nil {
		return err
	}

	// Ensure the buffer is exactly the size needed
	if cap(buf) < len(bz) {
		buf = make([]byte, len(bz))
	} else {
		buf = buf[:len(bz)]
	}

	// Copy marshalled data into the buffer
	copy(buf, bz)

	store.Set(types.ReceiptKey(txHash), buf)
	return nil
}
