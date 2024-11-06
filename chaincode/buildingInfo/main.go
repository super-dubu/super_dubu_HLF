package main

import (
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "buildingInfo/chaincode"
)

func main() {
    smartContract := new(buildingInfo.SmartContract)

    chaincode, err := contractapi.NewChaincode(smartContract)
    if err != nil {
        panic("Error creating buildingInfo chaincode: " + err.Error())
    }

    if err := chaincode.Start(); err != nil {
        panic("Error starting buildingInfo chaincode: " + err.Error())
    }
}

