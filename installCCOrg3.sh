#!/bin/bash

# Package the chaincode
peer lifecycle chaincode package nationalID.tar.gz --path chaincode/nationalID --lang golang --label nationalID_1.0.1

# Install the chaincode
peer lifecycle chaincode install nationalID.tar.gz

# Query installed chaincodes
peer lifecycle chaincode queryinstalled

# Calculate package ID and export it
export CC_PACKAGE_ID=$(peer lifecycle chaincode calculatepackageid nationalID.tar.gz)
echo "CC_PACKAGE_ID is set to: $CC_PACKAGE_ID"

# Approve the chaincode for the organization
peer lifecycle chaincode approveformyorg -o localhost:7050 \
--ordererTLSHostnameOverride orderer.dubu.com \
--tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" \
--channelID dubu --name nationalID --version 1.0.1 --package-id $CC_PACKAGE_ID --sequence 3

# Invoke the InitLedger function
peer chaincode invoke -o localhost:7050 \
--ordererTLSHostnameOverride orderer.dubu.com \
--tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" \
-C dubu -n nationalID \
--peerAddresses localhost:9051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org2.dubu.com/peers/peer0.org2.dubu.com/tls/ca.crt" \
--peerAddresses localhost:11051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org3.dubu.com/peers/peer0.org3.dubu.com/tls/ca.crt" \
-c '{"function":"InitLedger","Args":[]}'

