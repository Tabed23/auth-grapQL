.PHONY: all build test clean gqlgen run


GQLGEN_CMD=go run github.com/99designs/gqlgen generate
GQLGEN_INIT=go run github.com/99designs/gqlgen init
GOCMD=go
GOBUILD=$(GOCMD) build


gen:
	$(GQLGEN_CMD)

run:
	$(GOCMD) run server.go

init:
	$(GQLGEN_INIT)

package:
	$(GOCMD) mod tidy