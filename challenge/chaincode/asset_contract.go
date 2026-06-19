package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

// AssetExists checks if asset exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {

	// TODO:
	// 1. Use GetState(id)
	// 2. Return false and error if GetState fails
	// 3. Return true if asset exists
	// 4. Return false if asset does not exist

	return false, nil
}
