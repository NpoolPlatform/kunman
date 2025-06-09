package lock

import (
	"fmt"

	redis2 "github.com/NpoolPlatform/kunman/framework/redis"
)

const PrefixAccountLock = "AccountLock"

func key(id string) string {
	return fmt.Sprintf("%v:%v", PrefixAccountLock, id)
}

func Lock(id string) error {
	return redis2.TryLock(key(id), 0)
}

func Unlock(id string) error {
	return redis2.Unlock(key(id))
}
