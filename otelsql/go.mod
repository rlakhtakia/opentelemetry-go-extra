module github.com/rlakhtakia/opentelemetry-go-extra/otelsql

go 1.19

require (
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/metric v1.16.0
	go.opentelemetry.io/otel/trace v1.16.0
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
)

replace github.com/rlakhtakia/opentelemetry-go-extra => ./

replace github.com/uptrace/opentelemetry-go-extra => ./
