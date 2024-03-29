# Deployment steps

## Step 1: Get the image

- Build 

```bash
# cd the root dir of this project, then run
sh build.sh
```

OR

- Use the image in [dockerhub](https://hub.docker.com/repository/docker/hubdockername/argocd-sync-forcer/general)


```bash
docker pull hubdockername/argocd-sync-forcer:v0.0.1
```


## Step2: Create a service account with proper permission

```bash
# service account
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ServiceAccount
metadata:
  name: argocd-sync-forcer-sa
EOF
---
# service account secret(token)
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
type: kubernetes.io/service-account-token
metadata:
  name: argocd-sync-forcer-sa-secret
  annotations:
    kubernetes.io/service-account.name: argocd-sync-forcer-sa
EOF

kubectl edit serviceaccount argocd-sync-forcer-sa 

# edit and add following yaml to the end of the serviceaccount
secrets:
- name: argocd-sync-forcer-sa-secret

---
# cluster role
cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: event-reader-clusterrole
rules:
- apiGroups: [""]
  resources: ["events"]
  verbs: ["get", "list", "watch"]
EOF

---
# cluster role binding
cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: event-reader-argocd-sync-forcer-clusterrolebinding
subjects:
- kind: ServiceAccount
  name: argocd-sync-forcer-sa
  namespace: default
roleRef:
  kind: ClusterRole
  name: event-reader-clusterrole
  apiGroup: rbac.authorization.k8s.io
EOF

```

## Step 3: Create a secret on k8s
```bash
cat << EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: argocd-sync-forcer-secret
type: Opaque
data:
  ARGO_SERVER: {{ .server | base64 | quote }}
  ARGO_USER: {{ .user | base64 | quote }}
  ARGO_PASSWORD: {{ .password | base64 | quote }}
EOF
```

## Step 4: Deploy

```bash
kubectl apply -f deploy/deployment.yaml
```