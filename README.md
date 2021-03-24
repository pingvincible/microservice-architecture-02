# microservice-architecture-02
Kubernetes fundamentals

Deployment

Create namespace
    kubectl create namespace crudapp
    kubectl config set-context --current --namespace=crudapp

Run with skaffold:
    kubectl -f apply ./postgres.yaml
    skaffold run

Test:
    newman run ./crud-service.postman_collection.json