package testinit

import (
	"fmt"
	"path"
	"runtime"

	"github.com/NpoolPlatform/kunman/framework/app"

	"github.com/NpoolPlatform/kunman/middleware/agi/db"

	servicename "github.com/NpoolPlatform/kunman/middleware/agi/servicename"

	mysqlconst "github.com/NpoolPlatform/kunman/framework/mysql/const"
	rabbitmqconst "github.com/NpoolPlatform/kunman/framework/rabbitmq/const"
	redisconst "github.com/NpoolPlatform/kunman/framework/redis/const"
)

func Init() error {
	_, myPath, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("cannot get source file path")
	}

	appName := path.Base(path.Dir(path.Dir(path.Dir(myPath))))
	configPath := fmt.Sprintf("%s/../../cmd/%v", path.Dir(myPath), appName)

	err := app.Init(
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
	err = db.Init()
	if err != nil {
		return fmt.Errorf("cannot init database: %v", err)
	}

	return nil
}
