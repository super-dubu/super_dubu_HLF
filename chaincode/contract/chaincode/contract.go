package contract

import (
    "fmt"
    "encoding/json"
    "encoding/base64"
    "crypto/sha256"
    "time"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a Contract
type SmartContract struct {
    contractapi.Contract
}

// ItemInfo 구조체
type ItemInfo struct {
	TokenID         string `json:"tokenID"`         // PNU 코드
	BuildingName    string `json:"buildingName"`    // 건물 이름
	Hosu            string `json:"hosu"`           // 호수
	BuildingAddress string `json:"buildingAddress"` // 건물 주소
	Area            string `json:"area"`           // 면적 (단위: m^2)
	PriceRental     string `json:"priceRental"`     // 보증금 (단위: 천 원)
	PriceMonthly    string `json:"priceMonthly"`    // 월세 (단위: 천 원)
	BuildingType    string `json:"buildingType"`    // 건물 타입 (0,1,2,3,4,5)
	ItemType        string `json:"itemType"`        // 매물 유형 (0: 월세, 1: 전세)
	FloorInfo       string `json:"floorInfo"`       // 층 정보
	AvailableDate   string `json:"availableDate"`   // 입주 가능 날짜
	RoomCount       string `json:"roomCount"`       // 방 개수
	ConfirmDate     string `json:"confirmDate"`     // 건축 승인 날짜
	Parking         string `json:"parking"`         // 주차 가능 대수
	ManageFee       string `json:"manageFee"`       // 관리비 (단위: 천 원)
	Body            string `json:"body"`           // 상세 정보
	ItemID          string `json:"itemID"`         // 매물 등록 번호
	Status          string `json:"status"`         // 상태
}

// AgentInfo 구조체
type AgentInfo struct {
	AgentAddress string `json:"agentAddress"` // 공인중개사무소 주소
	AgentName    string `json:"agentName"`    // 중개업자명
	AgentPhone   string `json:"agentPhone"`   // 전화번호
	RegisterID   string `json:"registerID"`   // 등록번호
	OfficeName   string `json:"officeName"`   // 사업자 상호
	RegisterDate string `json:"registerDate"` // 등록일자
}

// Main 구조체
type Contract struct {
	IsExist     string    `json:"isExist"`     // 존재 여부
	TradeDate   string    `json:"tradeDate"`   // 계약 날짜
	StartDate   string    `json:"startDate"`   // 시작 날짜
	EndDate     string    `json:"endDate"`     // 종료 날짜
	LessorName  string    `json:"lessorName"`  // 임대인 이름
	LessorPhone string    `json:"lessorPhone"` // 임대인 전화번호
	LesseeName  string    `json:"lesseeName"`  // 임차인 이름
	LesseePhone string    `json:"lesseePhone"` // 임차인 전화번호
	ItemInfo    ItemInfo  `json:"itemInfo"`    // 매물 정보
	AgentInfo   AgentInfo `json:"agentInfo"`   // 공인중개사 정보
	TxID        string    `json:"txID"`        // 트랜잭션 고유 번호
}

// CreateContract issues a new asset to the world state with given details.
func (s *SmartContract) CreateContract(ctx contractapi.TransactionContextInterface, data Contract) error {
    // 입력 데이터
    temp := data.TxID

    // SHA256 해시 계산
    hash := sha256.New()
    str1 := []byte(temp)
    str2 := []byte(time.Now().Format(time.RFC3339))
    hash.Write(str1) // 문자열을 바이트 배열로 변환하여 입력
    hash.Write(str2) // 문자열을 바이트 배열로 변환하여 입력
	
    hashBytes := hash.Sum(nil)

    // 결과를 16진수 문자열로 출력
    hashBase64 := base64.StdEncoding.EncodeToString(hashBytes)

    // 앞 10자리 추출
    shortHash := hashBase64[:10]
    data.TxID = shortHash
    exists, err := s.ContractExists(ctx, shortHash)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("the asset %s already exists", shortHash)
    }

    assetJSON, err := json.Marshal(data)
    if err != nil {
        return err
    }

    err = ctx.GetStub().PutState(shortHash, assetJSON)
    if err != nil {
        return fmt.Errorf("Failed to create event: %v", err)
    }

    err = ctx.GetStub().SetEvent("CreateTxComplete", assetJSON)
    if err != nil {
        return fmt.Errorf("Failed to create event: %v", err)
    }

    return err
}

// ReadContract returns the asset stored in the world state with given id.
func (s *SmartContract) ReadContract(ctx contractapi.TransactionContextInterface, id string) (*Contract, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("the asset %s does not exist", id)
    }

    var asset Contract
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }

    return &asset, nil
}

// UpdateContract updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateContract(ctx contractapi.TransactionContextInterface, id string) error {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return err
    }
    if assetJSON == nil {
        return fmt.Errorf("the asset %s does not exist", id)
    }

    var asset Contract
    err = json.Unmarshal(assetJSON, &asset)
    asset.ItemInfo.Status = "COMMITED"

    assetJSON, err = json.Marshal(asset)
    if err != nil {
        return err
    }
    err = ctx.GetStub().SetEvent("CreatebankTxComplete", assetJSON)
    if err != nil {
        return fmt.Errorf("Failed to create event: %v", err)
    }

    return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteContract deletes a given asset from the world state.
func (s *SmartContract) DeleteContract(ctx contractapi.TransactionContextInterface, id string) error {
    exists, err := s.ContractExists(ctx, id)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("the asset %s does not exist", id)
    }

    return ctx.GetStub().DelState(id)
}

// ContractExists returns true when asset with given ID exists in world state
func (s *SmartContract) ContractExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return false, fmt.Errorf("failed to read from world state: %v", err)
    }

    return assetJSON != nil, nil
}

// GetAllContracts returns all assets found in world state
func (s *SmartContract) GetAllContracts(ctx contractapi.TransactionContextInterface) ([]*Contract, error) {
    // range query with empty string for startKey and endKey does an
    // open-ended query of all assets in the chaincode namespace.
    resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()

    var assets []*Contract
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return nil, err
        }

        var asset Contract
        err = json.Unmarshal(queryResponse.Value, &asset)
        if err != nil {
            return nil, err
        }
        assets = append(assets, &asset)
    }

    return assets, nil
}

