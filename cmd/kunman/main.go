package main

import (
	"fmt"
	"os"

	servicename "github.com/NpoolPlatform/kunman/pkg/servicename"

	app1 "github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/kunman/framework/app"
	"github.com/NpoolPlatform/kunman/framework/logger"

	mysqlconst "github.com/NpoolPlatform/kunman/framework/mysql/const"
	rabbitmqconst "github.com/NpoolPlatform/kunman/framework/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/kunman/framework/redis/const"
	accountmwsvcname "github.com/NpoolPlatform/kunman/middleware/account/servicename"
	appusermwsvcname "github.com/NpoolPlatform/kunman/middleware/appuser/servicename"
	basalservicename "github.com/NpoolPlatform/kunman/middleware/basal/servicename"
	chainmwsvcname "github.com/NpoolPlatform/kunman/middleware/chain/servicename"
	g11nmwsvcname "github.com/NpoolPlatform/kunman/middleware/g11n/servicename"
	goodmwsvcname "github.com/NpoolPlatform/kunman/middleware/good/servicename"
	inspiremwsvcname "github.com/NpoolPlatform/kunman/middleware/inspire/servicename"
	ledgermwsvcname "github.com/NpoolPlatform/kunman/middleware/ledger/servicename"
	miningpoolmwsvcname "github.com/NpoolPlatform/kunman/middleware/miningpool/servicename"
	notifmwsvcname "github.com/NpoolPlatform/kunman/middleware/notif/servicename"
	ordermwsvcname "github.com/NpoolPlatform/kunman/middleware/order/servicename"
	reviewmwsvcname "github.com/NpoolPlatform/kunman/middleware/review/servicename"
	sphinxproxysvcname "github.com/NpoolPlatform/sphinx-proxy/pkg/servicename"

	cli "github.com/urfave/cli/v2"
)

func main() {
	commands := cli.Commands{
		runCmd,
	}

	description := fmt.Sprintf("%v service cli\nFor help on any individual command run <%v COMMAND -h>\n",
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
		basalservicename.ServiceDomain,
		goodmwsvcname.ServiceDomain,
		ledgermwsvcname.ServiceDomain,
		inspiremwsvcname.ServiceDomain,
		ordermwsvcname.ServiceDomain,
		g11nmwsvcname.ServiceDomain,
		reviewmwsvcname.ServiceDomain,
		miningpoolmwsvcname.ServiceDomain,
		notifmwsvcname.ServiceDomain,
		accountmwsvcname.ServiceDomain,
		appusermwsvcname.ServiceDomain,
		chainmwsvcname.ServiceDomain,
		sphinxproxysvcname.ServiceDomain,
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Failed initialize",
			"ServiceName", servicename.ServiceName,
			"Error", err,
		)
		return
	}

	// Workaround for sphinx
	err = app1.Init(
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
		basalservicename.ServiceDomain,
		goodmwsvcname.ServiceDomain,
		ledgermwsvcname.ServiceDomain,
		inspiremwsvcname.ServiceDomain,
		ordermwsvcname.ServiceDomain,
		accountmwsvcname.ServiceDomain,
		appusermwsvcname.ServiceDomain,
		chainmwsvcname.ServiceDomain,
		sphinxproxysvcname.ServiceDomain,
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
