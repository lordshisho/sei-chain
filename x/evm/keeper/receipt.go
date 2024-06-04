package keeper

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sei-protocol/sei-chain/x/evm/types"
)

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

func (k *Keeper) pushSlice(b []byte) {
	k.sliceLock.Lock()
	defer k.sliceLock.Unlock()
	k.slices = append(k.slices, b)
}

func (k *Keeper) ReleaseBuffers() {
	k.sliceLock.Lock()
	defer k.sliceLock.Unlock()
	for _, b := range k.slices {
		k.slicePool.Put(b)
	}
	k.slices = make([][]byte, 0, cap(k.slices))
}

func (k *Keeper) SetReceipt(ctx sdk.Context, txHash common.Hash, receipt *types.Receipt) error {
	store := ctx.KVStore(k.storeKey)

	bz := k.slicePool.Get().([]byte)
	defer func() {
		k.pushSlice(bz)
	}()

	if cap(bz) < receipt.Size() {
		bz = make([]byte, receipt.Size())
	}
	bz = bz[:receipt.Size()]

	_, err := receipt.MarshalTo(bz)
	if err != nil {
		ctx.Logger().Error("error marshalling receipt", "err", err)
		return err
	}

	store.Set(types.ReceiptKey(txHash), bz)
	return nil
}
