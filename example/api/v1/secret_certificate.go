// api/v1/secret_certificate.go

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SecretCertificateSpec defines the desired state of SecretCertificate
type SecretCertificateSpec struct {
	VaultServer      string `json:"vaultServer"`
	VaultToken       string `json:"vaultToken"`
	VaultSecretPath  string `json:"vaultSecretPath"`
	SecretName       string `json:"secretName"`
	Namespace        string `json:"namespace"`
}

// SecretCertificateStatus defines the observed state of SecretCertificate
type SecretCertificateStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SecretCertificate is the Schema for the secretcertificates API
type SecretCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretCertificateSpec   `json:"spec,omitempty"`
	Status SecretCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretCertificateList contains a list of SecretCertificate
type SecretCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretCertificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecretCertificate{}, &SecretCertificateList{})
}

