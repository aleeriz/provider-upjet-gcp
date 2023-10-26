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

type OAuthIdPConfigInitParameters struct {

	// Human friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// For OIDC Idps, the issuer identifier.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The name of the OauthIdpConfig. Must start with oidc..
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type OAuthIdPConfigObservation struct {

	// Human friendly display name.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// an identifier for the resource with format projects/{{project}}/oauthIdpConfigs/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// For OIDC Idps, the issuer identifier.
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The name of the OauthIdpConfig. Must start with oidc..
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type OAuthIdPConfigParameters struct {

	// The client id of an OAuth client.
	// +kubebuilder:validation:Optional
	ClientIDSecretRef v1.SecretKeySelector `json:"clientIdSecretRef" tf:"-"`

	// The client secret of the OAuth client, to enable OIDC code flow.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef *v1.SecretKeySelector `json:"clientSecretSecretRef,omitempty" tf:"-"`

	// Human friendly display name.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// If this config allows users to sign in with the provider.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// For OIDC Idps, the issuer identifier.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// The name of the OauthIdpConfig. Must start with oidc..
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// OAuthIdPConfigSpec defines the desired state of OAuthIdPConfig
type OAuthIdPConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OAuthIdPConfigParameters `json:"forProvider"`
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
	InitProvider OAuthIdPConfigInitParameters `json:"initProvider,omitempty"`
}

// OAuthIdPConfigStatus defines the observed state of OAuthIdPConfig.
type OAuthIdPConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OAuthIdPConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OAuthIdPConfig is the Schema for the OAuthIdPConfigs API. OIDC IdP configuration for a Identity Toolkit project.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type OAuthIdPConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientIdSecretRef)",message="spec.forProvider.clientIdSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.issuer) || (has(self.initProvider) && has(self.initProvider.issuer))",message="spec.forProvider.issuer is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   OAuthIdPConfigSpec   `json:"spec"`
	Status OAuthIdPConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OAuthIdPConfigList contains a list of OAuthIdPConfigs
type OAuthIdPConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OAuthIdPConfig `json:"items"`
}

// Repository type metadata.
var (
	OAuthIdPConfig_Kind             = "OAuthIdPConfig"
	OAuthIdPConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OAuthIdPConfig_Kind}.String()
	OAuthIdPConfig_KindAPIVersion   = OAuthIdPConfig_Kind + "." + CRDGroupVersion.String()
	OAuthIdPConfig_GroupVersionKind = CRDGroupVersion.WithKind(OAuthIdPConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&OAuthIdPConfig{}, &OAuthIdPConfigList{})
}
