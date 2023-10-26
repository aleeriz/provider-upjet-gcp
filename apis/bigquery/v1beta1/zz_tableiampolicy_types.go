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

type TableIAMPolicyInitParameters struct {

	// The policy data generated by
	// a google_iam_policy data source.
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type TableIAMPolicyObservation struct {
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// (Computed) The etag of the IAM policy.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The policy data generated by
	// a google_iam_policy data source.
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`
}

type TableIAMPolicyParameters struct {

	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Dataset
	// +kubebuilder:validation:Optional
	DatasetID *string `json:"datasetId,omitempty" tf:"dataset_id,omitempty"`

	// Reference to a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDRef *v1.Reference `json:"datasetIdRef,omitempty" tf:"-"`

	// Selector for a Dataset in bigquery to populate datasetId.
	// +kubebuilder:validation:Optional
	DatasetIDSelector *v1.Selector `json:"datasetIdSelector,omitempty" tf:"-"`

	// The policy data generated by
	// a google_iam_policy data source.
	// +kubebuilder:validation:Optional
	PolicyData *string `json:"policyData,omitempty" tf:"policy_data,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/bigquery/v1beta1.Table
	// +kubebuilder:validation:Optional
	TableID *string `json:"tableId,omitempty" tf:"table_id,omitempty"`

	// Reference to a Table in bigquery to populate tableId.
	// +kubebuilder:validation:Optional
	TableIDRef *v1.Reference `json:"tableIdRef,omitempty" tf:"-"`

	// Selector for a Table in bigquery to populate tableId.
	// +kubebuilder:validation:Optional
	TableIDSelector *v1.Selector `json:"tableIdSelector,omitempty" tf:"-"`
}

// TableIAMPolicySpec defines the desired state of TableIAMPolicy
type TableIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TableIAMPolicyParameters `json:"forProvider"`
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
	InitProvider TableIAMPolicyInitParameters `json:"initProvider,omitempty"`
}

// TableIAMPolicyStatus defines the observed state of TableIAMPolicy.
type TableIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TableIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TableIAMPolicy is the Schema for the TableIAMPolicys API. Collection of resources to manage IAM policy for BigQuery Table
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TableIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyData) || (has(self.initProvider) && has(self.initProvider.policyData))",message="spec.forProvider.policyData is a required parameter"
	Spec   TableIAMPolicySpec   `json:"spec"`
	Status TableIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TableIAMPolicyList contains a list of TableIAMPolicys
type TableIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TableIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	TableIAMPolicy_Kind             = "TableIAMPolicy"
	TableIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TableIAMPolicy_Kind}.String()
	TableIAMPolicy_KindAPIVersion   = TableIAMPolicy_Kind + "." + CRDGroupVersion.String()
	TableIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(TableIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&TableIAMPolicy{}, &TableIAMPolicyList{})
}
