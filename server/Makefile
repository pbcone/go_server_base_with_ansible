# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=goServer
DEV_BINARY_NAME=goGerverDev
SSH_KEY=
HOST_IP=0.0.0.0
PID=

# all: test build
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
dev:
		${GOBUILD} -o ${DEV_BINARY_NAME} -v
		ENV=dev ./${DEV_BINARY_NAME}
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(DEV_BINARY_NAME)