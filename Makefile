BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOINSTALL=$(GO) install
GOBUILD=$(GO) build $(FLAGS)
GOCLEAN=$(GO) clean
GOGET=$(GO) get
GOCMD=GOPATH
PROJECT =neo_data
EXENAME=neo_data
MAINFILE=main

my_info:
	echo "Name : Marcio Venancio Batista"

makedir:
	@if [ ! -d $(BUILDPATH)/bin ] ; then mkdir -p $(BUILDPATH)/bin ; fi
	@if [ ! -d $(BUILDPATH)/pkg ] ; then mkdir -p $(BUILDPATH)/pkg ; fi

get:
	@$(GOGET) gopkg.in/mgo.v2/bson
	@$(GOGET) github.com/gorilla/mux
	@$(GOGET) github.com/stretchr/testify/assert

build:
	@echo "building..."
	@$(GOBUILD) -o bin/$(EXENAME) $(MAINFILE).go
	@$(GOINSTALL)
	@echo "build done..."

clean:
	@echo "cleaning..."
	@rm -rf $(BUILDPATH)/bin/$(EXENAME)
	@rm -rf $(BUILDPATH)/pkg
	@echo "done..."

all: makedir get build

run:
	@$(GO) run $(BUILDPATH)/main.go
