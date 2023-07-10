# OpenTelemetry instrumentations for Go (Compatible with Otel v0.20.0)

[![build workflow](https://github.com/rlakhtakia/opentelemetry-go-extra/actions/workflows/build.yml/badge.svg)](https://github.com/rlakhtakia/opentelemetry-go-extra/actions/workflows/build.yml)
[![Chat](https://img.shields.io/badge/-telegram-red?color=white&logo=telegram&logoColor=black)](https://t.me/uptrace)

| Instrumentation Package         | Metrics            | Traces             |
| ------------------------------- | ------------------ | ------------------ |
| [database/sql](/otelsql/)       | :heavy_check_mark: | :heavy_check_mark: |
| [GORM 1](/otelsql/#gorm-1)      | :heavy_check_mark: | :heavy_check_mark: |
| [GORM 2](/otelgorm/)            | :heavy_check_mark: | :heavy_check_mark: |
| [sqlx](/otelsqlx/)              | :heavy_check_mark: | :heavy_check_mark: |
| [sqlboler](/otelsql/#sqlboiler) | :heavy_check_mark: | :heavy_check_mark: |
| [logrus](/otellogrus/)          |                    | :heavy_check_mark: |
| [Zap](/otelzap/)                |                    | :heavy_check_mark: |
| [graphql-go](/otelgraphql/)     |                    | :heavy_check_mark: |

You can search for instrumentations [here](https://uptrace.dev/get/instrument/) and
[here](https://opentelemetry.io/registry/).

## Contributing

To simiplify maintenance, we use a single version and a shared [changelog](CHANGELOG.md) for all
instrumentations. The changelog is auto-generated from
[conventional commits](https://www.conventionalcommits.org/en/v1.0.0/).

If you want to contribute an instrumentation, please make sure to include tests and a runnable
example. Use Docker if you must but try to avoid it, for example, you can use SQLite instead of
MySQL to test database/sql instrumentation. Use [instrum-example](/instrum-example/) instrumentation
as a template.

To run all tests:

```shell
./scripts/test.sh
```
