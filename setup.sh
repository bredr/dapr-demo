#!/bin/bash

minikube start --cpus=4 --memory=4096 --kubernetes-version=1.16.2 --extra-config=apiserver.authorization-mode=RBAC
minikube addons enable dashboard
minikube addons enable ingress
dapr init -k

helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install redis bitnami/redis
