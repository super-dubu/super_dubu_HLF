package main

import (
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "bank/chaincode"
)

func main() {
    smartContract := new(bank.SmartContract)

    chaincode, err := contractapi.NewChaincode(smartContract)
    if err != nil {
        panic("Error creating bank chaincode: " + err.Error())
    }

    if err := chaincode.Start(); err != nil {
        panic("Error starting bank chaincode: " + err.Error())
    }
}

