# Kubernetes-argocd-sync-forcer

## Background

- When use istio to show the pod version on kiali, and need AUTO deploy on argocd, show the error `field is immutable`, could muanually force sync from argocd ui or argocd cli with `--force`, but it is hard to find a way to trigger the auto force deploy, tried many methods from Internet, still did not work, so made this program.

## Features

- Watch the deployment failed event from argocd controller
- Once get the failed event, call argocd cli to force sync
- The auth method on k8s is in-cluster, so need set proper permission for the serivce account, sample is [here](deploy/readme.md)
- Allow to set the log level in pod env, to easy debug
- Support set log format in pod env
- Support .env, it is optional
- Could read the argocd user/password from k8s secret
- [k8s Secret sample](deploy/deployment.yaml)
- Health check api is avalible

## Deployment

- [Steps in readme.md](deploy/readme.md)
- [Deployment yaml](deploy/deployment.yaml)

## Refered

- https://github.com/kubernetes/client-go

