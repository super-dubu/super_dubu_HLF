package bank

import (
    "encoding/json"
    "fmt"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a bankTx
type SmartContract struct {
    contractapi.Contract
}


type bankTx struct {
  TokenID              string         `json:"tokenID`
	MID                  string         `json:"mId"`
	LastTransactionKey   string         `json:"lastTransactionKey"`
	PaymentKey           string         `json:"paymentKey"`
	OrderID              string         `json:"orderId"`
	OrderName            string         `json:"orderName"`
	TaxExemptionAmount   int            `json:"taxExemptionAmount"`
	Status               string         `json:"status"`
	RequestedAt          string         `json:"requestedAt"`
	ApprovedAt           string         `json:"approvedAt"`
	UseEscrow            bool           `json:"useEscrow"`
	CultureExpense       bool           `json:"cultureExpense"`
	Card                 *CardInfo      `json:"card"`
	VirtualAccount       interface{}    `json:"virtualAccount"`
	Transfer             interface{}    `json:"transfer"`
	MobilePhone          interface{}    `json:"mobilePhone"`
	GiftCertificate      interface{}    `json:"giftCertificate"`
	CashReceipt          interface{}    `json:"cashReceipt"`
	CashReceipts         interface{}    `json:"cashReceipts"`
	Discount             interface{}    `json:"discount"`
	Cancels              interface{}    `json:"cancels"`
	Secret               string         `json:"secret"`
	Type                 string         `json:"type"`
	EasyPay              *EasyPayInfo   `json:"easyPay"`
	Country              string         `json:"country"`
	Failure              interface{}    `json:"failure"`
	IsPartialCancelable  bool           `json:"isPartialCancelable"`
	Receipt              *ReceiptInfo   `json:"receipt"`
	Checkout             *CheckoutInfo  `json:"checkout"`
	Currency             string         `json:"currency"`
	TotalAmount          int            `json:"totalAmount"`
	BalanceAmount        int            `json:"balanceAmount"`
	SuppliedAmount       int            `json:"suppliedAmount"`
	VAT                  int            `json:"vat"`
	TaxFreeAmount        int            `json:"taxFreeAmount"`
	Method               string         `json:"method"`
	Version              string         `json:"version"`
	Metadata             interface{}    `json:"metadata"`
}

type CardInfo struct {
	IssuerCode            string      `json:"issuerCode"`
	AcquirerCode          string      `json:"acquirerCode"`
	Number                string      `json:"number"`
	InstallmentPlanMonths int         `json:"installmentPlanMonths"`
	IsInterestFree        bool        `json:"isInterestFree"`
	InterestPayer         interface{} `json:"interestPayer"`
	ApproveNo             string      `json:"approveNo"`
	UseCardPoint          bool        `json:"useCardPoint"`
	CardType              string      `json:"cardType"`
	OwnerType             string      `json:"ownerType"`
	AcquireStatus         string      `json:"acquireStatus"`
	Amount                int         `json:"amount"`
}

type EasyPayInfo struct {
	Provider       string `json:"provider"`
	Amount         int    `json:"amount"`
	DiscountAmount int    `json:"discountAmount"`
}

type ReceiptInfo struct {
	URL string `json:"url"`
}

type CheckoutInfo struct {
	URL string `json:"url"`
}

// CreatebankTx issues a new asset to the world state with given details.
func (s *SmartContract) CreateBankTx(ctx contractapi.TransactionContextInterface, data bankTx) error {
    exists, err := s.bankTxExists(ctx, data.OrderID)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("the asset %s already exists", data.OrderID)
    }

    assetJSON, err := json.Marshal(data)
    if err != nil {
        return err
    }

    err = ctx.GetStub().PutState(data.OrderID, assetJSON)
    if err != nil {
        return fmt.Errorf("Failed to create state: %v", err)
    }

    err = ctx.GetStub().SetEvent("UpdateContract", assetJSON)
    if err != nil {
        return fmt.Errorf("Failed to create event: %v", err)
    }

    return err
}

// ReadbankTx returns the asset stored in the world state with given id.
func (s *SmartContract) ReadbankTx(ctx contractapi.TransactionContextInterface, id string) (*bankTx, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state: %v", err)
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("the asset %s does not exist", id)
    }

    var asset bankTx
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }

    return &asset, nil
}

// UpdatebankTx updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdatebankTx(ctx contractapi.TransactionContextInterface, data bankTx) error {
    exists, err := s.bankTxExists(ctx, data.OrderID)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("the asset %s does not exist", data.OrderID)
    }

    assetJSON, err := json.Marshal(data)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(data.OrderID, assetJSON)
}

// DeletebankTx deletes a given asset from the world state.
func (s *SmartContract) DeletebankTx(ctx contractapi.TransactionContextInterface, id string) error {
    exists, err := s.bankTxExists(ctx, id)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("the asset %s does not exist", id)
    }

    return ctx.GetStub().DelState(id)
}

// bankTxExists returns true when asset with given ID exists in world state
func (s *SmartContract) bankTxExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return false, fmt.Errorf("failed to read from world state: %v", err)
    }

    return assetJSON != nil, nil
}

// GetAllbankTxs returns all assets found in world state
func (s *SmartContract) GetAllbankTxs(ctx contractapi.TransactionContextInterface) ([]*bankTx, error) {
    // range query with empty string for startKey and endKey does an
    // open-ended query of all assets in the chaincode namespace.
    resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()

    var assets []*bankTx
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return nil, err
        }

        var asset bankTx
        err = json.Unmarshal(queryResponse.Value, &asset)
        if err != nil {
            return nil, err
        }
        assets = append(assets, &asset)
    }

    return assets, nil
}

