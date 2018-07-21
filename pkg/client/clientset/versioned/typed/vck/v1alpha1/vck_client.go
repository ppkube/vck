/*
<insert-license-here>
*/package v1alpha1

import (
	v1alpha1 "github.com/ppkube/vck/pkg/apis/vck/v1alpha1"
	"github.com/ppkube/vck/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type VckV1alpha1Interface interface {
	RESTClient() rest.Interface
	VolumeManagersGetter
}

// VckV1alpha1Client is used to interact with features provided by the vck.intelai.org group.
type VckV1alpha1Client struct {
	restClient rest.Interface
}

func (c *VckV1alpha1Client) VolumeManagers(namespace string) VolumeManagerInterface {
	return newVolumeManagers(c, namespace)
}

// NewForConfig creates a new VckV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*VckV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &VckV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new VckV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *VckV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new VckV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *VckV1alpha1Client {
	return &VckV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *VckV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
