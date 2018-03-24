PROJECT_NAME=silver-framework
CUR_DIR=$(CURDIR)
BIN_DIR=$(CUR_DIR)/bin

# Go parameters
GO_CMD=godep go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
GO_FMT=$(GO_CMD) fmt
GO_VET=$(GO_CMD) tool vet
GO_LINT=golint

# Current time
CUR_TIME=`date "+%Y/%m/%d %H:%M:%S"`

default: build

build:
	$(GO_BUILD) -o $(BIN_DIR)/$(PROJECT_NAME) -v $(CUR_DIR)/*.go
	@echo "$(CUR_TIME) [INFO ] Build completed"

clean:
	$(GO_CLEAN)
	rm $(BIN_DIR)/$(PROJECT_NAME)
	@echo "$(CUR_TIME) [INFO ] Clean completed"

fmt:
	$(GO_FMT) .
	@echo "$(CUR_TIME) [INFO ] Go fmt completed"

godep:
	godep save
	@echo "$(CUR_TIME) [INFO ] Godep saved"

ps:
	ps -ef | grep $(PROJECT_NAME)

run:
	@echo $(CUR_TIME) [INFO ] CUR_DIR=\"$(CUR_DIR)\"
	@echo $(CUR_TIME) [INFO ] BIN_DIR=\"$(BIN_DIR)\"
	$(BIN_DIR)/$(PROJECT_NAME)

stop:
	pgrep -f $(PROJECT_NAME) | xargs kill -9

# Test tools
test:
	$(GO_TEST)

cover:
	$(GO_TEST) -cover

# Check tools
check: vet lint

vet:
	$(GO_VET) $(shell find . -name "*.go" | egrep -v "vendor")
	@echo "$(CUR_TIME) [INFO ] Vet checked\n"

lint:
	@for f in `find . -type d -depth 1 | egrep -v "vendor"`; do \
		$(GO_LINT) $$f; \
	done
	@echo "$(CUR_TIME) [INFO ] Lint checked\n"