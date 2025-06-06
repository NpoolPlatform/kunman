package main

import (
	"github.com/NpoolPlatform/kunman/framework/config"
	cli "github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "enable-agi",
		Aliases: []string{"a"},
		Usage:   "Enable agi module",
		Value:   true,
	},
}

func InitializeFlags(ctx *cli.Context) {
	config.SetKeyValue("enable.agi", ctx.Bool("enable-agi"))
}
