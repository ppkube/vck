/*
<insert-license-here>
*/package v1alpha1

import (
	v1alpha1 "github.com/ppkube/vck/pkg/apis/vck/v1alpha1"
	scheme "github.com/ppkube/vck/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumeManagersGetter has a method to return a VolumeManagerInterface.
// A group's client should implement this interface.
type VolumeManagersGetter interface {
	VolumeManagers(namespace string) VolumeManagerInterface
}

// VolumeManagerInterface has methods to work with VolumeManager resources.
type VolumeManagerInterface interface {
	Create(*v1alpha1.VolumeManager) (*v1alpha1.VolumeManager, error)
	Update(*v1alpha1.VolumeManager) (*v1alpha1.VolumeManager, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VolumeManager, error)
	List(opts v1.ListOptions) (*v1alpha1.VolumeManagerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VolumeManager, err error)
	VolumeManagerExpansion
}

// volumeManagers implements VolumeManagerInterface
type volumeManagers struct {
	client rest.Interface
	ns     string
}

// newVolumeManagers returns a VolumeManagers
func newVolumeManagers(c *VckV1alpha1Client, namespace string) *volumeManagers {
	return &volumeManagers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the volumeManager, and returns the corresponding volumeManager object, and an error if there is any.
func (c *volumeManagers) Get(name string, options v1.GetOptions) (result *v1alpha1.VolumeManager, err error) {
	result = &v1alpha1.VolumeManager{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumemanagers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VolumeManagers that match those selectors.
func (c *volumeManagers) List(opts v1.ListOptions) (result *v1alpha1.VolumeManagerList, err error) {
	result = &v1alpha1.VolumeManagerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumeManagers.
func (c *volumeManagers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("volumemanagers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a volumeManager and creates it.  Returns the server's representation of the volumeManager, and an error, if there is any.
func (c *volumeManagers) Create(volumeManager *v1alpha1.VolumeManager) (result *v1alpha1.VolumeManager, err error) {
	result = &v1alpha1.VolumeManager{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("volumemanagers").
		Body(volumeManager).
		Do().
		Into(result)
	return
}

// Update takes the representation of a volumeManager and updates it. Returns the server's representation of the volumeManager, and an error, if there is any.
func (c *volumeManagers) Update(volumeManager *v1alpha1.VolumeManager) (result *v1alpha1.VolumeManager, err error) {
	result = &v1alpha1.VolumeManager{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumemanagers").
		Name(volumeManager.Name).
		Body(volumeManager).
		Do().
		Into(result)
	return
}

// Delete takes name of the volumeManager and deletes it. Returns an error if one occurs.
func (c *volumeManagers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumemanagers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumeManagers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumemanagers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched volumeManager.
func (c *volumeManagers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VolumeManager, err error) {
	result = &v1alpha1.VolumeManager{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("volumemanagers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
