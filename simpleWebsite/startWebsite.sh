#!/bin/sh

kubectl --version

kubectl apply -f manifest.yaml

kubectl get all

kubectl port-forward pod/simple-website 8080:80

