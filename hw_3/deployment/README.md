# Homework run

# Kubectl

In first. go to **./kube** directory

Run all services
```bash
kubectl apply -f .
```

wait to start postgres database pod and run migration job
```bash
kubectl apply -f ./jobs
```

or apply command to run all pods
```bash
kubectl apply -f . && kubectl apply -f ./jobs/
``` 

See ingress ip address
```bash
kubectl get ingress
```

add ingress ip and 80 port to postman environment **KUBE_HOST**
Run test collection

# Helm

go to **./service-chart** directory

Update dependencies
```bash
helm dependency update .
```

Dry run for check
```bash
helm install myapp . --dry-run
```

if not errors run app with command
```bash
helm install myapp .
```

For close app use command
```bash
helm uninstall myapp
```