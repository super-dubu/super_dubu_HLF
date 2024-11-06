package buildingInfo

import (
    "encoding/json"
    "fmt"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
    "os"
)

// SmartContract provides functions for managing a BuildingInfo
type SmartContract struct {
    contractapi.Contract
}

// BuildingInfo describes basic details of what makes up a simple asset
type BuildingInfo struct {
  TokenID         string `json:"tokenID"`
	BuildingName    string `json:"buildingName"`
	Hosu            string `json:"hosu"`
	BuildingAddress string `json:"buildingAddress"`
	Area            string `json:"area"`
	BuildingPrice   string `json:"buildingPrice"`
	BuildingType    int    `json:"buildingType"`
	FloorInfo       string `json:"floorInfo"`
	RoomCount       string `json:"roomCount"`
	ConfirmDate     string `json:"comfirmDate"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	data, err := os.ReadFile("buildingInfo.json")
	if err != nil {
		return fmt.Errorf("failed to read input JSON file: %v", err)
	}

	var buildings []BuildingInfo
	err = json.Unmarshal(data, &buildings)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}
	for _, building := range buildings {
		buildingJSON, err := json.Marshal(building)
		if err != nil {
      fmt.Println("in marshal : %v", err)
			return fmt.Errorf("failed to marshal building data: %v", err)
		}
		err = ctx.GetStub().PutState(building.TokenID, buildingJSON)
		if err != nil {
      fmt.Println("in putstate : %v", err)
			return fmt.Errorf("failed to put building data to world state: %v", err)
		}
    fmt.Println(building.TokenID)
	}

	return nil
}

// CreateBuildingInfo issues a new asset to the world state with given details.
func (s *SmartContract) CreateBuildingInfo(ctx contractapi.TransactionContextInterface, id, name, hosu, address, area, price, floor, room, date string, buildingType int) error {
    exists, err := s.BuildingInfoExists(ctx, id)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("the asset %s already exists", id)
    }

    asset := BuildingInfo{
      TokenID:         id,
      BuildingName:    name,
      Hosu:            hosu,
      BuildingAddress: address,
      Area:            area,
      BuildingPrice:   price,
      BuildingType:    buildingType,
      FloorInfo:       floor,
      RoomCount:      room,
      ConfirmDate:     date,
    }
    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }
        fmt.Println("추가되었습니다 : %s\n", asset.TokenID)

    return ctx.GetStub().PutState(id, assetJSON)
}

// ReadBuildingInfo returns the asset stored in the world state with given id.
func (s *SmartContract) ReadBuildingInfo(ctx contractapi.TransactionContextInterface, id string) (*BuildingInfo, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("the asset %s does not exist", id)
    }

    var asset BuildingInfo
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }

    return &asset, nil
}

// UpdateBuildingInfo updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateBuildingInfo(ctx contractapi.TransactionContextInterface, id, name, hosu, address, area, price, floor, room, date string, buildingType int) error {
    exists, err := s.BuildingInfoExists(ctx, id)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("the asset %s does not exist", id)
    }

    // overwriting original asset with new asset
    asset := BuildingInfo{
      TokenID:         id,
      BuildingName:    name,
      Hosu:            hosu,
      BuildingAddress: address,
      Area:            area,
      BuildingPrice:   price,
      BuildingType:    buildingType,
      FloorInfo:       floor,
      RoomCount:       room,
      ConfirmDate:     date,
    }
    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteBuildingInfo deletes a given asset from the world state.
func (s *SmartContract) DeleteBuildingInfo(ctx contractapi.TransactionContextInterface, id string) error {
    exists, err := s.BuildingInfoExists(ctx, id)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("the asset %s does not exist", id)
    }

    return ctx.GetStub().DelState(id)
}

// BuildingInfoExists returns true when asset with given ID exists in world state
func (s *SmartContract) BuildingInfoExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return false, fmt.Errorf("failed to read from world state: %v", err)
    }

    return assetJSON != nil, nil
}

// GetAllBuildingInfos returns all assets found in world state
func (s *SmartContract) GetAllBuildingInfos(ctx contractapi.TransactionContextInterface) ([]*BuildingInfo, error) {
    // range query with empty string for startKey and endKey does an
    // open-ended query of all assets in the chaincode namespace.
    resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()

    var assets []*BuildingInfo
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return nil, err
        }

        var asset BuildingInfo
        err = json.Unmarshal(queryResponse.Value, &asset)
        if err != nil {
            return nil, err
        }
        assets = append(assets, &asset)
    }

    return assets, nil
}

