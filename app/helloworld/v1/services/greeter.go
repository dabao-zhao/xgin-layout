package services

import (
	"context"

	"github.com/op/go-logging"

	"github.com/dabao-zhao/xgin-layout/app/helloworld/v1/biz"
	"github.com/dabao-zhao/xgin-layout/app/helloworld/v1/types"
)

type GreeterService struct {
	uc  *biz.GreeterUsecase
	log *logging.Logger
}

func NewGreeterService(uc *biz.GreeterUsecase, logger *logging.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: logger}
}

func (s *GreeterService) SayHello(ctx context.Context, in *types.GreeterRequest) (*types.GreeterResponse, error) {
	msg := s.uc.Hello(ctx)
	return &types.GreeterResponse{
		Msg: msg,
	}, nil
}
