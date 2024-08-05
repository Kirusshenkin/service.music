include .env
export $(shell sed 's/=.*//' .env)

.PHONY: build push deploy

# Переменные для Docker
GO_SERVER_IMAGE = $(DOCKER_USER)/go-server:latest
PYTHON_APP_IMAGE = $(DOCKER_USER)/python-app:latest

# Переменные для Kubernetes
K8S_DIR = k8s

# Команды для сборки Docker-образов
build:
	docker build -t $(GO_SERVER_IMAGE) -f docker/Dockerfile.go .
	docker build -t $(PYTHON_APP_IMAGE) -f docker/Dockerfile.python .

# Команды для пуша Docker-образов в Docker Registry
push:
	docker push $(GO_SERVER_IMAGE)
	docker push $(PYTHON_APP_IMAGE)

# Команда для деплоя в Kubernetes
deploy: build push
	kubectl apply -f $(K8S_DIR)

# Команда для удаления всех ресурсов в Kubernetes (на случай отката)
clean:
	kubectl delete -f $(K8S_DIR)
