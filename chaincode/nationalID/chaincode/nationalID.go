package nationalID

import (
    "encoding/json"
    "fmt"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "os"
)

// SmartContract provides functions for managing a NationalID
type SmartContract struct {
    contractapi.Contract
}

// NationalID describes basic details of what makes up a simple asset
type NationalID struct {
    Name    string `json:"name"`
    ID      string `json:"id"`
    Address string `json:"address"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	data, err := os.ReadFile("nationalID.json")
	if err != nil {
		return fmt.Errorf("failed to read input JSON file: %v", err)
	}

	var nationalIDs []NationalID
	err = json.Unmarshal(data, &nationalIDs)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}

	for _, nationalID := range nationalIDs {
		nationalIDJSON, err := json.Marshal(nationalID)
		if err != nil {
			return fmt.Errorf("failed to marshal national ID data: %v", err)
		}
		err = ctx.GetStub().PutState(nationalID.ID, nationalIDJSON)
		if err != nil {
			return fmt.Errorf("failed to put national ID data to world state: %v", err)
		}
	}

	return nil
}

// CreateNationalID issues a new asset to the world state with given details.
func (s *SmartContract) CreateNationalID(ctx contractapi.TransactionContextInterface, id, name, address string) error {
    exists, err := s.NationalIDExists(ctx, id)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("the asset %s already exists", id)
    }

    asset := NationalID{
        ID:      id,
        Name:    name,
        Address: address,
    }
    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }
        fmt.Println("추가되었습니다 : %s\n", asset.ID)

    return ctx.GetStub().PutState(id, assetJSON)
}

// ReadNationalID returns the asset stored in the world state with given id.
func (s *SmartContract) ReadNationalID(ctx contractapi.TransactionContextInterface, id string) (*NationalID, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("the asset %s does not exist", id)
    }

    var asset NationalID
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }

    return &asset, nil
}

// UpdateNationalID updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateNationalID(ctx contractapi.TransactionContextInterface, id, name, address string) error {
    exists, err := s.NationalIDExists(ctx, id)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("the asset %s does not exist", id)
    }

    // overwriting original asset with new asset
    asset := NationalID{
        ID:      id,
        Name:    name,
        Address: address,
    }
    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteNationalID deletes a given asset from the world state.
func (s *SmartContract) DeleteNationalID(ctx contractapi.TransactionContextInterface, id string) error {
    exists, err := s.NationalIDExists(ctx, id)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("the asset %s does not exist", id)
    }

    return ctx.GetStub().DelState(id)
}

// NationalIDExists returns true when asset with given ID exists in world state
func (s *SmartContract) NationalIDExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return false, fmt.Errorf("failed to read from world state: %v", err)
    }

    return assetJSON != nil, nil
}

// GetAllNationalIDs returns all assets found in world state
func (s *SmartContract) GetAllNationalIDs(ctx contractapi.TransactionContextInterface) ([]*NationalID, error) {
    // range query with empty string for startKey and endKey does an
    // open-ended query of all assets in the chaincode namespace.
    resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()

    var assets []*NationalID
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return nil, err
        }

        var asset NationalID
        err = json.Unmarshal(queryResponse.Value, &asset)
        if err != nil {
            return nil, err
        }
        assets = append(assets, &asset)
    }

    return assets, nil
}

