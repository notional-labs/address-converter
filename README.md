# Installation

Pull repo, have Golang 1.17 installed.

`go install ./...`

Once done, the address converter will be availalbe.

# Usage

In the terminal, use:

`addr-converter {cosmos_address} --chain-prefix {prefix}`

For example:

`addr-converter cosmos1h7j93mqeukv8vfuzwl3495c59sgsxvk2xv7e2n --chain-prefix 'sif'`
Will output:

`sif1h7j93mqeukv8vfuzwl3495c59sgsxvk2r3309c`