.PHONY: migrate-up migrate-down gen build local-env dev set-local-env set-heroku-env build-docker run

migrate-up: build
	bin/migrate up

migrate-down: build
	bin/migrate down

gen:
	@sqlboiler --wipe psql

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/migrate ./cmd/migrate
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server

local-env:
	docker-compose -f test_postgres/docker-compose.yml down
	docker-compose -f test_postgres/docker-compose.yml up -d

dev:
	go build -o bin/server ./cmd/server
	bin/server

test:
	go test -cover -v git.d.foundation/datcom/backend/src/service

test-output:
	go test -cover -v git.d.foundation/datcom/backend/src/service -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html

build-docker:
	docker build -t datcom/backend .

run: 
	bin/server
