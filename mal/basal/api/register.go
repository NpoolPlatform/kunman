package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unsafe"

	config "github.com/NpoolPlatform/kunman/framework/config"
	logger "github.com/NpoolPlatform/kunman/framework/logger"
	npool "github.com/NpoolPlatform/kunman/message/basal/middleware/v1/api"
	basalapi "github.com/NpoolPlatform/kunman/middleware/basal/api"

	"github.com/go-resty/resty/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func createAPIs(apis []*npool.APIReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	handler, err := basalapi.NewHandler(
		ctx,
		basalapi.WithReqs(apis),
	)
	if err != nil {
		return err
	}
	_, err = handler.CreateAPIs(ctx)
	return err
}

func muxAPIs(mux *runtime.ServeMux) []*npool.APIReq {
	var apis []*npool.APIReq
	serviceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	protocol := npool.Protocol_HTTP

	valueOfMux := reflect.ValueOf(mux).Elem()
	handlers := valueOfMux.FieldByName("handlers")
	methIter := handlers.MapRange()

	for methIter.Next() {
		for i := 0; i < methIter.Value().Len(); i++ {
			pat := methIter.Value().Index(i).FieldByName("pat")
			tmp := reflect.NewAt(pat.Type(), unsafe.Pointer(pat.UnsafeAddr())).Elem()
			str := tmp.MethodByName("String").Call(nil)[0].String()
			method, ok := npool.Method_value[methIter.Key().String()]
			if !ok {
				logger.Sugar().Warnw("muxAPIs", "Method", methIter.Key().String())
				continue
			}
			_method := npool.Method(method)

			apis = append(apis, &npool.APIReq{
				Protocol:    &protocol,
				ServiceName: &serviceName,
				Method:      &_method,
				Path:        &str,
			})
		}
	}

	return apis
}

func grpcAPIs(server grpc.ServiceRegistrar) []*npool.APIReq {
	srvInfo := server.(*grpc.Server).GetServiceInfo()

	var apis []*npool.APIReq
	serviceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	protocol := npool.Protocol_GRPC
	method := npool.Method_STREAM

	for key, info := range srvInfo {
		for _, _method := range info.Methods {
			path := fmt.Sprintf("%v/%v", key, _method.Name)
			methodName := _method.Name
			apis = append(apis, &npool.APIReq{
				Protocol:    &protocol,
				ServiceName: &serviceName,
				Method:      &method,
				MethodName:  &methodName,
				Path:        &path,
			})
		}
	}

	return apis
}

func getGatewayRouters(name string) ([]*EntryPoint, error) {
	const leastDomainLen = 2
	domain := strings.SplitN(name, ".", leastDomainLen)
	if len(domain) < leastDomainLen {
		return nil, errors.New("service name must like example.npool.top")
	}

	allRouters := []*EntryPoint{}

	page := 1
	perPage := 1000

	for {
		// provider can use kubernetes or k8s
		url := fmt.Sprintf(
			"http://traefik.kube-system.svc.cluster.local:38080/api/http/routers?provider=kubernetes&page=%v&per_page=%v&search=%v",
			page, perPage, domain[0],
		)
		// internal already set timeout
		resp, err := resty.New().R().Get(url)
		if err != nil {
			break
		}

		routers := make([]*EntryPoint, 0)
		err = json.Unmarshal(resp.Body(), &routers)
		if err != nil {
			break
		}
		allRouters = append(allRouters, routers...)
		page += 1
	}

	return allRouters, nil
}

func Register(mux *runtime.ServeMux) error {
	apis := muxAPIs(mux)
	return registerHttp(apis)
}

func registerHttp(apis []*npool.APIReq) error {
	serviceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	gatewayRouters, err := getGatewayRouters(serviceName)
	if err != nil {
		return err
	}

	if len(gatewayRouters) == 0 {
		return fmt.Errorf("invalid routers")
	}

	for _, router := range gatewayRouters {
		prefix, err := router.PathPrefix()
		if err != nil {
			return err
		}
		routerPath, err := router.Path()
		if err != nil {
			return err
		}

		exported := true
		for _, _api := range apis {
			if !strings.HasPrefix(*_api.Path, "/v1") {
				continue
			}
			if !strings.HasPrefix(*_api.Path, routerPath) {
				continue
			}
			_api.PathPrefix = &prefix
			_api.Exported = &exported
			_api.Domains = append(_api.Domains, router.Domain())
		}
	}

	return createAPIs(apis)
}

func RegisterGRPC(server grpc.ServiceRegistrar) error {
	apis := grpcAPIs(server)
	return createAPIs(apis)
}
