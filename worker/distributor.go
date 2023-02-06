package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

// TaskDistributor ...
type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(
		ctx context.Context,
		payload *PayloadSendVerifyEmail,
		opts ...asynq.Option,
	) error
}

// RedisTaskDistributor ...
type RedisTaskDistributor struct {
	client *asynq.Client
}

// NewRedisTaskDistributor ...
func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{client: client}
}

// TaskSendVerifyEmail ...
const TaskSendVerifyEmail = "task:send_verify_email"

// DistributeTaskSendVerifyEmail ...
func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(ctx context.Context, payload *PayloadSendVerifyEmail, opts ...asynq.Option) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task")
	}
	task := asynq.NewTask(TaskSendVerifyEmail, data, opts...)

	info, err := distributor.client.Enqueue(task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}
	log.Printf("type: %s, payload: %v, queue: %s, max_retry: %d\n", task.Type(), task.Payload(), info.Queue, info.MaxRetry)

	return nil
}
