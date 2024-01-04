// controllers/secret_certificate_controller.go

package controllers

import (
	"context"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"github.com/hashicorp/vault/api"
	v1 "github.com/your-namespace/vault-cert-operator/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// Add creates a new SecretCertificate Controller and adds it to the Manager.
func Add(mgr manager.Manager, l logr.Logger) error {
	return add(mgr, newReconciler(mgr, l))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, l logr.Logger) reconcile.Reconciler {
	return &ReconcileSecretCertificate{
		Client:   mgr.GetClient(),
		Log:      l,
		Scheme:   mgr.GetScheme(),
		Recorder: mgr.GetEventRecorderFor("secretcertificatereconciler"),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("secretcertificate-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource SecretCertificate
	err = c.Watch(&source.Kind{Type: &v1.SecretCertificate{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileSecretCertificate{}

// ReconcileSecretCertificate reconciles a SecretCertificate object
type ReconcileSecretCertificate struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// Reconcile reads that state of the cluster for a SecretCertificate object and makes changes based on the state read
// and what is in the SecretCertificate.Spec
func (r *ReconcileSecretCertificate) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = r.Log.WithValues("secretcertificate", request.NamespacedName)

	// Fetch the SecretCertificate instance
	instance := &v1.SecretCertificate{}
	err := r.Get(ctx, request.NamespacedName, instance)
	if err != nil {
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}

	// Fetch Vault secret
	vaultServer := instance.Spec.VaultServer
	vaultToken := instance.Spec.VaultToken
	vaultSecretPath := instance.Spec.VaultSecretPath

	client, err := getVaultClient(vaultServer, vaultToken)
	if err != nil {
		return reconcile.Result{}, err
	}

	secretData, err := getVaultSecret(client, vaultSecretPath)
	if err != nil {
		return reconcile.Result{}, err
	}

	// Create or update the Kubernetes Secret
	secret := newTLSSecret(instance.Spec.SecretName, instance.Spec.Namespace, secretData)

	if err := controllerutil.SetControllerReference(instance, secret, r.Scheme); err != nil {
		return reconcile.Result{}, err
	}

	err = r.Client.CreateOrUpdateSecret(ctx, secret)
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}

// getVaultClient returns a new Vault client
func getVaultClient(server, token string) (*api.Client, error) {
	config := api.DefaultConfig()
	config.Address = server
	config.Token = token

	return api.NewClient(config)
}

// getVaultSecret retrieves data from Vault secret path
func getVaultSecret(client *api.Client, path string) (map[string]interface{}, error) {
	secret, err := client.Logical().Read(path)
	if err != nil {
		return nil, err
	}

	if secret == nil {
		return nil, fmt.Errorf("Vault secret not found at path: %s", path)
	}

	return secret.Data, nil
}

// newTLSSecret creates a new Kubernetes TLS Secret
func newTLSSecret(name, namespace string, data map[string]interface{}) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Type: corev1.SecretTypeTLS,
		Data: map[string][]byte{
			corev1.TLSCertKey:       []byte(data["certificate"].(string)),
			corev1.TLSPrivateKeyKey: []byte(data["private_key"].(string)),
		},
	}
}
