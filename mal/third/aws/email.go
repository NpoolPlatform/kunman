package aws

import (
	"fmt"

	"github.com/NpoolPlatform/kunman/framework/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	CharSet = "UTF-8"
)

func SendEmail(subject, content, from, to string, replyTo ...string) error {
	myServiceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	region := config.GetStringValueWithNameSpace(myServiceName, Region)
	accessKey := config.GetStringValueWithNameSpace(myServiceName, AccessKey)
	secretKey := config.GetStringValueWithNameSpace(myServiceName, SecretKey)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return fmt.Errorf("new aws session error: %v", err)
	}

	svc := ses.New(sess)
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(content),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(from),
	}

	if len(replyTo) != 0 {
		input.ReplyToAddresses = aws.StringSlice(replyTo)
	}

	_, err = svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				return fmt.Errorf("%v, %v", ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				return fmt.Errorf("%v, %v", ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				return fmt.Errorf("%v, %v", ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				return aerr
			}
		} else {
			return err
		}
	}

	return nil
}
