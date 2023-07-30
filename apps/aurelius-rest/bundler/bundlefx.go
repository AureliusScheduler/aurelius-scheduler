package bundler

import (
	"aurelius/apps/aurelius-rest/controllers"
	"aurelius/apps/aurelius-rest/routes"
	"aurelius/apps/aurelius-rest/service"
	"aurelius/libs/aurelius-mysql/bundler"
	"aurelius/libs/aurelius-mysql/options"
	"aurelius/libs/aurelius_service"
	"context"
	"fmt"
	"go.uber.org/fx"
)

var Module = fx.Options(
	bundler.Module,
	controllers.Module,
	service.Module,
	routes.Module,
	fx.Provide(aurelius_service.NewJobService),
	fx.Provide(options.NewDbOptions),
	fx.Provide(func() aurelius_service.NowProvider {
		return aurelius_service.RealNowProvider{}
	}),
	fx.Invoke(runApplication),
)

func runApplication(lifecycle fx.Lifecycle, handler routes.Handler) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go handler.Gin.Run("localhost:8080")
			fmt.Println("Application started at http://localhost:8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping application")
			return nil
		},
	})

}
