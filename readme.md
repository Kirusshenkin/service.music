# Service Music

Service Music is a Go application deployed in Kubernetes using Minikube. This project demonstrates the deployment of a Go application with PostgreSQL and Redis as backend services.

## Prerequisites

- Docker
- Minikube
- kubectl
- Go (version 1.22 or later)

## Setup Instructions

### Step 1: Start Minikube

To start Minikube, use the following command:

```sh
make minikube-start
```

## Step 2: Build Docker Image
### To build the Docker image for the Go application, run:

```sh
make docker-build
```

### Step 3: Deploy the application
### To deploy the application and its services in Kubernetes, use:
    
```sh
make deploy
```

### Step 4: Check Status
### To check the status of pods and services, run:

```sh
make status
```

### Step 5: Clean Up Resources
### To delete all resources created by the deployment, use:

```sh
make clean
```

Miikube Management
## Start Minikube
### To start Minikube, run:

```sh
make minikube-start
```
## Stop Minikube
### To stop Minikube, run:

```sh
make minikube-stop
```
### Delete Minikube
### To delete Minikube, run:

```sh
make minikube-delete
```
### Restart Minikube
### To restart Minikube, use:

```sh
make minikube-restart
```

## Additional Information
 - Ensure Minikube is running with the Docker driver.
 - Use `kubectl config use-context minikube` to switch context to Minikube if necessary.
 - Use `eval $(minikube docker-env)` to switch Docker environment to Minikube if necessary.