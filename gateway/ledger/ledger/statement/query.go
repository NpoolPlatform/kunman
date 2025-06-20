package statement

import (
	"context"
	"fmt"

	appusermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	npool "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger/statement"
	statementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	statementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type queryHandler struct {
	*Handler
	statements []*statementmwpb.Statement
	appcoin    map[string]*appcoinmwpb.Coin
	appuser    map[string]*appusermwpb.User
	infos      []*npool.Statement
}

func (h *Handler) checkUser(ctx context.Context) error {
	if h.UserID == nil {
		return nil
	}

	handler, err := usermw.NewHandler(
		ctx,
		usermw.WithAppID(h.AppID, true),
		usermw.WithEntID(h.UserID, true),
	)
	if err != nil {
		return err
	}

	user, err := handler.GetUser(ctx)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("invalid user")
	}
	return nil
}

func (h *queryHandler) getAppCoins(ctx context.Context) error {
	coinTypeIDs := []string{}
	for _, val := range h.statements {
		coinTypeIDs = append(coinTypeIDs, val.CurrencyID)
	}

	conds := &appcoinmwpb.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}
	handler, err := appcoinmw.NewHandler(
		ctx,
		appcoinmw.WithConds(conds),
		appcoinmw.WithOffset(0),
		appcoinmw.WithLimit(int32(len(coinTypeIDs))),
	)
	if err != nil {
		return err
	}

	coins, _, err := handler.GetCoins(ctx)
	if err != nil {
		return err
	}

	for _, coin := range coins {
		h.appcoin[coin.CoinTypeID] = coin
	}
	return nil
}

func (h *queryHandler) getAppUsers(ctx context.Context) error {
	userIDs := []string{}
	for _, info := range h.statements {
		userIDs = append(userIDs, info.UserID)
	}

	conds := &appusermwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: userIDs},
	}
	handler, err := usermw.NewHandler(
		ctx,
		usermw.WithConds(conds),
		usermw.WithOffset(0),
		usermw.WithLimit(int32(len(userIDs))),
	)
	if err != nil {
		return err
	}

	users, _, err := handler.GetUsers(ctx)
	if err != nil {
		return err
	}

	for _, user := range users {
		h.appuser[user.EntID] = user
	}
	return nil
}

func (h *queryHandler) formalize() {
	for _, statement := range h.statements {
		coin, ok := h.appcoin[statement.CurrencyID]
		if !ok {
			continue
		}
		user, ok := h.appuser[statement.UserID]
		if !ok {
			continue
		}

		h.infos = append(h.infos, &npool.Statement{
			ID:           statement.ID,
			EntID:        statement.EntID,
			AppID:        statement.AppID,
			CurrencyID:   coin.CoinTypeID,
			CurrencyName: coin.CoinName,
			DisplayNames: coin.DisplayNames,
			CurrencyLogo: coin.Logo,
			CurrencyUnit: coin.Unit,
			IOType:       statement.IOType,
			IOSubType:    statement.IOSubType,
			IOExtra:      statement.IOExtra,
			Amount:       statement.Amount,
			CreatedAt:    statement.CreatedAt,
			UserID:       user.EntID,
			PhoneNO:      user.PhoneNO,
			EmailAddress: user.EmailAddress,
		})
	}
}

func (h *Handler) GetStatements(ctx context.Context) ([]*npool.Statement, uint32, error) {
	if err := h.checkUser(ctx); err != nil {
		return nil, 0, err
	}
	if err := h.CheckStartEndAt(); err != nil {
		return nil, 0, err
	}
	conds := &statementmwpb.Conds{}
	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}
	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	if h.StartAt != nil {
		conds.StartAt = &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.StartAt}
	}
	if h.EndAt != nil {
		conds.EndAt = &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.EndAt}
	}

	statementHandler, err := statementmw.NewHandler(
		ctx,
		statementmw.WithConds(conds),
		statementmw.WithOffset(h.Offset),
		statementmw.WithLimit(h.Limit),
	)
	if err != nil {
		return nil, 0, err
	}

	statements, total, err := statementHandler.GetStatements(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(statements) == 0 {
		return nil, total, nil
	}

	handler := &queryHandler{
		Handler:    h,
		statements: statements,
		appcoin:    map[string]*appcoinmwpb.Coin{},
		appuser:    map[string]*appusermwpb.User{},
	}

	if err := handler.getAppCoins(ctx); err != nil {
		return nil, 0, err
	}
	if err := handler.getAppUsers(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()
	return handler.infos, total, nil
}
