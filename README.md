# vault-cert-operator
The Vault-Cert-Operator is a Kubernetes Operator designed to facilitate the retrieval of domain certificates from the Vault Key-Value (KV) engine and subsequently generate Kubernetes secrets. This operator streamlines the process of managing SSL/TLS certificates for domain names within a Kubernetes environment.

# init vault-cert-operator project

```
go mod init svc.ink/m/v2
operator-sdk init --plugins go/v4 --domain svc.ink --owner "Haitao Pan"
```
