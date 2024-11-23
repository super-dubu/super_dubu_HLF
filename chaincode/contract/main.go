package main

import (
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "contract/chaincode"
)

func main() {
    smartContract := new(contract.SmartContract)

    chaincode, err := contractapi.NewChaincode(smartContract)
    if err != nil {
        panic("Error creating contract chaincode: " + err.Error())
    }

    if err := chaincode.Start(); err != nil {
        panic("Error starting contract chaincode: " + err.Error())
    }
}

