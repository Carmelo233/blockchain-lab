package main

import (
	"blockchain-lab/api"
	"blockchain-lab/util/conf"
)

// @title Fabric
// @version 1.0
// @description fabric-crud接口文档
// @termsOfService http://swagger.io/terms/

// @contact.name 叶浩辉
// @contact.url blog.yehaohui.com
// @contact.email yhh1934292134@163.com

// @host 119.29.53.176:8080
// @BasePath /
func main() {
	conf.InitConf()

	router := api.NewRouter()
	router.Run(":8080")

}
