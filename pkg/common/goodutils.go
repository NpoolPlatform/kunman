package common

import (
	"context"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodmwcli "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	goodtypes "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	"github.com/google/uuid"
)

func GetGoods(ctx context.Context, goodIDs []string) (map[string]*goodmwpb.Good, error) {
	for _, goodID := range goodIDs {
		if _, err := uuid.Parse(goodID); err != nil {
			return nil, wlog.WrapError(err)
		}
	}

	goods, _, err := goodmwcli.GetGoods(ctx, &goodmwpb.Conds{
		EntIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: goodIDs},
	}, int32(0), int32(len(goodIDs)))
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	goodMap := map[string]*goodmwpb.Good{}
	for _, good := range goods {
		goodMap[good.EntID] = good
	}
	return goodMap, nil
}

func GoodDurationDisplayType2Unit(_type goodtypes.GoodDurationType, seconds uint32) (units uint32, unit string) {
	switch _type {
	case goodtypes.GoodDurationType_GoodDurationByHour:
		units = seconds / timedef.SecondsPerHour
		unit = "MSG_HOUR"
	case goodtypes.GoodDurationType_GoodDurationByDay:
		units = seconds / timedef.SecondsPerDay
		unit = "MSG_DAY"
	case goodtypes.GoodDurationType_GoodDurationByMonth:
		units = seconds / timedef.SecondsPerMonth
		unit = "MSG_MONTH"
	case goodtypes.GoodDurationType_GoodDurationByYear:
		units = seconds / timedef.SecondsPerYear
		unit = "MSG_YEAR"
	}
	if units > 1 {
		unit += "S"
	}
	return units, unit
}

func GoodDurationDisplayType2Seconds(_type goodtypes.GoodDurationType) (units uint32) {
	switch _type {
	case goodtypes.GoodDurationType_GoodDurationByHour:
		return timedef.SecondsPerHour
	case goodtypes.GoodDurationType_GoodDurationByDay:
		return timedef.SecondsPerDay
	case goodtypes.GoodDurationType_GoodDurationByMonth:
		return timedef.SecondsPerMonth
	case goodtypes.GoodDurationType_GoodDurationByYear:
		return timedef.SecondsPerYear
	}
	return timedef.SecondsPerHour
}
