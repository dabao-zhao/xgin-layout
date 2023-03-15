//go:build wireinject
// +build wireinject

package v1

import (
	"github.com/dabao-zhao/xgin-layout/app/helloworld/v1/biz"
	"github.com/dabao-zhao/xgin-layout/app/helloworld/v1/data"
	"github.com/dabao-zhao/xgin-layout/app/helloworld/v1/handlers"
	"github.com/dabao-zhao/xgin-layout/app/helloworld/v1/services"
	"github.com/dabao-zhao/xgin-layout/config"
	"github.com/google/wire"
	"github.com/op/go-logging"
)

func InitializeGreeterHandler(*config.Config, *logging.Logger) *handlers.GreeterHandler {
	wire.Build(services.ProviderSet, data.ProviderSet, biz.ProviderSet, handlers.ProviderSet)
	return nil
}
