// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AppConnectionInitParameters struct {

	// Address of the remote application endpoint for the BeyondCorp AppConnection.
	// Structure is documented below.
	ApplicationEndpoint []ApplicationEndpointInitParameters `json:"applicationEndpoint,omitempty" tf:"application_endpoint,omitempty"`

	// List of AppConnectors that are authorised to be associated with this AppConnection
	Connectors []*string `json:"connectors,omitempty" tf:"connectors,omitempty"`

	// An arbitrary user-provided name for the AppConnection.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Gateway used by the AppConnection.
	// Structure is documented below.
	Gateway []GatewayInitParameters `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Resource labels to represent user provided metadata.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// ID of the AppConnection.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the AppConnection.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of hosting used by the gateway. Refer to
	// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1
	// for a list of possible values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AppConnectionObservation struct {

	// Address of the remote application endpoint for the BeyondCorp AppConnection.
	// Structure is documented below.
	ApplicationEndpoint []ApplicationEndpointObservation `json:"applicationEndpoint,omitempty" tf:"application_endpoint,omitempty"`

	// List of AppConnectors that are authorised to be associated with this AppConnection
	Connectors []*string `json:"connectors,omitempty" tf:"connectors,omitempty"`

	// An arbitrary user-provided name for the AppConnection.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// Gateway used by the AppConnection.
	// Structure is documented below.
	Gateway []GatewayObservation `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{region}}/appConnections/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource labels to represent user provided metadata.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// ID of the AppConnection.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the AppConnection.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// The type of hosting used by the gateway. Refer to
	// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1
	// for a list of possible values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AppConnectionParameters struct {

	// Address of the remote application endpoint for the BeyondCorp AppConnection.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ApplicationEndpoint []ApplicationEndpointParameters `json:"applicationEndpoint,omitempty" tf:"application_endpoint,omitempty"`

	// List of AppConnectors that are authorised to be associated with this AppConnection
	// +kubebuilder:validation:Optional
	Connectors []*string `json:"connectors,omitempty" tf:"connectors,omitempty"`

	// An arbitrary user-provided name for the AppConnection.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Gateway used by the AppConnection.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Gateway []GatewayParameters `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// Resource labels to represent user provided metadata.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// ID of the AppConnection.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the AppConnection.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The type of hosting used by the gateway. Refer to
	// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1
	// for a list of possible values.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ApplicationEndpointInitParameters struct {

	// Hostname or IP address of the remote application endpoint.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// Port of the remote application endpoint.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type ApplicationEndpointObservation struct {

	// Hostname or IP address of the remote application endpoint.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// Port of the remote application endpoint.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type ApplicationEndpointParameters struct {

	// Hostname or IP address of the remote application endpoint.
	// +kubebuilder:validation:Optional
	Host *string `json:"host" tf:"host,omitempty"`

	// Port of the remote application endpoint.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`
}

type GatewayInitParameters struct {

	// AppGateway name in following format: projects/{project_id}/locations/{locationId}/appgateways/{gateway_id}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/beyondcorp/v1beta1.AppGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	AppGateway *string `json:"appGateway,omitempty" tf:"app_gateway,omitempty"`

	// Reference to a AppGateway in beyondcorp to populate appGateway.
	// +kubebuilder:validation:Optional
	AppGatewayRef *v1.Reference `json:"appGatewayRef,omitempty" tf:"-"`

	// Selector for a AppGateway in beyondcorp to populate appGateway.
	// +kubebuilder:validation:Optional
	AppGatewaySelector *v1.Selector `json:"appGatewaySelector,omitempty" tf:"-"`

	// The type of hosting used by the gateway. Refer to
	// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1
	// for a list of possible values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type GatewayObservation struct {

	// AppGateway name in following format: projects/{project_id}/locations/{locationId}/appgateways/{gateway_id}.
	AppGateway *string `json:"appGateway,omitempty" tf:"app_gateway,omitempty"`

	// (Output)
	// Ingress port reserved on the gateways for this AppConnection, if not specified or zero, the default port is 19443.
	IngressPort *float64 `json:"ingressPort,omitempty" tf:"ingress_port,omitempty"`

	// The type of hosting used by the gateway. Refer to
	// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1
	// for a list of possible values.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Output)
	// Server-defined URI for this resource.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type GatewayParameters struct {

	// AppGateway name in following format: projects/{project_id}/locations/{locationId}/appgateways/{gateway_id}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/beyondcorp/v1beta1.AppGateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AppGateway *string `json:"appGateway,omitempty" tf:"app_gateway,omitempty"`

	// Reference to a AppGateway in beyondcorp to populate appGateway.
	// +kubebuilder:validation:Optional
	AppGatewayRef *v1.Reference `json:"appGatewayRef,omitempty" tf:"-"`

	// Selector for a AppGateway in beyondcorp to populate appGateway.
	// +kubebuilder:validation:Optional
	AppGatewaySelector *v1.Selector `json:"appGatewaySelector,omitempty" tf:"-"`

	// The type of hosting used by the gateway. Refer to
	// https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#Type_1
	// for a list of possible values.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// AppConnectionSpec defines the desired state of AppConnection
type AppConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppConnectionParameters `json:"forProvider"`
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
	InitProvider AppConnectionInitParameters `json:"initProvider,omitempty"`
}

// AppConnectionStatus defines the observed state of AppConnection.
type AppConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AppConnection is the Schema for the AppConnections API. A BeyondCorp AppConnection resource represents a BeyondCorp protected AppConnection to a remote application.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type AppConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.applicationEndpoint) || (has(self.initProvider) && has(self.initProvider.applicationEndpoint))",message="spec.forProvider.applicationEndpoint is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.region) || (has(self.initProvider) && has(self.initProvider.region))",message="spec.forProvider.region is a required parameter"
	Spec   AppConnectionSpec   `json:"spec"`
	Status AppConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppConnectionList contains a list of AppConnections
type AppConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppConnection `json:"items"`
}

// Repository type metadata.
var (
	AppConnection_Kind             = "AppConnection"
	AppConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppConnection_Kind}.String()
	AppConnection_KindAPIVersion   = AppConnection_Kind + "." + CRDGroupVersion.String()
	AppConnection_GroupVersionKind = CRDGroupVersion.WithKind(AppConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&AppConnection{}, &AppConnectionList{})
}
