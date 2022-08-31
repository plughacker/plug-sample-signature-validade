DOCKER_ARGS=--rm -it -v `pwd`:/app -w /app
DOCKER_IMAGE=mcr.microsoft.com/dotnet/sdk:6.0

run:
	@docker run $(DOCKER_ARGS) $(DOCKER_IMAGE) dotnet run .
