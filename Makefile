PROJECT_NAME=silver-framework
GIT_PATH=github.com/wuleying/$(PROJECT_NAME)

CUR_DIR=$(CURDIR)
BIN_DIR=$(CUR_DIR)/bin

BRANCH=`git rev-parse --abbrev-ref HEAD`
SHA1=`git rev-parse --short HEAD`
CUR_DATE=`date "+%Y%m%d"`
CUR_TIME=`date "+%Y/%m/%d %H:%M:%S"`

VERSION=$(BRANCH).$(CUR_DATE).$(SHA1)

LDFLAGS=-ldflags "-X \"$(GIT_PATH)/version.Version=$(VERSION)\""

# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_TEST=$(GO_CMD) test
GO_DOC=$(GO_CMD) doc
GO_GET=$(GO_CMD) get
GO_FMT=$(GO_CMD) fmt
GO_IMPORTS=goimports

# Tools
default: fmt build build-cli

dev: stop default run

build:
	$(GO_BUILD) $(LDFLAGS) -o $(BIN_DIR)/$(PROJECT_NAME) -v $(CUR_DIR)/*.go
	@echo "$(CUR_TIME) [INFO ] Build $(PROJECT_NAME) completed"

build-cli:
	$(GO_BUILD) $(LDFLAGS) -o $(BIN_DIR)/$(PROJECT_NAME)-cli -v $(CUR_DIR)/cli/*.go
	@echo "$(CUR_TIME) [INFO ] Build $(PROJECT_NAME)-cli completed"

clean:
	$(GO_CLEAN)
	rm $(BIN_DIR)/$(PROJECT_NAME)
	rm $(BIN_DIR)/$(PROJECT_NAME)-cli
	@echo "$(CUR_TIME) [INFO ] Clean completed"

fmt:
	$(GO_FMT) .
	@echo "$(CUR_TIME) [INFO ] Go fmt completed"

imports:
	$(GO_IMPORTS) $(shell find . -name "*.go" | egrep -v "vendor")
	@echo "$(CUR_TIME) [INFO ] Go imports completed"

ps:
	ps -ef | grep $(PROJECT_NAME)

run:
	@echo $(CUR_TIME) [INFO ] CUR_DIR=$(CUR_DIR)
	@echo $(CUR_TIME) [INFO ] BIN_DIR=$(BIN_DIR)
	@echo $(CUR_TIME) [INFO ] VERSION=$(VERSION)
	$(BIN_DIR)/$(PROJECT_NAME)

stop:
	pgrep -f $(PROJECT_NAME) | xargs kill -9

# Test tools
test:
	$(GO_TEST) -v ./utils

cover:
	$(GO_TEST) -cover ./utils

bench:
	$(GO_TEST) -bench=. ./utils

# Go docs
doc:
	$(GO_DOC) ./utils

# Check tools
check: vet lint gocyclo gosimple unused staticcheck ineffassign misspell

vet:
	@echo go vet
	@if test -n '$(shell go vet `glide nv` 2>&1)'; then \
		echo '$(shell go vet `glide nv` 2>&1)'; \
	fi

lint:
	@echo golint
	@if test -n '$(shell golint `glide nv` 2>&1)'; then \
		echo '$(shell golint `glide nv` 2>&1)'; \
	fi

gocyclo:
	gocyclo -over 20 $(shell find . -name "*.go" | egrep -v "vendor")

gosimple:
	gosimple $(shell glide nv)

unused:
	unused $(shell glide nv)

staticcheck:
	staticcheck $(shell glide nv)

ineffassign:
	@for f in `find . -type d -depth 1 | egrep -v "git|hook|vendor"`; do \
		ineffassign $$f; \
	done

misspell:
	misspell -i "unknwon" $(shell find . -maxdepth 1 -mindepth 1 -type d | egrep -v "vendor|doc|bin|.git|.idea")

goconst:
	goconst $(shell glide nv)

# Docker tools
docker:
	docker build -t $(PROJECT_NAME) .

docker_run:
	docker run -p 10088:10088 $(PROJECT_NAME)

docker_stop:
	docker stop $(PROJECT_NAME)

# Glide tools
glide:
	glide install
	glide update

# Get tools and third packages
get:
	$(GO_GET) github.com/Masterminds/glide
	$(GO_GET) github.com/gordonklaus/ineffassign
	$(GO_GET) github.com/fzipp/gocyclo
	$(GO_GET) github.com/golang/lint/golint
	$(GO_GET) github.com/pierrre/gotestcover
	$(GO_GET) github.com/smartystreets/goconvey
	$(GO_GET) github.com/client9/misspell/cmd/misspell
	$(GO_GET) github.com/jgautheron/goconst/cmd/goconst
	$(GO_GET) github.com/wgliang/goappmonitor
	$(GO_GET) github.com/petermattis/goid
	$(GO_GET) github.com/bwmarrin/snowflake
	$(GO_GET) github.com/go-redis/redis
	$(GO_GET) github.com/samuel/go-zookeeper/zk
	$(GO_GET) github.com/rcrowley/go-metrics
	$(GO_GET) github.com/vrischmann/go-metrics-influxdb
	$(GO_GET) honnef.co/go/tools/cmd/staticcheck
	$(GO_GET) honnef.co/go/tools/cmd/gosimple
	$(GO_GET) honnef.co/go/tools/cmd/unused
