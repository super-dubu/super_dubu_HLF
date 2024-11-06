#!/bin/bash
# Set environment variables for Hyperledger Fabric Peer

export PATH=/home/superDUBU/hyperledger/dubunet/bin:$PATH
export FABRIC_CFG_PATH=/home/superDUBU/hyperledger/dubunet/config
export GOPATH=/home/superDUBU/hyperledger/dubunet/dubu/chaincode
export CORE_PEER_TLS_ENABLED=true


export CORE_PEER_LOCALMSPID=Org1MSP
export CORE_PEER_TLS_ROOTCERT_FILE=/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org1.dubu.com/peers/peer0.org1.dubu.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org1.dubu.com/users/Admin@org1.dubu.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# 확인용 출력
echo "CORE_PEER_LOCALMSPID=$CORE_PEER_LOCALMSPID"
echo "CORE_PEER_TLS_ROOTCERT_FILE=$CORE_PEER_TLS_ROOTCERT_FILE"
echo "CORE_PEER_MSPCONFIGPATH=$CORE_PEER_MSPCONFIGPATH"
echo "CORE_PEER_ADDRESS=$CORE_PEER_ADDRESS"

echo "Set to Env for Org1"
