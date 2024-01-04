// main.go

package main

import (
	"flag"
	"os"

	"github.com/your-namespace/vault-cert-operator/controllers"
	"github.com/your-namespace/vault-cert-operator/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	cfg := config.GetConfigOrDie()

	scheme := runtime.NewScheme()
	_ = v1.AddToScheme(scheme)

	mgr, err := manager.New(cfg, manager.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
	})
	if err != nil {
		panic(err)
	}

	if err = controllers.Add(mgr, log); err != nil {
		panic(err)
	}

	if err = (&v1.SecretCertificate{}).SetupWebhookWithManager(mgr); err != nil {
		panic(err)
	}

	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		panic(err)
	}
}
