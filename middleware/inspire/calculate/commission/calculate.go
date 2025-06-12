package commission

import (
	"context"

	"github.com/NpoolPlatform/kunman/framework/wlog"
	types "github.com/NpoolPlatform/kunman/message/basetypes/inspire/v1"
	appcommissionconfig "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/commission/config"
	appgoodcommissionconfig "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/app/good/commission/config"
	npool "github.com/NpoolPlatform/kunman/message/inspire/middleware/v1/commission"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Commission struct {
	AppConfigID             string
	CommissionConfigID      string
	CommissionConfigType    types.CommissionConfigType
	AppID                   string
	UserID                  string
	DirectContributorUserID *string
	PaymentAmount           string
	Amount                  string
	CommissionAmountUSD     string
}

type calculateHandler struct {
	*Handler
}

//nolint:funlen
func (h *Handler) Calculate(ctx context.Context) ([]*Commission, error) {
	commMap := map[string]*npool.Commission{}
	for _, comm := range h.Commissions {
		commMap[comm.UserID] = comm
	}

	_comms := []*Commission{}

	for _, inviter := range h.Inviters {
		if inviter.InviterID == uuid.Nil.String() {
			break
		}

		percent1 := decimal.NewFromInt(0)
		percent2 := decimal.NewFromInt(0)

		var err error

		comm1, ok := commMap[inviter.InviteeID]
		if ok {
			percent1, err = decimal.NewFromString(comm1.GetAmountOrPercent())
			if err != nil {
				return nil, wlog.WrapError(err)
			}
		}

		comm2, ok := commMap[inviter.InviterID]
		if ok {
			percent2, err = decimal.NewFromString(comm2.GetAmountOrPercent())
			if err != nil {
				return nil, wlog.WrapError(err)
			}
		}

		if percent2.Cmp(percent1) < 0 {
			return nil, wlog.Errorf("%v/%v < %v/%v (%v)", inviter.InviterID, percent2, inviter.InviteeID, percent1, comm1.GetGoodID())
		}

		if percent2.Cmp(percent1) == 0 {
			commissionConfigID := uuid.Nil.String()
			if comm2 != nil {
				commissionConfigID = comm2.EntID
			}
			_comms = append(_comms, &Commission{
				AppConfigID:             h.AppConfig.EntID,
				CommissionConfigID:      commissionConfigID,
				CommissionConfigType:    types.CommissionConfigType_LegacyCommissionConfig,
				AppID:                   inviter.AppID,
				UserID:                  inviter.InviterID,
				DirectContributorUserID: &inviter.InviteeID,
				PaymentAmount:           h.PaymentAmount.String(),
				Amount:                  "0",
				CommissionAmountUSD:     "0",
			})
			continue
		}

		amount := h.PaymentAmount

		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			amount = decimal.NewFromInt(0)
		}

		_comms = append(_comms, &Commission{
			AppConfigID:             h.AppConfig.EntID,
			CommissionConfigID:      comm2.EntID,
			CommissionConfigType:    types.CommissionConfigType_LegacyCommissionConfig,
			AppID:                   inviter.AppID,
			UserID:                  inviter.InviterID,
			DirectContributorUserID: &inviter.InviteeID,
			PaymentAmount:           h.PaymentAmount.String(),
			Amount:                  amount.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(),             //nolint
			CommissionAmountUSD:     h.PaymentAmountUSD.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	commLast, ok := commMap[h.Inviters[len(h.Inviters)-1].InviteeID]
	if !ok {
		return _comms, nil
	}

	percent, err := decimal.NewFromString(commLast.GetAmountOrPercent())
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	amount := h.PaymentAmount

	if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		amount = decimal.NewFromInt(0)
	}

	amountLast := "0"
	if percent.Cmp(decimal.NewFromInt(0)) > 0 {
		amountLast = amount.Mul(percent).Div(decimal.NewFromInt(100)).String() //nolint
	}

	_comms = append(_comms, &Commission{
		AppConfigID:          h.AppConfig.EntID,
		CommissionConfigID:   commLast.EntID,
		CommissionConfigType: types.CommissionConfigType_LegacyCommissionConfig,
		AppID:                h.Inviters[len(h.Inviters)-1].AppID,
		UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
		PaymentAmount:        h.PaymentAmount.String(),
		Amount:               amountLast,
		CommissionAmountUSD:  h.PaymentAmountUSD.Mul(percent).Div(decimal.NewFromInt(100)).String(), //nolint
	})

	return _comms, nil
}

func (h *calculateHandler) getInvites(userID string) uint32 {
	achivmentUser, ok := h.AchievementUsers[userID]
	if !ok {
		return uint32(0)
	}

	return achivmentUser.DirectInvitees + achivmentUser.IndirectInvitees
}

//nolint:dupl
func (h *calculateHandler) getAppGoodCommLevelConf(userID string) (*appgoodcommissionconfig.AppGoodCommissionConfig, bool, error) {
	invites := h.getInvites(userID)
	_comm := &appgoodcommissionconfig.AppGoodCommissionConfig{}
	useful := false
	amount := h.PaymentAmountUSD
	consumeAmount := h.PaymentAmountUSD
	achivmentUser, ok := h.AchievementUsers[userID]
	if ok {
		directConsumeAmount, err := decimal.NewFromString(achivmentUser.DirectConsumeAmount)
		if err != nil {
			return nil, false, wlog.WrapError(err)
		}
		inviteeConsumeAmount, err := decimal.NewFromString(achivmentUser.InviteeConsumeAmount)
		if err != nil {
			return nil, false, wlog.WrapError(err)
		}
		consumeAmount = directConsumeAmount.Add(inviteeConsumeAmount).Add(amount)
	}

	percent := decimal.NewFromInt(0)
	for i, comm := range h.AppGoodCommissionConfigs {
		if i == 0 {
			_comm = comm
		}
		if invites < comm.Invites {
			break
		}
		thresholdAmount, err := decimal.NewFromString(comm.GetThresholdAmount())
		if err != nil {
			return nil, false, wlog.WrapError(err)
		}
		if consumeAmount.Cmp(thresholdAmount) < 0 {
			break
		}
		_percent, err := decimal.NewFromString(comm.GetAmountOrPercent())
		if err != nil {
			return nil, false, wlog.WrapError(err)
		}
		if _percent.Cmp(percent) > 0 {
			percent = _percent
			_comm = comm
			useful = true
		}
	}
	if percent.Cmp(decimal.NewFromInt(0)) < 0 {
		return _comm, false, nil
	}
	return _comm, useful, nil
}

//nolint:dupl
func (h *calculateHandler) getAppCommLevelConf(userID string) (*appcommissionconfig.AppCommissionConfig, bool, error) {
	invites := h.getInvites(userID)
	_comm := &appcommissionconfig.AppCommissionConfig{}
	useful := false
	amount := h.PaymentAmountUSD
	consumeAmount := h.PaymentAmountUSD

	achivmentUser, ok := h.AchievementUsers[userID]
	if ok {
		directConsumeAmount, err := decimal.NewFromString(achivmentUser.DirectConsumeAmount)
		if err != nil {
			return nil, false, wlog.WrapError(err)
		}
		inviteeConsumeAmount, err := decimal.NewFromString(achivmentUser.InviteeConsumeAmount)
		if err != nil {
			return nil, false, wlog.WrapError(err)
		}
		consumeAmount = directConsumeAmount.Add(inviteeConsumeAmount).Add(amount)
	}

	percent := decimal.NewFromInt(0)
	for i, comm := range h.AppCommissionConfigs {
		if i == 0 {
			_comm = comm
		}
		if invites < comm.Invites {
			break
		}
		thresholdAmount, err := decimal.NewFromString(comm.GetThresholdAmount())
		if err != nil {
			return nil, false, wlog.WrapError(err)
		}
		if consumeAmount.Cmp(thresholdAmount) < 0 {
			break
		}
		_percent, err := decimal.NewFromString(comm.GetAmountOrPercent())
		if err != nil {
			return nil, false, wlog.WrapError(err)
		}
		if _percent.Cmp(percent) > 0 {
			percent = _percent
			_comm = comm
			useful = true
		}
	}
	if percent.Cmp(decimal.NewFromInt(0)) < 0 {
		return _comm, false, nil
	}
	return _comm, useful, nil
}

//nolint:dupl,funlen
func (h *Handler) CalculateByAppCommConfig(ctx context.Context) ([]*Commission, error) {
	_comms := []*Commission{}
	handler := &calculateHandler{
		Handler: h,
	}
	for _, inviter := range h.Inviters {
		if inviter.InviterID == uuid.Nil.String() {
			break
		}

		percent1 := decimal.NewFromInt(0)
		percent2 := decimal.NewFromInt(0)

		var err error

		comm1, comm1Useful, err := handler.getAppCommLevelConf(inviter.InviteeID)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		if comm1 != nil && comm1Useful {
			percent1, err = decimal.NewFromString(comm1.GetAmountOrPercent())
			if err != nil {
				return nil, wlog.WrapError(err)
			}
		}

		if h.AppConfig.CommissionType == types.CommissionType_DirectCommission {
			percent1 = decimal.NewFromInt(0)
		}

		comm2, comm2Useful, err := handler.getAppCommLevelConf(inviter.InviterID)
		if err != nil {
			return nil, wlog.WrapError(err)
		}

		if comm2 != nil && comm2Useful {
			percent2, err = decimal.NewFromString(comm2.GetAmountOrPercent())
			if err != nil {
				return nil, wlog.WrapError(err)
			}
		}

		if percent2.Cmp(percent1) < 0 {
			return nil, wlog.Errorf("%v/%v < %v/%v", inviter.InviterID, percent2, inviter.InviteeID, percent1)
		}

		if percent2.Cmp(percent1) == 0 {
			_comms = append(_comms, &Commission{
				AppConfigID:             h.AppConfig.EntID,
				CommissionConfigID:      comm2.EntID,
				CommissionConfigType:    types.CommissionConfigType_AppCommissionConfig,
				AppID:                   inviter.AppID,
				UserID:                  inviter.InviterID,
				DirectContributorUserID: &inviter.InviteeID,
				PaymentAmount:           h.PaymentAmount.String(),
				Amount:                  "0",
				CommissionAmountUSD:     "0",
			})
			continue
		}

		amount := h.PaymentAmount

		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			amount = decimal.NewFromInt(0)
		}

		_comms = append(_comms, &Commission{
			AppConfigID:             h.AppConfig.EntID,
			CommissionConfigID:      comm2.EntID,
			CommissionConfigType:    types.CommissionConfigType_AppCommissionConfig,
			AppID:                   inviter.AppID,
			UserID:                  inviter.InviterID,
			DirectContributorUserID: &inviter.InviteeID,
			PaymentAmount:           h.PaymentAmount.String(),
			Amount:                  amount.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(),             //nolint
			CommissionAmountUSD:     h.PaymentAmountUSD.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	if h.AppConfig.CommissionType == types.CommissionType_DirectCommission {
		_comms = append(_comms, &Commission{
			AppConfigID:          h.AppConfig.EntID,
			CommissionConfigID:   uuid.Nil.String(),
			CommissionConfigType: types.CommissionConfigType_WithoutCommissionConfig,
			AppID:                h.Inviters[len(h.Inviters)-1].AppID,
			UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
			PaymentAmount:        h.PaymentAmount.String(),
			Amount:               "0",
			CommissionAmountUSD:  "0",
		})
		return _comms, nil
	}

	commLast, commLastUseful, err := handler.getAppCommLevelConf(h.Inviters[len(h.Inviters)-1].InviteeID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}

	if commLast == nil {
		return _comms, nil
	}

	percent, err := decimal.NewFromString(commLast.GetAmountOrPercent())
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if !commLastUseful {
		percent = decimal.NewFromInt(0)
	}

	amount := h.PaymentAmount

	if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		amount = decimal.NewFromInt(0)
	}

	amountLast := "0"
	if percent.Cmp(decimal.NewFromInt(0)) > 0 {
		amountLast = amount.Mul(percent).Div(decimal.NewFromInt(100)).String() //nolint
	}

	_comms = append(_comms, &Commission{
		AppConfigID:          h.AppConfig.EntID,
		CommissionConfigID:   commLast.EntID,
		CommissionConfigType: types.CommissionConfigType_AppCommissionConfig,
		AppID:                h.Inviters[len(h.Inviters)-1].AppID,
		UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
		PaymentAmount:        h.PaymentAmount.String(),
		Amount:               amountLast,
		CommissionAmountUSD:  h.PaymentAmountUSD.Mul(percent).Div(decimal.NewFromInt(100)).String(), //nolint
	})

	return _comms, nil
}

//nolint:dupl,funlen
func (h *Handler) CalculateByAppGoodCommConfig(ctx context.Context) ([]*Commission, error) {
	_comms := []*Commission{}
	handler := &calculateHandler{
		Handler: h,
	}

	for _, inviter := range h.Inviters {
		if inviter.InviterID == uuid.Nil.String() {
			break
		}

		percent1 := decimal.NewFromInt(0)
		percent2 := decimal.NewFromInt(0)

		var err error

		comm1, comm1Useful, err := handler.getAppGoodCommLevelConf(inviter.InviteeID)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		if comm1 != nil && comm1Useful {
			percent1, err = decimal.NewFromString(comm1.GetAmountOrPercent())
			if err != nil {
				return nil, wlog.WrapError(err)
			}
		}

		if h.AppConfig.CommissionType == types.CommissionType_DirectCommission {
			percent1 = decimal.NewFromInt(0)
		}

		comm2, comm2Useful, err := handler.getAppGoodCommLevelConf(inviter.InviterID)
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		if comm2 != nil && comm2Useful {
			percent2, err = decimal.NewFromString(comm2.GetAmountOrPercent())
			if err != nil {
				return nil, wlog.WrapError(err)
			}
		}

		if percent2.Cmp(percent1) < 0 {
			return nil, wlog.Errorf("%v/%v < %v/%v", inviter.InviterID, percent2, inviter.InviteeID, percent1)
		}

		if percent2.Cmp(percent1) == 0 {
			_comms = append(_comms, &Commission{
				AppConfigID:             h.AppConfig.EntID,
				CommissionConfigID:      comm2.EntID,
				CommissionConfigType:    types.CommissionConfigType_AppGoodCommissionConfig,
				AppID:                   inviter.AppID,
				UserID:                  inviter.InviterID,
				DirectContributorUserID: &inviter.InviteeID,
				PaymentAmount:           h.PaymentAmount.String(),
				Amount:                  "0",
				CommissionAmountUSD:     "0",
			})
			continue
		}

		amount := h.PaymentAmount

		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			amount = decimal.NewFromInt(0)
		}

		_comms = append(_comms, &Commission{
			AppConfigID:             h.AppConfig.EntID,
			CommissionConfigID:      comm2.EntID,
			CommissionConfigType:    types.CommissionConfigType_AppGoodCommissionConfig,
			AppID:                   inviter.AppID,
			UserID:                  inviter.InviterID,
			DirectContributorUserID: &inviter.InviteeID,
			PaymentAmount:           h.PaymentAmount.String(),
			Amount:                  amount.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(),             //nolint
			CommissionAmountUSD:     h.PaymentAmountUSD.Mul(percent2.Sub(percent1)).Div(decimal.NewFromInt(100)).String(), //nolint
		})
	}

	if h.AppConfig.CommissionType == types.CommissionType_DirectCommission {
		_comms = append(_comms, &Commission{
			AppConfigID:          h.AppConfig.EntID,
			CommissionConfigID:   uuid.Nil.String(),
			CommissionConfigType: types.CommissionConfigType_WithoutCommissionConfig,
			AppID:                h.Inviters[len(h.Inviters)-1].AppID,
			UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
			PaymentAmount:        h.PaymentAmount.String(),
			Amount:               "0",
			CommissionAmountUSD:  "0",
		})
		return _comms, nil
	}

	commLast, commLastUseful, err := handler.getAppGoodCommLevelConf(h.Inviters[len(h.Inviters)-1].InviteeID)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if commLast == nil {
		return _comms, nil
	}

	percent, err := decimal.NewFromString(commLast.GetAmountOrPercent())
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if !commLastUseful {
		percent = decimal.NewFromInt(0)
	}

	amount := h.PaymentAmount

	if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		amount = decimal.NewFromInt(0)
	}

	amountLast := "0"
	if percent.Cmp(decimal.NewFromInt(0)) > 0 {
		amountLast = amount.Mul(percent).Div(decimal.NewFromInt(100)).String() //nolint
	}

	_comms = append(_comms, &Commission{
		AppConfigID:          h.AppConfig.EntID,
		CommissionConfigID:   commLast.EntID,
		CommissionConfigType: types.CommissionConfigType_AppGoodCommissionConfig,
		AppID:                h.Inviters[len(h.Inviters)-1].AppID,
		UserID:               h.Inviters[len(h.Inviters)-1].InviteeID,
		PaymentAmount:        h.PaymentAmount.String(),
		Amount:               amountLast,
		CommissionAmountUSD:  h.PaymentAmountUSD.Mul(percent).Div(decimal.NewFromInt(100)).String(), //nolint
	})

	return _comms, nil
}
