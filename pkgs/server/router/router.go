package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "greport/docs"
	"greport/pkgs/server/router/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	//gin.Default()
	r.Use(middlewares.Logger(), middlewares.Recovery())
	r.POST("/v1/apikey", GenerateApiKey)
	r.POST("/v1/template/docx/render", RenderDocxTemplate)
	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
