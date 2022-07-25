package router

import (
	"github.com/gin-gonic/gin"
	"greport/pkgs/server/router/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	//gin.Default()
	r.Use(middlewares.Logger(), middlewares.Recovery())
	r.POST("/v1/apikey", GenerateApiKey)
	r.POST("/v1/template/docx/render", RenderDocxTemplate)
	return r
}
