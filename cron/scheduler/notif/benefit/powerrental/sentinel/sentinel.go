package sentinel

import (
	"context"
	"fmt"
	"math"
	"os"
	"time"

	cancelablefeed "github.com/NpoolPlatform/kunman/cron/scheduler/base/cancelablefeed"
	basesentinel "github.com/NpoolPlatform/kunman/cron/scheduler/base/sentinel"
	timedef "github.com/NpoolPlatform/kunman/framework/const/time"
	goodtypes "github.com/NpoolPlatform/kunman/message/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/kunman/message/basetypes/v1"
	notifbenefitmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/goodbenefit"
	notifbenefitmw "github.com/NpoolPlatform/kunman/middleware/notif/notif/goodbenefit"
	constant "github.com/NpoolPlatform/kunman/pkg/const"
	"github.com/NpoolPlatform/kunman/pkg/cruder/cruder"

	"github.com/google/uuid"
)

type handler struct {
	ID              string
	nextBenefitAt   uint32
	benefitInterval uint32
}

func NewSentinel() basesentinel.Scanner {
	_interval := timedef.SecondsPerHour
	if interval, err := time.ParseDuration(
		fmt.Sprintf("%vm", os.Getenv("ENV_BENEFIT_NOTIFY_INTERVAL_MINS"))); err == nil && math.Round(interval.Seconds()) > 0 {
		_interval = int(math.Round(interval.Seconds()))
	}
	return &handler{
		ID:              uuid.NewString(),
		nextBenefitAt:   uint32((int(time.Now().Unix()) + _interval) / _interval * _interval),
		benefitInterval: uint32(_interval),
	}
}

func (h *handler) scanGoodBenefits(ctx context.Context, exec chan interface{}) error {
	offset := int32(0)
	limit := constant.DefaultRowLimit
	benefits := []*notifbenefitmwpb.GoodBenefit{}

	conds := &notifbenefitmwpb.Conds{
		Generated: &basetypes.BoolVal{Op: cruder.EQ, Value: false},
		GoodTypes: &basetypes.Uint32SliceVal{
			Op: cruder.IN,
			Value: []uint32{
				uint32(goodtypes.GoodType_PowerRental),
				uint32(goodtypes.GoodType_LegacyPowerRental),
			},
		},
	}

	for {
		handler, err := notifbenefitmw.NewHandler(
			ctx,
			notifbenefitmw.WithConds(conds),
			notifbenefitmw.WithOffset(offset),
			notifbenefitmw.WithLimit(limit),
		)
		if err != nil {
			return err
		}

		_benefits, _, err := handler.GetGoodBenefits(ctx)
		if err != nil {
			return err
		}
		if len(_benefits) == 0 {
			break
		}
		benefits = append(benefits, _benefits...)
		offset += limit
	}
	if len(benefits) > 0 {
		cancelablefeed.CancelableFeed(ctx, benefits, exec)
	}
	return nil
}

func (h *handler) Scan(ctx context.Context, exec chan interface{}) error {
	if uint32(time.Now().Unix()) < h.nextBenefitAt {
		return nil
	}
	if err := h.scanGoodBenefits(ctx, exec); err != nil {
		return err
	}
	h.nextBenefitAt = (uint32(time.Now().Unix()) + h.benefitInterval) / h.benefitInterval * h.benefitInterval
	return nil
}

func (h *handler) InitScan(ctx context.Context, exec chan interface{}) error {
	return h.scanGoodBenefits(ctx, exec)
}

func (h *handler) TriggerScan(ctx context.Context, cond interface{}, exec chan interface{}) error {
	return nil
}

func (h *handler) ObjectID(ent interface{}) string {
	return h.ID
}
