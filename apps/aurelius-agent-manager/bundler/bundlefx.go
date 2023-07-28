package bundler

import (
	"aurelius/apps/aurelius-agent-manager/pb_server"
	"aurelius/apps/aurelius-agent-manager/service"
	"aurelius/libs/aurelius-mysql/db_context"
	"aurelius/libs/aurelius-mysql/options"
	"context"
	"fmt"
	"go.uber.org/fx"
)

var Module = fx.Options(
	pb_server.Module,
	service.Module,
	fx.Provide(options.NewDbOptions),
	fx.Provide(db_context.NewDbContext),
	fx.Invoke(runApplication),
)

func runApplication(lifecycle fx.Lifecycle, pbServer *pb_server.PbServer) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go pbServer.Listen()

			fmt.Println("Server started")

			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping application")
			return nil
		},
	})

}
