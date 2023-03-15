package router

import (
	xlog "github.com/dabao-zhao/xgin/log"
	"github.com/gin-gonic/gin"

	"github.com/dabao-zhao/xgin-layout/app/helloworld/v1/server"
	"github.com/dabao-zhao/xgin-layout/config"
)

func Register(r *gin.Engine, conf *config.Config) {

	logger, err := xlog.NewLogger(&conf.Log)
	if err != nil {
		panic(err)
	}

	server.Register(r, conf, logger)
}
