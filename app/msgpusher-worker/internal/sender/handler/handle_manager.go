package handler

import (
	"austin-v2/pkg/manager"
	"austin-v2/pkg/types"
)

type HandleManager struct {
	manager *manager.Manager
}

func NewHandleManager(
	sms *SmsHandler,
	email *EmailHandler,
	officialAccount *OfficialAccountHandler,
	dingDingRobot *DingDingRobotHandler,
	dingDingWorkNotice *DingDingWorkNoticeHandler,
	miniProgramH *MiniProgramHandler,
) *HandleManager {
	return &HandleManager{
		manager: manager.NewManager(
			sms,
			email,
			officialAccount,
			dingDingRobot,
			dingDingWorkNotice,
			miniProgramH,
		),
	}
}

func (hm *HandleManager) Get(key string) (resp types.IHandler, err error) {
	if h, err := hm.manager.Get(key); err == nil {
		return h.(types.IHandler), nil
	}
	return nil, err
}
