package vo

type GenerateApiKeyRequest struct {
	AccessKeyID     string `json:"accessKeyID" binding:"required"`
	SecretAccessKey string `json:"secretAccessKey" binding:"required"`
}

type GenerateApiKeyResponse struct {
	ApiKey string `json:"apiKey"`
}
