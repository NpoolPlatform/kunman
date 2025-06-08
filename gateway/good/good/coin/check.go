package goodcoin

import (
	"context"
	"fmt"

	goodcoinmwcli "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkGoodCoin(ctx context.Context) error {
	exist, err := goodcoinmwcli.ExistGoodCoinConds(ctx, &goodcoinmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid goodcoin")
	}
	return nil
}
