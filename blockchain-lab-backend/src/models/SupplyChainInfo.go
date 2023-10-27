package models

type SupplyChainInfo struct {
	ProductID         string `json:"product_id"`
	Producer          string `json:"producer"`
	Distributor       string `json:"distributor"`
	Retailer          string `json:"retailer"`
	ManufacturingDate string `json:"manufacturing_date"`
	Describe          string `json:"describe"`
	ExpiryDate        string `json:"expiry_date"`
}
type CreateReq struct {
	Producer          string `json:"producer"`
	Distributor       string `json:"distributor"`
	Retailer          string `json:"retailer"`
	ManufacturingDate string `json:"manufacturing_date"`
	Describe          string `json:"describe"`
	ExpiryDate        string `json:"expiry_date"`
}
type Response struct {
	Code    int    `json:"code"`    // 业务响应状态码
	Message string `json:"message"` // 提示信息

}
