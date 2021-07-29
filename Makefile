#!make

POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/walletsuro?sslmode=disable'


migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down

migrate-create:
	migrate create -ext sql -dir migrations -seq create_event_table

# install tools binary: linter, mockgen, etc.
.PHONY: tools
tools:
	cd tools && go generate -tags tools

# SWAGGER
gen-server:
	bin/swagger generate server -f ./swagger.json -t internal/generated -A walletsuro \
		--main-package ../../../cmd/web

#
test:
	go test -short -race ./...

test-full:
	go test -race ./...

lint:
	bin/golangci-lint run

