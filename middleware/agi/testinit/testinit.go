package testinit

import (
	"fmt"
	"path"
	"runtime"

	"github.com/NpoolPlatform/kunman/framework/app"

	servicename "github.com/NpoolPlatform/kunman/pkg/servicename"

	mysqlconst "github.com/NpoolPlatform/kunman/framework/mysql/const"
	rabbitmqconst "github.com/NpoolPlatform/kunman/framework/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/kunman/framework/redis/const"
)

func Init() error {
	_, myPath, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("cannot get source file path")
	}

	configPath := fmt.Sprintf("%s/../../../cmd/kunman", path.Dir(myPath))

	err := app.Initialize(
		servicename.ServiceName,
		"",
		"",
		"",
		configPath,
		nil,
		nil,
		mysqlconst.MysqlServiceName,
		rabbitmqconst.RabbitMQServiceName,
		redisconst.RedisServiceName,
	)
	if err != nil {
		return fmt.Errorf("cannot init app stub: %v", err)
	}

	return nil
}
