# Prer-un

**HOSTNAME** - arch.homework
add etc/hosts INGRESS IP = HOSTNAME

Start minikube and enable ingress
```bash
minikube start --cpus=6 --memory=6g
minikube addons enable ingress
```

Run hw pods
```bash
make auth-hw-run
```

Wait for lunch (see pods status ```watch kubectl get all```)

Run test postman collection
```bash
newman run ./postman-collection/auth-hw.json
```

Clear pods and stop minikube
```bash
make clean
minikube delete
```
