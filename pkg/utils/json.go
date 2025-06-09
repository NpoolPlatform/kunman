package utils

import (
	"github.com/bytedance/sonic"
)

func PrettyStruct(data interface{}) string {
	val, err := sonic.ConfigFastest.MarshalIndent(data, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(val)
}
