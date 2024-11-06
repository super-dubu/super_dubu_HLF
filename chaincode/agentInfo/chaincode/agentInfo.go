package agentInfo

import (
    "encoding/json"
    "fmt"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "os"
)

// SmartContract provides functions for managing a AgentInfo
type SmartContract struct {
    contractapi.Contract
}

// AgentInfo describes basic details of what makes up a simple asset
type AgentInfo struct {
    AgentName    string `json:"agentName"`
    RegisterID  string `json:"registerID"`
    OfficeName  string `json:"officeName"`
    RegisterDate  string `json:"registerDate"`
    AgentPhone    string `json:"agentPhone"`
    AgentAddress    string `json:"agentAddress"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	data, err := os.ReadFile("agentInfo.json")
	if err != nil {
		return fmt.Errorf("failed to read input JSON file: %v", err)
	}

	var agentInfos []AgentInfo
	err = json.Unmarshal(data, &agentInfos)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}

	for _, agentInfo := range agentInfos {
		agentInfoJSON, err := json.Marshal(agentInfo)
		if err != nil {
			return fmt.Errorf("failed to marshal national ID data: %v", err)
		}
		err = ctx.GetStub().PutState(agentInfo.RegisterID, agentInfoJSON)
		if err != nil {
			return fmt.Errorf("failed to put national ID data to world state: %v", err)
		}
	}

	return nil
}

// CreateAgentInfo issues a new asset to the world state with given details.
func (s *SmartContract) CreateAgentInfo(ctx contractapi.TransactionContextInterface, officeName, date, phone, id, agentName, address string) error {
    exists, err := s.AgentInfoExists(ctx, id)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("the asset %s already exists", id)
    }

    asset := AgentInfo{
        RegisterID:      id,
        AgentName:    agentName,
        AgentAddress: address,
        OfficeName: officeName,
        AgentPhone: phone,
        RegisterDate: date,
    }
    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }
        fmt.Println("추가되었습니다 : %s\n", asset.RegisterID)

    return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAgentInfo returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAgentInfo(ctx contractapi.TransactionContextInterface, registerID string) (*AgentInfo, error) {
    assetJSON, err := ctx.GetStub().GetState(registerID)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("the asset %s does not exist", registerID)
    }

    var asset AgentInfo
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }

    return &asset, nil
}

// UpdateAgentInfo updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAgentInfo(ctx contractapi.TransactionContextInterface, officeName, date, phone, id, agentName, address string) error {
    exists, err := s.AgentInfoExists(ctx, id)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("the asset %s does not exist", id)
    }

    // overwriting original asset with new asset
    asset := AgentInfo{
        RegisterID:      id,
        AgentName:    agentName,
        AgentAddress: address,
        OfficeName: officeName,
        AgentPhone: phone,
        RegisterDate: date,
    }
    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteAgentInfo deletes a given asset from the world state.
func (s *SmartContract) DeleteAgentInfo(ctx contractapi.TransactionContextInterface, id string) error {
    exists, err := s.AgentInfoExists(ctx, id)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("the asset %s does not exist", id)
    }

    return ctx.GetStub().DelState(id)
}

// AgentInfoExists returns true when asset with given ID exists in world state
func (s *SmartContract) AgentInfoExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return false, fmt.Errorf("failed to read from world state: %v", err)
    }

    return assetJSON != nil, nil
}

// GetAllAgentInfos returns all assets found in world state
func (s *SmartContract) GetAllAgentInfos(ctx contractapi.TransactionContextInterface) ([]*AgentInfo, error) {
    // range query with empty string for startKey and endKey does an
    // open-ended query of all assets in the chaincode namespace.
    resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()

    var assets []*AgentInfo
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return nil, err
        }

        var asset AgentInfo
        err = json.Unmarshal(queryResponse.Value, &asset)
        if err != nil {
            return nil, err
        }
        assets = append(assets, &asset)
    }

    return assets, nil
}

