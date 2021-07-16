#!make

POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/walletsuro?sslmode=disable'
MIGRATE_PLATFORM=darwin
MIGRATE_VERSION=v4.14.1

install-migrate:
	curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.${MIGRATE_PLATFORM}-amd64.tar.gz | tar xvz -C bin
	mv bin/migrate.${MIGRATE_PLATFORM}-amd64 bin/migrate

gen-server:
	bin/swagger generate server -f ./swagger.json -t internal/generated -A walletsuro \
		--main-package ../../../cmd/web

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path migrations down
