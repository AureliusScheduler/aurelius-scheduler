package bundler

import (
	"aurelius/apps/aurelius-rest/controllers"
	"aurelius/apps/aurelius-rest/db"
	"aurelius/apps/aurelius-rest/routes"
	"aurelius/apps/aurelius-rest/service"
	"aurelius/libs/aurelius-mysql/bundler"
	"context"
	"fmt"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(db.NewDbOptions),
	bundler.Module,
	controllers.Module,
	service.Module,
	routes.Module,
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
