module github.com/NpoolPlatform/kunman

go 1.24

require (
	entgo.io/ent v0.14.4
	github.com/NpoolPlatform/account-middleware v0.0.0-20250115034541-5f3f60ceee9a
	github.com/NpoolPlatform/appuser-middleware v0.0.0-20240926060520-1949246cf711
	github.com/NpoolPlatform/chain-middleware v0.0.0-20240731030645-1b4ce8fcac22
	github.com/NpoolPlatform/good-gateway v0.0.0-20250303022721-b1e88c6dc0af
	github.com/NpoolPlatform/good-middleware v0.0.0-20250115034449-9d9d053987a2
	github.com/NpoolPlatform/libent-cruder v0.0.0-20240323085036-7679b4818ce7
	github.com/NpoolPlatform/message v0.0.0-20250604112546-25213bae6bb9
	github.com/NpoolPlatform/miningpool-middleware v0.0.0-20241106020300-03d147aeb913
	github.com/NpoolPlatform/order-middleware v0.0.0-20241106020324-b91339b3d20e
	github.com/NpoolPlatform/sphinx-proxy v0.0.0-20231020115438-bd66d3d334af
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
	github.com/go-resty/resty/v2 v2.7.0
	github.com/go-sql-driver/mysql v1.7.0
	github.com/google/uuid v1.6.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/hashicorp/consul/api v1.12.0
	github.com/prometheus/client_golang v1.14.0
	github.com/shopspring/decimal v1.4.0
	github.com/spf13/viper v1.12.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.10.0
	github.com/urfave/cli/v2 v2.17.2-0.20221006022127-8f469abc00aa
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.31.0
	go.opentelemetry.io/otel v1.10.0
	go.opentelemetry.io/otel/exporters/jaeger v1.6.3
	go.opentelemetry.io/otel/sdk v1.6.3
	go.uber.org/zap v1.24.0
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.36.5
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	ariga.io/atlas v0.31.1-0.20250212144724-069be8033e83 // indirect
	github.com/NpoolPlatform/go-service-framework v0.0.0-20250123031703-66a23c81e8e0 // indirect
	github.com/Shonminh/apollo-client v0.4.0 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.7 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.21 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.15 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.22 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.21 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.17 // indirect
	github.com/aws/smithy-go v1.13.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coocood/freecache v1.0.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.2.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/hashicorp/serf v0.9.7 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.7 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rabbitmq/amqp091-go v1.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	github.com/ugorji/go/codec v1.2.14 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	github.com/zclconf/go-cty v1.14.4 // indirect
	github.com/zclconf/go-cty-yaml v1.1.0 // indirect
	go.opentelemetry.io/otel/trace v1.10.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	go.uber.org/ratelimit v0.2.0 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/tools v0.31.0 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
