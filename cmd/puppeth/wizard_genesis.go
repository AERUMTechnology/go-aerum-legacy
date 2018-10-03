// Copyright 2017 The go-aerum Authors
// This file is part of go-aerum.
//
// go-aerum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-aerum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-aerum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/AERUMTechnology/go-aerum/common"
	"github.com/AERUMTechnology/go-aerum/accounts/abi/bind"
	guvnor "github.com/AERUMTechnology/go-aerum/contracts/atmosGovernance"
	"github.com/AERUMTechnology/go-aerum/core"
	"github.com/AERUMTechnology/go-aerum/ethclient"
	"github.com/AERUMTechnology/go-aerum/log"
	"github.com/AERUMTechnology/go-aerum/params"
)

// Values for AERUMS Genesis related to ATMOS Consensus
var (
	atmosMinDelegateNo = 3
	atmosNetID               = 418313827693
	atmosGovernanceAddress   = "0x5df4f6cde8f209ef8f152b59cabce70db21787b3"
	atmosBlockInterval       = uint64(2)
	atmosEpochInterval       = uint64(1000)
	atmosGasLimit            = uint64(25000000)
	atmosEthereumRPCProvider = "https://rinkeby.infura.io"
)

func getBootstrapDelegates() ([]common.Address, error) {

	fmt.Println("\n\n[aerDEV] --------------------------------------------------------------------------------------------------------- [aerDEV]")
	fmt.Println("[aerDEV] --- We are calling our Governance Contract on Ethereum to add our bootstrap signers to this genesis --- [aerDEV]")
	fmt.Println("[aerDEV] --------------------------------------------------------------------------------------------------------- [aerDEV]\n\n")

	contractAddress := common.HexToAddress(atmosGovernanceAddress)

	bootstrapDelegates := make([]common.Address, 0)

	ethclient, err := ethclient.Dial(atmosEthereumRPCProvider)
	if err != nil {
		fmt.Println(err)
	}

	caller, err := guvnor.NewAtmosCaller(contractAddress, ethclient)
	if err != nil {
		fmt.Println(err)
	}

	addresses, err := caller.ShowBootstrapDelegates(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
	}

	if len(addresses) < atmosMinDelegateNo {
		log.Error("Failed to save genesis file", "err", "Not enough Delegates to continue. Contact the aerum team to report this issue.")
	}
	
	if len(addresses) >= atmosMinDelegateNo {
		log.Info( fmt.Sprintf("Fantastic! we found %d delegates. you may proceed in generating a genesis.", len(addresses) )  )
	}
 
	for _, address := range addresses {
		bootstrapDelegates = append(bootstrapDelegates, address)
	}

	return bootstrapDelegates, nil
}

// makeGenesis creates a new genesis struct based on some user input.
func (w *wizard) makeGenesis() {

	boostrapDelegate, err := getBootstrapDelegates()

	if err != nil {
		fmt.Println("ISSUE WITH GENERATION, Please take note: ", err)
	}

	// Construct a default genesis block
	genesis := &core.Genesis{
		Timestamp:  uint64(time.Now().Unix()),
		GasLimit:   atmosGasLimit,
		Difficulty: big.NewInt(1),
		Alloc:      make(core.GenesisAlloc),
		Config: &params.ChainConfig{
			HomesteadBlock: big.NewInt(1),
			EIP150Block:    big.NewInt(2),
			EIP155Block:    big.NewInt(3),
			EIP158Block:    big.NewInt(3),
			ByzantiumBlock: big.NewInt(4),
			Atmos: &params.AtmosConfig{
				Period:                     atmosBlockInterval,
				Epoch:                      atmosEpochInterval,
				GovernanceAddress:          common.HexToAddress(atmosGovernanceAddress),
				AERUMTechnologyApiEndpoint: atmosEthereumRPCProvider,
			},
		},
	}

	// Figure out which consensus engine to choose
	fmt.Println()
	fmt.Println("Which consensus engine to use? As if you have a choice... Please type 1 or simply click ENTER.")
	fmt.Println(" 1. ATMOS - DxPoS consensus (delegated `cross-chain proof-of-stake`)")

	choice := w.read()
	switch {
	case len(choice) < 1 || choice == "1":

		genesis.Config.ChainID = new(big.Int).SetUint64(uint64(atmosNetID))

		var signers []common.Address

		for _, signer := range boostrapDelegate {
			signers = append(signers, signer)
		}
		// Sort the signers and embed into the extra-data section
		for i := 0; i < len(signers); i++ {
			for j := i + 1; j < len(signers); j++ {
				if bytes.Compare(signers[i][:], signers[j][:]) > 0 {
					signers[i], signers[j] = signers[j], signers[i]
				}
			}
		}
		genesis.ExtraData = make([]byte, 32+len(signers)*common.AddressLength+65)
		for i, signer := range signers {
			copy(genesis.ExtraData[32+i*common.AddressLength:], signer[:])
		}
	default:
		log.Crit("Invalid consensus engine choice", "choice", choice)
	}

	fmt.Println("\n\n[aerDEV] ----------------------------------------------------------- [aerDEV]")
	fmt.Println("[aerDEV] --- We have just preallocated some Aerum Coin to hard coded accounts --- [aerDEV]")
	fmt.Println("[aerDEV] ----------------------------------------------------------- [aerDEV]\n\n")

	aerumTeamAddress := map[string]string{
		"52c47938be22aab6f22b6608d9fe7f1e42aa8c61": "50000000000000000000000000",
		"56220873fb32f35a27f5e0f6604fda2aef439a5f": "10000000000000000000000000",
		"5ff4b0211b11d83b59d168bf1110a223a6e669fd": "10000000000000000000000000",
		"827720d8c7d3bab2566eee1ecc2207139dfb25af": "10000000000000000000000000",
		"8a3eefdbf626ae336272a379bddeb8dcad91d07b": "10000000000000000000000000",
		"b65b8c2376293fea13f6ed7d2a8467f0949c312d": "10000000000000000000000000",
	}

	for aerumTeamAddress, aerumTeamBalance := range aerumTeamAddress {
		bigaddr, _ := new(big.Int).SetString(aerumTeamAddress, 16)
		address := common.BigToAddress(bigaddr)

		bignum := new(big.Int)
		bignum.SetString(aerumTeamBalance, 10)

		genesis.Alloc[address] = core.GenesisAccount{
			Balance: bignum,
		}
	}

	// Add a batch of precompile balances to avoid them getting deleted
	for i := int64(0); i < 256; i++ {
		genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
	}

	// All done, store the genesis and flush to disk
	log.Info("Configured new genesis block")

	w.conf.Genesis = genesis
	w.conf.flush()
}

// manageGenesis permits the modification of chain configuration parameters in
// a genesis config and the export of the entire genesis spec.
func (w *wizard) manageGenesis() {
	// Figure out whether to modify or export the genesis
	fmt.Println()
	// fmt.Println(" 1. Modify existing fork rules")
	fmt.Println(" 1. Export genesis configuration")
	fmt.Println(" 2. Remove genesis configuration")

	choice := w.read()
	switch {
	// case choice == "1":
	// 	// Fork rule updating requested, iterate over each fork
	// 	fmt.Println()
	// 	fmt.Printf("Which block should Homestead come into effect? (default = %v)\n", w.conf.Genesis.Config.HomesteadBlock)
	// 	w.conf.Genesis.Config.HomesteadBlock = w.readDefaultBigInt(w.conf.Genesis.Config.HomesteadBlock)

	// 	fmt.Println()
	// 	fmt.Printf("Which block should EIP150 come into effect? (default = %v)\n", w.conf.Genesis.Config.EIP150Block)
	// 	w.conf.Genesis.Config.EIP150Block = w.readDefaultBigInt(w.conf.Genesis.Config.EIP150Block)

	// 	fmt.Println()
	// 	fmt.Printf("Which block should EIP155 come into effect? (default = %v)\n", w.conf.Genesis.Config.EIP155Block)
	// 	w.conf.Genesis.Config.EIP155Block = w.readDefaultBigInt(w.conf.Genesis.Config.EIP155Block)

	// 	fmt.Println()
	// 	fmt.Printf("Which block should EIP158 come into effect? (default = %v)\n", w.conf.Genesis.Config.EIP158Block)
	// 	w.conf.Genesis.Config.EIP158Block = w.readDefaultBigInt(w.conf.Genesis.Config.EIP158Block)

	// 	fmt.Println()
	// 	fmt.Printf("Which block should Byzantium come into effect? (default = %v)\n", w.conf.Genesis.Config.ByzantiumBlock)
	// 	w.conf.Genesis.Config.ByzantiumBlock = w.readDefaultBigInt(w.conf.Genesis.Config.ByzantiumBlock)

	// 	out, _ := json.MarshalIndent(w.conf.Genesis.Config, "", "  ")
	// 	fmt.Printf("Chain configuration updated:\n\n%s\n", out)

	case choice == "1":
		// Save whatever genesis configuration we currently have
		fmt.Println()
		fmt.Printf("Which file to save the genesis into? (default = %s.json)\n", w.network)
		out, _ := json.MarshalIndent(w.conf.Genesis, "", "  ")
		if err := ioutil.WriteFile(w.readDefaultString(fmt.Sprintf("%s.json", w.network)), out, 0644); err != nil {
			log.Error("Failed to save genesis file", "err", err)
		}
		log.Info("Exported existing genesis block")

	case choice == "2":
		// Make sure we don't have any services running
		if len(w.conf.servers()) > 0 {
			log.Error("Genesis reset requires all services and servers torn down")
			return
		}
		log.Info("Genesis block destroyed")

		w.conf.Genesis = nil
		w.conf.flush()

	default:
		log.Error("That's not something I can do")
	}
}
