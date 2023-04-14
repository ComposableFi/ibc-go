package keeper

import (
	"bytes"
	"crypto/sha256"
	"math"
	"strings"

	cosmwasm "github.com/CosmWasm/wasmvm"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmos/ibc-go/v7/modules/light-clients/08-wasm/types"
)

type Keeper struct {
	// implements gRPC QueryServer interface
	types.QueryServer

	storeKey  storetypes.StoreKey
	cdc       codec.BinaryCodec
	wasmVM    *cosmwasm.VM
	authority string
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey) Keeper {
	// Wasm VM
	wasmDataDir := "ibc_08-wasm_client_data"
	wasmSupportedFeatures := strings.Join([]string{"storage", "iterator"}, ",")
	wasmMemoryLimitMb := uint32(math.Pow(2, 12))
	wasmPrintDebug := true
	wasmCacheSizeMb := uint32(math.Pow(2, 8))

	vm, err := cosmwasm.NewVM(wasmDataDir, wasmSupportedFeatures, wasmMemoryLimitMb, wasmPrintDebug, wasmCacheSizeMb)
	if err != nil {
		panic(err)
	}
	types.WasmVM = vm

	// governance authority
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	return Keeper{
		cdc:       cdc,
		storeKey:  key,
		wasmVM:    vm,
		authority: authority.String(),
	}
}

func generateWasmCodeHash(code []byte) []byte {
	hash := sha256.Sum256(code)
	return hash[:]
}

func (k Keeper) storeWasmCode(ctx sdk.Context, code []byte) ([]byte, error) {
	store := ctx.KVStore(k.storeKey)

	var err error
	if IsGzip(code) {
		ctx.GasMeter().ConsumeGas(types.VMGasRegister.UncompressCosts(len(code)), "Uncompress gzip bytecode")
		code, err = Uncompress(code, uint64(types.MaxWasmSize))
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrCreateContractFailed, err.Error())
		}
	}

	// Check to see if the store has a code with the same code it
	codeHash := generateWasmCodeHash(code)
	codeIDKey := types.CodeIDKey(codeHash)
	if store.Has(codeIDKey) {
		return nil, types.ErrWasmCodeExists
	}

	// run the code through the wasm light client validation process
	if isValidWasmCode, err := types.ValidateWasmCode(code); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrWasmCodeValidation, err.Error())
	} else if !isValidWasmCode {
		return nil, types.ErrWasmInvalidCode
	}

	// create the code in the vm
	ctx.GasMeter().ConsumeGas(types.VMGasRegister.CompileCosts(len(code)), "Compiling wasm bytecode")
	codeID, err := k.wasmVM.Create(code)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCreateContractFailed, err.Error())
	}

	// safety check to assert that code ID returned by WasmVM equals to code hash
	if !bytes.Equal(codeID, codeHash) {
		return nil, types.ErrWasmInvalidCodeID
	}

	store.Set(codeIDKey, code)
	return codeID, nil
}

func (k Keeper) importWasmCode(ctx sdk.Context, codeIDKey, wasmCode []byte) error {
	store := ctx.KVStore(k.storeKey)
	if IsGzip(wasmCode) {
		var err error
		wasmCode, err = Uncompress(wasmCode, uint64(types.MaxWasmSize))
		if err != nil {
			return sdkerrors.Wrap(types.ErrCreateContractFailed, err.Error())
		}
	}

	codeID, err := k.wasmVM.Create(wasmCode)
	if err != nil {
		return sdkerrors.Wrap(types.ErrCreateContractFailed, err.Error())
	}
	if !bytes.Equal(codeIDKey, types.CodeIDKey(codeID)) {
		return sdkerrors.Wrap(types.ErrInvalid, "code hashes not same")
	}

	store.Set(codeIDKey, wasmCode)
	return nil
}
