# Название образа
IMAGE_NAME = service-music

# Тег образа
IMAGE_TAG = latest

# Переменные для kubectl
NAMESPACE = default

.PHONY: all build docker-build deploy status clean

all: build

# Сборка Go-приложения
build:
	@echo "Building Go application..."
	go build -o bin/$(IMAGE_NAME) cmd/service.music/main.go

# Сборка Docker-образа
docker-build:
	@echo "Building Docker image..."
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

# Деплой в Kubernetes
deploy:
	@echo "Deploying to Kubernetes..."
	kubectl apply -f configs/kubernetes/postgres.yaml --validate=false
	kubectl apply -f configs/kubernetes/redis.yaml --validate=false
	kubectl apply -f configs/kubernetes/deployment.yaml --validate=false

# Получение статуса подов и сервисов
status:
	@echo "Getting Kubernetes status..."
	kubectl get pods
	kubectl get services

# Удаление всех ресурсов
clean:
	@echo "Deleting Kubernetes resources..."
	kubectl delete -f configs/kubernetes/postgres.yaml
	kubectl delete -f configs/kubernetes/redis.yaml
	kubectl delete -f configs/kubernetes/deployment.yaml
