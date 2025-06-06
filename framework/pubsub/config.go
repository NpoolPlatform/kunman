package pubsub

import (
	"fmt"

	"github.com/NpoolPlatform/kunman/framework/config"
	constant "github.com/NpoolPlatform/kunman/framework/rabbitmq/const"

	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
)

func DurablePubSubConfig() (*amqp.Config, error) {
	service, err := config.PeekService(constant.RabbitMQServiceName)
	if err != nil {
		return nil, fmt.Errorf("fail to query rabbitmq service: %v", err)
	}

	username := config.GetStringValueWithNameSpace(
		constant.RabbitMQServiceName,
		config.KeyUsername,
	)
	password := config.GetStringValueWithNameSpace(
		constant.RabbitMQServiceName,
		config.KeyPassword,
	)

	if username == "" {
		return nil, fmt.Errorf("invalid username")
	}
	if password == "" {
		return nil, fmt.Errorf("invalid password")
	}

	rsl := fmt.Sprintf(
		"amqp://%v:%v@%v:%v/%v",
		username,
		password,
		service.Address,
		service.Port,
		GlobalPubsubTopic,
	)

	amqpConfig := amqp.NewDurablePubSubConfig(rsl, func(topic string) string {
		return config.ServiceNameToNamespace(
			config.GetStringValueWithNameSpace(
				"",
				config.KeyHostname,
			),
		)
	})

	amqpConfig.Publish.ConfirmDelivery = true

	return &amqpConfig, nil
}

func Sender() string {
	return config.ServiceNameToNamespace(
		config.GetStringValueWithNameSpace(
			"",
			config.KeyHostname,
		),
	)
}
