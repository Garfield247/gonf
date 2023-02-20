package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Garfield247/gonf.git/config"
	"github.com/Garfield247/gonf.git/models"
	"github.com/Garfield247/gonf.git/routers"
	"github.com/Garfield247/gonf.git/utils/rediscli"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("start")
	config.SetUp()
	fmt.Println(config.LogCfg)
	models.SetUp()
	rediscli.SetUp()
}

func main() {
	gin.SetMode(config.ServerCfg.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := config.ServerCfg.ReadTimeout
	writeTimeout := config.ServerCfg.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.ServerCfg.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
