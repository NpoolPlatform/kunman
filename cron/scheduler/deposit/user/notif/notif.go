package notif

import (
	"context"
	"encoding/json"
	"fmt"

	basenotif "github.com/NpoolPlatform/kunman/cron/scheduler/base/notif"
	retry1 "github.com/NpoolPlatform/kunman/cron/scheduler/base/retry"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/deposit/user/types"
	"github.com/NpoolPlatform/kunman/framework/logger"
	"github.com/NpoolPlatform/kunman/framework/pubsub"
	ledgertypes "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	statementmwpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"

	"github.com/google/uuid"
)

type handler struct{}

func NewNotif() basenotif.Notify {
	return &handler{}
}

func (p *handler) statement(account *types.PersistentAccount) *statementmwpb.StatementReq {
	id := uuid.NewString()
	ioType := ledgertypes.IOType_Incoming
	ioSubType := ledgertypes.IOSubType_Deposit

	return &statementmwpb.StatementReq{
		EntID:      &id,
		AppID:      &account.AppID,
		UserID:     &account.UserID,
		CurrencyID: &account.CoinTypeID,
		IOType:     &ioType,
		IOSubType:  &ioSubType,
		Amount:     &account.DepositAmount,
		IOExtra:    &account.Extra,
	}
}

func (p *handler) notifyDeposit(account *types.PersistentAccount) error {
	return pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		msgID := basetypes.MsgID_DepositReceivedReq.String()
		if account.Error != nil {
			msgID = basetypes.MsgID_DepositCheckFailReq.String()
		}
		if account.Error == nil {
			return publisher.Update(msgID, nil, nil, nil, p.statement(account))
		}
		req := &basetypes.MsgError{
			Error: account.Error.Error(),
		}
		value, _ := json.Marshal(p.statement(account))
		req.Value = string(value)
		return publisher.Update(msgID, nil, nil, nil, req)
	})
}

func (p *handler) Notify(ctx context.Context, account interface{}, retry chan interface{}) error {
	_account, ok := account.(*types.PersistentAccount)
	if !ok {
		return fmt.Errorf("invalid account")
	}
	if err := p.notifyDeposit(_account); err != nil {
		logger.Sugar().Errorw(
			"notifDeposit",
			"AppID", _account.AppID,
			"UserID", _account.UserID,
			"Account", _account.CoinTypeID,
			"AccountType", _account.DepositAmount,
			"Error", err,
		)
		retry1.Retry(_account.EntID, _account, retry)
		return err
	}
	return nil
}
