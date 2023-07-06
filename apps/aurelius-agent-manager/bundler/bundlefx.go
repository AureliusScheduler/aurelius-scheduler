package bundler

import (
	"aurelius/apps/aurelius-agent-manager/pb_server"
	"context"
	"fmt"
	"go.uber.org/fx"
)

var Module = fx.Options(
	pb_server.Module,
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
