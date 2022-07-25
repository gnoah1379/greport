package router

import (
	"github.com/gin-gonic/gin"
	"greport/pkgs/apikey"
	"greport/pkgs/server/router/vo"
	"net/http"
)

func GenerateApiKey(c *gin.Context) {
	var req vo.GenerateApiKeyRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.ErrorResponse{
			Message: err.Error(),
			Detail:  nil,
		})
		return
	}
	apiKey, err := apikey.GetApiKey(req.AccessKeyID, req.SecretAccessKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.ErrorResponse{
			Message: err.Error(),
			Detail:  nil,
		})
		return
	}
	c.JSON(http.StatusOK, vo.GenerateApiKeyResponse{ApiKey: apiKey})
}
