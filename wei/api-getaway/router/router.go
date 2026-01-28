package router

import (
	"week3/wei/api-getaway/handler/server"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/goods/add", server.GoodsAdd)
	r.GET("/goods/get", server.GoodsShow)
	r.GET("/es", server.Es)
	return r
}
