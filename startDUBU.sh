#!/bin/bash
export PATH=/home/superDUBU/hyperledger/dubunet/bin:$PATH
export FABRIC_CFG_PATH=/home/superDUBU/hyperledger/dubunet/config
export GOPATH=/home/superDUBU/hyperledger/dubunet/dubu/chaincode
export CORE_PEER_TLS_ENABLED=true

./network.sh up createChannel
cd addOrg3
./addOrg3.sh up -c dubu
cd ..

# Set env for Org1
export CORE_PEER_LOCALMSPID=Org1MSP
export CORE_PEER_TLS_ROOTCERT_FILE=/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org1.dubu.com/peers/peer0.org1.dubu.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org1.dubu.com/users/Admin@org1.dubu.com/msp
export CORE_PEER_ADDRESS=localhost:7051


# Install CC Org1, Org2
./network.sh deployCC -ccn nationalID -ccp chaincode/nationalID -ccl go
./network.sh deployCC -ccn agentInfo -ccp chaincode/agentInfo -ccl go
./network.sh deployCC -ccn buildingInfo -ccp chaincode/buildingInfo -ccl go
./network.sh deployCC -ccn bank -ccp chaincode/bank -ccl go


# Set env for Org3
export CORE_PEER_LOCALMSPID=Org3MSP
export CORE_PEER_TLS_ROOTCERT_FILE=/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org3.dubu.com/peers/peer0.org3.dubu.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org3.dubu.com/users/Admin@org3.dubu.com/msp
export CORE_PEER_ADDRESS=localhost:11051

# Install the chaincode(nationalID)
peer lifecycle chaincode package nationalID.tar.gz --path chaincode/nationalID --lang golang --label nationalID_1.0.1
peer lifecycle chaincode install nationalID.tar.gz
peer lifecycle chaincode queryinstalled
export CC_PACKAGE_ID=$(peer lifecycle chaincode calculatepackageid nationalID.tar.gz)
echo "CC_PACKAGE_ID is set to: $CC_PACKAGE_ID"
peer lifecycle chaincode approveformyorg -o localhost:7050 \
--ordererTLSHostnameOverride orderer.dubu.com \
--tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" \
--channelID dubu --name nationalID --version 1.0.1 --package-id $CC_PACKAGE_ID --sequence 1


# Install the chaincode(agentInfo)
peer lifecycle chaincode package agentInfo.tar.gz --path chaincode/agentInfo --lang golang --label agentInfo_1.0.1
peer lifecycle chaincode install agentInfo.tar.gz
peer lifecycle chaincode queryinstalled
export CC_PACKAGE_ID=$(peer lifecycle chaincode calculatepackageid agentInfo.tar.gz)
echo "CC_PACKAGE_ID is set to: $CC_PACKAGE_ID"
peer lifecycle chaincode approveformyorg -o localhost:7050 \
--ordererTLSHostnameOverride orderer.dubu.com \
--tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" \
--channelID dubu --name agentInfo --version 1.0.1 --package-id $CC_PACKAGE_ID --sequence 1


# Install the chaincode(buildingInfo)
peer lifecycle chaincode package buildingInfo.tar.gz --path chaincode/buildingInfo --lang golang --label buildingInfo_1.0.1
peer lifecycle chaincode install buildingInfo.tar.gz
peer lifecycle chaincode queryinstalled
export CC_PACKAGE_ID=$(peer lifecycle chaincode calculatepackageid buildingInfo.tar.gz)
echo "CC_PACKAGE_ID is set to: $CC_PACKAGE_ID"
peer lifecycle chaincode approveformyorg -o localhost:7050 \
--ordererTLSHostnameOverride orderer.dubu.com \
--tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" \
--channelID dubu --name buildingInfo --version 1.0.1 --package-id $CC_PACKAGE_ID --sequence 1


# Install the chaincode(bank)
peer lifecycle chaincode package bank.tar.gz --path chaincode/bank --lang golang --label bank_1.0.1
peer lifecycle chaincode install bank.tar.gz
peer lifecycle chaincode queryinstalled
export CC_PACKAGE_ID=$(peer lifecycle chaincode calculatepackageid bank.tar.gz)
echo "CC_PACKAGE_ID is set to: $CC_PACKAGE_ID"
peer lifecycle chaincode approveformyorg -o localhost:7050 \
--ordererTLSHostnameOverride orderer.dubu.com \
--tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" \
--channelID dubu --name bank --version 1.0.1 --package-id $CC_PACKAGE_ID --sequence 1

# Copy json file
ID1=$(docker ps -f "name=^dev-peer0.org1.dubu.com-nationalID" --format "{{.ID}}")
ID2=$(docker ps -f "name=^dev-peer0.org2.dubu.com-nationalID" --format "{{.ID}}")
ID3=$(docker ps -f "name=^dev-peer0.org3.dubu.com-nationalID" --format "{{.ID}}")
docker cp ./nationalID.json $ID1:/
docker cp ./nationalID.json $ID2:/
docker cp ./nationalID.json $ID3:/

ID1=$(docker ps -f "name=^dev-peer0.org1.dubu.com-agentInfo" --format "{{.ID}}")
ID2=$(docker ps -f "name=^dev-peer0.org2.dubu.com-agentInfo" --format "{{.ID}}")
ID3=$(docker ps -f "name=^dev-peer0.org3.dubu.com-agentInfo" --format "{{.ID}}")
docker cp ./agentInfo.json $ID1:/
docker cp ./agentInfo.json $ID2:/
docker cp ./agentInfo.json $ID3:/


# Set env for Org1
export CORE_PEER_LOCALMSPID=Org1MSP
export CORE_PEER_TLS_ROOTCERT_FILE=/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org1.dubu.com/peers/peer0.org1.dubu.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org1.dubu.com/users/Admin@org1.dubu.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# Init nationalID
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dubu.com --tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" -C dubu -n nationalID --peerAddresses localhost:7051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org1.dubu.com/peers/peer0.org1.dubu.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org2.dubu.com/peers/peer0.org2.dubu.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'

# Init agentInfo
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dubu.com --tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" -C dubu -n agentInfo --peerAddresses localhost:7051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org1.dubu.com/peers/peer0.org1.dubu.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org2.dubu.com/peers/peer0.org2.dubu.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'

ID1=$(docker ps -f "name=^dev-peer0.org1.dubu.com-buildingInfo" --format "{{.ID}}")
ID2=$(docker ps -f "name=^dev-peer0.org2.dubu.com-buildingInfo" --format "{{.ID}}")
ID3=$(docker ps -f "name=^dev-peer0.org3.dubu.com-buildingInfo" --format "{{.ID}}")
docker cp ./buildingInfo.json $ID1:/
docker cp ./buildingInfo.json $ID2:/
docker cp ./buildingInfo.json $ID3:/

# Init buildingInfo
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dubu.com --tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" -C dubu -n buildingInfo --peerAddresses localhost:7051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org1.dubu.com/peers/peer0.org1.dubu.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org2.dubu.com/peers/peer0.org2.dubu.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'

# Set env for Org3
export CORE_PEER_LOCALMSPID=Org3MSP
export CORE_PEER_TLS_ROOTCERT_FILE=/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org3.dubu.com/peers/peer0.org3.dubu.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org3.dubu.com/users/Admin@org3.dubu.com/msp
export CORE_PEER_ADDRESS=localhost:11051

# Init nationalID
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dubu.com --tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" -C dubu -n nationalID --peerAddresses localhost:9051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org2.dubu.com/peers/peer0.org2.dubu.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org3.dubu.com/peers/peer0.org3.dubu.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'

# Init agentInfo
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dubu.com --tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" -C dubu -n agentInfo --peerAddresses localhost:9051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org2.dubu.com/peers/peer0.org2.dubu.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org3.dubu.com/peers/peer0.org3.dubu.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'

# Init buildingInfo
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.dubu.com --tls --cafile "/home/superDUBU/hyperledger/dubunet/dubu/organizations/ordererOrganizations/dubu.com/orderers/orderer.dubu.com/msp/tlscacerts/tlsca.dubu.com-cert.pem" -C dubu -n buildingInfo --peerAddresses localhost:9051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org2.dubu.com/peers/peer0.org2.dubu.com/tls/ca.crt" --peerAddresses localhost:11051 --tlsRootCertFiles "/home/superDUBU/hyperledger/dubunet/dubu/organizations/peerOrganizations/org3.dubu.com/peers/peer0.org3.dubu.com/tls/ca.crt" -c '{"function":"InitLedger","Args":[]}'

chmod 774 -R organizations
