.PHONY: all build test clean gqlgen run

GQLGEN_CMD=go run github.com/99designs/gqlgen generate
GQLGEN_INIT=go run github.com/99designs/gqlgen init
GOCMD=go
GOBUILD=$(GOCMD) build
GOFMT=find . -name '*.go' -exec go fmt {} \;

gen:
	$(GQLGEN_CMD)

run:
	$(GOCMD) run server.go

init:
	$(GQLGEN_INIT)

package:
	$(GOCMD) mod tidy

fmt:
	$(GOFMT)

publish:
	git add .
	git commit -m "$(MESSAGE)"
	git push origin main


db:
	docker compose up -d

downdb:
	docker compose down