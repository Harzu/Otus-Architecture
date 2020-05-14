clean:
	@helm uninstall myapp
	@helm uninstall myapp-api

auth-hw-run: auth-hw-myapp auth-hw-myapp-api

auth-hw-myapp:
	@helm install myapp ./app/deployment

auth-hw-myapp-api:
	@helm install myapp-api ./auth/deployment