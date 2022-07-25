package vo

type DocxRequestGenerate struct {
	Type       string         `json:"type" binding:"required,oneof=pdf docx"`
	Template   string         `json:"template" binding:"required"`
	Parameters map[string]any `json:"parameters" binding:"required"`
}
