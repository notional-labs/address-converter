package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tmlibs/bech32"
)

// AccAddressFromBech32 creates an AccAddress from a Bech32 string.
func AccAddressFromBech32(address string, prefix string) (addr types.AccAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return types.AccAddress{}, errors.New("empty address string is not allowed")
	}
	bz, err := types.GetFromBech32(address, prefix)
	if err != nil {
		return nil, err
	}

	err = types.VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return types.AccAddress(bz), nil
}

// String implements the Stringer interface.
func String(aa types.AccAddress, prefix string) string {
	if aa.Empty() {
		return ""
	}

	bech32Addr, err := bech32.ConvertAndEncode(prefix, aa.Bytes())
	if err != nil {
		panic(err)
	}

	return bech32Addr
}

func main() {
	cosmosAddr := "cosmos1fmv0nj9fgaplam537u46twdjvzcgjqmqrxjaln"

	addrBz, _ := AccAddressFromBech32(cosmosAddr, "cosmos")
	osmoAddr := String(addrBz, "osmo")
	fmt.Println(osmoAddr)

}
