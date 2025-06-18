package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/deposit/user/types"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	depositaccmw "github.com/NpoolPlatform/kunman/middleware/account/deposit"
	ledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"

	"github.com/google/uuid"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) withUpdateAccount(ctx context.Context, account *types.PersistentAccount) error {
	handler, err := depositaccmw.NewHandler(
		ctx,
		depositaccmw.WithID(&account.ID, true),
		depositaccmw.WithIncoming(&account.DepositAmount, true),
	)
	if err != nil {
		return err
	}

	_, err = handler.AddBalance(ctx)
	return err
}

func (p *handler) withCreateStatement(ctx context.Context, account *types.PersistentAccount) error {
	id := uuid.NewString()
	ioType := ledgertypes.IOType_Incoming
	ioSubType := ledgertypes.IOSubType_Deposit

	handler, err := ledgerstatementmw.NewHandler(
		ctx,
		ledgerstatementmw.WithEntID(&id, true),
		ledgerstatementmw.WithAppID(&account.AppID, true),
		ledgerstatementmw.WithUserID(&account.UserID, true),
		ledgerstatementmw.WithCoinTypeID(&account.CoinTypeID, true),
		ledgerstatementmw.WithIOType(&ioType, true),
		ledgerstatementmw.WithIOSubType(&ioSubType, true),
		ledgerstatementmw.WithAmount(&account.DepositAmount, true),
		ledgerstatementmw.WithIOExtra(&account.Extra, true),
	)
	if err != nil {
		return err
	}

	_, err = handler.CreateStatement(ctx)
	return err
}

func (p *handler) Update(ctx context.Context, account interface{}, reward, notif, done chan interface{}) error {
	_account, ok := account.(*types.PersistentAccount)
	if !ok {
		return fmt.Errorf("invalid account")
	}

	defer asyncfeed.AsyncFeed(ctx, _account, reward)

	if err := p.withUpdateAccount(ctx, _account); err != nil {
		return err
	}
	if err := p.withCreateStatement(ctx, _account); err != nil {
		return err
	}

	asyncfeed.AsyncFeed(ctx, _account, notif)

	return nil
}
