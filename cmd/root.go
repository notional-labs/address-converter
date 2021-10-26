/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"os"
	"errors"
	"strings"
	"github.com/tendermint/tmlibs/bech32"
	"github.com/spf13/viper"
)

var cfgFile string
var chainPrefix string


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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "address-converter",
	Short: "Convert a cosmos address to another tendermint chain",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		cosmosAddr := args[0];
		prefix, _ := cmd.Flags().GetString("chain-prefix")
		addrBz, _ := AccAddressFromBech32(cosmosAddr, "cosmos")
		osmoAddr := String(addrBz, prefix)
		fmt.Println(osmoAddr)
	},
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.address-converter.yaml)")

	rootCmd.PersistentFlags().StringVar(&chainPrefix, "chain-prefix", "c", "The chain prefix, e.g. 'osmo' or 'sif'")
	rootCmd.MarkFlagRequired("chain-prefix")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".address-converter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".address-converter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
