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

type AllowInitParameters struct {

	// An optional list of ports to which this rule applies. This field
	// is only applicable for UDP or TCP protocol. Each entry must be
	// either an integer or a range. If not specified, this rule
	// applies to connections through any port.
	// Example inputs include: ["22"], ["80","443"], and
	// ["12345-12349"].
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The IP protocol to which this rule applies. The protocol type is
	// required when creating a firewall rule. This value can either be
	// one of the following well known protocol strings (tcp, udp,
	// icmp, esp, ah, sctp, ipip, all), or the IP protocol number.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type AllowObservation struct {

	// An optional list of ports to which this rule applies. This field
	// is only applicable for UDP or TCP protocol. Each entry must be
	// either an integer or a range. If not specified, this rule
	// applies to connections through any port.
	// Example inputs include: ["22"], ["80","443"], and
	// ["12345-12349"].
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The IP protocol to which this rule applies. The protocol type is
	// required when creating a firewall rule. This value can either be
	// one of the following well known protocol strings (tcp, udp,
	// icmp, esp, ah, sctp, ipip, all), or the IP protocol number.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type AllowParameters struct {

	// An optional list of ports to which this rule applies. This field
	// is only applicable for UDP or TCP protocol. Each entry must be
	// either an integer or a range. If not specified, this rule
	// applies to connections through any port.
	// Example inputs include: ["22"], ["80","443"], and
	// ["12345-12349"].
	// +kubebuilder:validation:Optional
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The IP protocol to which this rule applies. The protocol type is
	// required when creating a firewall rule. This value can either be
	// one of the following well known protocol strings (tcp, udp,
	// icmp, esp, ah, sctp, ipip, all), or the IP protocol number.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type DenyInitParameters struct {

	// An optional list of ports to which this rule applies. This field
	// is only applicable for UDP or TCP protocol. Each entry must be
	// either an integer or a range. If not specified, this rule
	// applies to connections through any port.
	// Example inputs include: ["22"], ["80","443"], and
	// ["12345-12349"].
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The IP protocol to which this rule applies. The protocol type is
	// required when creating a firewall rule. This value can either be
	// one of the following well known protocol strings (tcp, udp,
	// icmp, esp, ah, sctp, ipip, all), or the IP protocol number.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type DenyObservation struct {

	// An optional list of ports to which this rule applies. This field
	// is only applicable for UDP or TCP protocol. Each entry must be
	// either an integer or a range. If not specified, this rule
	// applies to connections through any port.
	// Example inputs include: ["22"], ["80","443"], and
	// ["12345-12349"].
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The IP protocol to which this rule applies. The protocol type is
	// required when creating a firewall rule. This value can either be
	// one of the following well known protocol strings (tcp, udp,
	// icmp, esp, ah, sctp, ipip, all), or the IP protocol number.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type DenyParameters struct {

	// An optional list of ports to which this rule applies. This field
	// is only applicable for UDP or TCP protocol. Each entry must be
	// either an integer or a range. If not specified, this rule
	// applies to connections through any port.
	// Example inputs include: ["22"], ["80","443"], and
	// ["12345-12349"].
	// +kubebuilder:validation:Optional
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The IP protocol to which this rule applies. The protocol type is
	// required when creating a firewall rule. This value can either be
	// one of the following well known protocol strings (tcp, udp,
	// icmp, esp, ah, sctp, ipip, all), or the IP protocol number.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type FirewallInitParameters struct {

	// The list of ALLOW rules specified by this firewall. Each rule
	// specifies a protocol and port-range tuple that describes a permitted
	// connection.
	// Structure is documented below.
	Allow []AllowInitParameters `json:"allow,omitempty" tf:"allow,omitempty"`

	// The list of DENY rules specified by this firewall. Each rule specifies
	// a protocol and port-range tuple that describes a denied connection.
	// Structure is documented below.
	Deny []DenyInitParameters `json:"deny,omitempty" tf:"deny,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If destination ranges are specified, the firewall will apply only to
	// traffic that has destination IP address in these ranges. These ranges
	// must be expressed in CIDR format. IPv4 or IPv6 ranges are supported.
	// +listType=set
	DestinationRanges []*string `json:"destinationRanges,omitempty" tf:"destination_ranges,omitempty"`

	// Direction of traffic to which this firewall applies; default is
	// INGRESS. Note: For INGRESS traffic, one of source_ranges,
	// source_tags or source_service_accounts is required.
	// Possible values are: INGRESS, EGRESS.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Denotes whether the firewall rule is disabled, i.e not applied to the
	// network it is associated with. When set to true, the firewall rule is
	// not enforced and the network behaves as if it did not exist. If this
	// is unspecified, the firewall rule will be enabled.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// This field denotes whether to enable logging for a particular firewall rule.
	// If logging is enabled, logs will be exported to Stackdriver. Deprecated in favor of log_config
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// This field denotes the logging options for a particular firewall rule.
	// If defined, logging is enabled, and logs will be exported to Cloud Logging.
	// Structure is documented below.
	LogConfig []FirewallLogConfigInitParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The name or self_link of the network to attach this firewall to.
	// +crossplane:generate:reference:type=Network
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-gcp/config/common.SelfLinkExtractor()
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Reference to a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// Selector for a Network to populate network.
	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// Priority for this rule. This is an integer between 0 and 65535, both
	// inclusive. When not specified, the value assumed is 1000. Relative
	// priorities determine precedence of conflicting rules. Lower value of
	// priority implies higher precedence (eg, a rule with priority 0 has
	// higher precedence than a rule with priority 1). DENY rules take
	// precedence over ALLOW rules having equal priority.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// If source ranges are specified, the firewall will apply only to
	// traffic that has source IP address in these ranges. These ranges must
	// be expressed in CIDR format. One or both of sourceRanges and
	// sourceTags may be set. If both properties are set, the firewall will
	// apply to traffic that has source IP address within sourceRanges OR the
	// source IP that belongs to a tag listed in the sourceTags property. The
	// connection does not need to match both properties for the firewall to
	// apply. IPv4 or IPv6 ranges are supported. For INGRESS traffic, one of
	// source_ranges, source_tags or source_service_accounts is required.
	// +listType=set
	SourceRanges []*string `json:"sourceRanges,omitempty" tf:"source_ranges,omitempty"`

	// If source service accounts are specified, the firewall will apply only
	// to traffic originating from an instance with a service account in this
	// list. Source service accounts cannot be used to control traffic to an
	// instance's external IP address because service accounts are associated
	// with an instance, not an IP address. sourceRanges can be set at the
	// same time as sourceServiceAccounts. If both are set, the firewall will
	// apply to traffic that has source IP address within sourceRanges OR the
	// source IP belongs to an instance with service account listed in
	// sourceServiceAccount. The connection does not need to match both
	// properties for the firewall to apply. sourceServiceAccounts cannot be
	// used at the same time as sourceTags or targetTags. For INGRESS traffic,
	// one of source_ranges, source_tags or source_service_accounts is required.
	// +listType=set
	SourceServiceAccounts []*string `json:"sourceServiceAccounts,omitempty" tf:"source_service_accounts,omitempty"`

	// If source tags are specified, the firewall will apply only to traffic
	// with source IP that belongs to a tag listed in source tags. Source
	// tags cannot be used to control traffic to an instance's external IP
	// address. Because tags are associated with an instance, not an IP
	// address. One or both of sourceRanges and sourceTags may be set. If
	// both properties are set, the firewall will apply to traffic that has
	// source IP address within sourceRanges OR the source IP that belongs to
	// a tag listed in the sourceTags property. The connection does not need
	// to match both properties for the firewall to apply. For INGRESS traffic,
	// one of source_ranges, source_tags or source_service_accounts is required.
	// +listType=set
	SourceTags []*string `json:"sourceTags,omitempty" tf:"source_tags,omitempty"`

	// A list of service accounts indicating sets of instances located in the
	// network that may make network connections as specified in allowed[].
	// targetServiceAccounts cannot be used at the same time as targetTags or
	// sourceTags. If neither targetServiceAccounts nor targetTags are
	// specified, the firewall rule applies to all instances on the specified
	// network.
	// +listType=set
	TargetServiceAccounts []*string `json:"targetServiceAccounts,omitempty" tf:"target_service_accounts,omitempty"`

	// A list of instance tags indicating sets of instances located in the
	// network that may make network connections as specified in allowed[].
	// If no targetTags are specified, the firewall rule applies to all
	// instances on the specified network.
	// +listType=set
	TargetTags []*string `json:"targetTags,omitempty" tf:"target_tags,omitempty"`
}

type FirewallLogConfigInitParameters struct {

	// This field denotes whether to include or exclude metadata for firewall logs.
	// Possible values are: EXCLUDE_ALL_METADATA, INCLUDE_ALL_METADATA.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

type FirewallLogConfigObservation struct {

	// This field denotes whether to include or exclude metadata for firewall logs.
	// Possible values are: EXCLUDE_ALL_METADATA, INCLUDE_ALL_METADATA.
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`
}

type FirewallLogConfigParameters struct {

	// This field denotes whether to include or exclude metadata for firewall logs.
	// Possible values are: EXCLUDE_ALL_METADATA, INCLUDE_ALL_METADATA.
	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata" tf:"metadata,omitempty"`
}

type FirewallObservation struct {

	// The list of ALLOW rules specified by this firewall. Each rule
	// specifies a protocol and port-range tuple that describes a permitted
	// connection.
	// Structure is documented below.
	Allow []AllowObservation `json:"allow,omitempty" tf:"allow,omitempty"`

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// The list of DENY rules specified by this firewall. Each rule specifies
	// a protocol and port-range tuple that describes a denied connection.
	// Structure is documented below.
	Deny []DenyObservation `json:"deny,omitempty" tf:"deny,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If destination ranges are specified, the firewall will apply only to
	// traffic that has destination IP address in these ranges. These ranges
	// must be expressed in CIDR format. IPv4 or IPv6 ranges are supported.
	// +listType=set
	DestinationRanges []*string `json:"destinationRanges,omitempty" tf:"destination_ranges,omitempty"`

	// Direction of traffic to which this firewall applies; default is
	// INGRESS. Note: For INGRESS traffic, one of source_ranges,
	// source_tags or source_service_accounts is required.
	// Possible values are: INGRESS, EGRESS.
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Denotes whether the firewall rule is disabled, i.e not applied to the
	// network it is associated with. When set to true, the firewall rule is
	// not enforced and the network behaves as if it did not exist. If this
	// is unspecified, the firewall rule will be enabled.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// This field denotes whether to enable logging for a particular firewall rule.
	// If logging is enabled, logs will be exported to Stackdriver. Deprecated in favor of log_config
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// an identifier for the resource with format projects/{{project}}/global/firewalls/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// This field denotes the logging options for a particular firewall rule.
	// If defined, logging is enabled, and logs will be exported to Cloud Logging.
	// Structure is documented below.
	LogConfig []FirewallLogConfigObservation `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The name or self_link of the network to attach this firewall to.
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// Priority for this rule. This is an integer between 0 and 65535, both
	// inclusive. When not specified, the value assumed is 1000. Relative
	// priorities determine precedence of conflicting rules. Lower value of
	// priority implies higher precedence (eg, a rule with priority 0 has
	// higher precedence than a rule with priority 1). DENY rules take
	// precedence over ALLOW rules having equal priority.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// If source ranges are specified, the firewall will apply only to
	// traffic that has source IP address in these ranges. These ranges must
	// be expressed in CIDR format. One or both of sourceRanges and
	// sourceTags may be set. If both properties are set, the firewall will
	// apply to traffic that has source IP address within sourceRanges OR the
	// source IP that belongs to a tag listed in the sourceTags property. The
	// connection does not need to match both properties for the firewall to
	// apply. IPv4 or IPv6 ranges are supported. For INGRESS traffic, one of
	// source_ranges, source_tags or source_service_accounts is required.
	// +listType=set
	SourceRanges []*string `json:"sourceRanges,omitempty" tf:"source_ranges,omitempty"`

	// If source service accounts are specified, the firewall will apply only
	// to traffic originating from an instance with a service account in this
	// list. Source service accounts cannot be used to control traffic to an
	// instance's external IP address because service accounts are associated
	// with an instance, not an IP address. sourceRanges can be set at the
	// same time as sourceServiceAccounts. If both are set, the firewall will
	// apply to traffic that has source IP address within sourceRanges OR the
	// source IP belongs to an instance with service account listed in
	// sourceServiceAccount. The connection does not need to match both
	// properties for the firewall to apply. sourceServiceAccounts cannot be
	// used at the same time as sourceTags or targetTags. For INGRESS traffic,
	// one of source_ranges, source_tags or source_service_accounts is required.
	// +listType=set
	SourceServiceAccounts []*string `json:"sourceServiceAccounts,omitempty" tf:"source_service_accounts,omitempty"`

	// If source tags are specified, the firewall will apply only to traffic
	// with source IP that belongs to a tag listed in source tags. Source
	// tags cannot be used to control traffic to an instance's external IP
	// address. Because tags are associated with an instance, not an IP
	// address. One or both of sourceRanges and sourceTags may be set. If
	// both properties are set, the firewall will apply to traffic that has
	// source IP address within sourceRanges OR the source IP that belongs to
	// a tag listed in the sourceTags property. The connection does not need
	// to match both properties for the firewall to apply. For INGRESS traffic,
	// one of source_ranges, source_tags or source_service_accounts is required.
	// +listType=set
	SourceTags []*string `json:"sourceTags,omitempty" tf:"source_tags,omitempty"`

	// A list of service accounts indicating sets of instances located in the
	// network that may make network connections as specified in allowed[].
	// targetServiceAccounts cannot be used at the same time as targetTags or
	// sourceTags. If neither targetServiceAccounts nor targetTags are
	// specified, the firewall rule applies to all instances on the specified
	// network.
	// +listType=set
	TargetServiceAccounts []*string `json:"targetServiceAccounts,omitempty" tf:"target_service_accounts,omitempty"`

	// A list of instance tags indicating sets of instances located in the
	// network that may make network connections as specified in allowed[].
	// If no targetTags are specified, the firewall rule applies to all
	// instances on the specified network.
	// +listType=set
	TargetTags []*string `json:"targetTags,omitempty" tf:"target_tags,omitempty"`
}

type FirewallParameters struct {

	// The list of ALLOW rules specified by this firewall. Each rule
	// specifies a protocol and port-range tuple that describes a permitted
	// connection.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Allow []AllowParameters `json:"allow,omitempty" tf:"allow,omitempty"`

	// The list of DENY rules specified by this firewall. Each rule specifies
	// a protocol and port-range tuple that describes a denied connection.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Deny []DenyParameters `json:"deny,omitempty" tf:"deny,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// If destination ranges are specified, the firewall will apply only to
	// traffic that has destination IP address in these ranges. These ranges
	// must be expressed in CIDR format. IPv4 or IPv6 ranges are supported.
	// +kubebuilder:validation:Optional
	// +listType=set
	DestinationRanges []*string `json:"destinationRanges,omitempty" tf:"destination_ranges,omitempty"`

	// Direction of traffic to which this firewall applies; default is
	// INGRESS. Note: For INGRESS traffic, one of source_ranges,
	// source_tags or source_service_accounts is required.
	// Possible values are: INGRESS, EGRESS.
	// +kubebuilder:validation:Optional
	Direction *string `json:"direction,omitempty" tf:"direction,omitempty"`

	// Denotes whether the firewall rule is disabled, i.e not applied to the
	// network it is associated with. When set to true, the firewall rule is
	// not enforced and the network behaves as if it did not exist. If this
	// is unspecified, the firewall rule will be enabled.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// This field denotes whether to enable logging for a particular firewall rule.
	// If logging is enabled, logs will be exported to Stackdriver. Deprecated in favor of log_config
	// +kubebuilder:validation:Optional
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging,omitempty"`

	// This field denotes the logging options for a particular firewall rule.
	// If defined, logging is enabled, and logs will be exported to Cloud Logging.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	LogConfig []FirewallLogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The name or self_link of the network to attach this firewall to.
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

	// Priority for this rule. This is an integer between 0 and 65535, both
	// inclusive. When not specified, the value assumed is 1000. Relative
	// priorities determine precedence of conflicting rules. Lower value of
	// priority implies higher precedence (eg, a rule with priority 0 has
	// higher precedence than a rule with priority 1). DENY rules take
	// precedence over ALLOW rules having equal priority.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// If source ranges are specified, the firewall will apply only to
	// traffic that has source IP address in these ranges. These ranges must
	// be expressed in CIDR format. One or both of sourceRanges and
	// sourceTags may be set. If both properties are set, the firewall will
	// apply to traffic that has source IP address within sourceRanges OR the
	// source IP that belongs to a tag listed in the sourceTags property. The
	// connection does not need to match both properties for the firewall to
	// apply. IPv4 or IPv6 ranges are supported. For INGRESS traffic, one of
	// source_ranges, source_tags or source_service_accounts is required.
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceRanges []*string `json:"sourceRanges,omitempty" tf:"source_ranges,omitempty"`

	// If source service accounts are specified, the firewall will apply only
	// to traffic originating from an instance with a service account in this
	// list. Source service accounts cannot be used to control traffic to an
	// instance's external IP address because service accounts are associated
	// with an instance, not an IP address. sourceRanges can be set at the
	// same time as sourceServiceAccounts. If both are set, the firewall will
	// apply to traffic that has source IP address within sourceRanges OR the
	// source IP belongs to an instance with service account listed in
	// sourceServiceAccount. The connection does not need to match both
	// properties for the firewall to apply. sourceServiceAccounts cannot be
	// used at the same time as sourceTags or targetTags. For INGRESS traffic,
	// one of source_ranges, source_tags or source_service_accounts is required.
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceServiceAccounts []*string `json:"sourceServiceAccounts,omitempty" tf:"source_service_accounts,omitempty"`

	// If source tags are specified, the firewall will apply only to traffic
	// with source IP that belongs to a tag listed in source tags. Source
	// tags cannot be used to control traffic to an instance's external IP
	// address. Because tags are associated with an instance, not an IP
	// address. One or both of sourceRanges and sourceTags may be set. If
	// both properties are set, the firewall will apply to traffic that has
	// source IP address within sourceRanges OR the source IP that belongs to
	// a tag listed in the sourceTags property. The connection does not need
	// to match both properties for the firewall to apply. For INGRESS traffic,
	// one of source_ranges, source_tags or source_service_accounts is required.
	// +kubebuilder:validation:Optional
	// +listType=set
	SourceTags []*string `json:"sourceTags,omitempty" tf:"source_tags,omitempty"`

	// A list of service accounts indicating sets of instances located in the
	// network that may make network connections as specified in allowed[].
	// targetServiceAccounts cannot be used at the same time as targetTags or
	// sourceTags. If neither targetServiceAccounts nor targetTags are
	// specified, the firewall rule applies to all instances on the specified
	// network.
	// +kubebuilder:validation:Optional
	// +listType=set
	TargetServiceAccounts []*string `json:"targetServiceAccounts,omitempty" tf:"target_service_accounts,omitempty"`

	// A list of instance tags indicating sets of instances located in the
	// network that may make network connections as specified in allowed[].
	// If no targetTags are specified, the firewall rule applies to all
	// instances on the specified network.
	// +kubebuilder:validation:Optional
	// +listType=set
	TargetTags []*string `json:"targetTags,omitempty" tf:"target_tags,omitempty"`
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallParameters `json:"forProvider"`
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
	InitProvider FirewallInitParameters `json:"initProvider,omitempty"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Firewall is the Schema for the Firewalls API. Each network has its own firewall controlling access to and from the instances.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec"`
	Status            FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	Firewall_Kind             = "Firewall"
	Firewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Firewall_Kind}.String()
	Firewall_KindAPIVersion   = Firewall_Kind + "." + CRDGroupVersion.String()
	Firewall_GroupVersionKind = CRDGroupVersion.WithKind(Firewall_Kind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
