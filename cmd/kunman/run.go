package main

import (
	"context"

	// "github.com/NpoolPlatform/kunman/api"
	"github.com/NpoolPlatform/kunman/framework/action"
	"github.com/NpoolPlatform/kunman/framework/logger"
	// "github.com/NpoolPlatform/kunman/framework/wlog"
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

		return action.Run(
			c.Context,
			run,
			rpcRegister,
			rpcGatewayRegister,
			watch,
		)
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
	return nil
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	// api.Register(server)

	// apicli.RegisterGRPC(server)

	return nil
}

func rpcGatewayRegister(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	// err := api.RegisterGateway(mux, endpoint, opts)
	// if err != nil {
	// 	return wlog.WrapError(err)
	// }

	// _ = apicli.Register(mux)
	return nil
}
