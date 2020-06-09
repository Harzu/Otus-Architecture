DOCKER_NAME="fanyshu/otus-cache-hw"
VERSION="v0.1"

clean:
	@echo "Removing generated files..."
	@rm -rf ./app/generated/models
	@rm -rf ./app/generated/restapi/operations
	@rm -rf ./app/generated/restapi/doc.go
	@rm -rf ./app/generated/restapi/embedded_spec.go
	@rm -rf ./app/generated/restapi/server.go

generate:
	@echo "Generating server..."
	@swagger generate server -f ./swagger/swagger.yml -t ./app/generated --exclude-main

build-docker: build-docker-app build-docker-migrations
build-docker-app:
	@docker build -f ./Dockerfile -t "$(DOCKER_NAME):$(VERSION)" .
build-docker-migrations:
	@docker build -f ./migrations.Dockerfile -t "$(DOCKER_NAME)-migrations:$(VERSION)" .

push-docker:
	@docker push "$(DOCKER_NAME):$(VERSION)"
	@docker push "$(DOCKER_NAME)-migrations:$(VERSION)"