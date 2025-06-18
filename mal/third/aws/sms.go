package aws

import (
	"github.com/NpoolPlatform/kunman/framework/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendSMS(msg, to string) error {
	myServiceName := config.GetStringValueWithNameSpace("", config.KeyHostname)
	region := config.GetStringValueWithNameSpace(myServiceName, Region)
	accessKey := config.GetStringValueWithNameSpace(myServiceName, AccessKey)
	secretKey := config.GetStringValueWithNameSpace(myServiceName, SecretKey)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return err
	}

	pin := &sns.PublishInput{}
	pin.SetMessage(msg)
	pin.SetPhoneNumber(to)

	svc := sns.New(sess)
	_, err = svc.Publish(pin)
	if err != nil {
		return err
	}

	return nil
}
