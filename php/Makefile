DOCKER_ARGS=--rm -it -v `pwd`:/app -w /app

run:
	@echo "Running PHP verify signature ..."
	@docker run $(DOCKER_ARGS) composer install
	@clear
	@docker run $(DOCKER_ARGS) composer php main.php

	