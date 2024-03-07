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

type CertificateMapInitParameters struct {

	// A human-readable description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set of labels associated with a Certificate Map resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type CertificateMapObservation struct {

	// Creation timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds with up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// A human-readable description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// for all of the labels present on the resource.
	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// A list of target proxies that use this Certificate Map
	// Structure is documented below.
	GclbTargets []GclbTargetsObservation `json:"gclbTargets,omitempty" tf:"gclb_targets,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/global/certificateMaps/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of labels associated with a Certificate Map resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// Update timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds with up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type CertificateMapParameters struct {

	// A human-readable description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Set of labels associated with a Certificate Map resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type GclbTargetsInitParameters struct {
}

type GclbTargetsObservation struct {

	// An IP configuration where this Certificate Map is serving
	// Structure is documented below.
	IPConfigs []IPConfigsObservation `json:"ipConfigs,omitempty" tf:"ip_configs,omitempty"`

	// Proxy name must be in the format projects//locations//targetHttpsProxies/*.
	// This field is part of a union field target_proxy: Only one of targetHttpsProxy or
	// targetSslProxy may be set.
	TargetHTTPSProxy *string `json:"targetHttpsProxy,omitempty" tf:"target_https_proxy,omitempty"`

	// Proxy name must be in the format projects//locations//targetSslProxies/*.
	// This field is part of a union field target_proxy: Only one of targetHttpsProxy or
	// targetSslProxy may be set.
	TargetSSLProxy *string `json:"targetSslProxy,omitempty" tf:"target_ssl_proxy,omitempty"`
}

type GclbTargetsParameters struct {
}

type IPConfigsInitParameters struct {
}

type IPConfigsObservation struct {

	// An external IP address
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// A list of ports
	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`
}

type IPConfigsParameters struct {
}

// CertificateMapSpec defines the desired state of CertificateMap
type CertificateMapSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateMapParameters `json:"forProvider"`
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
	InitProvider CertificateMapInitParameters `json:"initProvider,omitempty"`
}

// CertificateMapStatus defines the observed state of CertificateMap.
type CertificateMapStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateMapObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CertificateMap is the Schema for the CertificateMaps API. CertificateMap defines a collection of certificate configurations, which are usable by any associated target proxies
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type CertificateMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateMapSpec   `json:"spec"`
	Status            CertificateMapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateMapList contains a list of CertificateMaps
type CertificateMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateMap `json:"items"`
}

// Repository type metadata.
var (
	CertificateMap_Kind             = "CertificateMap"
	CertificateMap_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateMap_Kind}.String()
	CertificateMap_KindAPIVersion   = CertificateMap_Kind + "." + CRDGroupVersion.String()
	CertificateMap_GroupVersionKind = CRDGroupVersion.WithKind(CertificateMap_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateMap{}, &CertificateMapList{})
}
