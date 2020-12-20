package utility

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-xray-sdk-go/xray"
)

var sqsSvc *sqs.SQS

func init() {
	if os.Getenv("ENVIRONMENT") == "production" {
		sess, _ := session.NewSession()
		sqsSvc = sqs.New(sess)
	} else {
		keyID := "local"
		keySecret := "local"

		creds := credentials.NewStaticCredentials(keyID, keySecret, "")

		sess := session.Must(session.NewSession(&aws.Config{
			Region:      aws.String("ap-southeast-1"),
			Endpoint:    aws.String("http://sqs:9324"),
			Credentials: creds,
		}))

		sqsSvc = sqs.New(sess)
	}
	// enable sqs within xray
	xray.AWS(sqsSvc.Client)
}

/*
SendSQSMsg ...
*/
func SendSQSMsg(ctx context.Context, message string) (*sqs.SendMessageOutput, error) {
	sendMsgInput := &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(getQueueURL()),
	}
	return sqsSvc.SendMessageWithContext(ctx, sendMsgInput)
}

func getQueueURL() string {
	SQSQueueURL := os.Getenv("SQS_QUEUE_URL")
	if SQSQueueURL != "" {
		return SQSQueueURL
	}
	return "http://sqs:9324/queue/default"
}
