package persistent

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	notifmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	notifbenefitmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/goodbenefit"
	tmplmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template"
	notifmwcli "github.com/NpoolPlatform/notif-middleware/pkg/client/notif"
	notifbenefitmwcli "github.com/NpoolPlatform/notif-middleware/pkg/client/notif/goodbenefit"
	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/notif/benefit/powerrental/types"
)

type handler struct{}

func NewPersistent() basepersistent.Persistenter {
	return &handler{}
}

func (p *handler) Update(ctx context.Context, benefit interface{}, reward, notif, done chan interface{}) error {
	_benefit, ok := benefit.(*types.PersistentGoodBenefit)
	if !ok {
		return fmt.Errorf("invalid benefit")
	}

	defer asyncfeed.AsyncFeed(ctx, _benefit, done)

	for _, content := range _benefit.NotifContents {
		if _, err := notifmwcli.GenerateNotifs(ctx, &notifmwpb.GenerateNotifsRequest{
			AppID:     content.AppID,
			EventType: basetypes.UsedFor_GoodBenefit1,
			NotifType: basetypes.NotifType_NotifMulticast,
			Vars: &tmplmwpb.TemplateVars{
				Message: &content.Content,
			},
		}); err != nil {
			continue
		}
	}

	generated := true
	for _, benefit := range _benefit.Benefits {
		if _, err := notifbenefitmwcli.UpdateGoodBenefit(ctx, &notifbenefitmwpb.GoodBenefitReq{
			ID:        &benefit.ID,
			Generated: &generated,
		}); err != nil {
			return err
		}
	}

	return nil
}
