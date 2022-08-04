package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"greport/pkgs/server/router"
	"os"
	"time"
)

// @title Document Template API
// @version 1.0
// @description C08 Document Template API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email nguyenhai.hoang1@etc.vn

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKey
// @in header
// @name Authorization
func main() {
	//pwd, _ := os.Getwd()
	//log.Debug().Msg("pwd: " + pwd)
	//data, err := ioutil.ReadFile("./TEST_MQD02.docx")
	//if err != nil {
	//	log.Error().Err(err).Msgf("open file error")
	//	return
	//}
	//template, err := docx.ParseBytes(data)
	//if err != nil {
	//	log.Error().Err(err).Msgf("parse template file error")
	//	return
	//}
	//pdf, err := template.RenderPdf(map[string]any{
	//	"donViCapTren":     "Phong Canh Sat Giao Thong Ha Noi",
	//	"donViRaQuyetDinh": "Doi tuan tra kiem soat so 1",
	//	"diaDanhHanhChinh": "Ha Noi",
	//	"soQD":             "MQD02-00000000000000001",
	//})
	//if err != nil {
	//	log.Error().Err(err).Msgf("render template failed")
	//	return
	//}
	//err = ioutil.WriteFile("./MQD02.pdf", pdf, 0666)
	//if err != nil {
	//	log.Error().Err(err).Msgf("write pdf file error")
	//	return
	//}
	//client, err := minio.New("minio.local:9000", &minio.Options{
	//	Creds:  credentials.NewStaticV4("admin", "Etc@12341", ""),
	//	Secure: false,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//_, err = client.HealthCheck(time.Second * 5)
	//if err != nil {
	//	panic(err)
	//
	//}

	//objInfo, err := client.StatObject(context.Background(), "abc", "hoang/test", minio.GetObjectOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(objInfo)
	viper.SetConfigFile("config.yml")
	viper.SetConfigType("yaml")
	_ = viper.ReadInConfig()

	r := router.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func init() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log.Logger = log.Output(consoleWriter).With().Caller().Logger()
}
