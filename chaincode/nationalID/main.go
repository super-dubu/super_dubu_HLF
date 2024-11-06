package main

import (
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "nationalID/chaincode"
)

func main() {
    smartContract := new(nationalID.SmartContract)

    chaincode, err := contractapi.NewChaincode(smartContract)
    if err != nil {
        panic("Error creating nationalID chaincode: " + err.Error())
    }

    if err := chaincode.Start(); err != nil {
        panic("Error starting nationalID chaincode: " + err.Error())
    }
}

