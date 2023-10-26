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

type ContactInitParameters struct {

	// The email address to send notifications to. This does not need to be a Google account.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
	LanguageTag *string `json:"languageTag,omitempty" tf:"language_tag,omitempty"`

	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions []*string `json:"notificationCategorySubscriptions,omitempty" tf:"notification_category_subscriptions,omitempty"`

	// The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`
}

type ContactObservation struct {

	// The email address to send notifications to. This does not need to be a Google account.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
	LanguageTag *string `json:"languageTag,omitempty" tf:"language_tag,omitempty"`

	// The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions []*string `json:"notificationCategorySubscriptions,omitempty" tf:"notification_category_subscriptions,omitempty"`

	// The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`
}

type ContactParameters struct {

	// The email address to send notifications to. This does not need to be a Google account.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.
	// +kubebuilder:validation:Optional
	LanguageTag *string `json:"languageTag,omitempty" tf:"language_tag,omitempty"`

	// The categories of notifications that the contact will receive communications for.
	// +kubebuilder:validation:Optional
	NotificationCategorySubscriptions []*string `json:"notificationCategorySubscriptions,omitempty" tf:"notification_category_subscriptions,omitempty"`

	// The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`
}

// ContactSpec defines the desired state of Contact
type ContactSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContactParameters `json:"forProvider"`
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
	InitProvider ContactInitParameters `json:"initProvider,omitempty"`
}

// ContactStatus defines the observed state of Contact.
type ContactStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContactObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Contact is the Schema for the Contacts API. A contact that will receive notifications from Google Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Contact struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.languageTag) || (has(self.initProvider) && has(self.initProvider.languageTag))",message="spec.forProvider.languageTag is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.notificationCategorySubscriptions) || (has(self.initProvider) && has(self.initProvider.notificationCategorySubscriptions))",message="spec.forProvider.notificationCategorySubscriptions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.parent) || (has(self.initProvider) && has(self.initProvider.parent))",message="spec.forProvider.parent is a required parameter"
	Spec   ContactSpec   `json:"spec"`
	Status ContactStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContactList contains a list of Contacts
type ContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Contact `json:"items"`
}

// Repository type metadata.
var (
	Contact_Kind             = "Contact"
	Contact_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Contact_Kind}.String()
	Contact_KindAPIVersion   = Contact_Kind + "." + CRDGroupVersion.String()
	Contact_GroupVersionKind = CRDGroupVersion.WithKind(Contact_Kind)
)

func init() {
	SchemeBuilder.Register(&Contact{}, &ContactList{})
}
