package chaincode

import (
	"encoding/json"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SupplyChainContract struct {
}

type SupplyChainInfo struct {
	ProductID         string `json:"product_id"`
	Producer          string `json:"producer"`
	Distributor       string `json:"distributor"`
	Retailer          string `json:"retailer"`
	ManufacturingDate string `json:"manufacturing_date"`
	ExpiryDate        string `json:"expiry_date"`
}

func (s *SupplyChainContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return pb.Response(shim.Success(nil))
}

func (s *SupplyChainContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()

	if function == "recordSupplyChainInfo" {
		return s.recordSupplyChainInfo(APIstub, args)
	} else if function == "querySupplyChainInfo" {
		return s.querySupplyChainInfo(APIstub, args)
	}

	return pb.Response(shim.Error("Invalid function name"))
}

func (s *SupplyChainContract) recordSupplyChainInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return pb.Response(shim.Error("Incorrect number of arguments. Expecting 6"))
	}

	supplyChainInfo := SupplyChainInfo{
		ProductID:         args[0],
		Producer:          args[1],
		Distributor:       args[2],
		Retailer:          args[3],
		ManufacturingDate: args[4],
		ExpiryDate:        args[5],
	}

	supplyChainJSON, err := json.Marshal(supplyChainInfo)
	if err != nil {
		return pb.Response(shim.Error("Error creating supply chain record: " + err.Error()))
	}

	err = APIstub.PutState(args[0], supplyChainJSON)
	if err != nil {
		return pb.Response(shim.Error("Error recording supply chain information: " + err.Error()))
	}

	return pb.Response(shim.Success(nil))
}

func (s *SupplyChainContract) querySupplyChainInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return pb.Response(shim.Error("Incorrect number of arguments. Expecting ProductID to query"))
	}

	supplyChainInfoBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return pb.Response(shim.Error("Failed to get supply chain information: " + err.Error()))
	}

	if supplyChainInfoBytes == nil {
		return pb.Response(shim.Error("Supply chain information does not exist"))
	}

	return pb.Response(shim.Success(supplyChainInfoBytes))
}
