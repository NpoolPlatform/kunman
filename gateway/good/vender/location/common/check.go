package common

import (
	"context"
	"fmt"

	vendorlocationmwcli "github.com/NpoolPlatform/kunman/middleware/good/vender/location"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	vendorlocationmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/vender/location"
)

type LocationCheckHandler struct {
	LocationID *string
}

func (h *LocationCheckHandler) CheckLocationWithLocationID(ctx context.Context, vendorLocationID string) error {
	exist, err := vendorlocationmwcli.ExistLocationConds(ctx, &vendorlocationmwpb.Conds{
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: vendorLocationID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid vendorlocation")
	}
	return nil
}

func (h *LocationCheckHandler) CheckLocation(ctx context.Context) error {
	return h.CheckLocationWithLocationID(ctx, *h.LocationID)
}
