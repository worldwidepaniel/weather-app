REPO_DIR=`pwd`
.PHONY: build
build:
	@docker build -f docker/Dockerfile --build-arg "PWD=$(REPO_DIR)" -t weather-app .
.PHONY: run
run:
	@echo REPO_DIR=${REPO_DIR} > ./docker/compose/.env
	@docker-compose -f ./docker/compose/dev.yaml up
