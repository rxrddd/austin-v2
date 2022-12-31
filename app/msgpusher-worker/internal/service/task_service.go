package service

type TaskService struct {
	DiscardMessageService    *DiscardMessageService
	ShieldService            *ShieldService
	DeduplicationRuleService *DeduplicationRuleService
}

func NewTaskService(
	discardMessageService *DiscardMessageService,
	shieldService *ShieldService,
	deduplicationRuleService *DeduplicationRuleService,
) *TaskService {
	return &TaskService{
		DiscardMessageService:    discardMessageService,
		ShieldService:            shieldService,
		DeduplicationRuleService: deduplicationRuleService,
	}
}
