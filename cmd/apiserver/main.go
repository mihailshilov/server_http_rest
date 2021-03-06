package main

import (
	"github.com/mihailshilov/server_http_rest/app/apiserver"
	"github.com/mihailshilov/server_http_rest/app/apiserver/model"

	logger "github.com/mihailshilov/server_http_rest/app/apiserver/logger"
)

func main() {

	config, err := model.NewConfig()
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	if err := apiserver.Start(config); err != nil {
		logger.ErrorLogger.Println(err)
	}

}
