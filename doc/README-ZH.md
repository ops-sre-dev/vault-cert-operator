Vault-Cert-Operator 概览:

Vault-Cert-Operator 是一个 Kubernetes Operator，旨在简化从 Vault Key-Value (KV) 引擎中检索域名证书并生成 Kubernetes secrets 的过程。该操作符简化了在 Kubernetes 环境中管理 SSL/TLS 证书的流程。

特点:

Vault 集成: 连接到 HashiCorp Vault 实例，安全地检索 SSL/TLS 证书。

Kubernetes Secret 生成: 自动创建包含检索到的 SSL/TLS 证书的 Kubernetes secrets。

CRD 支持: 使用自定义资源定义 (CRD)（命名为 SecretCertificate）定义所需的状态，包括 Vault 服务器详细信息、令牌、密钥路径和 Kubernetes secret 信息。

自动化: 实现整个过程的自动化，减少在 Kubernetes 中管理 SSL/TLS 证书所需的手动工作。

使用方式:

定义一个 SecretCertificate 自定义资源 (CR)，指定 Vault 服务器详细信息、令牌、密钥路径和所需的 Kubernetes secret 信息。

Vault-Cert-Operator 在检测到 CR 后，连接到指定的 Vault 实例，从提供的密钥路径检索 SSL/TLS 证书，并生成一个包含获得的证书的 Kubernetes secret。

Kubernetes 应用程序随后可以使用生成的 secret 进行 SSL/TLS 安全通信。
