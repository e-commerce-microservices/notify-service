package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// A list of task types
const (
	TypeEmailDelivery = "email:deliver"
	TypeImageReSize   = "image:resize"
)

// EmailDeliveryPayload ...
type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
}

// ImageResizePayload ...
type ImageResizePayload struct {
	SourceURL string
}

// NewEmailDeliveryTask ...
func NewEmailDeliveryTask(userID int, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{
		UserID:     userID,
		TemplateID: tmplID,
	})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

// NewImageResizeTask ...
func NewImageResizeTask(src string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageResizePayload{
		SourceURL: src,
	})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeImageReSize, payload, asynq.Timeout(20*time.Minute)), nil
}

// HandleEmailDeliveryTask ...
func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Sending Email to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID)
	// Email delivery code ...
	return nil
}

// ImageProcessor implements asynq.Handler interface
type ImageProcessor struct {
	// ...fields for struct
	SourceURL string
}

func (processor *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p ImageResizePayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Resizing image: src=%s", processor.SourceURL)

	return nil
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}
