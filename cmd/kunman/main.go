package main

import (
	"fmt"
	"os"

	servicename "github.com/NpoolPlatform/kunman/pkg/servicename"

	"github.com/NpoolPlatform/kunman/framework/app"
	"github.com/NpoolPlatform/kunman/framework/logger"

	mysqlconst "github.com/NpoolPlatform/kunman/framework/mysql/const"
	rabbitmqconst "github.com/NpoolPlatform/kunman/framework/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/kunman/framework/redis/const"

	cli "github.com/urfave/cli/v2"
)

func main() {
	commands := cli.Commands{
		runCmd,
	}

	description := fmt.Sprintf("my %v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
		servicename.ServiceName, servicename.ServiceName)
	err := app.Initialize(
		servicename.ServiceName,
		description,
		"",
		"",
		"./",
		nil,
		commands,
		mysqlconst.MysqlServiceName,
		rabbitmqconst.RabbitMQServiceName,
		redisconst.RedisServiceName,
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Failed initialize",
			"ServiceName", servicename.ServiceName,
			"Error", err,
		)
		return
	}
	err = app.Run(os.Args)
	if err != nil {
		logger.Sugar().Errorw(
			"Failed run",
			"ServiceName", servicename.ServiceName,
			"Error", err,
		)
	}
}
