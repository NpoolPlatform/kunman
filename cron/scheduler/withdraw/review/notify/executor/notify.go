package executor

import (
	"context"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/withdraw/review/notify/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	accountmwpb "github.com/NpoolPlatform/kunman/message/account/middleware/v1/account"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	coinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/coin"
	ledgerwithdrawmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/withdraw"
	withdrawreviewnotifypb "github.com/NpoolPlatform/kunman/message/scheduler/middleware/v1/withdraw/review/notify"
	accountmw "github.com/NpoolPlatform/kunman/middleware/account/account"
	appmw "github.com/NpoolPlatform/kunman/middleware/appuser/app"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	coinmw "github.com/NpoolPlatform/kunman/middleware/chain/coin"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type withdrawReviewNotifyHandler struct {
	withdraws        []*ledgerwithdrawmwpb.Withdraw
	accounts         map[string]*accountmwpb.Account
	apps             map[string]*appmwpb.App
	users            map[string]*usermwpb.User
	coins            map[string]*coinmwpb.Coin
	appWithdrawInfos []*withdrawreviewnotifypb.AppWithdrawInfos
	persistent       chan interface{}
	notif            chan interface{}
	done             chan interface{}
}

func (h *withdrawReviewNotifyHandler) getAccounts(ctx context.Context) error {
	accountIDs := []string{}
	for _, withdraw := range h.withdraws {
		accountIDs = append(accountIDs, withdraw.AccountID)
	}

	conds := &accountmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: accountIDs},
	}
	handler, err := accountmw.NewHandler(
		ctx,
		accountmw.WithConds(conds),
		accountmw.WithOffset(0),
		accountmw.WithLimit(int32(len(accountIDs))),
	)
	if err != nil {
		return err
	}

	accounts, _, err := handler.GetAccounts(ctx)
	if err != nil {
		return err
	}

	h.accounts = map[string]*accountmwpb.Account{}
	for _, account := range accounts {
		h.accounts[account.EntID] = account
	}

	return nil
}

func (h *withdrawReviewNotifyHandler) getApps(ctx context.Context) error {
	appIDs := []string{}
	for _, withdraw := range h.withdraws {
		appIDs = append(appIDs, withdraw.AppID)
	}

	conds := &appmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: appIDs},
	}
	handler, err := appmw.NewHandler(
		ctx,
		appmw.WithConds(conds),
		appmw.WithOffset(0),
		appmw.WithLimit(int32(len(appIDs))),
	)
	if err != nil {
		return err
	}

	apps, _, err := handler.GetApps(ctx)
	if err != nil {
		return err
	}

	h.apps = map[string]*appmwpb.App{}
	for _, app := range apps {
		h.apps[app.EntID] = app
	}

	return nil
}

func (h *withdrawReviewNotifyHandler) getAppUsers(ctx context.Context) error {
	userIDs := []string{}
	for _, withdraw := range h.withdraws {
		userIDs = append(userIDs, withdraw.UserID)
	}

	conds := &usermwpb.Conds{
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

	h.users = map[string]*usermwpb.User{}
	for _, user := range users {
		h.users[user.EntID] = user
	}

	return nil
}

func (h *withdrawReviewNotifyHandler) getCoins(ctx context.Context) error {
	coinTypeIDs := []string{}
	for _, withdraw := range h.withdraws {
		coinTypeIDs = append(coinTypeIDs, withdraw.CoinTypeID)
	}

	conds := &coinmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: coinTypeIDs},
	}
	handler, err := coinmw.NewHandler(
		ctx,
		coinmw.WithConds(conds),
		coinmw.WithOffset(0),
		coinmw.WithLimit(int32(len(coinTypeIDs))),
	)
	if err != nil {
		return err
	}

	coins, _, err := handler.GetCoins(ctx)
	if err != nil {
		return err
	}

	h.coins = map[string]*coinmwpb.Coin{}
	for _, coin := range coins {
		h.coins[coin.EntID] = coin
	}

	return nil
}

func (h *withdrawReviewNotifyHandler) resolveWithdrawInfos() {
	appWithdrawInfos := map[string]*withdrawreviewnotifypb.AppWithdrawInfos{}

	for _, withdraw := range h.withdraws {
		account, ok := h.accounts[withdraw.AccountID]
		if !ok {
			continue
		}
		user, ok := h.users[withdraw.UserID]
		if !ok {
			continue
		}
		coin, ok := h.coins[withdraw.CoinTypeID]
		if !ok {
			continue
		}
		app, ok := h.apps[withdraw.AppID]
		if !ok {
			continue
		}

		withdrawInfos, ok := appWithdrawInfos[withdraw.AppID]
		if !ok {
			withdrawInfos = &withdrawreviewnotifypb.AppWithdrawInfos{
				AppID:   withdraw.AppID,
				AppName: app.Name,
			}
		}

		withdrawInfos.Withdraws = append(withdrawInfos.Withdraws, &withdrawreviewnotifypb.WithdrawInfo{
			Withdraw: withdraw,
			Account:  account,
			User:     user,
			Coin:     coin,
		})

		appWithdrawInfos[withdraw.AppID] = withdrawInfos
	}

	for _, _appWithdrawInfos := range appWithdrawInfos {
		h.appWithdrawInfos = append(h.appWithdrawInfos, _appWithdrawInfos)
	}
}

//nolint:gocritic
func (h *withdrawReviewNotifyHandler) final(ctx context.Context, err *error) {
	if *err != nil {
		logger.Sugar().Errorw(
			"final",
			"Withdraws", h.withdraws,
			"AppWithdraws", h.appWithdrawInfos,
			"Error", *err,
		)
	}
	persistentWithdrawReviewNotify := &types.PersistentWithdrawReviewNotify{
		AppWithdraws: h.appWithdrawInfos,
	}
	if *err == nil {
		asyncfeed.AsyncFeed(ctx, persistentWithdrawReviewNotify, h.notif)
		asyncfeed.AsyncFeed(ctx, persistentWithdrawReviewNotify, h.persistent)
		return
	}
	asyncfeed.AsyncFeed(ctx, persistentWithdrawReviewNotify, h.done)
}

//nolint:gocritic
func (h *withdrawReviewNotifyHandler) exec(ctx context.Context) error {
	var err error
	defer h.final(ctx, &err)

	if err = h.getAccounts(ctx); err != nil {
		return err
	}
	if err = h.getApps(ctx); err != nil {
		return err
	}
	if err = h.getAppUsers(ctx); err != nil {
		return err
	}
	if err = h.getCoins(ctx); err != nil {
		return err
	}
	h.resolveWithdrawInfos()

	return nil
}
