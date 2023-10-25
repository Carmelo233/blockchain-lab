package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type SupplyChainController struct {
	// 控制器的属性和依赖
	contract *gateway.Contract
}

type SupplyChainInfo struct {
	ProductID         string `json:"product_id"`
	Producer          string `json:"producer"`
	Distributor       string `json:"distributor"`
	Retailer          string `json:"retailer"`
	ManufacturingDate string `json:"manufacturing_date"`
	ExpiryDate        string `json:"expiry_date"`
}

func NewSupplyChainController() *SupplyChainController {
	// 创建新的控制器实例
	return &SupplyChainController{}
}

func (c *SupplyChainController) RecordSupplyChainInfo(w http.ResponseWriter, r *http.Request) {
	// 处理供应链信息记录请求
	vars := mux.Vars(r)
	productID := vars["productID"]
	producer := r.FormValue("producer")
	distributor := r.FormValue("distributor")
	retailer := r.FormValue("retailer")
	manufacturingDate := r.FormValue("manufacturingDate")
	expiryDate := r.FormValue("expiryDate")

	supplyChainInfo := SupplyChainInfo{
		ProductID:         productID,
		Producer:          producer,
		Distributor:       distributor,
		Retailer:          retailer,
		ManufacturingDate: manufacturingDate,
		ExpiryDate:        expiryDate,
	}

	supplyChainInfoJSON, err := json.Marshal(supplyChainInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := c.contract.SubmitTransaction("recordSupplyChainInfo", string(supplyChainInfoJSON))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	print(response)
	w.WriteHeader(http.StatusNoContent)
}

func (c *SupplyChainController) QuerySupplyChainInfo(w http.ResponseWriter, r *http.Request) {
	// 处理供应链信息查询请求
	vars := mux.Vars(r)
	productID := vars["productID"]

	response, err := c.contract.EvaluateTransaction("querySupplyChainInfo", productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
