package app

import (
	"blockchain-lab/app/fabric"
	"blockchain-lab/models"
	"blockchain-lab/util"
	"blockchain-lab/util/conf"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// HandleCreate
// @Summary 创建一个信息
// @Description 根据上传json创建
// @Accept application/json
// @Produce application/json
// @Param object query models.CreateReq false "查询参数"
// @Success 200 {object} models.Response
// @Router /supplychain/create [post]
func HandleCreate(c *gin.Context) {

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := util.NewGrpcConnection(conf.Peer1)
	defer clientConnection.Close()

	id := util.NewIdentity(conf.Peer1)
	sign := util.NewSign(conf.Peer1)

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	// Override default values for chaincode and channel name as they may differ in testing contexts.
	chaincodeName := "basic"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	req := models.CreateReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": http.StatusUnprocessableEntity, "msg": err.Error()})
		return
	}
	productId := fmt.Sprintf("asset%8v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000000))
	info := models.SupplyChainInfo{
		ProductID:         productId,
		Producer:          req.Producer,
		Distributor:       req.Distributor,
		Retailer:          req.Retailer,
		Describe:          req.Describe,
		ManufacturingDate: req.ManufacturingDate,
		ExpiryDate:        req.ExpiryDate,
	}
	err = fabric.CreateAsset(contract, info)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"id": productId, "code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": productId, "code": http.StatusOK, "msg": "success"})

}

// HandleQueryById
// @Summary 查询一个信息
// @Description 根据id查询一个信息
// @Produce application/json
// @Param id query string false "查询ID"
// @Success 200 {object} models.Response
// @Router /supplychain/:id [get]
func HandleQueryById(c *gin.Context) {

	clientConnection := util.NewGrpcConnection(conf.Peer1)
	defer clientConnection.Close()

	id := util.NewIdentity(conf.Peer1)
	sign := util.NewSign(conf.Peer1)

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	// Override default values for chaincode and channel name as they may differ in testing contexts.
	chaincodeName := "basic"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	product_id := c.Param("id")

	infobyte, err := fabric.ReadAssetByID(contract, product_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}
	c.Writer.Write(infobyte)
	c.Header("Content-Type", "application/json")
}

// HandleQueryById
// @Summary 查询所有信息
// @Description 查询所有信息
// @Produce application/json
// @Success 200 {object} models.Response
// @Router /supplychains [get]
func HandleQueryAll(c *gin.Context) {

	clientConnection := util.NewGrpcConnection(conf.Peer1)
	defer clientConnection.Close()

	id := util.NewIdentity(conf.Peer1)
	sign := util.NewSign(conf.Peer1)

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	// Override default values for chaincode and channel name as they may differ in testing contexts.
	chaincodeName := "basic"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	infobyte, err := fabric.GetAllAssets(contract)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}
	c.Writer.Write(infobyte)
	c.Header("Content-Type", "application/json")
}

// HandleCreate
// @Summary 创建一个信息
// @Description 根据上传json创建
// @Accept application/json
// @Produce application/json
// @Param object query models.SupplyChainInfo false "查询参数"
// @Success 200 {object} models.Response
// @Router /supplychain/:id [put]
func HandleUpdate(c *gin.Context) {

	clientConnection := util.NewGrpcConnection(conf.Peer1)
	defer clientConnection.Close()

	id := util.NewIdentity(conf.Peer1)
	sign := util.NewSign(conf.Peer1)

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	// Override default values for chaincode and channel name as they may differ in testing contexts.
	chaincodeName := "basic"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	info := models.SupplyChainInfo{}
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": http.StatusUnprocessableEntity, "msg": err.Error()})
		return
	}
	err = fabric.UpdateAsset(contract, info)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
}

// HandleDelete
// @Summary 删除一个信息
// @Description 根据id删除一个信息
// @Produce application/json
// @Param id query string false "要删除信息的ID"
// @Success 200 {object} models.Response
// @Router /supplychain/:id [delete]
func HandleDelete(c *gin.Context) {
	clientConnection := util.NewGrpcConnection(conf.Peer1)
	defer clientConnection.Close()

	id := util.NewIdentity(conf.Peer1)
	sign := util.NewSign(conf.Peer1)

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	// Override default values for chaincode and channel name as they may differ in testing contexts.
	chaincodeName := "basic"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	product_idid := c.Param("id")
	fmt.Println(product_idid)
	err = fabric.DeleteAsset(contract, product_idid)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
}
