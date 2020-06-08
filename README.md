# go mod
cd ~/work/github/kubernetes/staging/src/k8s.io
ls | xargs -n 1 -I{} bash -c "curl -s https://proxy.golang.org/k8s.io/{}/@v/kubernetes-1.16.10.info | jq -r '\"k8s.io/{} => k8s.io/{} \(.Version)\"'"

# sample-scheduler-framework

This repo is a sample for Kubernetes scheduler framework.

## Deploy

```shell
$ kubectl apply -f deploy/sample-scheduler.yaml
```

## Test
```shell
$ kubectl apply -f deploy/test-scheduler.yaml
```

Then watch sample-scheduler pod logs.
