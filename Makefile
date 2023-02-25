SWAG = $(shell which swag)
GOLINT = $(shell which golangci-lint)
GCI = $(shell which gci)
GOIMPORTS = $(shell which goimports)
GOFUMPT = $(shell which gofumpt)

swagger: ## Runs 'swagger' in the local build environment
	@${SWAG} init -g cmd/http/server.go -o swagger; \
    [ -e ./swagger/doc.json ] && rm -v ./swagger/doc.json; \
    [ -s ./swagger/swagger.json ] && mv -v ./swagger/swagger.json ./swagger/doc.json


lint-fix:
	@echo "Running gofumpt:"
	${GOFUMPT} -l -extra -w .
	@echo Running goimports:
	goimports -w -local "onlyCloudBackend" .
	@echo Running gci:
	${GCI} write --skip-generated -s standard,default,"prefix(onlyCloudBackend)" ./
	@echo Running golangci-lint:
	${GOLINT} run -c golangci.yml
	@echo Running staticcheck:
	staticcheck ./...

lint:
	@echo Running golangci-lint:
	${GOLINT} run -c golangci.yml
	@echo Running staticcheck:
	staticcheck ./...
