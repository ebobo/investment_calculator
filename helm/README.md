# helm install (kubectl get pods / kubectl get all for check deployment)

# run this in project root dir

helm install investment-service helm

# helm upgrade to a new version

helm upgrade investment-service helm

# helm rollback

helm rollback investment-service 1

# helm history (get revision number)

helm history investment-service

# helm uninstall

helm uninstall investment-service

# how to use local docker image with Minikube

Set the environment variables with (check with "dcoker images")
eval $(minikube docker-env)
Build the image with the Docker daemon of Minikube (eg docker build -t my-image .)

Set the imagePullPolicy to Never, otherwise Kubernetes will try to download the image.

Important note: You have to run eval $(minikube docker-env) on each terminal you want to use, since it only sets the environment variables for the current shell session.
