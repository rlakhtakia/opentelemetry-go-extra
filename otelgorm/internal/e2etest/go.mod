module github.com/rlakhtakia/opentelemetry-go-extra/otelgorm/internal/e2e

go 1.19

replace github.com/rlakhtakia/opentelemetry-go-extra/otelgorm => ./../..

replace github.com/rlakhtakia/opentelemetry-go-extra/otelsql => ./../../../otelsql

require (
	github.com/rlakhtakia/opentelemetry-go-extra/otelgorm v0.2.2
	github.com/stretchr/testify v1.8.3
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/sdk v1.15.1
	go.opentelemetry.io/otel/trace v1.16.0
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gorm v1.25.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rlakhtakia/opentelemetry-go-extra/otelsql v0.2.2 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/rlakhtakia/opentelemetry-go-extra => ./

replace github.com/uptrace/opentelemetry-go-extra => ./
