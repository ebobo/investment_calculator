# use minikube for development enviorment

# install minicube

https://minikube.sigs.k8s.io/docs/start/

# basic minicube commands

minikube version (Print the version of minikube)

minikube start (Starts a local Kubernetes cluster)

minikube status (Gets the status of a local Kubernetes cluster)

minikube stop (Stops a running local Kubernetes cluster)

minikube delete --all (Delete all of the minikube clusters)

# kubectl

kubectl get pods / kubectl get all for check deployment

kubectl apply -f ic-server.yaml

kubectl apply -f record-ms.yaml

kubectl describe deployments
