package testinit

import (
	accountmwsvcname "github.com/NpoolPlatform/kunman/middleware/account/servicename"
	appusermwsvcname "github.com/NpoolPlatform/kunman/middleware/appuser/servicename"
	chainmwsvcname "github.com/NpoolPlatform/kunman/middleware/chain/servicename"
	goodmwsvcname "github.com/NpoolPlatform/kunman/middleware/good/servicename"
	inspiremwsvcname "github.com/NpoolPlatform/kunman/middleware/inspire/servicename"
	ledgermwsvcname "github.com/NpoolPlatform/kunman/middleware/ledger/servicename"
	ordermwsvcname "github.com/NpoolPlatform/kunman/middleware/order/servicename"
	testinit "github.com/NpoolPlatform/kunman/pkg/testinit"
	sphinxproxysvcname "github.com/NpoolPlatform/sphinx-proxy/pkg/servicename"
)

func Init() error {
	return testinit.Initialize(
		goodmwsvcname.ServiceDomain,
		ledgermwsvcname.ServiceDomain,
		inspiremwsvcname.ServiceDomain,
		ordermwsvcname.ServiceDomain,
		accountmwsvcname.ServiceDomain,
		appusermwsvcname.ServiceDomain,
		chainmwsvcname.ServiceDomain,
		sphinxproxysvcname.ServiceDomain,
	)
}
