package record

import (
	"context"
	"fmt"

	gwcommon "github.com/NpoolPlatform/kunman/gateway/billing/common"
	recordmwcli "github.com/NpoolPlatform/kunman/middleware/billing/client/user/credit/record"
	cruder "github.com/NpoolPlatform/kunman/pkg/cruder/cruder"
	appmwpb "github.com/NpoolPlatform/kunman/message/appuser/mw/v1/app"
	usermwpb "github.com/NpoolPlatform/kunman/message/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	npool "github.com/NpoolPlatform/kunman/message/billing/gw/v1/user/credit/record"
	recordmwpb "github.com/NpoolPlatform/kunman/message/billing/mw/v1/user/credit/record"
)

type queryHandler struct {
	*Handler
	records []*recordmwpb.Record
	infos   []*npool.UserCreditRecord
	apps    map[string]*appmwpb.App
	users   map[string]*usermwpb.User
}

func (h *queryHandler) getApps(ctx context.Context) (err error) {
	h.apps, err = gwcommon.GetApps(ctx, func() (appIDs []string) {
		for _, record := range h.records {
			appIDs = append(appIDs, record.AppID)
		}
		return
	}())
	return err
}

func (h *queryHandler) getUsers(ctx context.Context) (err error) {
	h.users, err = gwcommon.GetUsers(ctx, func() (userIDs []string) {
		for _, record := range h.records {
			userIDs = append(userIDs, record.UserID)
		}
		return
	}())
	return err
}

func (h *queryHandler) formalize() {
	for _, recommend := range h.records {
		info := &npool.UserCreditRecord{
			ID:            recommend.ID,
			EntID:         recommend.EntID,
			AppID:         recommend.AppID,
			OperationType: recommend.OperationType,
			CreditsChange: recommend.CreditsChange,
			Extra:         recommend.Extra,
			CreatedAt:     recommend.CreatedAt,
			UpdatedAt:     recommend.UpdatedAt,
		}

		app, ok := h.apps[recommend.AppID]
		if ok {
			info.AppName = app.Name
		}
		user, ok := h.users[recommend.UserID]
		if ok {
			if user.Username != "" {
				info.Username = &user.Username
			}
			if user.EmailAddress != "" {
				info.EmailAddress = &user.EmailAddress
			}
			if user.PhoneNO != "" {
				info.PhoneNO = &user.PhoneNO
			}
		}
		h.infos = append(h.infos, info)
	}
}

func (h *Handler) GetUserCreditRecord(ctx context.Context) (*npool.UserCreditRecord, error) {
	record, err := recordmwcli.GetRecord(ctx, *h.EntID)
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, fmt.Errorf("invalid record")
	}

	handler := &queryHandler{
		Handler: h,
		records: []*recordmwpb.Record{record},
		apps:    map[string]*appmwpb.App{},
		users:   map[string]*usermwpb.User{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}

func (h *Handler) GetUserCreditRecords(ctx context.Context) ([]*npool.UserCreditRecord, error) {
	conds := &recordmwpb.Conds{}

	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}

	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}

	records, err := recordmwcli.GetRecords(ctx, conds, h.Offset, h.Limit)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, nil
	}

	handler := &queryHandler{
		Handler: h,
		records: records,
		apps:    map[string]*appmwpb.App{},
		users:   map[string]*usermwpb.User{},
	}
	if err := handler.getApps(ctx); err != nil {
		return nil, err
	}
	if err := handler.getUsers(ctx); err != nil {
		return nil, err
	}

	handler.formalize()
	return handler.infos, nil
}

func (h *Handler) GetUserCreditRecordsCount(ctx context.Context) (uint32, error) {
	conds := &recordmwpb.Conds{}

	if h.AppID != nil {
		conds.AppID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID}
	}

	if h.UserID != nil {
		conds.UserID = &basetypes.StringVal{Op: cruder.EQ, Value: *h.UserID}
	}
	return recordmwcli.GetRecordsCount(ctx, conds)
}
