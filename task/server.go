package task

import (
	"log"

	"github.com/hibiken/asynq"
)

func Process() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// workers
			Concurrency: 10,
			// Optionally specify multiple queues with differnt priority
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeEmailDelivery, HandleEmailDeliveryTask)
	mux.Handle(TypeImageReSize, NewImageProcessor())

	if err := srv.Run(mux); err != nil {
		log.Fatalf("cound not run server: %v", err)
	}
}
