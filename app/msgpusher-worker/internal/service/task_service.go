package service

import "austin-v2/app/msgpusher-worker/internal/service/srv"

type TaskService struct {
	DiscardMessageService *srv.DiscardMessageService
	ShieldService         *srv.ShieldService
	DeduplicationService  *srv.DeduplicationRuleService
}

func NewTaskService(
	discardMessageService *srv.DiscardMessageService,
	shieldService *srv.ShieldService,
	deduplicationService *srv.DeduplicationRuleService,
) *TaskService {
	return &TaskService{
		DiscardMessageService: discardMessageService,
		ShieldService:         shieldService,
		DeduplicationService:  deduplicationService,
	}
}
