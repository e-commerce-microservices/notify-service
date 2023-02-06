package task

import (
	"log"

	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

// SendTask ...
func SendTask() {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: redisAddr,
	})
	defer client.Close()

	//
	// Example 1: Enqueue task to be processed immediately
	// Use (*Client).Enqueue method

	task, err := NewEmailDeliveryTask(1304, "Hello world")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
