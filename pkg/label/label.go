package label

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/kolikons/label-watch/cmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type patchStringValue struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

var updateErr error

// Kubeconfig flag of path to kube config
var Kubeconfig *string

// APISet alias to connection k8s
var APISet func() v1.CoreV1Interface

func getClient(p string) (*kubernetes.Clientset, error) {
	var c *rest.Config
	var e error
	// in cluster
	c, e = rest.InClusterConfig()
	if e != nil {
		// out cluster
		c, e = clientcmd.BuildConfigFromFlags("", p)
		if e != nil {
			return nil, e
		}
	}
	return kubernetes.NewForConfig(c)
}

// APICore connect to k8s
func APICore() (api v1.CoreV1Interface) {
	config, err := getClient(*Kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset := *config

	api = clientset.CoreV1()
	return api
}

// PatchNodeLabel add label to node-role.kubernetes.io/$STRING=true
func PatchNodeLabel(s, node string) (bool, error) {
	var payload []patchStringValue
	var pathLabel string
	// ~1 is character '/'
	pathLabel = "/metadata/labels/node-role.kubernetes.io~1" + s

	payload = []patchStringValue{{
		Op:    "add",
		Path:  pathLabel,
		Value: "true",
	}}

	payloadBytes, _ := json.Marshal(payload)
	_, updateErr = APISet().Nodes().Patch(context.TODO(), node, types.JSONPatchType, payloadBytes, metav1.PatchOptions{})
	if updateErr == nil {
		return true, nil
	}

	return false, updateErr
}

// RunLabel function to updated k8s object
func RunLabel(c *cmd.Command) {

	// covert string to []string
	labels := regexp.MustCompile(` *, *`).Split(c.Label, -1)
	Kubeconfig = &c.Kubeconfig
	// set poiner to ApiCore function
	APISet = APICore

	// get nodes without master
	nodes, err := APISet().Nodes().List(context.TODO(), metav1.ListOptions{LabelSelector: "!node-role.kubernetes.io/master"})
	// check error
	if err != nil {
		panic(err.Error())
	}

	var nodePatched bool
	// loop throught all nodes items
	for _, node := range nodes.Items {

		// get all labels for each node
		labelsMap := node.GetLabels()

		// start loop with labels what recived from cli
		for _, l := range labels {

			// started find if label what recived present in nodeLabels
			for k, v := range labelsMap {
				if l == k {
					// write key from label in cli
					nodePatched, updateErr = PatchNodeLabel(v, node.GetName())
					if updateErr == nil {
						fmt.Printf("Node %s has been labelled successfully. Label: node-role.kubernetes.io/%s=true \n", node.GetName(), v)
					}
				}

			}

			if nodePatched == false {
				fmt.Printf("Node %s wasn't patched because missed Label: %s\n", node.GetName(), l)
			}
		}
	}
}
