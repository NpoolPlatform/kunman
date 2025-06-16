package constant

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/NpoolPlatform/kunman/mal/sphinx/plugin/config"
	"google.golang.org/grpc/metadata"
)

const (
	ServiceName       = "sphinx-plugin.npool.top"
	GrpcTimeout       = time.Second * 10
	WaitMsgOutTimeout = time.Second * 40
)

func SetPluginInfo(ctx context.Context) context.Context {
	md := metadata.New(
		map[string]string{
			"_pluginwanip":    config.GetENV().WanIP,
			"_pluginposition": config.GetENV().Position,
		})
	return metadata.NewOutgoingContext(ctx, md)
}

func GetPluginInfo(ctx context.Context) string {
	pluginInfo := ""
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if v, ok := md["_pluginposition"]; ok {
			pluginInfo = strings.Join(v, "_")
		}
		if v, ok := md["_pluginwanip"]; ok {
			pluginInfo = fmt.Sprintf("%v-%v", pluginInfo, strings.Join(v, "_"))
		}
		return pluginInfo
	}
	return "pluginInfo-not-set"
}
