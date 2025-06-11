package goodcoin

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	goodcoinmwpb "github.com/NpoolPlatform/kunman/message/good/middleware/v1/good/coin"
	goodcoinmw "github.com/NpoolPlatform/kunman/middleware/good/good/coin"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type checkHandler struct {
	*Handler
}

func (h *checkHandler) checkGoodCoin(ctx context.Context) error {
	conds := &goodcoinmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.EntID},
	}
	handler, err := goodcoinmw.NewHandler(
		ctx,
		goodcoinmw.WithConds(conds),
	)
	if err != nil {
		return err
	}

	exist, err := handler.ExistGoodCoinConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid goodcoin")
	}
	return nil
}
