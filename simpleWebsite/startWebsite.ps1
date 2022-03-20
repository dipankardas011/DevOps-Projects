kubectl --version

kubectl apply -f manifest.yaml

kubectl get all

sleep 20

kubectl port-forward pod/simple-website 8080:80
