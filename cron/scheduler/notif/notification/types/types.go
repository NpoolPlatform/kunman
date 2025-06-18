package types

import (
	notifmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/notif"
	sendmwpb "github.com/NpoolPlatform/kunman/message/third/middleware/v1/send"
)

type PersistentNotif struct {
	*notifmwpb.Notif
	EventNotifs    []*notifmwpb.Notif
	MessageRequest *sendmwpb.SendMessageInput
}
