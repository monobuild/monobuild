.PHONY: test

test: ## Run all the tests
	rm -f coverage.tmp && rm -f coverage.txt
	echo 'mode: atomic' > coverage.txt && go list ./... | xargs -n1 -I{} sh -c 'go test -race -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp

cover: test ## Run all the tests and opens the coverage report
	go tool cover -html=coverage.txt

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

fmt-run: ## run gofmt all go files without changing
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -l -s "$$file"; done

test-file:
	make build; mv .build/git-hook-commit ~/go/bin; git-hook-commit test/commit-message

snapshot: ## Create snapshot build
	goreleaser --skip=publish --clean --snapshot

release: ## Create release build
	goreleaser --clean

build: ## build binary to .build folder
	-rm -f .build/monobuild
	go build -o ".build/monobuild" cmd/monobuild/main.go

install: build ## install binary to $GOPATH/bin
	cp .build/monobuild ${GOPATH}/bin/monobuild

# Self-Documented Makefile see https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
