package persistent

import (
	"context"
	"fmt"

	asyncfeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/asyncfeed"
	basepersistent "github.com/NpoolPlatform/kunman/cron/scheduler/base/persistent"
	types "github.com/NpoolPlatform/kunman/cron/scheduler/notif/benefit/powerrental/types"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	tmplmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template"
	notifmw "github.com/NpoolPlatform/kunman/middleware/notif/notif"
	notifbenefitmw "github.com/NpoolPlatform/kunman/middleware/notif/notif/goodbenefit"
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
		handler, err := notifmw.NewHandler(
			ctx,
			notifmw.WithAppID(&content.AppID, true),
			notifmw.WithEventType(basetypes.UsedFor_GoodBenefit1.Enum(), true),
			notifmw.WithNotifType(basetypes.NotifType_NotifMulticast.Enum(), true),
			notifmw.WithVars(&tmplmwpb.TemplateVars{
				Message: &content.Content,
			}, true),
		)
		if err != nil {
			return err
		}

		if _, err := handler.GenerateNotifs(ctx); err != nil {
			continue
		}
	}

	generated := true
	for _, benefit := range _benefit.Benefits {
		handler, err := notifbenefitmw.NewHandler(
			ctx,
			notifbenefitmw.WithID(&benefit.ID, true),
			notifbenefitmw.WithGenerated(&generated, true),
		)
		if err != nil {
			return err
		}

		if _, err := handler.UpdateGoodBenefit(ctx); err != nil {
			return err
		}
	}

	return nil
}
