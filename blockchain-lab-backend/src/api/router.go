package api

import (
	"blockchain-lab/app"
	_ "blockchain-lab/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
	router.POST("/supplychain/create", app.HandleCreate)
	router.GET("/supplychain/:id", app.HandleQueryById)
	router.GET("/supplychains", app.HandleQueryAll)
	router.PUT("/supplychain/:id", app.HandleUpdate)
	router.DELETE("/supplychain/:id", app.HandleDelete)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
