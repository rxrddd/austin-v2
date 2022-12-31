package service

import (
	"austin-v2/app/msgpusher-worker/internal/service/deduplication"
	limit "austin-v2/app/msgpusher-worker/internal/service/limiter"
	"austin-v2/app/msgpusher-worker/internal/service/srv"
	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	srv.NewDiscardMessageService,
	srv.NewShieldService,
	srv.NewDeduplicationRuleService,
	NewTaskService,
	limit.NewSimpleLimitService,
	limit.NewSlideWindowLimitService,
	limit.NewLimiterManager,
	deduplication.NewContentDeduplicationService,
	deduplication.NewFrequencyDeduplicationService,
	deduplication.NewDeduplicationManager,
)
