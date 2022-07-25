package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"greport/pkgs/server/router"
	"os"
	"time"
)

func main() {
	//pwd, _ := os.Getwd()
	//log.Debug().Msg("pwd: " + pwd)
	//data, err := ioutil.ReadFile("./MQD02.docx")
	//if err != nil {
	//	log.Error().Err(err).Msgf("open file error")
	//	return
	//}
	//template, err := docx.ParseBytes(data)
	//if err != nil {
	//	log.Error().Err(err).Msgf("parse template file error")
	//	return
	//}
	//pdf, err := template.RenderPdf(map[string]any{})
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
