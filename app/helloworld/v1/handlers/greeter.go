package handlers

import (
	"github.com/dabao-zhao/xgin-layout/app/helloworld/v1/services"
	"github.com/dabao-zhao/xgin-layout/app/helloworld/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

type GreeterHandler struct {
	s   *services.GreeterService
	log *logging.Logger
}

func NewGreeterHandler(s *services.GreeterService, logger *logging.Logger) *GreeterHandler {
	return &GreeterHandler{
		s:   s,
		log: logger,
	}
}

func (h *GreeterHandler) Hello() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hello, err := h.s.SayHello(ctx, &types.GreeterRequest{})
		if err != nil {
			return
		}
		ctx.JSON(200, hello)
	}
}
