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

type ConditionInitParameters struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`

	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type ConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Optional
	Title *string `json:"title" tf:"title,omitempty"`
}

type SubscriptionIAMMemberInitParameters struct {
	Condition []ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type SubscriptionIAMMemberObservation struct {
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	Subscription *string `json:"subscription,omitempty" tf:"subscription,omitempty"`
}

type SubscriptionIAMMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +crossplane:generate:reference:type=Subscription
	// +kubebuilder:validation:Optional
	Subscription *string `json:"subscription,omitempty" tf:"subscription,omitempty"`

	// Reference to a Subscription to populate subscription.
	// +kubebuilder:validation:Optional
	SubscriptionRef *v1.Reference `json:"subscriptionRef,omitempty" tf:"-"`

	// Selector for a Subscription to populate subscription.
	// +kubebuilder:validation:Optional
	SubscriptionSelector *v1.Selector `json:"subscriptionSelector,omitempty" tf:"-"`
}

// SubscriptionIAMMemberSpec defines the desired state of SubscriptionIAMMember
type SubscriptionIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionIAMMemberParameters `json:"forProvider"`
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
	InitProvider SubscriptionIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// SubscriptionIAMMemberStatus defines the observed state of SubscriptionIAMMember.
type SubscriptionIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionIAMMember is the Schema for the SubscriptionIAMMembers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type SubscriptionIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.member) || (has(self.initProvider) && has(self.initProvider.member))",message="spec.forProvider.member is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   SubscriptionIAMMemberSpec   `json:"spec"`
	Status SubscriptionIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionIAMMemberList contains a list of SubscriptionIAMMembers
type SubscriptionIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubscriptionIAMMember `json:"items"`
}

// Repository type metadata.
var (
	SubscriptionIAMMember_Kind             = "SubscriptionIAMMember"
	SubscriptionIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubscriptionIAMMember_Kind}.String()
	SubscriptionIAMMember_KindAPIVersion   = SubscriptionIAMMember_Kind + "." + CRDGroupVersion.String()
	SubscriptionIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(SubscriptionIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&SubscriptionIAMMember{}, &SubscriptionIAMMemberList{})
}
