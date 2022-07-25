package router

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
	"github.com/xeipuuv/gojsonschema"
	"greport/pkgs/docx"
	"greport/pkgs/server/router/vo"
	"greport/pkgs/storage"
	"io/ioutil"
	"net/http"
)

func RenderDocxTemplate(c *gin.Context) {
	var request vo.DocxRequestGenerate
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, vo.ErrorResponse{
			Message: err.Error(),
			Detail:  nil,
		})
	}
	client, err := storage.GetClient(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, vo.ErrorResponse{Message: "can't connect to minio"})
		return
	}
	templateObj, err := client.GetObject(c, "docx", request.Template+".docx", minio.GetObjectOptions{})
	if err != nil {
		return
	}
	templateObjStat, err := templateObj.Stat()
	if err != nil {
		return
	}
	template, err := docx.Parse(templateObj, templateObjStat.Size)
	if err != nil {
		return
	}
	schemaObj, err := client.GetObject(c, "docx", request.Template+".schema.json", minio.GetObjectOptions{})
	if err == nil {
		schemaData, err := ioutil.ReadAll(schemaObj)
		if err == nil {
			schema, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(schemaData))
			if err == nil {
				result, err := schema.Validate(gojsonschema.NewGoLoader(request.Parameters))
				if err == nil && !result.Valid() {
					log.Debug().
						Str("template", request.Template).
						Interface("errors", result.Errors()).
						Msg("invalid parameter")
					c.JSON(http.StatusBadRequest, vo.ErrorResponse{Message: "invalid parameters"})
					return
				}
			}
		}
		return
	}

	var (
		data        []byte
		contentType string
	)
	switch request.Type {
	case "pdf":
		contentType = "application/pdf"
		data, err = template.RenderPdf(request.Parameters)
	case "docx":
		contentType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
		data, err = template.Render(request.Parameters)
	}
	c.Data(http.StatusOK, contentType, data)
}
