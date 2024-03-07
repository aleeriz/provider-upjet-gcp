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

type AuxiliaryVersionsInitParameters struct {

	// A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml).
	// The mappings override system defaults (some keys cannot be overridden)
	// +mapType=granular
	ConfigOverrides map[string]*string `json:"configOverrides,omitempty" tf:"config_overrides,omitempty"`

	// The identifier for this object. Format specified above.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The Hive metastore schema version.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type AuxiliaryVersionsObservation struct {

	// A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml).
	// The mappings override system defaults (some keys cannot be overridden)
	// +mapType=granular
	ConfigOverrides map[string]*string `json:"configOverrides,omitempty" tf:"config_overrides,omitempty"`

	// The identifier for this object. Format specified above.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The Hive metastore schema version.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type AuxiliaryVersionsParameters struct {

	// A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml).
	// The mappings override system defaults (some keys cannot be overridden)
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ConfigOverrides map[string]*string `json:"configOverrides,omitempty" tf:"config_overrides,omitempty"`

	// The identifier for this object. Format specified above.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// The Hive metastore schema version.
	// +kubebuilder:validation:Optional
	Version *string `json:"version" tf:"version,omitempty"`
}

type ConsumersInitParameters struct {

	// The subnetwork of the customer project from which an IP address is reserved and used as the Dataproc Metastore service's endpoint.
	// It is accessible to hosts in the subnet and to all hosts in a subnet in the same region and same network.
	// There must be at least one IP address available in the subnet's primary range. The subnet is specified in the following form:
	// `projects/{projectNumber}/regions/{region_id}/subnetworks/{subnetwork_id}
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Reference to a Subnetwork in compute to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// Selector for a Subnetwork in compute to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`
}

type ConsumersObservation struct {

	// The URI of the endpoint used to access the metastore service.
	EndpointURI *string `json:"endpointUri,omitempty" tf:"endpoint_uri,omitempty"`

	// The subnetwork of the customer project from which an IP address is reserved and used as the Dataproc Metastore service's endpoint.
	// It is accessible to hosts in the subnet and to all hosts in a subnet in the same region and same network.
	// There must be at least one IP address available in the subnet's primary range. The subnet is specified in the following form:
	// `projects/{projectNumber}/regions/{region_id}/subnetworks/{subnetwork_id}
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`
}

type ConsumersParameters struct {

	// The subnetwork of the customer project from which an IP address is reserved and used as the Dataproc Metastore service's endpoint.
	// It is accessible to hosts in the subnet and to all hosts in a subnet in the same region and same network.
	// There must be at least one IP address available in the subnet's primary range. The subnet is specified in the following form:
	// `projects/{projectNumber}/regions/{region_id}/subnetworks/{subnetwork_id}
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/compute/v1beta1.Subnetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// Reference to a Subnetwork in compute to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// Selector for a Subnetwork in compute to populate subnetwork.
	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`
}

type DataCatalogConfigInitParameters struct {

	// Defines whether the metastore metadata should be synced to Data Catalog. The default value is to disable syncing metastore metadata to Data Catalog.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DataCatalogConfigObservation struct {

	// Defines whether the metastore metadata should be synced to Data Catalog. The default value is to disable syncing metastore metadata to Data Catalog.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DataCatalogConfigParameters struct {

	// Defines whether the metastore metadata should be synced to Data Catalog. The default value is to disable syncing metastore metadata to Data Catalog.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type HiveMetastoreConfigInitParameters struct {

	// A mapping of Hive metastore version to the auxiliary version configuration.
	// When specified, a secondary Hive metastore service is created along with the primary service.
	// All auxiliary versions must be less than the service's primary version.
	// The key is the auxiliary service name and it must match the regular expression a-z?.
	// This means that the first character must be a lowercase letter, and all the following characters must be hyphens, lowercase letters, or digits, except the last character, which cannot be a hyphen.
	// Structure is documented below.
	AuxiliaryVersions []AuxiliaryVersionsInitParameters `json:"auxiliaryVersions,omitempty" tf:"auxiliary_versions,omitempty"`

	// A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml).
	// The mappings override system defaults (some keys cannot be overridden)
	// +mapType=granular
	ConfigOverrides map[string]*string `json:"configOverrides,omitempty" tf:"config_overrides,omitempty"`

	// The protocol to use for the metastore service endpoint. If unspecified, defaults to THRIFT.
	// Default value is THRIFT.
	// Possible values are: THRIFT, GRPC.
	EndpointProtocol *string `json:"endpointProtocol,omitempty" tf:"endpoint_protocol,omitempty"`

	// Information used to configure the Hive metastore service as a service principal in a Kerberos realm.
	// Structure is documented below.
	KerberosConfig []HiveMetastoreConfigKerberosConfigInitParameters `json:"kerberosConfig,omitempty" tf:"kerberos_config,omitempty"`

	// The Hive metastore schema version.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type HiveMetastoreConfigKerberosConfigInitParameters struct {

	// A Kerberos keytab file that can be used to authenticate a service principal with a Kerberos Key Distribution Center (KDC).
	// Structure is documented below.
	Keytab []KeytabInitParameters `json:"keytab,omitempty" tf:"keytab,omitempty"`

	// A Cloud Storage URI that specifies the path to a krb5.conf file. It is of the form gs://{bucket_name}/path/to/krb5.conf, although the file does not need to be named krb5.conf explicitly.
	Krb5ConfigGcsURI *string `json:"krb5ConfigGcsUri,omitempty" tf:"krb5_config_gcs_uri,omitempty"`

	// A Kerberos principal that exists in the both the keytab the KDC to authenticate as. A typical principal is of the form "primary/instance@REALM", but there is no exact format.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`
}

type HiveMetastoreConfigKerberosConfigObservation struct {

	// A Kerberos keytab file that can be used to authenticate a service principal with a Kerberos Key Distribution Center (KDC).
	// Structure is documented below.
	Keytab []KeytabObservation `json:"keytab,omitempty" tf:"keytab,omitempty"`

	// A Cloud Storage URI that specifies the path to a krb5.conf file. It is of the form gs://{bucket_name}/path/to/krb5.conf, although the file does not need to be named krb5.conf explicitly.
	Krb5ConfigGcsURI *string `json:"krb5ConfigGcsUri,omitempty" tf:"krb5_config_gcs_uri,omitempty"`

	// A Kerberos principal that exists in the both the keytab the KDC to authenticate as. A typical principal is of the form "primary/instance@REALM", but there is no exact format.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`
}

type HiveMetastoreConfigKerberosConfigParameters struct {

	// A Kerberos keytab file that can be used to authenticate a service principal with a Kerberos Key Distribution Center (KDC).
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Keytab []KeytabParameters `json:"keytab" tf:"keytab,omitempty"`

	// A Cloud Storage URI that specifies the path to a krb5.conf file. It is of the form gs://{bucket_name}/path/to/krb5.conf, although the file does not need to be named krb5.conf explicitly.
	// +kubebuilder:validation:Optional
	Krb5ConfigGcsURI *string `json:"krb5ConfigGcsUri" tf:"krb5_config_gcs_uri,omitempty"`

	// A Kerberos principal that exists in the both the keytab the KDC to authenticate as. A typical principal is of the form "primary/instance@REALM", but there is no exact format.
	// +kubebuilder:validation:Optional
	Principal *string `json:"principal" tf:"principal,omitempty"`
}

type HiveMetastoreConfigObservation struct {

	// A mapping of Hive metastore version to the auxiliary version configuration.
	// When specified, a secondary Hive metastore service is created along with the primary service.
	// All auxiliary versions must be less than the service's primary version.
	// The key is the auxiliary service name and it must match the regular expression a-z?.
	// This means that the first character must be a lowercase letter, and all the following characters must be hyphens, lowercase letters, or digits, except the last character, which cannot be a hyphen.
	// Structure is documented below.
	AuxiliaryVersions []AuxiliaryVersionsObservation `json:"auxiliaryVersions,omitempty" tf:"auxiliary_versions,omitempty"`

	// A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml).
	// The mappings override system defaults (some keys cannot be overridden)
	// +mapType=granular
	ConfigOverrides map[string]*string `json:"configOverrides,omitempty" tf:"config_overrides,omitempty"`

	// The protocol to use for the metastore service endpoint. If unspecified, defaults to THRIFT.
	// Default value is THRIFT.
	// Possible values are: THRIFT, GRPC.
	EndpointProtocol *string `json:"endpointProtocol,omitempty" tf:"endpoint_protocol,omitempty"`

	// Information used to configure the Hive metastore service as a service principal in a Kerberos realm.
	// Structure is documented below.
	KerberosConfig []HiveMetastoreConfigKerberosConfigObservation `json:"kerberosConfig,omitempty" tf:"kerberos_config,omitempty"`

	// The Hive metastore schema version.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type HiveMetastoreConfigParameters struct {

	// A mapping of Hive metastore version to the auxiliary version configuration.
	// When specified, a secondary Hive metastore service is created along with the primary service.
	// All auxiliary versions must be less than the service's primary version.
	// The key is the auxiliary service name and it must match the regular expression a-z?.
	// This means that the first character must be a lowercase letter, and all the following characters must be hyphens, lowercase letters, or digits, except the last character, which cannot be a hyphen.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	AuxiliaryVersions []AuxiliaryVersionsParameters `json:"auxiliaryVersions,omitempty" tf:"auxiliary_versions,omitempty"`

	// A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml).
	// The mappings override system defaults (some keys cannot be overridden)
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ConfigOverrides map[string]*string `json:"configOverrides,omitempty" tf:"config_overrides,omitempty"`

	// The protocol to use for the metastore service endpoint. If unspecified, defaults to THRIFT.
	// Default value is THRIFT.
	// Possible values are: THRIFT, GRPC.
	// +kubebuilder:validation:Optional
	EndpointProtocol *string `json:"endpointProtocol,omitempty" tf:"endpoint_protocol,omitempty"`

	// Information used to configure the Hive metastore service as a service principal in a Kerberos realm.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	KerberosConfig []HiveMetastoreConfigKerberosConfigParameters `json:"kerberosConfig,omitempty" tf:"kerberos_config,omitempty"`

	// The Hive metastore schema version.
	// +kubebuilder:validation:Optional
	Version *string `json:"version" tf:"version,omitempty"`
}

type KeytabInitParameters struct {

	// The relative resource name of a Secret Manager secret version, in the following form:
	// "projects/{projectNumber}/secrets/{secret_id}/versions/{version_id}".
	CloudSecret *string `json:"cloudSecret,omitempty" tf:"cloud_secret,omitempty"`
}

type KeytabObservation struct {

	// The relative resource name of a Secret Manager secret version, in the following form:
	// "projects/{projectNumber}/secrets/{secret_id}/versions/{version_id}".
	CloudSecret *string `json:"cloudSecret,omitempty" tf:"cloud_secret,omitempty"`
}

type KeytabParameters struct {

	// The relative resource name of a Secret Manager secret version, in the following form:
	// "projects/{projectNumber}/secrets/{secret_id}/versions/{version_id}".
	// +kubebuilder:validation:Optional
	CloudSecret *string `json:"cloudSecret" tf:"cloud_secret,omitempty"`
}

type MaintenanceWindowInitParameters struct {

	// The day of week, when the window starts.
	// Possible values are: MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY.
	DayOfWeek *string `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// The hour of day (0-23) when the window starts.
	HourOfDay *float64 `json:"hourOfDay,omitempty" tf:"hour_of_day,omitempty"`
}

type MaintenanceWindowObservation struct {

	// The day of week, when the window starts.
	// Possible values are: MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY.
	DayOfWeek *string `json:"dayOfWeek,omitempty" tf:"day_of_week,omitempty"`

	// The hour of day (0-23) when the window starts.
	HourOfDay *float64 `json:"hourOfDay,omitempty" tf:"hour_of_day,omitempty"`
}

type MaintenanceWindowParameters struct {

	// The day of week, when the window starts.
	// Possible values are: MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY.
	// +kubebuilder:validation:Optional
	DayOfWeek *string `json:"dayOfWeek" tf:"day_of_week,omitempty"`

	// The hour of day (0-23) when the window starts.
	// +kubebuilder:validation:Optional
	HourOfDay *float64 `json:"hourOfDay" tf:"hour_of_day,omitempty"`
}

type MetadataIntegrationInitParameters struct {

	// The integration config for the Data Catalog service.
	// Structure is documented below.
	DataCatalogConfig []DataCatalogConfigInitParameters `json:"dataCatalogConfig,omitempty" tf:"data_catalog_config,omitempty"`
}

type MetadataIntegrationObservation struct {

	// The integration config for the Data Catalog service.
	// Structure is documented below.
	DataCatalogConfig []DataCatalogConfigObservation `json:"dataCatalogConfig,omitempty" tf:"data_catalog_config,omitempty"`
}

type MetadataIntegrationParameters struct {

	// The integration config for the Data Catalog service.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DataCatalogConfig []DataCatalogConfigParameters `json:"dataCatalogConfig" tf:"data_catalog_config,omitempty"`
}

type MetastoreServiceEncryptionConfigInitParameters struct {

	// The fully qualified customer provided Cloud KMS key name to use for customer data encryption.
	// Use the following format: projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// Reference to a CryptoKey in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeyRef *v1.Reference `json:"kmsKeyRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeySelector *v1.Selector `json:"kmsKeySelector,omitempty" tf:"-"`
}

type MetastoreServiceEncryptionConfigObservation struct {

	// The fully qualified customer provided Cloud KMS key name to use for customer data encryption.
	// Use the following format: projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`
}

type MetastoreServiceEncryptionConfigParameters struct {

	// The fully qualified customer provided Cloud KMS key name to use for customer data encryption.
	// Use the following format: projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/kms/v1beta1.CryptoKey
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// Reference to a CryptoKey in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeyRef *v1.Reference `json:"kmsKeyRef,omitempty" tf:"-"`

	// Selector for a CryptoKey in kms to populate kmsKey.
	// +kubebuilder:validation:Optional
	KMSKeySelector *v1.Selector `json:"kmsKeySelector,omitempty" tf:"-"`
}

type MetastoreServiceInitParameters struct {

	// The database type that the Metastore service stores its data.
	// Default value is MYSQL.
	// Possible values are: MYSQL, SPANNER.
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type,omitempty"`

	// Information used to configure the Dataproc Metastore service to encrypt
	// customer data at rest.
	// Structure is documented below.
	EncryptionConfig []MetastoreServiceEncryptionConfigInitParameters `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// Configuration information specific to running Hive metastore software as the metastore service.
	// Structure is documented below.
	HiveMetastoreConfig []HiveMetastoreConfigInitParameters `json:"hiveMetastoreConfig,omitempty" tf:"hive_metastore_config,omitempty"`

	// User-defined labels for the metastore service.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The one hour maintenance window of the metastore service.
	// This specifies when the service can be restarted for maintenance purposes in UTC time.
	// Maintenance window is not needed for services with the SPANNER database type.
	// Structure is documented below.
	MaintenanceWindow []MaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The setting that defines how metastore metadata should be integrated with external services and systems.
	// Structure is documented below.
	MetadataIntegration []MetadataIntegrationInitParameters `json:"metadataIntegration,omitempty" tf:"metadata_integration,omitempty"`

	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
	// "projects/{projectNumber}/global/networks/{network_id}".
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The configuration specifying the network settings for the Dataproc Metastore service.
	// Structure is documented below.
	NetworkConfig []NetworkConfigInitParameters `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// The TCP port at which the metastore service is reached. Default: 9083.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The release channel of the service. If unspecified, defaults to STABLE.
	// Default value is STABLE.
	// Possible values are: CANARY, STABLE.
	ReleaseChannel *string `json:"releaseChannel,omitempty" tf:"release_channel,omitempty"`

	// Represents the scaling configuration of a metastore service.
	// Structure is documented below.
	ScalingConfig []ScalingConfigInitParameters `json:"scalingConfig,omitempty" tf:"scaling_config,omitempty"`

	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	// Structure is documented below.
	TelemetryConfig []TelemetryConfigInitParameters `json:"telemetryConfig,omitempty" tf:"telemetry_config,omitempty"`

	// The tier of the service.
	// Possible values are: DEVELOPER, ENTERPRISE.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type MetastoreServiceObservation struct {

	// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
	ArtifactGcsURI *string `json:"artifactGcsUri,omitempty" tf:"artifact_gcs_uri,omitempty"`

	// The database type that the Metastore service stores its data.
	// Default value is MYSQL.
	// Possible values are: MYSQL, SPANNER.
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type,omitempty"`

	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// Information used to configure the Dataproc Metastore service to encrypt
	// customer data at rest.
	// Structure is documented below.
	EncryptionConfig []MetastoreServiceEncryptionConfigObservation `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// The URI of the endpoint used to access the metastore service.
	EndpointURI *string `json:"endpointUri,omitempty" tf:"endpoint_uri,omitempty"`

	// Configuration information specific to running Hive metastore software as the metastore service.
	// Structure is documented below.
	HiveMetastoreConfig []HiveMetastoreConfigObservation `json:"hiveMetastoreConfig,omitempty" tf:"hive_metastore_config,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/services/{{service_id}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// User-defined labels for the metastore service.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location where the metastore service should reside.
	// The default value is global.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The one hour maintenance window of the metastore service.
	// This specifies when the service can be restarted for maintenance purposes in UTC time.
	// Maintenance window is not needed for services with the SPANNER database type.
	// Structure is documented below.
	MaintenanceWindow []MaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The setting that defines how metastore metadata should be integrated with external services and systems.
	// Structure is documented below.
	MetadataIntegration []MetadataIntegrationObservation `json:"metadataIntegration,omitempty" tf:"metadata_integration,omitempty"`

	// The relative resource name of the metastore service.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
	// "projects/{projectNumber}/global/networks/{network_id}".
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The configuration specifying the network settings for the Dataproc Metastore service.
	// Structure is documented below.
	NetworkConfig []NetworkConfigObservation `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// The TCP port at which the metastore service is reached. Default: 9083.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The release channel of the service. If unspecified, defaults to STABLE.
	// Default value is STABLE.
	// Possible values are: CANARY, STABLE.
	ReleaseChannel *string `json:"releaseChannel,omitempty" tf:"release_channel,omitempty"`

	// Represents the scaling configuration of a metastore service.
	// Structure is documented below.
	ScalingConfig []ScalingConfigObservation `json:"scalingConfig,omitempty" tf:"scaling_config,omitempty"`

	// The current state of the metastore service.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Additional information about the current state of the metastore service, if available.
	StateMessage *string `json:"stateMessage,omitempty" tf:"state_message,omitempty"`

	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	// Structure is documented below.
	TelemetryConfig []TelemetryConfigObservation `json:"telemetryConfig,omitempty" tf:"telemetry_config,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// The tier of the service.
	// Possible values are: DEVELOPER, ENTERPRISE.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`

	// The globally unique resource identifier of the metastore service.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type MetastoreServiceParameters struct {

	// The database type that the Metastore service stores its data.
	// Default value is MYSQL.
	// Possible values are: MYSQL, SPANNER.
	// +kubebuilder:validation:Optional
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type,omitempty"`

	// Information used to configure the Dataproc Metastore service to encrypt
	// customer data at rest.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	EncryptionConfig []MetastoreServiceEncryptionConfigParameters `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// Configuration information specific to running Hive metastore software as the metastore service.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	HiveMetastoreConfig []HiveMetastoreConfigParameters `json:"hiveMetastoreConfig,omitempty" tf:"hive_metastore_config,omitempty"`

	// User-defined labels for the metastore service.
	// Note: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field effective_labels for all of the labels present on the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// The location where the metastore service should reside.
	// The default value is global.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The one hour maintenance window of the metastore service.
	// This specifies when the service can be restarted for maintenance purposes in UTC time.
	// Maintenance window is not needed for services with the SPANNER database type.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The setting that defines how metastore metadata should be integrated with external services and systems.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	MetadataIntegration []MetadataIntegrationParameters `json:"metadataIntegration,omitempty" tf:"metadata_integration,omitempty"`

	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
	// "projects/{projectNumber}/global/networks/{network_id}".
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// The configuration specifying the network settings for the Dataproc Metastore service.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	NetworkConfig []NetworkConfigParameters `json:"networkConfig,omitempty" tf:"network_config,omitempty"`

	// The TCP port at which the metastore service is reached. Default: 9083.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The release channel of the service. If unspecified, defaults to STABLE.
	// Default value is STABLE.
	// Possible values are: CANARY, STABLE.
	// +kubebuilder:validation:Optional
	ReleaseChannel *string `json:"releaseChannel,omitempty" tf:"release_channel,omitempty"`

	// Represents the scaling configuration of a metastore service.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ScalingConfig []ScalingConfigParameters `json:"scalingConfig,omitempty" tf:"scaling_config,omitempty"`

	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TelemetryConfig []TelemetryConfigParameters `json:"telemetryConfig,omitempty" tf:"telemetry_config,omitempty"`

	// The tier of the service.
	// Possible values are: DEVELOPER, ENTERPRISE.
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type NetworkConfigInitParameters struct {

	// The consumer-side network configuration for the Dataproc Metastore instance.
	// Structure is documented below.
	Consumers []ConsumersInitParameters `json:"consumers,omitempty" tf:"consumers,omitempty"`
}

type NetworkConfigObservation struct {

	// The consumer-side network configuration for the Dataproc Metastore instance.
	// Structure is documented below.
	Consumers []ConsumersObservation `json:"consumers,omitempty" tf:"consumers,omitempty"`
}

type NetworkConfigParameters struct {

	// The consumer-side network configuration for the Dataproc Metastore instance.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Consumers []ConsumersParameters `json:"consumers" tf:"consumers,omitempty"`
}

type ScalingConfigInitParameters struct {

	// Metastore instance sizes.
	// Possible values are: EXTRA_SMALL, SMALL, MEDIUM, LARGE, EXTRA_LARGE.
	InstanceSize *string `json:"instanceSize,omitempty" tf:"instance_size,omitempty"`

	// Scaling factor, in increments of 0.1 for values less than 1.0, and increments of 1.0 for values greater than 1.0.
	ScalingFactor *float64 `json:"scalingFactor,omitempty" tf:"scaling_factor,omitempty"`
}

type ScalingConfigObservation struct {

	// Metastore instance sizes.
	// Possible values are: EXTRA_SMALL, SMALL, MEDIUM, LARGE, EXTRA_LARGE.
	InstanceSize *string `json:"instanceSize,omitempty" tf:"instance_size,omitempty"`

	// Scaling factor, in increments of 0.1 for values less than 1.0, and increments of 1.0 for values greater than 1.0.
	ScalingFactor *float64 `json:"scalingFactor,omitempty" tf:"scaling_factor,omitempty"`
}

type ScalingConfigParameters struct {

	// Metastore instance sizes.
	// Possible values are: EXTRA_SMALL, SMALL, MEDIUM, LARGE, EXTRA_LARGE.
	// +kubebuilder:validation:Optional
	InstanceSize *string `json:"instanceSize,omitempty" tf:"instance_size,omitempty"`

	// Scaling factor, in increments of 0.1 for values less than 1.0, and increments of 1.0 for values greater than 1.0.
	// +kubebuilder:validation:Optional
	ScalingFactor *float64 `json:"scalingFactor,omitempty" tf:"scaling_factor,omitempty"`
}

type TelemetryConfigInitParameters struct {

	// The output format of the Dataproc Metastore service's logs.
	// Default value is JSON.
	// Possible values are: LEGACY, JSON.
	LogFormat *string `json:"logFormat,omitempty" tf:"log_format,omitempty"`
}

type TelemetryConfigObservation struct {

	// The output format of the Dataproc Metastore service's logs.
	// Default value is JSON.
	// Possible values are: LEGACY, JSON.
	LogFormat *string `json:"logFormat,omitempty" tf:"log_format,omitempty"`
}

type TelemetryConfigParameters struct {

	// The output format of the Dataproc Metastore service's logs.
	// Default value is JSON.
	// Possible values are: LEGACY, JSON.
	// +kubebuilder:validation:Optional
	LogFormat *string `json:"logFormat,omitempty" tf:"log_format,omitempty"`
}

// MetastoreServiceSpec defines the desired state of MetastoreService
type MetastoreServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetastoreServiceParameters `json:"forProvider"`
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
	InitProvider MetastoreServiceInitParameters `json:"initProvider,omitempty"`
}

// MetastoreServiceStatus defines the observed state of MetastoreService.
type MetastoreServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetastoreServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MetastoreService is the Schema for the MetastoreServices API. A managed metastore service that serves metadata queries.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type MetastoreService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MetastoreServiceSpec   `json:"spec"`
	Status            MetastoreServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetastoreServiceList contains a list of MetastoreServices
type MetastoreServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetastoreService `json:"items"`
}

// Repository type metadata.
var (
	MetastoreService_Kind             = "MetastoreService"
	MetastoreService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MetastoreService_Kind}.String()
	MetastoreService_KindAPIVersion   = MetastoreService_Kind + "." + CRDGroupVersion.String()
	MetastoreService_GroupVersionKind = CRDGroupVersion.WithKind(MetastoreService_Kind)
)

func init() {
	SchemeBuilder.Register(&MetastoreService{}, &MetastoreServiceList{})
}
