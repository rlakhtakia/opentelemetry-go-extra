module github.com/rlakhtakia/opentelemetry-go-extra/otelgorm/example

go 1.19

replace github.com/rlakhtakia/opentelemetry-go-extra/otelsql => ../../otelsql

replace github.com/rlakhtakia/opentelemetry-go-extra/otelgorm => ./..

replace github.com/rlakhtakia/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/rlakhtakia/opentelemetry-go-extra/otelgorm v0.2.2
	gorm.io/driver/sqlite v1.5.2
	gorm.io/gorm v1.25.2
)

require (
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/rlakhtakia/opentelemetry-go-extra/otelsql v0.2.2 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.42.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.39.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.39.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.16.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.16.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk v1.16.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.39.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

require (
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/rlakhtakia/opentelemetry-go-extra/otelplay v0.2.2
	github.com/uptrace/uptrace-go v1.16.0 // indirect
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/proto/otlp v0.20.0 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	google.golang.org/grpc v1.56.1 // indirect
)

replace github.com/rlakhtakia/opentelemetry-go-extra => /usr/local/google/home/rlakhtakia/go/src/opentelemetry-go-extra

replace github.com/uptrace/opentelemetry-go-extra => /usr/local/google/home/rlakhtakia/go/src/opentelemetry-go-extra
