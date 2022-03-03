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
