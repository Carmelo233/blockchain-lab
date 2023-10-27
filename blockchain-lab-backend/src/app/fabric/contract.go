package fabric

import (
	"blockchain-lab/models"
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// This type of transaction would typically only be run once by an application the first time it was started after its
// initial deployment. A new version of the model deployed later would likely not need to run an "init" function.
func InitLedger(contract *client.Contract) error {
	fmt.Printf("\n--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger \n")

	_, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		return err
	}

	fmt.Printf("*** Transaction committed successfully\n")
	return nil
}

// Evaluate a transaction to query ledger state.
func GetAllAssets(contract *client.Contract) ([]byte, error) {
	fmt.Println("\n--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		return nil, err
	}
	result := formatJSON(evaluateResult)

	return []byte(result), nil
}

// Submit a transaction synchronously, blocking until it has been committed to the ledger.
func CreateAsset(contract *client.Contract, asset models.SupplyChainInfo) error {
	fmt.Printf("\n--> Submit Transaction: CreateAsset, creates new asset with ID, Color, Size, Owner and AppraisedValue arguments \n")

	_, err := contract.SubmitTransaction("CreateAsset", asset.ProductID, asset.Producer, asset.Distributor, asset.Retailer, asset.ManufacturingDate, asset.Describe, asset.ExpiryDate)
	if err != nil {
		return err
	}

	fmt.Printf("*** Transaction committed successfully\n")
	return nil
}

// Evaluate a transaction by assetID to query ledger state.
func ReadAssetByID(contract *client.Contract, assetId string) ([]byte, error) {
	fmt.Printf("\n--> Evaluate Transaction: ReadAsset, function returns asset attributes\n")

	evaluateResult, err := contract.EvaluateTransaction("ReadAsset", assetId)
	if err != nil {
		return []byte(""), nil
	}
	result := formatJSON(evaluateResult)

	return []byte(result), nil

}

func UpdateAsset(contract *client.Contract, asset models.SupplyChainInfo) error {
	_, err := contract.SubmitTransaction("UpdateAsset", asset.ProductID, asset.Producer, asset.Distributor, asset.Retailer, asset.ManufacturingDate, asset.Describe, asset.ExpiryDate)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAsset(contract *client.Contract, assetId string) error {
	_, err := contract.SubmitTransaction("DeleteAsset", assetId)
	if err != nil {
		return err
	}

	return nil
}

// Submit transaction asynchronously, blocking until the transaction has been sent to the orderer, and allowing
// this thread to process the model response (e.g. update a UI) without waiting for the commit notification
// func transferAssetAsync(contract *client.Contract) {
// 	fmt.Printf("\n--> Async Submit Transaction: TransferAsset, updates existing asset owner")

// 	submitResult, commit, err := contract.SubmitAsync("TransferAsset", client.WithArguments(assetId, "Mark"))
// 	if err != nil {
// 		panic(fmt.Errorf("failed to submit transaction asynchronously: %w", err))
// 	}

// 	fmt.Printf("\n*** Successfully submitted transaction to transfer ownership from %s to Mark. \n", string(submitResult))
// 	fmt.Println("*** Waiting for transaction commit.")

// 	if commitStatus, err := commit.Status(); err != nil {
// 		panic(fmt.Errorf("failed to get commit status: %w", err))
// 	} else if !commitStatus.Successful {
// 		panic(fmt.Errorf("transaction %s failed to commit with status: %d", commitStatus.TransactionID, int32(commitStatus.Code)))
// 	}

// 	fmt.Printf("*** Transaction committed successfully\n")
// }

// // Submit transaction, passing in the wrong number of arguments ,expected to throw an error containing details of any error responses from the smart contract.
// func exampleErrorHandling(contract *client.Contract) {
// 	fmt.Println("\n--> Submit Transaction: UpdateAsset asset70, asset70 does not exist and should return an error")

// 	_, err := contract.SubmitTransaction("UpdateAsset", "asset70", "blue", "5", "Tomoko", "300")
// 	if err == nil {
// 		panic("******** FAILED to return an error")
// 	}

// 	fmt.Println("*** Successfully caught the error:")

// 	switch err := err.(type) {
// 	case *client.EndorseError:
// 		fmt.Printf("Endorse error for transaction %s with gRPC status %v: %s\n", err.TransactionID, status.Code(err), err)
// 	case *client.SubmitError:
// 		fmt.Printf("Submit error for transaction %s with gRPC status %v: %s\n", err.TransactionID, status.Code(err), err)
// 	case *client.CommitStatusError:
// 		if errors.Is(err, context.DeadlineExceeded) {
// 			fmt.Printf("Timeout waiting for transaction %s commit status: %s", err.TransactionID, err)
// 		} else {
// 			fmt.Printf("Error obtaining commit status for transaction %s with gRPC status %v: %s\n", err.TransactionID, status.Code(err), err)
// 		}
// 	case *client.CommitError:
// 		fmt.Printf("Transaction %s failed to commit with status %d: %s\n", err.TransactionID, int32(err.Code), err)
// 	default:
// 		panic(fmt.Errorf("unexpected error type %T: %w", err, err))
// 	}

// 	// Any error that originates from a peer or orderer node external to the gateway will have its details
// 	// embedded within the gRPC status error. The following code shows how to extract that.
// 	statusErr := status.Convert(err)

// 	details := statusErr.Details()
// 	if len(details) > 0 {
// 		fmt.Println("Error Details:")

// 		for _, detail := range details {
// 			switch detail := detail.(type) {
// 			case *gateway.ErrorDetail:
// 				fmt.Printf("- address: %s, mspId: %s, message: %s\n", detail.Address, detail.MspId, detail.Message)
// 			}
// 		}
// 	}
// }

// Format JSON data
func formatJSON(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		fmt.Println(data)
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}
