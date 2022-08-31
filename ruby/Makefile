DOCKER_IMAGE=ruby:3.0
RUBY_COMMAND=bundle install && clear && ruby main.rb
DOCKER_ARGS=--rm -it -v `pwd`:/app -w /app

run:
	@echo "Running ruby verify signature ..."
	@docker run $(DOCKER_ARGS) $(DOCKER_IMAGE) bash -c '$(RUBY_COMMAND)'
	