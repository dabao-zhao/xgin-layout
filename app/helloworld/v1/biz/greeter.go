package biz

import (
	"context"

	"github.com/op/go-logging"
)

type (
	GreeterRepo interface {
		Hello(ctx context.Context) string
	}
	GreeterUsecase struct {
		repo GreeterRepo
		log  *logging.Logger
	}
)

func NewGreeterUsecase(repo GreeterRepo, logger *logging.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: logger}
}

func (uc *GreeterUsecase) Hello(ctx context.Context) string {
	return uc.repo.Hello(ctx)
}
