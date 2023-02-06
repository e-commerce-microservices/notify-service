package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  any
}

// Start implements TaskProcessor
func (r RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, r.ProcessTaskSendVerifyEmail)

	return r.server.Start(mux)

}

// ProcessTaskSendVerifyEmail implements TaskProcessor
func (r RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	// TODO: send email to user
	log.Printf("type: %s payload: %v email: %s\n", task.Type(), task.Payload(), payload.Username)
	log.Println("processed task")

	return nil
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store any) TaskProcessor {
	server := asynq.NewServer(redisOpt, asynq.Config{})

	return RedisTaskProcessor{server: server, store: store}
}
