version ?= latest
cov = .coverage.out

pkg ?= ./...
testtimeout ?= 30s
testflag ?= -race -timeout $(testtimeout) -coverprofile=$(cov) $(flag)
gotest = go test -failfast $(pkg) $(testflag) $(if $(testcase),-run "$(testcase)")

ldflags = -w -s -X main.version=${version}

all: static-analysis test

.PHONY: help
help: ## display this help
	@ echo "Please use 'make <target>' where <target> is one of:"
	@ echo
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-16s\033[0m - %s\n", $$1, $$2}'
	@ echo

build: ## Build the binaries
	@go build -v -ldflags "$(ldflags)" -o ./cmd/gbanalytics/gbanalytics ./cmd/gbanalytics

.PHONY: test
test: ## Run unit tests, set testcase=<testcase> or flag=-v if you need them
	$(gotest)

.PHONY: fmt
fmt: ## run gofmt
	gofmt -w -s -l .

.PHONY: static-analysis
static-analysis: fmt ## run gofmt
