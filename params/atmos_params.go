package params

import (
	"github.com/AERUMTechnology/go-aerum/common"
	"math/big"
)

// Values for AERUMS Genesis related to ATMOS Consensus
var (
	atmosMinDelegateNo       = 10
	atmosNetID               = 418313827693
	atmosGovernanceAddress   = "0x5df4f6cde8f209ef8f152b59cabce70db21787b3"
	atmosBlockInterval       = uint64(2)
	atmosEpochInterval       = uint64(1000)
	atmosGasLimit            = uint64(25000000)
	atmosEthereumRPCProvider = "https://rinkeby.infura.io"
	atmosBlockRewards = big.NewInt(0.487e+18)
)

func NewAtmosMinDelegateNo() int {
	return atmosMinDelegateNo
}

func NewAtmosNetID() int {
	return atmosNetID
}

func NewAtmosGovernanceAddress() common.Address {
	return common.HexToAddress( atmosGovernanceAddress )
}

func NewAtmosBlockInterval() uint64 {
	return atmosBlockInterval
}

func NewAtmosEpochInterval() uint64 {
	return atmosEpochInterval
}

func NewAtmosGasLimit() uint64 {
	return atmosGasLimit
}

func NewAtmosEthereumRPCProvider() string {
	return atmosEthereumRPCProvider
}

func NewAtmosBlockRewards() *big.Int {
	return atmosBlockRewards
}


func NewAerumPreAlloc() map[string]string {
	aerumPreAlloc := map[string]string{
		"52c47938be22aab6f22b6608d9fe7f1e42aa8c61": "50000000000000000000000000",
		"56220873fb32f35a27f5e0f6604fda2aef439a5f": "10000000000000000000000000",
		"5ff4b0211b11d83b59d168bf1110a223a6e669fd": "10000000000000000000000000",
		"827720d8c7d3bab2566eee1ecc2207139dfb25af": "10000000000000000000000000",
		"8a3eefdbf626ae336272a379bddeb8dcad91d07b": "10000000000000000000000000",
		"b65b8c2376293fea13f6ed7d2a8467f0949c312d": "10000000000000000000000000",
	}	
	return aerumPreAlloc
}