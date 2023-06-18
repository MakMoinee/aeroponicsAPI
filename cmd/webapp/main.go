package main

import (
	"aeroponicsAPI/cmd/webapp/config"
	"aeroponicsAPI/cmd/webapp/routes"
	"aeroponicsAPI/internal/aeroponicsapi/common"
	"log"

	"github.com/MakMoinee/go-mith/pkg/goserve"
)

func main() {
	config.Set()
	httpService := goserve.NewService(common.SERVER_PORT)
	routes.Set(httpService)
	log.Println("Server Started At Port ", common.SERVER_PORT)
	if err := httpService.Start(); err != nil {
		panic(err)
	}
}
