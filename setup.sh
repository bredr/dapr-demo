#!/bin/bash

minikube start --cpus=4 --memory=4096 --kubernetes-version=1.16.2 --extra-config=apiserver.authorization-mode=RBAC
minikube addons enable dashboard
minikube addons enable ingress
dapr init -k

helm repo add bitnami https://charts.bitnami.com/bitnami --force-update
helm repo update
helm upgrade --install redis bitnami/redis

kubectl create deployment zipkin --image openzipkin/zipkin --dry-run=client -o yaml | kubectl apply -f -
kubectl expose deployment zipkin --type ClusterIP --port 9411 --dry-run=client -o yaml | kubectl apply -f -
