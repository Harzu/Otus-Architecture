DOCKER_NAME="fanyshu/otus-auth-api:v2"

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

build-docker:
	@docker build -t $(DOCKER_NAME) .

push-docker:
	@docker push $(DOCKER_NAME)