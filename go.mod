module github.com/NpoolPlatform/kunman

go 1.24

require (
	entgo.io/ent v0.14.4
	github.com/NpoolPlatform/basal-middleware v0.0.0-20240731030616-5ed3dac01fec
	github.com/NpoolPlatform/billing-middleware v0.0.0-20250604121136-abce5f7aadb9
	github.com/NpoolPlatform/go-service-framework v0.0.0-20250123031703-66a23c81e8e0
	github.com/ThreeDotsLabs/watermill v1.2.0
	github.com/ThreeDotsLabs/watermill-amqp/v2 v2.0.7
	github.com/aws/aws-sdk-go-v2 v1.16.14
	github.com/aws/aws-sdk-go-v2/config v1.17.5
	github.com/aws/aws-sdk-go-v2/credentials v1.12.18
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.9
	github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
	github.com/go-chassis/go-archaius v1.5.3
	github.com/go-chassis/openlog v1.1.3
	github.com/go-chi/chi/v5 v5.0.8
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/hashicorp/consul/api v1.12.0
	github.com/prometheus/client_golang v1.14.0
	github.com/shopspring/decimal v1.4.0
	github.com/spf13/viper v1.11.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.8.2
	github.com/urfave/cli/v2 v2.4.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.31.0
	go.opentelemetry.io/otel v1.8.0
	go.opentelemetry.io/otel/exporters/jaeger v1.6.3
	go.opentelemetry.io/otel/sdk v1.6.3
	go.uber.org/zap v1.19.1
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2
	google.golang.org/grpc v1.55.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0
)

require github.com/ugorji/go/codec v1.2.14 // indirect
