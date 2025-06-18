package sentinel

import (
	"fmt"

	npool "github.com/NpoolPlatform/kunman/message/scheduler/middleware/v1/sentinel"
	benefitwait "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/wait"
	benefitwaittypes "github.com/NpoolPlatform/kunman/cron/scheduler/benefit/powerrental/wait/types"
)

func triggerBenefitWait(req *npool.BenefitWait) {
	benefitwait.Trigger(&benefitwaittypes.TriggerCond{
		GoodIDs:  req.GetGoodIDs(),
		RewardAt: req.GetRewardAt(),
	})
}

func Trigger(req *npool.TriggerRequest) error {
	switch req.Subsystem {
	case "benefitwait":
		triggerBenefitWait(req.GetBenefitWait())
	default:
		return fmt.Errorf("invalid subsystem")
	}
	return nil
}
