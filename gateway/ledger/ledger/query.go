package ledger

import (
	"context"
	"fmt"

	appusermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	npool "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger"
	ledgermwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	ledgermw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	ledgers      map[string]map[string]*ledgermwpb.Ledger
	appCoins     map[string]*appcoinmwpb.Coin
	appUsers     map[string]*appusermwpb.User
	infos        []*npool.Ledger
	totalLedgers uint32
	totalCoins   uint32
}

func (h *queryHandler) checkUser(ctx context.Context) error {
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

func (h *queryHandler) getAppCoins(ctx context.Context, conds *appcoinmwpb.Conds, offset, limit int32) error {
	handler, err := appcoinmw.NewHandler(
		ctx,
		appcoinmw.WithConds(conds),
		appcoinmw.WithOffset(offset),
		appcoinmw.WithLimit(limit),
	)
	if err != nil {
		return err
	}

	coins, total, err := handler.GetCoins(ctx)
	if err != nil {
		return err
	}
	for _, coin := range coins {
		h.appCoins[coin.CoinTypeID] = coin
	}
	h.totalCoins = total
	return nil
}

func (h *queryHandler) getLedgers(ctx context.Context, conds *ledgermwpb.Conds, offset, limit int32) error {
	handler, err := ledgermw.NewHandler(
		ctx,
		ledgermw.WithConds(conds),
		ledgermw.WithOffset(offset),
		ledgermw.WithLimit(limit),
	)
	if err != nil {
		return err
	}

	ledgers, total, err := handler.GetLedgers(ctx)
	if err != nil {
		return err
	}
	for _, ledger := range ledgers {
		ledgers, ok := h.ledgers[ledger.UserID]
		if !ok {
			ledgers = map[string]*ledgermwpb.Ledger{}
		}
		ledgers[ledger.CurrencyID] = ledger
		h.ledgers[ledger.UserID] = ledgers
	}
	h.totalLedgers = total
	return nil
}

func (h *queryHandler) getAppUsers(ctx context.Context) error {
	userIDs := []string{}
	if h.UserID != nil {
		userIDs = append(userIDs, *h.UserID)
	}
	for userID := range h.ledgers {
		userIDs = append(userIDs, userID)
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
		h.appUsers[user.EntID] = user
	}
	return nil
}

func (h *queryHandler) prepareAppLedgers(ctx context.Context) error {
	// Get offset/limit ledgers
	if err := h.getLedgers(ctx, &ledgermwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}, h.Offset, h.Limit); err != nil {
		return err
	}
	coinTypeIDs := []string{}
	for _, ledgers := range h.ledgers {
		for coinTypeID := range ledgers {
			coinTypeIDs = append(coinTypeIDs, coinTypeID)
		}
	}
	conds := &appcoinmwpb.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}
	// Get ledger coins
	if err := h.getAppCoins(ctx, conds, 0, int32(len(coinTypeIDs))); err != nil {
		return err
	}
	return nil
}

func (h *queryHandler) prepareUserLedgers(ctx context.Context) error {
	// Get offset/limit coins
	if err := h.getAppCoins(ctx, &appcoinmwpb.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	}, h.Offset, h.Limit); err != nil {
		return err
	}
	coinTypeIDs := []string{}
	for coinTypeID := range h.appCoins {
		coinTypeIDs = append(coinTypeIDs, coinTypeID)
	}
	// Get coin ledgers
	conds := &ledgermwpb.Conds{
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
		UserID:      &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID},
		CurrencyIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}
	if err := h.getLedgers(ctx, conds, 0, int32(len(coinTypeIDs))); err != nil {
		return err
	}
	return nil
}

func (h *queryHandler) formalize(ledger *ledgermwpb.Ledger, coin *appcoinmwpb.Coin, user *appusermwpb.User) {
	h.infos = append(h.infos, &npool.Ledger{
		CurrencyID:       coin.CoinTypeID,
		CurrencyName:     coin.Name,
		DisplayNames:     coin.DisplayNames,
		CurrencyLogo:     coin.Logo,
		CurrencyUnit:     coin.Unit,
		CurrencyDisabled: coin.Disabled,
		CurrencyDisplay:  coin.Display,
		Incoming:         ledger.Incoming,
		Locked:           ledger.Locked,
		Outcoming:        ledger.Outcoming,
		Spendable:        ledger.Spendable,
		PhoneNO:          user.PhoneNO,
		EmailAddress:     user.EmailAddress,
		UserID:           user.EntID,
		AppID:            user.AppID,
		ID:               ledger.ID,
		EntID:            ledger.EntID,
	})
}

func (h *queryHandler) formalizeAppLedgers() {
	for userID, ledgers := range h.ledgers {
		user, ok := h.appUsers[userID]
		if !ok {
			continue
		}
		for coinTypeID, ledger := range ledgers {
			coin, ok := h.appCoins[coinTypeID]
			if !ok {
				continue
			}
			h.formalize(ledger, coin, user)
		}
	}
}

func (h *queryHandler) formalizeUserLedgers() {
	user, ok := h.appUsers[*h.UserID]
	if !ok {
		return
	}

	ledgers, _ := h.ledgers[*h.UserID] //nolint
	for coinTypeID, coin := range h.appCoins {
		if ledgers != nil {
			ledger, ok := ledgers[coinTypeID]
			if ok {
				h.formalize(ledger, coin, user)
				continue
			}
		}
		h.infos = append(h.infos, &npool.Ledger{
			CurrencyID:       coin.CoinTypeID,
			CurrencyName:     coin.Name,
			DisplayNames:     coin.DisplayNames,
			CurrencyLogo:     coin.Logo,
			CurrencyUnit:     coin.Unit,
			CurrencyDisabled: coin.Disabled,
			CurrencyDisplay:  coin.Display,
			Incoming:         decimal.NewFromInt(0).String(),
			Locked:           decimal.NewFromInt(0).String(),
			Outcoming:        decimal.NewFromInt(0).String(),
			Spendable:        decimal.NewFromInt(0).String(),
			PhoneNO:          user.PhoneNO,
			EmailAddress:     user.EmailAddress,
			UserID:           user.EntID,
			AppID:            user.AppID,
		})
	}
}

func (h *Handler) GetLedgers(ctx context.Context) ([]*npool.Ledger, uint32, error) {
	handler := &queryHandler{
		Handler:  h,
		ledgers:  map[string]map[string]*ledgermwpb.Ledger{},
		appCoins: map[string]*appcoinmwpb.Coin{},
		appUsers: map[string]*appusermwpb.User{},
	}
	if err := handler.checkUser(ctx); err != nil {
		return nil, 0, err
	}
	if h.UserID == nil {
		if err := handler.prepareAppLedgers(ctx); err != nil {
			return nil, 0, err
		}
	} else {
		if err := handler.prepareUserLedgers(ctx); err != nil {
			return nil, 0, err
		}
	}
	if err := handler.getAppUsers(ctx); err != nil {
		return nil, 0, err
	}

	if h.UserID == nil {
		handler.formalizeAppLedgers()
		return handler.infos, handler.totalLedgers, nil
	}

	handler.formalizeUserLedgers()
	return handler.infos, handler.totalCoins, nil
}
