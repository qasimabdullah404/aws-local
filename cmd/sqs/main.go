package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {
	ctx := context.Background()
	queueName := os.Getenv("SQS_QUEUE")
	endpoint := os.Getenv("LOCALSTACK_ENDPOINT")

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("queueName: %s endpoint: %s\n", queueName, endpoint)

	client := sqs.NewFromConfig(cfg, func(o *sqs.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})

	queue, err := client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	})
	if err != nil {
		log.Fatal(err)
	}

	client.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody: aws.String("Hello, World!"),
		QueueUrl:    queue.QueueUrl,
	})
}
