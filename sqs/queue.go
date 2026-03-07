package sqs

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type IQueue interface {
	SendMessage(ctx context.Context, queueURL string, message string) error
}

type queue struct {
	sqsClient *sqs.Client
}

func NewQueue(awsCfg *aws.Config) IQueue {
	return &queue{sqsClient: sqs.NewFromConfig(*awsCfg)}
}

func (q *queue) SendMessage(ctx context.Context, queueURL string, message string) error {
	output, err := q.sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(queueURL),
	})
	if err != nil {
		return err
	}
	if output.MessageId == nil {
		return errors.New("message id is nil")
	}
	return nil
}
