package common

import "github.com/spf13/pflag"

type Options struct {
	// Refresh interval in seconds.
	RefreshInterval int

	PromURL string

	// Path to the kube-config file.
	Kubeconfig string

	CertFile string
	CertKey  string

	// Only to be used to for testing.
	DisableAuthForTesting bool

	// Disable components health evaluation
	DisableComponentsHealth bool

	// Disable incident detection
	DisableIncidents bool

	// path to the components yaml file
	ComponentsPath string
}

// flags returns supported cli flags for the options.
func (o *Options) Flags() *pflag.FlagSet {
	fs := &pflag.FlagSet{}
	fs.IntVarP(&o.RefreshInterval, "refresh-interval", "i", o.RefreshInterval,
		"Refresh interval in seconds")
	fs.StringVarP(&o.PromURL, "prom-url", "u", o.PromURL,
		"URL of the Prometheus server")
	fs.StringVar(&o.Kubeconfig, "kubeconfig", o.Kubeconfig,
		"The path to the kubeconfig (defaults to in-cluster config)")

	fs.StringVar(&o.CertFile, "tls-cert-file", "", "The path to the server certificate")
	fs.StringVar(&o.CertKey, "tls-private-key-file", "", "The path to the server key")

	fs.BoolVar(&o.DisableAuthForTesting, "disable-auth-for-testing", o.DisableAuthForTesting,
		"Flag for testing purposes to disable auth")
	fs.BoolVar(&o.DisableComponentsHealth, "disable-components-health", o.DisableComponentsHealth,
		"Flag to disable components health evaluation based on alerts and kube-health evaluation")
	fs.BoolVar(&o.DisableIncidents, "disable-incidents", o.DisableIncidents,
		"Flag to disable incident detection and related metrics")
	fs.StringVar(&o.ComponentsPath, "components", o.ComponentsPath,
		"The path to the components yaml file - for testing purposes")
	return fs
}
