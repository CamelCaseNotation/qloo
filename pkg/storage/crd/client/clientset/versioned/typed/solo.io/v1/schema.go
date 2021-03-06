/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	scheme "github.com/solo-io/qloo/pkg/storage/crd/client/clientset/versioned/scheme"
	v1 "github.com/solo-io/qloo/pkg/storage/crd/solo.io/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SchemasGetter has a method to return a SchemaInterface.
// A group's client should implement this interface.
type SchemasGetter interface {
	Schemas(namespace string) SchemaInterface
}

// SchemaInterface has methods to work with Schema resources.
type SchemaInterface interface {
	Create(*v1.Schema) (*v1.Schema, error)
	Update(*v1.Schema) (*v1.Schema, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Schema, error)
	List(opts meta_v1.ListOptions) (*v1.SchemaList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Schema, err error)
	SchemaExpansion
}

// schemas implements SchemaInterface
type schemas struct {
	client rest.Interface
	ns     string
}

// newSchemas returns a Schemas
func newSchemas(c *QlooV1Client, namespace string) *schemas {
	return &schemas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the schema, and returns the corresponding schema object, and an error if there is any.
func (c *schemas) Get(name string, options meta_v1.GetOptions) (result *v1.Schema, err error) {
	result = &v1.Schema{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schemas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Schemas that match those selectors.
func (c *schemas) List(opts meta_v1.ListOptions) (result *v1.SchemaList, err error) {
	result = &v1.SchemaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("schemas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested schemas.
func (c *schemas) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("schemas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a schema and creates it.  Returns the server's representation of the schema, and an error, if there is any.
func (c *schemas) Create(schema *v1.Schema) (result *v1.Schema, err error) {
	result = &v1.Schema{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("schemas").
		Body(schema).
		Do().
		Into(result)
	return
}

// Update takes the representation of a schema and updates it. Returns the server's representation of the schema, and an error, if there is any.
func (c *schemas) Update(schema *v1.Schema) (result *v1.Schema, err error) {
	result = &v1.Schema{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("schemas").
		Name(schema.Name).
		Body(schema).
		Do().
		Into(result)
	return
}

// Delete takes name of the schema and deletes it. Returns an error if one occurs.
func (c *schemas) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schemas").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *schemas) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("schemas").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched schema.
func (c *schemas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Schema, err error) {
	result = &v1.Schema{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("schemas").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
