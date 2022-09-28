module github.com/ZQCard/kratos-base-project

go 1.16

require (
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20220414054820-d0b704b8f38d
	github.com/go-kratos/kratos/v2 v2.2.1
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.1
	github.com/hashicorp/consul/api v1.12.0
	go.opentelemetry.io/otel v1.10.0
	go.opentelemetry.io/otel/exporters/jaeger v1.6.3
	go.opentelemetry.io/otel/sdk v1.6.3
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/genproto v0.0.0-20220421151946-72621c1f0bd3
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
)

require (
	github.com/bwmarrin/snowflake v0.3.0
	github.com/casbin/casbin/v2 v2.52.1
	github.com/casbin/gorm-adapter/v3 v3.7.4
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.11.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/lib/pq v1.10.6 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.18.1 // indirect
)
