package types

import (
	withdrawreviewnotifypb "github.com/NpoolPlatform/kunman/message/scheduler/middleware/v1/withdraw/review/notify"
)

type PersistentWithdrawReviewNotify struct {
	AppWithdraws []*withdrawreviewnotifypb.AppWithdrawInfos
}
