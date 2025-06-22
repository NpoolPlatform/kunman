package main

import (
	"context"

	"github.com/NpoolPlatform/kunman/api"
	"github.com/NpoolPlatform/kunman/cron/scheduler/scheduler"
	"github.com/NpoolPlatform/kunman/framework/action"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/wlog"
	"github.com/NpoolPlatform/kunman/gateway/webhook"
	basalapi "github.com/NpoolPlatform/kunman/mal/basal/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	cli "github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Flags:   Flags,
	Action: func(c *cli.Context) error {
		InitializeFlags(c)

		err := action.Run(
			c.Context,
			run,
			rpcRegister,
			rpcGatewayRegister,
			watch,
		)

		scheduler.Finalize(c.Context)

		return err
	},
}

func run(ctx context.Context) error {
	return nil
}

func shutdown(ctx context.Context) {
	<-ctx.Done()
	logger.Sugar().Infow(
		"Watch",
		"State", "Done",
		"Error", ctx.Err(),
	)
}

func watch(ctx context.Context, cancel context.CancelFunc) error {
	go shutdown(ctx)

	scheduler.Initialize(ctx, cancel)
	webhook.Initialize()

	return nil
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)

	_ = basalapi.RegisterGRPC(server)
	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	err := api.RegisterGateway(mux, endpoint, opts)
	if err != nil {
		return wlog.WrapError(err)
	}

	_ = basalapi.Register(mux)
	return nil
}
