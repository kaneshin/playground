package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	config *aws.Config
	name   = os.Getenv("AWS_SQS_QUEUE_NAME")
	region = os.Getenv("AWS_SQS_QUEUE_REGION")
)

func init() {
	// Create shared configuration of AWS.
	config = aws.NewConfig().
		WithCredentials(credentials.NewEnvCredentials()).
		WithHTTPClient(http.DefaultClient).
		WithMaxRetries(aws.UseServiceDefaultRetries).
		WithLogger(aws.NewDefaultLogger()).
		WithLogLevel(aws.LogOff)
}

// NewSQS creates a new instance of the SQS client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a SQS client with additional configuration
//     svc := NewSQS(aws.NewConfig().WithRegion("us-west-2"))
func NewSQS(c *aws.Config) *sqs.SQS {
	sess := session.New(config, c)
	return sqs.New(sess, config, c)
}

func main() {
	// Initialize the SQS service.
	svc := NewSQS(aws.NewConfig().WithRegion(region))

	var url string
	{
		// Get Queue URL
		resp, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
			QueueName: aws.String(name),
		})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Queue URL: %v\n", *resp.QueueUrl)
		url = *resp.QueueUrl
	}

	{
		// Enqueue
		resp, err := svc.SendMessage(&sqs.SendMessageInput{
			MessageBody:  aws.String(time.Now().String()),
			DelaySeconds: aws.Int64(1),
			QueueUrl:     aws.String(url),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("Enqueue:", resp)
	}

	{
		// Bulk enquene
		input := &sqs.SendMessageBatchInput{
			QueueUrl: aws.String(url),
		}
		for i := 0; i < 2; i++ {
			ent := &sqs.SendMessageBatchRequestEntry{
				Id:          aws.String(strconv.Itoa(i)),
				MessageBody: aws.String(time.Now().String()),
			}
			input.Entries = append(input.Entries, ent)
		}
		resp, err := svc.SendMessageBatch(input)
		if err != nil {
			panic(err)
		}
		fmt.Println("Bulk Enqueue:", resp)
	}

	var messages []*sqs.Message
	{
		// Dequeue
		resp, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			MaxNumberOfMessages: aws.Int64(3),
			WaitTimeSeconds:     aws.Int64(10),
			QueueUrl:            aws.String(url),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("Dequeue:", resp)
		messages = resp.Messages
	}

	{
		// Delete
		for _, msg := range messages {
			resp, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
				ReceiptHandle: aws.String(*msg.ReceiptHandle),
				QueueUrl:      aws.String(url),
			})
			if err != nil {
				panic(err)
			}
			fmt.Println("Delete:", resp)
		}
	}
}
