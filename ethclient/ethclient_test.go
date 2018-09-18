// Copyright 2016 The go-aerum Authors
// This file is part of the go-aerum library.
//
// The go-aerum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-aerum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-aerum library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import "github.com/AERUMTechnology/go-aerum"

// Verify that Client implements the AERUMTechnology interfaces.
var (
	_ = AERUMTechnology.ChainReader(&Client{})
	_ = AERUMTechnology.TransactionReader(&Client{})
	_ = AERUMTechnology.ChainStateReader(&Client{})
	_ = AERUMTechnology.ChainSyncReader(&Client{})
	_ = AERUMTechnology.ContractCaller(&Client{})
	_ = AERUMTechnology.GasEstimator(&Client{})
	_ = AERUMTechnology.GasPricer(&Client{})
	_ = AERUMTechnology.LogFilterer(&Client{})
	_ = AERUMTechnology.PendingStateReader(&Client{})
	// _ = AERUMTechnology.PendingStateEventer(&Client{})
	_ = AERUMTechnology.PendingContractCaller(&Client{})
)
