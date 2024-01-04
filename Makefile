all:
	operator-sdk build your-namespace/vault-cert-operator
	docker push your-namespace/vault-cert-operator
deploy:
	operator-sdk run --local --watch-namespace=<your-namespace>

