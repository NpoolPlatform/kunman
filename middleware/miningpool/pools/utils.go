package pools

import (
	"fmt"
	"strings"

	"github.com/NpoolPlatform/kunman/middleware/miningpool/pools/f2pool"
)

const (
	TestNamePrefix = "testset"
)

func IsTestPoolUserName(name string) bool {
	return strings.HasPrefix(name, TestNamePrefix)
}

func RandomPoolUserNameForTest() string {
	name := TestNamePrefix
	_name, err := f2pool.RandomF2PoolUser(6)
	if err != nil {
		panic(err)
	}
	name += _name
	fmt.Println("error ssdfasfasdf", name)
	return name
}
