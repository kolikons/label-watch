[![Release](https://img.shields.io/github/v/release/kolikons/label-watch.svg)](https://github.com/kolikons/label-watch/releases/latest)
![Release](https://github.com/kolikons/label-watch/actions/workflows/release.yaml/badge.svg)
![helm release](https://github.com/kolikons/label-watch/actions/workflows/helm-release.yaml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# label-watch

When Kubernetes cluster's created, worker nodes is tagged as none.

`label-watch` checks a specific label on worker node then create an label `node-role.kubernetes.io/***`

---

## Usage of label-watch

label-watch supports two mode of running. The first one is outside kubernetes cluster and inside

#### Example label-watch outside kuberntes cluster:

1. You must have `kube config` that uses for connecting `kubectl`
2. Run command with the following flags:

```sh
$ kubectl get node
AME                 STATUS     ROLES                  AGE   VERSION
kind-control-plane   Ready      control-plane,master   39s   v1.20.2
kind-worker          NotReady   <none>                 8s    v1.20.2
kind-worker2         NotReady   <none>                 8s    v1.20.2
$ kubectl get node --show label
kind-control-plane   Ready    control-plane,master   54s   v1.20.2   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=kind-control-plane,kubernetes.io/os=linux,node-role.kubernetes.io/control-plane=,node-role.kubernetes.io/master=
kind-worker          Ready    <none>                 23s   v1.20.2   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=kind-worker,kubernetes.io/os=linux
kind-worker2         Ready    <none>                 23s   v1.20.2   beta.kubernetes.io/arch=amd64,beta.kubernetes.io/os=linux,kubernetes.io/arch=amd64,kubernetes.io/hostname=kind-worker2,kubernetes.io/os=linux
$ kubectl label node kind-worker node-type="worker"
node/kind-worker labeled
$ kubectl label node kind-worker group="infra"
node/kind-worker labeled
$ kubectl label node kind-worker type="backend"
node/kind-worker labeled
$ label-watch -kubeconfig ~/.kube/config -label node-type,group,type
Running label-watch
Node kind-worker has been labelled successfully. Label: node-role.kubernetes.io/worker=true
Node kind-worker has been labelled successfully. Label: node-role.kubernetes.io/infra=true
Node kind-worker has been labelled successfully. Label: node-role.kubernetes.io/backend=true
$ kubectl get node
NAME                 STATUS   ROLES                  AGE     VERSION
kind-control-plane   Ready    control-plane,master   3m7s    v1.20.2
kind-worker          Ready    backend,infra,worker   2m36s   v1.20.2
kind-worker2         Ready    <none>                 2m36s   v1.20.2
```

#### Example label-watch inside kubernetes cluster:

1. Modify ARGs in [scripts/deployment.yml](scripts/deployment.yml#22)
2. Deploy the kubernetes manifest from [scripts/deployment.yml](scripts/deployment.yml)

```sh
$ kubectl get node
NAME                 STATUS   ROLES                  AGE   VERSION
kind-control-plane   Ready    control-plane,master   13m   v1.20.2
kind-worker          Ready    <none>                 12m   v1.20.2
kind-worker2         Ready    <none>                 12m   v1.20.2
$ kubectl apply -f deployment.yml
deployment.apps/label-watch configured
serviceaccount/label-watch created
clusterrole.rbac.authorization.k8s.io/label-watch created
clusterrolebinding.rbac.authorization.k8s.io/label-watch created
$ kubectl get node
NAME                 STATUS   ROLES                  AGE   VERSION
kind-control-plane   Ready    control-plane,master   14m   v1.20.2
kind-worker          Ready    worker                 13m   v1.20.2
kind-worker2         Ready    <none>                 13m   v1.20.2
```

# Helm

Add kolikons repository to Helm repos:

```bash
helm repo add kolikons https://kolikons.github.io/charts/
```

Install label-watch

```bash
helm install label-watch kolikons/label-watch \
--namespace kube-system
```

---

# Docker

```sh
docker run kolikons/label-watch:latest
```

---

# Homebrew

```sh
brew install kolikons/tap/label-watch
```

---

## label-watch ARGS

```sh
label-watch --help
Usage of label-watch:
  -interval string
    	(optional) Start application in daemon mode. Supports format: 's', 'm', 'h'.
  -kubeconfig string
    	(optional) absolute path to the kubeconfig file
  -label string
    	Label that's checking on worker nodes then set label in format node-role.kubernetes.io/VALUE_FROM_LABEL=true.
    	Supports multiple labels: -label node-type,type,etc
    	Example:
    	$ kubectl get node NODE -o jsonpath='{.metadata.labels}' | jq
    	{
    		"beta.kubernetes.io/arch": "amd64",
    		....
    		"node-type": "worker"
    	}
    	$ label-watch -label node-type
    	$ kubectl get node NODE -o jsonpath='{.metadata.labels}' | jq
    	{
    		"beta.kubernetes.io/arch": "amd64",
    		....
    		"node-type": "worker",
    		"node-role.kubernetes.io/worker": "true"
    	}
  -v	Makes verbose output
```

---
