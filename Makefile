PROJECTNAME=gotomation

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

build:
	@go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go || exit

production:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go

start:
	@go build -o $(GOBIN)/$(PROJECTNAME) ./cmd/$(PROJECTNAME)/main.go || exit
	@./bin/$(PROJECTNAME)
