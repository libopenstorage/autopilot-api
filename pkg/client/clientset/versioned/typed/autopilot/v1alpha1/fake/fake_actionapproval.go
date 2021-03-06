/*
Copyright 2019 Openstorage.org

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

package fake

import (
	v1alpha1 "github.com/libopenstorage/autopilot-api/pkg/apis/autopilot/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeActionApprovals implements ActionApprovalInterface
type FakeActionApprovals struct {
	Fake *FakeAutopilotV1alpha1
	ns   string
}

var actionapprovalsResource = schema.GroupVersionResource{Group: "autopilot.libopenstorage.org", Version: "v1alpha1", Resource: "actionapprovals"}

var actionapprovalsKind = schema.GroupVersionKind{Group: "autopilot.libopenstorage.org", Version: "v1alpha1", Kind: "ActionApproval"}

// Get takes name of the actionApproval, and returns the corresponding actionApproval object, and an error if there is any.
func (c *FakeActionApprovals) Get(name string, options v1.GetOptions) (result *v1alpha1.ActionApproval, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(actionapprovalsResource, c.ns, name), &v1alpha1.ActionApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ActionApproval), err
}

// List takes label and field selectors, and returns the list of ActionApprovals that match those selectors.
func (c *FakeActionApprovals) List(opts v1.ListOptions) (result *v1alpha1.ActionApprovalList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(actionapprovalsResource, actionapprovalsKind, c.ns, opts), &v1alpha1.ActionApprovalList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ActionApprovalList{ListMeta: obj.(*v1alpha1.ActionApprovalList).ListMeta}
	for _, item := range obj.(*v1alpha1.ActionApprovalList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested actionApprovals.
func (c *FakeActionApprovals) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(actionapprovalsResource, c.ns, opts))

}

// Create takes the representation of a actionApproval and creates it.  Returns the server's representation of the actionApproval, and an error, if there is any.
func (c *FakeActionApprovals) Create(actionApproval *v1alpha1.ActionApproval) (result *v1alpha1.ActionApproval, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(actionapprovalsResource, c.ns, actionApproval), &v1alpha1.ActionApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ActionApproval), err
}

// Update takes the representation of a actionApproval and updates it. Returns the server's representation of the actionApproval, and an error, if there is any.
func (c *FakeActionApprovals) Update(actionApproval *v1alpha1.ActionApproval) (result *v1alpha1.ActionApproval, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(actionapprovalsResource, c.ns, actionApproval), &v1alpha1.ActionApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ActionApproval), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeActionApprovals) UpdateStatus(actionApproval *v1alpha1.ActionApproval) (*v1alpha1.ActionApproval, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(actionapprovalsResource, "status", c.ns, actionApproval), &v1alpha1.ActionApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ActionApproval), err
}

// Delete takes name of the actionApproval and deletes it. Returns an error if one occurs.
func (c *FakeActionApprovals) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(actionapprovalsResource, c.ns, name), &v1alpha1.ActionApproval{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeActionApprovals) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(actionapprovalsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ActionApprovalList{})
	return err
}

// Patch applies the patch and returns the patched actionApproval.
func (c *FakeActionApprovals) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ActionApproval, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(actionapprovalsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ActionApproval{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ActionApproval), err
}
