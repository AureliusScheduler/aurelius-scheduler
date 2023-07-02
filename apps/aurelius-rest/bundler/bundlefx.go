package bundler

import (
	"aurelius/apps/aurelius-rest/controllers"
	"aurelius/apps/aurelius-rest/db"
	"aurelius/apps/aurelius-rest/routes"
	"aurelius/libs/aurelius-mysql"
	"context"
	"fmt"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(db.NewDbOptions),
	aurelius_mysql.Module,
	controllers.Module,
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
