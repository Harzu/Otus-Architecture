# Otus-cache homework

### Install and run
you need make and another binaries for run this project for check use **make check-need-bin**

For install use **make install-dep**
For start project  in minikube use **make start-dev**
For stop project use **make stop-dev**
For run before make stop-dev use **make run-dev**
For shutdown project use **make shutdown**

INGRESS HOSTNAME - **arch.homework**

### Tests
For run tests use postman/newman(cli)
postman collections ../postman-collections/cache-hw.json

example newman
```bash
newman run ../postman-collections/cache-hw.json
```