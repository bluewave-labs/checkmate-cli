.PHONY: format, format-check, lint, build, build-with-version, clean
.DEFAULT_GOAL := build

APP_NAME = "Checkmate CLI"
BINARY_NAME = "checkmate"
GO_MODULE_NAME = "github.com/bluewave-labs/checkmate-cli"
GO_MAIN_DIR = "./"
GO_DIST_DIR = "./dist"

# Build, CI and Release
tag = $(shell git describe --tags --abbrev=0)

format:
	@gofmt -w ./

format-check:
	@gofmt -l ./

lint:
	@echo "Linting..."
	@golangci-lint run

build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(GO_DIST_DIR)/$(BINARY_NAME) $(GO_MAIN_DIR)

# build-with-version:
# 	@echo "Building $(APP_NAME) with version info"
# 	@go build -o $(GO_DIST_DIR)/$(BINARY_NAME) \
# 		-ldflags "-X cmd.Version=$(tag)" \
# 		$(GO_MAIN_DIR)

clean:
	rm -r $(GO_DIST_DIR)
