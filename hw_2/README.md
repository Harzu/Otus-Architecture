# HW 2 Minikube

## install tools
```bash
brew install minicube
brew install kubectl
brew install watch
```

## Run minikube
```bash
minikube start --cpus=6 --memory=6g
```

## Check cluster info
```bash
kubectl cluster-info
```

## Resource watch
```bash
watch kubectl get all
```

## Set Docker env (minikube)
```bash
minikube docker-env
```

## Pull container in hw_1
```bash
docker pull fanyshu/otus-arch-hw-1
```

## Create namespace
```bash
kubectl create namespace hw2app
```

## Set config
```bash
kubectl config set-context --current --namespace=hw2app
```

## Apply pod
```bash
kubectl apply -f ./pod.yml
```

## Check status
```bash
kubectl describe pod/homework-2
```

## Stop and delete pod
```bash
kubectl delete pod homework-2
```

## Enable nginx-ingress
```bash
minikube addons enable ingress
```