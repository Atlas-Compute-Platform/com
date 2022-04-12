help:
	@printf "Atlas Make Options\n"
	@printf	"build: Builds repository\n"
	@printf "clean: Cleans repository\n"
	@printf "docker: Builds from Dockerfile\n"
build:
	@go fmt
	@go build
clean:
	@go clean
docker:build
	@docker build
