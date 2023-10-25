package main

import (
	//"food-scanning/controller"
	"food-supply-chain/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// 创建和配置供应链记录和扫描验证的控制器
	supplyChainController := controller.NewSupplyChainController()
	//scanningController := controller.NewScanningController()

	r := mux.NewRouter()
	r.HandleFunc("/supplychain/record", supplyChainController.RecordSupplyChainInfo).Methods("POST")
	r.HandleFunc("/supplychain/query", supplyChainController.QuerySupplyChainInfo).Methods("GET")
	//r.HandleFunc("/scanning/scan", scanningController.ScanAndVerifyFood).Methods("POST")

	// 启动HTTP服务器
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
