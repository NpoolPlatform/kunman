package types

import (
	notifbenefitmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif/goodbenefit"
)

type NotifContent struct {
	AppID   string
	Content string
}

type PersistentGoodBenefit struct {
	Benefits      []*notifbenefitmwpb.GoodBenefit
	NotifContents []*NotifContent
}
