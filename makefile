version := $(shell /bin/date "+%Y-%m-%d %H:%M")
# must be [adminx, auth, files, standalone]
BUILD_NAME=knockout

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

new-ent:
	GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent --target codegen/entgen/schema new $(NAME)
migrate-init:
	GOWORK=off go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration ./ent/schema
migrate:
	GOWORK=off go run -mod=mod ent/migrate/main.go -dsn="$(DSN)" -name=$(NAME)
migrate-lint:
	atlas migrate lint --dev-url="$(DSN)" --dir="file://ent/migrate/migrations" --latest=$(LATEST)
migrate-apply:
	atlas migrate apply --dev-url="$(DSN)" --dir="file://ent/migrate/migrations" --latest=$(LATEST)

.PHONY: db
db: initialize-schema initialize-basedata initialize-qql-action initialize-qql-menu initialize-app-policy
initialize-schema:
	GOWORK=off go run -mod=mod script/initschema.go
initialize-basedata:
	GOWORK=off go run -mod=mod script/initdata.go
initialize-qql-action:
	kocli res gql-action -a resource -g codegen/gqlgen/gqlgen.yaml -f codegen/knockout.yaml
initialize-qql-menu:
	# todo move adminx-ui to project
	echo "kocli res menu -a resource -d {adminui}/src/components/layout/menu.json -f codegen/knockout.yaml"
initialize-app-policy:
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
	woco oasgen -c ./codegen/oasgen/auth-config.yaml
