PROJECT_NAME=silver-framework

CUR_DIR=$(CURDIR)
BIN_DIR=$(CUR_DIR)/bin

# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_GET=$(GO_CMD) get
GO_FMT=$(GO_CMD) fmt

# Current time
CUR_TIME=`date "+%Y/%m/%d %H:%M:%S"`

# Tools
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

imports:
	$(GO_IMPORTS) $(shell find . -name "*.go" | egrep -v "vendor")
	@echo "$(CUR_TIME) [INFO ] Go imports completed"

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
check: vet lint gocyclo gosimple

vet:
	@echo "$(CUR_TIME) [INFO ] Check: vet begin"
	@if test -n '$(shell go vet `glide nv` 2>&1)'; then \
		echo '$(shell go vet `glide nv` 2>&1)'; \
	fi
	@echo "$(CUR_TIME) [INFO ] Check: vet checked\n"

lint:
	@echo "$(CUR_TIME) [INFO ] Check: lint begin"
	@if test -n '$(shell golint `glide nv` 2>&1)'; then \
		echo '$(shell golint `glide nv` 2>&1)'; \
	fi
	@echo "$(CUR_TIME) [INFO ] Check: lint checked\n"

gocyclo:
	@echo "$(CUR_TIME) [INFO ] Check: gocyclo begin"
	gocyclo -over 10 $(shell find . -name "*.go" | egrep -v "vendor")
	@echo "$(CUR_TIME) [INFO ] Check: gocyclo checked\n"

gosimple:
	@echo "$(CUR_TIME) [INFO ] Check: gosimple begin"
	gosimple $(shell glide nv)
	@echo "$(CUR_TIME) [INFO ] Check: gosimple checked\n"