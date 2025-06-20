package deposit

import (
	"context"
	"fmt"
	"time"

	appusermwpb "github.com/NpoolPlatform/kunman/message/appuser/middleware/v1/user"
	types "github.com/NpoolPlatform/kunman/message/basetypes/ledger/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	appcoinmwpb "github.com/NpoolPlatform/kunman/message/chain/middleware/v1/app/coin"
	npool "github.com/NpoolPlatform/kunman/message/ledger/gateway/v1/ledger/statement"
	statementpb "github.com/NpoolPlatform/kunman/message/ledger/middleware/v2/ledger/statement"
	usermw "github.com/NpoolPlatform/kunman/middleware/appuser/user"
	appcoinmw "github.com/NpoolPlatform/kunman/middleware/chain/app/coin"
	ledgerstatementmw "github.com/NpoolPlatform/kunman/middleware/ledger/ledger/statement"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
)

type createHandler struct {
	*Handler
	user *appusermwpb.User
}

func (h *createHandler) checkUser(ctx context.Context) error {
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
	h.user = user
	return nil
}

func (h *Handler) CreateDeposit(ctx context.Context) (*npool.Statement, error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.checkUser(ctx); err != nil {
		return nil, err
	}

	appCoinConds := &appcoinmwpb.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: *h.TargetAppID},
		CoinTypeID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.CoinTypeID},
	}
	appCoinHandler, err := appcoinmw.NewHandler(
		ctx,
		appcoinmw.WithConds(appCoinConds),
	)
	if err != nil {
		return nil, err
	}

	coin, err := appCoinHandler.GetCoinOnly(ctx)
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, fmt.Errorf("invalid coin")
	}

	ioExtra := fmt.Sprintf(
		`{"AppID":"%v","UserID":"%v","TargetAppID":"%v","TargetUserID":"%v","CoinName":"%v","Amount":"%v","Date":"%v"}`,
		*h.AppID,
		*h.UserID,
		*h.TargetAppID,
		*h.TargetUserID,
		coin.Name,
		*h.Amount,
		time.Now(),
	)

	ioType := types.IOType_Incoming
	ioSubtype := types.IOSubType_Deposit

	ledgerStatementHandler, err := ledgerstatementmw.NewHandler(
		ctx,
		ledgerstatementmw.WithAppID(h.TargetAppID, true),
		ledgerstatementmw.WithUserID(h.TargetUserID, true),
		ledgerstatementmw.WithCurrencyID(h.CoinTypeID, true),
		ledgerstatementmw.WithIOType(&ioType, true),
		ledgerstatementmw.WithIOSubType(&ioSubtype, true),
		ledgerstatementmw.WithAmount(h.Amount, true),
		ledgerstatementmw.WithIOExtra(&ioExtra, true),
	)
	if err != nil {
		return nil, err
	}

	if _, err := ledgerStatementHandler.CreateStatement(ctx); err != nil {
		return nil, err
	}

	ledgerStatementConds := &statementpb.Conds{
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: *h.TargetAppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: *h.TargetUserID},
		CurrencyID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.CoinTypeID},
		IOExtra:    &basetypes.StringVal{Op: cruder.LIKE, Value: ioExtra},
	}
	ledgerStatementHandler, err = ledgerstatementmw.NewHandler(
		ctx,
		ledgerstatementmw.WithConds(ledgerStatementConds),
	)
	if err != nil {
		return nil, err
	}

	statement, err := ledgerStatementHandler.GetStatementOnly(ctx)
	if err != nil {
		return nil, err
	}
	if statement == nil {
		return nil, fmt.Errorf("fail get statement")
	}

	return &npool.Statement{
		ID:           statement.ID,
		EntID:        statement.EntID,
		UserID:       statement.UserID,
		EmailAddress: handler.user.EmailAddress,
		CurrencyID:   *h.CoinTypeID,
		CurrencyType: types.CurrencyType_CurrencyCrypto,
		CurrencyName: coin.Name,
		DisplayNames: coin.DisplayNames,
		CurrencyLogo: coin.Logo,
		CurrencyUnit: coin.Unit,
		IOType:       statement.IOType,
		IOSubType:    statement.IOSubType,
		Amount:       statement.Amount,
		IOExtra:      statement.IOExtra,
		CreatedAt:    statement.CreatedAt,
	}, nil
}
