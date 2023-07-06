PACKAGE_DIRS := $(shell find . -type f -name 'go.mod' -exec dirname {} \; | sort)

test:
	go test ./...
	go test ./... -short -race
	go test ./... -run=NONE -bench=. -benchmem
	env GOOS=linux GOARCH=386 CGO_ENABLED=1 go test ./...
	go vet ./...

fmt:
	gofmt -w -s ./
	gofumports -w  -local github.com/rlakhtakia/opentelemetry-go-extra ./

go_mod_tidy:
	set -e; for dir in $(PACKAGE_DIRS); do \
	  echo "go mod tidy in $${dir}"; \
	  (cd "$${dir}" && go clean -cache && go clean -modcache && go get -u && go mod tidy -go=1.19); \
	done

go_mod_replace:
	set -e; for dir in $(PACKAGE_DIRS); do \
	  echo "replace github.com/uptrace/opentelemetry-go-extra  => ./ in $${dir}"; \
	  (cd "$${dir}" && go mod edit -replace github.com/rlakhtakia/opentelemetry-go-extra=./ && go mod edit -replace github.com/uptrace/opentelemetry-go-extra=./); \
	done

go_mod_all_with_vendor: go_mod_replace go_get_dependencies go_mod_tidy go_mod_download go_mod_vendor

go_mod_all: go_mod_replace go_get_dependencies go_mod_tidy go_mod_download

go_mod_download:
	set -e; for dir in $(PACKAGE_DIRS); do \
	  echo "go mod download in $${dir}"; \
	  (cd "$${dir}" && go mod download); \
	done


go_mod_rm:
	set -e; for dir in $(PACKAGE_DIRS); do \
	  echo "rm -rf vendor in $${dir}"; \
	  (cd "$${dir}" && rm -rf vendor); \
	done

go_mod_vendor:
	set -e; for dir in $(PACKAGE_DIRS); do \
	  echo "go mod vendor in $${dir}"; \
	  (cd "$${dir}" && go mod vendor); \
	done

go_test:
	set -e; for dir in $(PACKAGE_DIRS); do \
	  echo "go build in $${dir}"; \
	  (cd "$${dir}" && go build); \
	done

go_get_dependencies:
	set -e; for dir in $(PACKAGE_DIRS); do \
	  echo "go get -d ./... in $${dir}"; \
	  (cd "$${dir}" && go get -d ./...); \
	done
