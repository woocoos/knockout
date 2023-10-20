version := $(shell /bin/date "+%Y-%m-%d %H:%M")
# must be [adminx, auth, files, standalone]
BUILD_NAME=adminx

build:
	go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME)/$(BUILD_NAME) ./cmd/$(BUILD_NAME)/main.go
mac:
	GOOS=darwin go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME)-darwin ./cmd/$(BUILD_NAME)/main.go
	$(if $(shell command -v upx), upx $(BUILD_NAME)-darwin)
win:
	GOOS=windows go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME)/$(BUILD_NAME).exe ./cmd/$(BUILD_NAME)/main.go
	$(if $(shell command -v upx), upx $(BUILD_NAME).exe)
linux:
	GOOS=linux go build -ldflags="-s -w" -ldflags="-X 'main.BuildTime=$(version)'" -o ./cmd/$(BUILD_NAME)/$(BUILD_NAME)-linux ./cmd/$(BUILD_NAME)/main.go
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

.PHONY: db db-init db-base db-apppolicy
db: db-init db-base cli-gql-actions cli-gql-menu db-apppolicy
db-init:
	GOWORK=off go run -mod=mod script/initdb.go
db-base:
	GOWORK=off go run -mod=mod script/initdata.go
db-apppolicy:
	GOWORK=off go run -mod=mod script/initapppolicy.go

.PHONY: gen genent gengql genoas
gen: genent gengql
genent:
	go run codegen/entgen/entc.go
gengql:
	go run codegen/gqlgen/gqlgen.go
genoas:
	# go run codegen/oasgen/oasgen.go
	# go install github.com/tsingsun/woocoo/cmd/woco@main
	woco oasgen -c ./codegen/oasgen/config.yaml

.PHONY: cli-gql-actions
cli-gql-actions:
	kocli res gql-action -a resource -g codegen/gqlgen/gqlgen.yaml -f codegen/knockout.yaml
cli-gql-menu:
	# todo move adminx-ui to project
	echo "kocli res menu -a resource -d {adminui}/src/components/layout/menu.json -f codegen/knockout.yaml"