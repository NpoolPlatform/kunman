package exchange

import (
	"context"
	"fmt"

	exchangemwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/credit/exchange"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	exchangemwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/credit/exchange"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkExchange(ctx context.Context) error {
	exist, err := exchangemwcli.ExistExchangeConds(ctx, &exchangemwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid exchange")
	}
	return nil
}
