#!make

POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/walletsuro?sslmode=disable'
MIGRATE_PLATFORM=darwin
MIGRATE_VERSION=v4.14.1
MOCKGEN_VERSION=1.6.0
MOCKGEN_BIN=mock_${MOCKGEN_VERSION}_${MIGRATE_PLATFORM}_amd64

# INSTALL TOOLS

install-migrate:
	curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.${MIGRATE_PLATFORM}-amd64.tar.gz | tar xvz -C bin
	mv bin/migrate.${MIGRATE_PLATFORM}-amd64 bin/migrate

install-mockgen:
	curl -L https://github.com/golang/mock/releases/download/v${MOCKGEN_VERSION}/${MOCKGEN_BIN}.tar.gz | tar xvz -C bin
	mv bin/${MOCKGEN_BIN}/mockgen bin/mockgen
	rm -rf bin/${MOCKGEN_BIN}

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down

migrate-create:
	migrate create -ext sql -dir migrations -seq create_event_table

# SWAGGER
gen-server:
	bin/swagger generate server -f ./swagger.json -t internal/generated -A walletsuro \
		--main-package ../../../cmd/web

#
test:
	go test -race ./...

