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
GO_IMPORTS=goimports

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
	misspell $(shell find . -maxdepth 1 -mindepth 1 -type d | egrep -v "vendor|doc|bin|.git|.idea")

goconst:
	goconst $(shell glide nv)

# Get tools and third packages
get:
	go get github.com/Masterminds/glide
	go get honnef.co/go/tools/cmd/staticcheck
	go get honnef.co/go/tools/cmd/gosimple
	go get honnef.co/go/tools/cmd/unused
	go get github.com/gordonklaus/ineffassign
	go get github.com/fzipp/gocyclo
	go get github.com/golang/lint/golint
	go get github.com/pierrre/gotestcover
	go get github.com/client9/misspell/cmd/misspell
	go get github.com/jgautheron/goconst/cmd/goconst