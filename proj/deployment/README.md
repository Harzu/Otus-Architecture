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

for checks and test change /etc/hosts (add host-name ingress to ip ingress)

For close app use command
```bash
helm uninstall myapp
```

# Prometheus
```bash
helm repo add stable https://kubernetes-charts.storage.googleapis.com
helm repo update
helm install prom stable/prometheus-operator -f ./prometheus.yml --atomic
```

# Nginx-ingress
```bash
helm install nginx stable/nginx-ingress -f ./nginx-ingress.yml --atomic
```

# Postgres exporter
```bash
helm install postgres-exporter stable/prometheus-postgres-exporter -f ./postgres-exporter.yml
```

# Grafana dashboard postgres
```bash
kubectl apply -f ./psql-grafana.yml
```

# Uninstall
```bash
kubectl delete -f ./stresstest.yml
helm uninstall postgres-exporter
helm uninstall nginx
helm uninstall myapp
helm uninstall prom
```