#!make

gen-server:
	bin/swagger generate server -f ./swagger.json -t internal/generated -A walletsuro \
		--main-package ../../../cmd/web

