DOCKER_ARGS=--rm -it -v `pwd`:/app -w /app
DOCKER_IMAGE=maven:3.8.6-jdk-11

build:
	@docker run $(DOCKER_ARGS) $(DOCKER_IMAGE) mvn clean compile package

run:
	@docker run $(DOCKER_ARGS) $(DOCKER_IMAGE) java -jar target/plug-signature-verify-1.0-SNAPSHOT-jar-with-dependencies.jar
