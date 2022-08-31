PIP_COMMAND=pip install cryptography==37.0.4
PYTHON_COMMAND=python main.py
DOCKER_IMAGE=python:3.10.6-alpine3.16
DOCKER_ARGS=--rm -it -v `pwd`:/app -w /app

run:
	@echo "Running python verify signature ..."
	@docker run $(DOCKER_ARGS) $(DOCKER_IMAGE) sh -c '$(PIP_COMMAND) && clear && $(PYTHON_COMMAND)'
