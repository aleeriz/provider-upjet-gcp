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

type TargetPoolInitParameters struct {

	// URL to the backup target pool. Must also set
	// failover_ratio.
	BackupPool *string `json:"backupPool,omitempty" tf:"backup_pool,omitempty"`

	// Textual description field.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio *float64 `json:"failoverRatio,omitempty" tf:"failover_ratio,omitempty"`

	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name".
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT_IP_PROTO" also includes the protocol (default "NONE").
	SessionAffinity *string `json:"sessionAffinity,omitempty" tf:"session_affinity,omitempty"`
}

type TargetPoolObservation struct {

	// URL to the backup target pool. Must also set
	// failover_ratio.
	BackupPool *string `json:"backupPool,omitempty" tf:"backup_pool,omitempty"`

	// Textual description field.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio *float64 `json:"failoverRatio,omitempty" tf:"failover_ratio,omitempty"`

	// List of zero or one health check name or self_link. Only
	// legacy google_compute_http_health_check is supported.
	HealthChecks []*string `json:"healthChecks,omitempty" tf:"health_checks,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/targetPools/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name".
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Where the target pool resides. Defaults to project
	// region.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT_IP_PROTO" also includes the protocol (default "NONE").
	SessionAffinity *string `json:"sessionAffinity,omitempty" tf:"session_affinity,omitempty"`
}

type TargetPoolParameters struct {

	// URL to the backup target pool. Must also set
	// failover_ratio.
	// +kubebuilder:validation:Optional
	BackupPool *string `json:"backupPool,omitempty" tf:"backup_pool,omitempty"`

	// Textual description field.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	// +kubebuilder:validation:Optional
	FailoverRatio *float64 `json:"failoverRatio,omitempty" tf:"failover_ratio,omitempty"`

	// List of zero or one health check name or self_link. Only
	// legacy google_compute_http_health_check is supported.
	// +crossplane:generate:reference:type=HTTPHealthCheck
	// +kubebuilder:validation:Optional
	HealthChecks []*string `json:"healthChecks,omitempty" tf:"health_checks,omitempty"`

	// References to HTTPHealthCheck to populate healthChecks.
	// +kubebuilder:validation:Optional
	HealthChecksRefs []v1.Reference `json:"healthChecksRefs,omitempty" tf:"-"`

	// Selector for a list of HTTPHealthCheck to populate healthChecks.
	// +kubebuilder:validation:Optional
	HealthChecksSelector *v1.Selector `json:"healthChecksSelector,omitempty" tf:"-"`

	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name".
	// +kubebuilder:validation:Optional
	Instances []*string `json:"instances,omitempty" tf:"instances,omitempty"`

	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Where the target pool resides. Defaults to project
	// region.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT_IP_PROTO" also includes the protocol (default "NONE").
	// +kubebuilder:validation:Optional
	SessionAffinity *string `json:"sessionAffinity,omitempty" tf:"session_affinity,omitempty"`
}

// TargetPoolSpec defines the desired state of TargetPool
type TargetPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetPoolParameters `json:"forProvider"`
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
	InitProvider TargetPoolInitParameters `json:"initProvider,omitempty"`
}

// TargetPoolStatus defines the observed state of TargetPool.
type TargetPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TargetPool is the Schema for the TargetPools API. Manages a Target Pool within GCE.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type TargetPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetPoolSpec   `json:"spec"`
	Status            TargetPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetPoolList contains a list of TargetPools
type TargetPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetPool `json:"items"`
}

// Repository type metadata.
var (
	TargetPool_Kind             = "TargetPool"
	TargetPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetPool_Kind}.String()
	TargetPool_KindAPIVersion   = TargetPool_Kind + "." + CRDGroupVersion.String()
	TargetPool_GroupVersionKind = CRDGroupVersion.WithKind(TargetPool_Kind)
)

func init() {
	SchemeBuilder.Register(&TargetPool{}, &TargetPoolList{})
}
