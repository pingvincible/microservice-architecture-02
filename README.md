# microservice-architecture-02
Kubernetes fundamentals

Deployment

Create namespace:
    kubectl create namespace crudapp
    kubectl config set-context --current --namespace=crudapp

Run:
    helm install crud-service ./crud-chart

Test:
    newman run ./crud-service.postman_collection.json