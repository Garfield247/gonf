package routers

import (
	"github.com/Garfield247/gonf.git/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.POST("/test", controller.InpectHeartBeat)
	return engine
}
