package types

import (
	ancmwpb "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/announcement"
	sendmwpb "github.com/NpoolPlatform/kunman/message/third/middleware/v1/send"
)

type PersistentAnnouncement struct {
	*ancmwpb.Announcement
	SendAppID      string
	SendUserID     string
	MessageRequest *sendmwpb.SendMessageInput
}
