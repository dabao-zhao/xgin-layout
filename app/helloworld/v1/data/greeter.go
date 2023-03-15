package data

import (
	"context"

	"github.com/op/go-logging"

	"github.com/dabao-zhao/xgin-layout/app/helloworld/v1/biz"
)

type (
	greeterRepo struct {
		data *Data
		log  *logging.Logger
	}
)

func NewGreeterRepo(data *Data, logger *logging.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  logger,
	}
}

func (r *greeterRepo) Hello(ctx context.Context) string {
	return "HELLO"
}
