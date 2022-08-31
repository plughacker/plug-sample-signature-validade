DOCKER_IMAGE=golang:1.18-alpine
DOCKER_ARGS=--rm -it -v `pwd`:/app -w /app
GO_COMMAND=go mod tidy && go run .

run:
	@docker run $(DOCKER_ARGS) $(DOCKER_IMAGE) $(GO_COMMAND)
