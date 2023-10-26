// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceGroupInitParameters struct {

	// An optional textual description of the instance
	// group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The list of instances in the group, in self_link format.
	// When adding instances they must all be in the same network and zone as the instance group.
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// The named port configuration. See the section below
	// for details on configuration. Structure is documented below.
	NamedPort []NamedPortInitParameters `json:"namedPort,omitempty" tf:"named_port,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type InstanceGroupObservation struct {

	// An optional textual description of the instance
	// group.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}/zones/{{zone}}/instanceGroups/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The list of instances in the group, in self_link format.
	// When adding instances they must all be in the same network and zone as the instance group.
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// The named port configuration. See the section below
	// for details on configuration. Structure is documented below.
	NamedPort []NamedPortObservation `json:"namedPort,omitempty" tf:"named_port,omitempty"`

	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// network nor instances is specified, this field will be blank).
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The number of instances in the group.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The zone that this instance group should be created in.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type InstanceGroupParameters struct {

	// An optional textual description of the instance
	// group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The list of instances in the group, in self_link format.
	// When adding instances they must all be in the same network and zone as the instance group.
	// +kubebuilder:validation:Optional
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// The named port configuration. See the section below
	// for details on configuration. Structure is documented below.
	// +kubebuilder:validation:Optional
	NamedPort []NamedPortParameters `json:"namedPort,omitempty" tf:"named_port,omitempty"`

	// The URL of the network the instance group is in. If
	// this is different from the network where the instances are in, the creation
	// fails. Defaults to the network where the instances are in (if neither
	// network nor instances is specified, this field will be blank).
	// +crossplane:generate:reference:type=Network
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The zone that this instance group should be created in.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type NamedPortInitParameters struct {

	// The name which the port will be mapped to.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The port number to map the name to.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type NamedPortObservation struct {

	// The name which the port will be mapped to.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The port number to map the name to.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type NamedPortParameters struct {

	// The name which the port will be mapped to.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The port number to map the name to.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`
}

// InstanceGroupSpec defines the desired state of InstanceGroup
type InstanceGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceGroupParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider InstanceGroupInitParameters `json:"initProvider,omitempty"`
}

// InstanceGroupStatus defines the observed state of InstanceGroup.
type InstanceGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceGroup is the Schema for the InstanceGroups API. Manages an Instance Group within GCE.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type InstanceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceGroupSpec   `json:"spec"`
	Status            InstanceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceGroupList contains a list of InstanceGroups
type InstanceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceGroup `json:"items"`
}

// Repository type metadata.
var (
	InstanceGroup_Kind             = "InstanceGroup"
	InstanceGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceGroup_Kind}.String()
	InstanceGroup_KindAPIVersion   = InstanceGroup_Kind + "." + CRDGroupVersion.String()
	InstanceGroup_GroupVersionKind = CRDGroupVersion.WithKind(InstanceGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceGroup{}, &InstanceGroupList{})
}
