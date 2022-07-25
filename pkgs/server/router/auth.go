package router

import (
	"github.com/gin-gonic/gin"
	"greport/pkgs/apikey"
	"greport/pkgs/server/router/vo"
	"net/http"
)

// GenerateApiKey godoc
// @Summary API generate apikey
// @Description API generate apikey
// @Tags Template
// @Accept  json
// @Produce  json
// @Param	data	body	vo.GenerateApiKeyRequest	true	"data"
// @Success 200 {object} vo.GenerateApiKeyResponse
// @Router /v1/apikey [post]
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
