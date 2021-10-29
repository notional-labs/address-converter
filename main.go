package main

import (
	"github.com/notional-labs/addr-converter/cmd"
)

func main() {
	cmd.Execute()
}

//
//func main() {
//	cosmosAddr := "cosmos1fmv0nj9fgaplam537u46twdjvzcgjqmqrxjaln"
//
//	addrBz, _ := AccAddressFromBech32(cosmosAddr, "cosmos")
//	osmoAddr := String(addrBz, "sif")
//	fmt.Println(osmoAddr)
//
//}
