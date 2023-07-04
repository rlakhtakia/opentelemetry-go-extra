module github.com/rlakhtakia/opentelemetry-go-extra/otellogrus

go 1.19

replace github.com/rlakhtakia/opentelemetry-go-extra/otelutil => ../otelutil

require (
	github.com/rlakhtakia/opentelemetry-go-extra/otelutil v0.2.2
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.3
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/sdk v1.15.1
	go.opentelemetry.io/otel/trace v1.16.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/rlakhtakia/opentelemetry-go-extra => /usr/local/google/home/rlakhtakia/go/src/opentelemetry-go-extra

replace github.com/uptrace/opentelemetry-go-extra => /usr/local/google/home/rlakhtakia/go/src/opentelemetry-go-extra
