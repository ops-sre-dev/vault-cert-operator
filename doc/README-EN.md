Vault-Cert-Operator Overview:

The Vault-Cert-Operator is a Kubernetes Operator designed to facilitate the retrieval of domain certificates from the Vault Key-Value (KV) engine and subsequently generate Kubernetes secrets. This operator streamlines the process of managing SSL/TLS certificates for domain names within a Kubernetes environment.

Features:

Vault Integration: Connects to a HashiCorp Vault instance to securely retrieve SSL/TLS certificates.

Kubernetes Secret Generation: Automatically generates Kubernetes secrets containing the fetched SSL/TLS certificates.

CRD Support: Utilizes a Custom Resource Definition (CRD) named SecretCertificate to define the desired state, including Vault server details, token, secret path, and Kubernetes secret information.

Automation: Enables automation of the entire process, reducing the manual effort required to manage SSL/TLS certificates in Kubernetes.

Usage:

Define a SecretCertificate Custom Resource (CR) specifying Vault server details, token, secret path, and desired Kubernetes secret information.

The Vault-Cert-Operator, upon detecting the CR, connects to the specified Vault instance, retrieves SSL/TLS certificates from the provided secret path, and generates a Kubernetes secret with the obtained certificates.

Kubernetes applications can then use the generated secret for secure communication with SSL/TLS.
