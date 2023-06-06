version := $(shell /bin/date "+%Y-%m-%d %H:%M")
BUILD_NAME=adminx

build:
	go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME) ./cmd/main.go
mac:
	GOOS=darwin go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME)-darwin ./cmd/main.go
	$(if $(shell command -v upx), upx $(BUILD_NAME)-darwin)
win:
	GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME).exe ./cmd/main.go
	$(if $(shell command -v upx), upx $(BUILD_NAME).exe)
linux:
	GOOS=linux go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME)-linux ./cmd/main.go
	$(if $(shell command -v upx), upx $(BUILD_NAME)-linux)
ent-new:
	GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent --target codegen/entgen/schema new $(NAME)
migration-init:
	GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration ./ent/schema
migration:
	GOWORK=off go run -mod=mod ent/migrate/main.go -dsn="$(DSN)" -name=$(NAME)
migration-lint:
	atlas migrate lint --dev-url="$(DSN)" --dir="file://ent/migrate/migrations" --latest=$(LATEST)
migration-apply:
	atlas migrate apply --dev-url="$(DSN)" --dir="file://ent/migrate/migrations" --latest=$(LATEST)
test-db:
	GOWORK=off go run -mod=mod test/initdb.go
test-data:
	GOWORK=off go run -mod=mod test/testdata/initdata.go

.PHONY: gen genent gengql genoas
gen: genent gengql
genent:
	go run codegen/entgen/entc.go
gengql:
	go run codegen/gqlgen/gqlgen.go
gengqlfile:
	go run github.com/woocoos/entco/cmd/gqltools
genoas:
	# go run codegen/oasgen/oasgen.go
	# go install github.com/tsingsun/woocoo/cmd/woco
	woco oasgen -c ./codegen/oasgen/config.yaml

.PHONY: cli-gql-actions
cli-gql-actions:
	go run cmd/tools/main.go res gql-action -a resource