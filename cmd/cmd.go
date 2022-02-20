package cmd

import (
	"errors"
	"flag"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

type Command struct {
	Verbose    bool
	Label      string
	Kubeconfig string
	Interval   string
	Args       *[]string
}

// opts is creating flag for cli
func opts() (c *Command, e error) {
	c = new(Command)
	// Detailed output info
	flag.BoolVar(
		&c.Verbose,
		"v",
		false,
		"Makes verbose output",
	)

	// Interval if set the run util like daemon.
	flag.StringVar(
		&c.Interval,
		"interval",
		"",
		"(optional) Start application in deamon mode. Supports format: 's', 'm', 'h'.",
	)

	// Uses for finding label and conver to node-role. It's requered
	flag.StringVar(
		&c.Label,
		"label",
		"",
		`Label that's checking on worker nodes then set label in format node-role.kubernetes.io/VALUE_FROM_LABEL=true.
Supports multiple labels: -label node-type,type,etc
Example: 
$ kubectl get node NODE -o jsonpath='{.metadata.labels}' | jq
{
	"beta.kubernetes.io/arch": "amd64",
	....
	"node-type": "worker"
}
$ label-watch -label node-type
$ kubectl get node NODE -o jsonpath='{.metadata.labels}' | jq
{
	"beta.kubernetes.io/arch": "amd64",
	....
	"node-type": "worker",
	"node-role.kubernetes.io/worker": "true"
}`,
	)
	//Flag for kubeconfig if u run from your workstation
	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(
			&c.Kubeconfig,
			"kubeconfig",
			filepath.Join(home, ".kube", "config"),
			"(optional) absolute path to the kubeconfig file",
		)
	} else {
		flag.StringVar(
			&c.Kubeconfig,
			"kubeconfig",
			"",
			"absolute path to the kubeconfig file",
		)
	}

	flag.Parse()
	a := flag.Args()
	c.Args = &a

	return c, e

}

// ParseFlags is
func ParseFlags() (c *Command, e error) {
	c, e = opts()
	if c.Label == "" {
		e = errors.New("flag is empty: run with -h or --help")
	}

	return c, e
}
