package push

import (
	"context"
	"encoding/json"
	"transmitter/sqs"
	targetservice "transmitter/target_service"
)

type IPushService interface {
	Push(ctx context.Context, req map[string]interface{}) error
}

type pushService struct {
	qService sqs.IQueue
	tService targetservice.ITargetService
}

func NewPushService(qService sqs.IQueue, tService targetservice.ITargetService) IPushService {
	return &pushService{qService: qService, tService: tService}
}

func (p *pushService) Push(ctx context.Context, req map[string]interface{}) error {
	queueURL, err := p.tService.GetSQSUrlForBot(ctx)
	if err != nil {
		return err
	}
	jsonString, _ := json.Marshal(req)
	err = p.qService.SendMessage(ctx, queueURL, string(jsonString))
	if err != nil {
		return err
	}
	return nil
}
