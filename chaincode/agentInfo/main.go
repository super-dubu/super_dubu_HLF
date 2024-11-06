package main

import (
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "agentInfo/chaincode"
)

func main() {
    smartContract := new(agentInfo.SmartContract)

    chaincode, err := contractapi.NewChaincode(smartContract)
    if err != nil {
        panic("Error creating agentInfo chaincode: " + err.Error())
    }

    if err := chaincode.Start(); err != nil {
        panic("Error starting agentInfo chaincode: " + err.Error())
    }
}

