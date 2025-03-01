ifneq (,$(wildcard .env.local))
    include .env.local
    export
endif

_defaut: dev

dev: $(LOCAL_ENV_FILE)
	docker compose -f docker-compose.dev.yml up || true

clean-dev: $(LOCAL_ENV_FILE)
	docker compose -f docker-compose.dev.yml --profile "*" down -v

install:
	docker build -t pokedex-builder -f docker/Dockerfile.build .

build-assets:
	docker run -v "./:/app" -it pokedex-builder make ent && yarn install && make gen-api && yarn build

swag:
	swag init

gen-api:
	./node_modules/.bin/openapi-generator-cli generate -i ./docs/swagger.json -g typescript-fetch -o ./assets/typescript/api

.PHONY: ci-backend
ci-backend: fmt vet

.PHONY: vet
vet:
	go vet ./main.go

.PHONY: fmt
fmt:
	gofmt -w internal
	gofmt -w pkg

ent: ./ent/schema
	go run cmd/entc/main.go