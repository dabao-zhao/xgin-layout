package server

import (
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	
	"github.com/dabao-zhao/xgin-layout/app/helloworld/v1"
	"github.com/dabao-zhao/xgin-layout/config"
)

func Register(r *gin.Engine, conf *config.Config, logger *logging.Logger) {
	registerGreeter(r, conf, logger)
}

func registerGreeter(r *gin.Engine, conf *config.Config, logger *logging.Logger) {
	handler := v1.InitializeGreeterHandler(conf, logger)
	r.GET("/", handler.Hello())
}
