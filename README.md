# microservice-architecture-02
Kubernetes fundamentals

Deployment

Create namespace
    kubectl create namespace crudapp
    kubectl config set-context --current --namespace=crudapp

Run for base version:
    kubectl -f apply ./manifests_base


Test:
    newman run ./crud-service.postman_collection.json