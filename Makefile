GOLANGCI_LINT_VERSION = "v2.5.0"
GOTESTSUM_VERSION = "v1.11.0"

.PHONY: golangci-lint-version
golangci-lint-version:
	@echo $(GOLANGCI_LINT_VERSION)

.PHONY: test
test:
	go run gotest.tools/gotestsum@$(GOTESTSUM_VERSION) --format testname -- -coverprofile=cover.out ./...
.PHONY: lint
lint:
	@go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run ./...
