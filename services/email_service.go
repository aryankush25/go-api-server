package services

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type EmailService struct {
    Client *ses.Client
    Sender string
}

func NewEmailService() (*EmailService, error) {
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
    if err != nil {
        return nil, fmt.Errorf("unable to load SDK config, %v", err)
    }

    client := ses.NewFromConfig(cfg)
    sender := os.Getenv("EMAIL_FROM")

    return &EmailService{
        Client: client,
        Sender: sender,
    }, nil
}

func (s *EmailService) SendEmail(to string, subject string, body string) error {
    input := &ses.SendEmailInput{
        Source: aws.String(s.Sender),
        Destination: &types.Destination{
            ToAddresses: []string{to},
        },
        Message: &types.Message{
            Body: &types.Body{
                Text: &types.Content{
                    Data: aws.String(body),
                },
            },
            Subject: &types.Content{
                Data: aws.String(subject),
            },
        },
    }

    _, err := s.Client.SendEmail(context.TODO(), input)
    if err != nil {
        return fmt.Errorf("failed to send email, %v", err)
    }

    return nil
}
