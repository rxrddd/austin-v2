package initDB

import (
	"github.com/ZQCard/kratos-base-project/app/jobs/service/conf"
	"github.com/hibiken/asynq"
	"time"
)

func CallInitDB(data *conf.Data) {
	cstSh, _ := time.LoadLocation("Asia/Shanghai")

	redisConnOpt := asynq.RedisClientOpt{
		Addr:     data.Redis.Addr,
		Password: data.Redis.Password,
	}
	scheduler := asynq.NewScheduler(
		redisConnOpt,
		&asynq.SchedulerOpts{Location: cstSh},
	)
	payload, _ := NewInitDbTask(data)
	// 10分钟一次恢复数据库
	if _, err := scheduler.Register("@every 10m", payload); err != nil {
		panic(err)
	}

	// Run blocks and waits for os signal to terminate the program.
	if err := scheduler.Run(); err != nil {
		panic(err)
	}
}
