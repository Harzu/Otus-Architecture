DOCKER_NAME="fanyshu/otus-cache-hw"
VERSION="v0.7"

all: check-need-bin clean generate build-docker push-docker

check-need-bin:
	. ./scripts/check-bin

install-dep: check-need-bin
	@helm repo add bitnami https://charts.bitnami.com/bitnami
	@helm dependency update ./deployment

start-dev: check-need-bin
	@minikube start --cpus=2 --memory=6g
	@minikube addons enable ingress
	@kubectl delete validatingwebhookconfiguration ingress-nginx-admission
	@helm install myapp ./deployment
	@watch kubectl get all

run-dev:
	@helm install myapp ./deployment

stop-dev:
	@helm delete myapp

shutdown: stop-dev
	@minikube stop
	@minikube delete

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