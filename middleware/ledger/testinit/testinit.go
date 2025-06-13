package testinit

import (
	myservicename "github.com/NpoolPlatform/kunman/middleware/ledger/servicename"
	testinit "github.com/NpoolPlatform/kunman/pkg/testinit"
)

func Init() error {
	return testinit.Initialize(myservicename.ServiceDomain)
}
