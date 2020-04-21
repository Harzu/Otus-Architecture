# Homework run

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
