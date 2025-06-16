package testinit

import (
	"fmt"
	"os"
	"path"
	"runtime"

	app1 "github.com/NpoolPlatform/go-service-framework/pkg/app"
	"github.com/NpoolPlatform/kunman/framework/app"

	servicename "github.com/NpoolPlatform/kunman/pkg/servicename"

	mysqlconst "github.com/NpoolPlatform/kunman/framework/mysql/const"
	rabbitmqconst "github.com/NpoolPlatform/kunman/framework/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/kunman/framework/redis/const"
)

func init() {
	os.Setenv("RUN_IN_UNIT_TEST", "true")
}

func Initialize(serviceNames ...string) error {
	_, myPath, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("cannot get source file path")
	}

	configPath := fmt.Sprintf("%s/../../cmd/kunman", path.Dir(myPath))
	serviceNames = append(
		serviceNames,
		mysqlconst.MysqlServiceName,
		rabbitmqconst.RabbitMQServiceName,
		redisconst.RedisServiceName,
	)

	err := app.Initialize(
		servicename.ServiceName,
		"",
		"",
		"",
		configPath,
		nil,
		nil,
		serviceNames...,
	)
	if err != nil {
		return fmt.Errorf("cannot init app stub: %v", err)
	}

	err = app1.Init(
		servicename.ServiceName,
		"",
		"",
		"",
		configPath,
		nil,
		nil,
		serviceNames...,
	)
	if err != nil {
		return fmt.Errorf("cannot init app stub: %v", err)
	}

	return nil
}
