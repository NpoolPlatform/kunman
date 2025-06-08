package goodcoin

import (
	"context"
	"fmt"

	goodcoinmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
