package testinit

import (
	myservicename "github.com/NpoolPlatform/kunman/gateway/order/servicename"
	appusermwsvcname "github.com/NpoolPlatform/kunman/middleware/appuser/servicename"
	goodmwsvcname "github.com/NpoolPlatform/kunman/middleware/good/servicename"
	ordermwsvcname "github.com/NpoolPlatform/kunman/middleware/order/servicename"
	testinit "github.com/NpoolPlatform/kunman/pkg/testinit"
)

func Init() error {
	return testinit.Initialize(
		myservicename.ServiceDomain,
		goodmwsvcname.ServiceDomain,
		ordermwsvcname.ServiceDomain,
		appusermwsvcname.ServiceDomain,
	)
}
