
SRC_GO = $(shell find . -not -path "*vendor*" -not -path "*node_modules*" -type f -name '*.go')

SRC_PROTOS = $(shell find apps -type f -name *.proto)
COMPILED_PROTOS := $(SRC_PROTOS:.proto=.pb.go)
.PHONY: data

build:
	actools go mod download
	sudo python infra/hosts.py

serve:
	@actools start \
		auth \
		masters \
		settings \

test:
	@actools go test ./$(pkg)

data:
	actools rm database redis
	actools start database redis
	bash -c "until actools mysql -h database -u dev-user -pdev-password -e ';' 2> /dev/null ; do sleep 1; done"

	mkdir -p tmp
	curl https://storage.googleapis.com/altipla-sql-schema/dump-adoquier.sql -o tmp/dump-adoquier.sql
	bash -c "actools mysql -h database -u root -pdev-root < tmp/dump-adoquier.sql"

	actools go install ./cmd/fill-data
	actools run go fill-data

gofmt:
	@gofmt -w $(SRC_GO)
	@gofmt -r '&a{} -> new(a)' -w $(SRC_GO)

update-deps:
	actools go get -u
	actools go mod download