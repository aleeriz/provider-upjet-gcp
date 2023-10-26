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

type DailyScheduleInitParameters struct {

	// Defines a schedule with units measured in days. The value determines how many days pass between the start of each cycle. Days in cycle for snapshot schedule policy must be 1.
	DaysInCycle *float64 `json:"daysInCycle,omitempty" tf:"days_in_cycle,omitempty"`

	// Time within the window to start the operations.
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type DailyScheduleObservation struct {

	// Defines a schedule with units measured in days. The value determines how many days pass between the start of each cycle. Days in cycle for snapshot schedule policy must be 1.
	DaysInCycle *float64 `json:"daysInCycle,omitempty" tf:"days_in_cycle,omitempty"`

	// Time within the window to start the operations.
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type DailyScheduleParameters struct {

	// Defines a schedule with units measured in days. The value determines how many days pass between the start of each cycle. Days in cycle for snapshot schedule policy must be 1.
	// +kubebuilder:validation:Optional
	DaysInCycle *float64 `json:"daysInCycle" tf:"days_in_cycle,omitempty"`

	// Time within the window to start the operations.
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type DayOfWeeksInitParameters struct {

	// The day of the week to create the snapshot. e.g. MONDAY
	// Possible values are: MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Time within the window to start the operations.
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type DayOfWeeksObservation struct {

	// The day of the week to create the snapshot. e.g. MONDAY
	// Possible values are: MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Time within the window to start the operations.
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type DayOfWeeksParameters struct {

	// The day of the week to create the snapshot. e.g. MONDAY
	// Possible values are: MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY.
	// +kubebuilder:validation:Optional
	Day *string `json:"day" tf:"day,omitempty"`

	// Time within the window to start the operations.
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type DiskConsistencyGroupPolicyInitParameters struct {

	// Enable disk consistency on the resource policy.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DiskConsistencyGroupPolicyObservation struct {

	// Enable disk consistency on the resource policy.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type DiskConsistencyGroupPolicyParameters struct {

	// Enable disk consistency on the resource policy.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type GroupPlacementPolicyInitParameters struct {

	// The number of availability domains instances will be spread across. If two instances are in different
	// availability domain, they will not be put in the same low latency network
	AvailabilityDomainCount *float64 `json:"availabilityDomainCount,omitempty" tf:"availability_domain_count,omitempty"`

	// Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.
	// Specify COLLOCATED to enable collocation. Can only be specified with vm_count. If compute instances are created
	// with a COLLOCATED policy, then exactly vm_count instances must be created at the same time with the resource policy
	// attached.
	// Possible values are: COLLOCATED.
	Collocation *string `json:"collocation,omitempty" tf:"collocation,omitempty"`

	// Number of VMs in this placement group. Google does not recommend that you use this field
	// unless you use a compact policy and you want your policy to work only if it contains this
	// exact number of VMs.
	VMCount *float64 `json:"vmCount,omitempty" tf:"vm_count,omitempty"`
}

type GroupPlacementPolicyObservation struct {

	// The number of availability domains instances will be spread across. If two instances are in different
	// availability domain, they will not be put in the same low latency network
	AvailabilityDomainCount *float64 `json:"availabilityDomainCount,omitempty" tf:"availability_domain_count,omitempty"`

	// Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.
	// Specify COLLOCATED to enable collocation. Can only be specified with vm_count. If compute instances are created
	// with a COLLOCATED policy, then exactly vm_count instances must be created at the same time with the resource policy
	// attached.
	// Possible values are: COLLOCATED.
	Collocation *string `json:"collocation,omitempty" tf:"collocation,omitempty"`

	// Number of VMs in this placement group. Google does not recommend that you use this field
	// unless you use a compact policy and you want your policy to work only if it contains this
	// exact number of VMs.
	VMCount *float64 `json:"vmCount,omitempty" tf:"vm_count,omitempty"`
}

type GroupPlacementPolicyParameters struct {

	// The number of availability domains instances will be spread across. If two instances are in different
	// availability domain, they will not be put in the same low latency network
	// +kubebuilder:validation:Optional
	AvailabilityDomainCount *float64 `json:"availabilityDomainCount,omitempty" tf:"availability_domain_count,omitempty"`

	// Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.
	// Specify COLLOCATED to enable collocation. Can only be specified with vm_count. If compute instances are created
	// with a COLLOCATED policy, then exactly vm_count instances must be created at the same time with the resource policy
	// attached.
	// Possible values are: COLLOCATED.
	// +kubebuilder:validation:Optional
	Collocation *string `json:"collocation,omitempty" tf:"collocation,omitempty"`

	// Number of VMs in this placement group. Google does not recommend that you use this field
	// unless you use a compact policy and you want your policy to work only if it contains this
	// exact number of VMs.
	// +kubebuilder:validation:Optional
	VMCount *float64 `json:"vmCount,omitempty" tf:"vm_count,omitempty"`
}

type HourlyScheduleInitParameters struct {

	// The number of hours between snapshots.
	HoursInCycle *float64 `json:"hoursInCycle,omitempty" tf:"hours_in_cycle,omitempty"`

	// Time within the window to start the operations.
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type HourlyScheduleObservation struct {

	// The number of hours between snapshots.
	HoursInCycle *float64 `json:"hoursInCycle,omitempty" tf:"hours_in_cycle,omitempty"`

	// Time within the window to start the operations.
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type HourlyScheduleParameters struct {

	// The number of hours between snapshots.
	// +kubebuilder:validation:Optional
	HoursInCycle *float64 `json:"hoursInCycle" tf:"hours_in_cycle,omitempty"`

	// Time within the window to start the operations.
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type InstanceSchedulePolicyInitParameters struct {

	// The expiration time of the schedule. The timestamp is an RFC3339 string.
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// The start time of the schedule. The timestamp is an RFC3339 string.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Specifies the time zone to be used in interpreting the schedule. The value of this field must be a time zone name
	// from the tz database: http://en.wikipedia.org/wiki/Tz_database.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// Specifies the schedule for starting instances.
	// Structure is documented below.
	VMStartSchedule []VMStartScheduleInitParameters `json:"vmStartSchedule,omitempty" tf:"vm_start_schedule,omitempty"`

	// Specifies the schedule for stopping instances.
	// Structure is documented below.
	VMStopSchedule []VMStopScheduleInitParameters `json:"vmStopSchedule,omitempty" tf:"vm_stop_schedule,omitempty"`
}

type InstanceSchedulePolicyObservation struct {

	// The expiration time of the schedule. The timestamp is an RFC3339 string.
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// The start time of the schedule. The timestamp is an RFC3339 string.
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Specifies the time zone to be used in interpreting the schedule. The value of this field must be a time zone name
	// from the tz database: http://en.wikipedia.org/wiki/Tz_database.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`

	// Specifies the schedule for starting instances.
	// Structure is documented below.
	VMStartSchedule []VMStartScheduleObservation `json:"vmStartSchedule,omitempty" tf:"vm_start_schedule,omitempty"`

	// Specifies the schedule for stopping instances.
	// Structure is documented below.
	VMStopSchedule []VMStopScheduleObservation `json:"vmStopSchedule,omitempty" tf:"vm_stop_schedule,omitempty"`
}

type InstanceSchedulePolicyParameters struct {

	// The expiration time of the schedule. The timestamp is an RFC3339 string.
	// +kubebuilder:validation:Optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// The start time of the schedule. The timestamp is an RFC3339 string.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Specifies the time zone to be used in interpreting the schedule. The value of this field must be a time zone name
	// from the tz database: http://en.wikipedia.org/wiki/Tz_database.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone" tf:"time_zone,omitempty"`

	// Specifies the schedule for starting instances.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VMStartSchedule []VMStartScheduleParameters `json:"vmStartSchedule,omitempty" tf:"vm_start_schedule,omitempty"`

	// Specifies the schedule for stopping instances.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VMStopSchedule []VMStopScheduleParameters `json:"vmStopSchedule,omitempty" tf:"vm_stop_schedule,omitempty"`
}

type ResourcePolicyInitParameters struct {

	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Replication consistency group for asynchronous disk replication.
	// Structure is documented below.
	DiskConsistencyGroupPolicy []DiskConsistencyGroupPolicyInitParameters `json:"diskConsistencyGroupPolicy,omitempty" tf:"disk_consistency_group_policy,omitempty"`

	// Resource policy for instances used for placement configuration.
	// Structure is documented below.
	GroupPlacementPolicy []GroupPlacementPolicyInitParameters `json:"groupPlacementPolicy,omitempty" tf:"group_placement_policy,omitempty"`

	// Resource policy for scheduling instance operations.
	// Structure is documented below.
	InstanceSchedulePolicy []InstanceSchedulePolicyInitParameters `json:"instanceSchedulePolicy,omitempty" tf:"instance_schedule_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy []SnapshotSchedulePolicyInitParameters `json:"snapshotSchedulePolicy,omitempty" tf:"snapshot_schedule_policy,omitempty"`
}

type ResourcePolicyObservation struct {

	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Replication consistency group for asynchronous disk replication.
	// Structure is documented below.
	DiskConsistencyGroupPolicy []DiskConsistencyGroupPolicyObservation `json:"diskConsistencyGroupPolicy,omitempty" tf:"disk_consistency_group_policy,omitempty"`

	// Resource policy for instances used for placement configuration.
	// Structure is documented below.
	GroupPlacementPolicy []GroupPlacementPolicyObservation `json:"groupPlacementPolicy,omitempty" tf:"group_placement_policy,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource policy for scheduling instance operations.
	// Structure is documented below.
	InstanceSchedulePolicy []InstanceSchedulePolicyObservation `json:"instanceSchedulePolicy,omitempty" tf:"instance_schedule_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where resource policy resides.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	SnapshotSchedulePolicy []SnapshotSchedulePolicyObservation `json:"snapshotSchedulePolicy,omitempty" tf:"snapshot_schedule_policy,omitempty"`
}

type ResourcePolicyParameters struct {

	// An optional description of this resource. Provide this property when you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Replication consistency group for asynchronous disk replication.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DiskConsistencyGroupPolicy []DiskConsistencyGroupPolicyParameters `json:"diskConsistencyGroupPolicy,omitempty" tf:"disk_consistency_group_policy,omitempty"`

	// Resource policy for instances used for placement configuration.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GroupPlacementPolicy []GroupPlacementPolicyParameters `json:"groupPlacementPolicy,omitempty" tf:"group_placement_policy,omitempty"`

	// Resource policy for scheduling instance operations.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	InstanceSchedulePolicy []InstanceSchedulePolicyParameters `json:"instanceSchedulePolicy,omitempty" tf:"instance_schedule_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where resource policy resides.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SnapshotSchedulePolicy []SnapshotSchedulePolicyParameters `json:"snapshotSchedulePolicy,omitempty" tf:"snapshot_schedule_policy,omitempty"`
}

type RetentionPolicyInitParameters struct {

	// Maximum age of the snapshot that is allowed to be kept.
	MaxRetentionDays *float64 `json:"maxRetentionDays,omitempty" tf:"max_retention_days,omitempty"`

	// Specifies the behavior to apply to scheduled snapshots when
	// the source disk is deleted.
	// Default value is KEEP_AUTO_SNAPSHOTS.
	// Possible values are: KEEP_AUTO_SNAPSHOTS, APPLY_RETENTION_POLICY.
	OnSourceDiskDelete *string `json:"onSourceDiskDelete,omitempty" tf:"on_source_disk_delete,omitempty"`
}

type RetentionPolicyObservation struct {

	// Maximum age of the snapshot that is allowed to be kept.
	MaxRetentionDays *float64 `json:"maxRetentionDays,omitempty" tf:"max_retention_days,omitempty"`

	// Specifies the behavior to apply to scheduled snapshots when
	// the source disk is deleted.
	// Default value is KEEP_AUTO_SNAPSHOTS.
	// Possible values are: KEEP_AUTO_SNAPSHOTS, APPLY_RETENTION_POLICY.
	OnSourceDiskDelete *string `json:"onSourceDiskDelete,omitempty" tf:"on_source_disk_delete,omitempty"`
}

type RetentionPolicyParameters struct {

	// Maximum age of the snapshot that is allowed to be kept.
	// +kubebuilder:validation:Optional
	MaxRetentionDays *float64 `json:"maxRetentionDays" tf:"max_retention_days,omitempty"`

	// Specifies the behavior to apply to scheduled snapshots when
	// the source disk is deleted.
	// Default value is KEEP_AUTO_SNAPSHOTS.
	// Possible values are: KEEP_AUTO_SNAPSHOTS, APPLY_RETENTION_POLICY.
	// +kubebuilder:validation:Optional
	OnSourceDiskDelete *string `json:"onSourceDiskDelete,omitempty" tf:"on_source_disk_delete,omitempty"`
}

type ScheduleInitParameters struct {

	// The policy will execute every nth day at the specified time.
	// Structure is documented below.
	DailySchedule []DailyScheduleInitParameters `json:"dailySchedule,omitempty" tf:"daily_schedule,omitempty"`

	// The policy will execute every nth hour starting at the specified time.
	// Structure is documented below.
	HourlySchedule []HourlyScheduleInitParameters `json:"hourlySchedule,omitempty" tf:"hourly_schedule,omitempty"`

	// Allows specifying a snapshot time for each day of the week.
	// Structure is documented below.
	WeeklySchedule []WeeklyScheduleInitParameters `json:"weeklySchedule,omitempty" tf:"weekly_schedule,omitempty"`
}

type ScheduleObservation struct {

	// The policy will execute every nth day at the specified time.
	// Structure is documented below.
	DailySchedule []DailyScheduleObservation `json:"dailySchedule,omitempty" tf:"daily_schedule,omitempty"`

	// The policy will execute every nth hour starting at the specified time.
	// Structure is documented below.
	HourlySchedule []HourlyScheduleObservation `json:"hourlySchedule,omitempty" tf:"hourly_schedule,omitempty"`

	// Allows specifying a snapshot time for each day of the week.
	// Structure is documented below.
	WeeklySchedule []WeeklyScheduleObservation `json:"weeklySchedule,omitempty" tf:"weekly_schedule,omitempty"`
}

type ScheduleParameters struct {

	// The policy will execute every nth day at the specified time.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DailySchedule []DailyScheduleParameters `json:"dailySchedule,omitempty" tf:"daily_schedule,omitempty"`

	// The policy will execute every nth hour starting at the specified time.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	HourlySchedule []HourlyScheduleParameters `json:"hourlySchedule,omitempty" tf:"hourly_schedule,omitempty"`

	// Allows specifying a snapshot time for each day of the week.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	WeeklySchedule []WeeklyScheduleParameters `json:"weeklySchedule,omitempty" tf:"weekly_schedule,omitempty"`
}

type SnapshotPropertiesInitParameters struct {

	// Creates the new snapshot in the snapshot chain labeled with the
	// specified name. The chain name must be 1-63 characters long and comply
	// with RFC1035.
	ChainName *string `json:"chainName,omitempty" tf:"chain_name,omitempty"`

	// Whether to perform a 'guest aware' snapshot.
	GuestFlush *bool `json:"guestFlush,omitempty" tf:"guest_flush,omitempty"`

	// A set of key-value pairs.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Cloud Storage bucket location to store the auto snapshot
	// (regional or multi-regional)
	StorageLocations []*string `json:"storageLocations,omitempty" tf:"storage_locations,omitempty"`
}

type SnapshotPropertiesObservation struct {

	// Creates the new snapshot in the snapshot chain labeled with the
	// specified name. The chain name must be 1-63 characters long and comply
	// with RFC1035.
	ChainName *string `json:"chainName,omitempty" tf:"chain_name,omitempty"`

	// Whether to perform a 'guest aware' snapshot.
	GuestFlush *bool `json:"guestFlush,omitempty" tf:"guest_flush,omitempty"`

	// A set of key-value pairs.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Cloud Storage bucket location to store the auto snapshot
	// (regional or multi-regional)
	StorageLocations []*string `json:"storageLocations,omitempty" tf:"storage_locations,omitempty"`
}

type SnapshotPropertiesParameters struct {

	// Creates the new snapshot in the snapshot chain labeled with the
	// specified name. The chain name must be 1-63 characters long and comply
	// with RFC1035.
	// +kubebuilder:validation:Optional
	ChainName *string `json:"chainName,omitempty" tf:"chain_name,omitempty"`

	// Whether to perform a 'guest aware' snapshot.
	// +kubebuilder:validation:Optional
	GuestFlush *bool `json:"guestFlush,omitempty" tf:"guest_flush,omitempty"`

	// A set of key-value pairs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Cloud Storage bucket location to store the auto snapshot
	// (regional or multi-regional)
	// +kubebuilder:validation:Optional
	StorageLocations []*string `json:"storageLocations,omitempty" tf:"storage_locations,omitempty"`
}

type SnapshotSchedulePolicyInitParameters struct {

	// Retention policy applied to snapshots created by this resource policy.
	// Structure is documented below.
	RetentionPolicy []RetentionPolicyInitParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Contains one of an hourlySchedule, dailySchedule, or weeklySchedule.
	// Structure is documented below.
	Schedule []ScheduleInitParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Properties with which the snapshots are created, such as labels.
	// Structure is documented below.
	SnapshotProperties []SnapshotPropertiesInitParameters `json:"snapshotProperties,omitempty" tf:"snapshot_properties,omitempty"`
}

type SnapshotSchedulePolicyObservation struct {

	// Retention policy applied to snapshots created by this resource policy.
	// Structure is documented below.
	RetentionPolicy []RetentionPolicyObservation `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Contains one of an hourlySchedule, dailySchedule, or weeklySchedule.
	// Structure is documented below.
	Schedule []ScheduleObservation `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// Properties with which the snapshots are created, such as labels.
	// Structure is documented below.
	SnapshotProperties []SnapshotPropertiesObservation `json:"snapshotProperties,omitempty" tf:"snapshot_properties,omitempty"`
}

type SnapshotSchedulePolicyParameters struct {

	// Retention policy applied to snapshots created by this resource policy.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// Contains one of an hourlySchedule, dailySchedule, or weeklySchedule.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule" tf:"schedule,omitempty"`

	// Properties with which the snapshots are created, such as labels.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SnapshotProperties []SnapshotPropertiesParameters `json:"snapshotProperties,omitempty" tf:"snapshot_properties,omitempty"`
}

type VMStartScheduleInitParameters struct {

	// Specifies the frequency for the operation, using the unix-cron format.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type VMStartScheduleObservation struct {

	// Specifies the frequency for the operation, using the unix-cron format.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type VMStartScheduleParameters struct {

	// Specifies the frequency for the operation, using the unix-cron format.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type VMStopScheduleInitParameters struct {

	// Specifies the frequency for the operation, using the unix-cron format.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type VMStopScheduleObservation struct {

	// Specifies the frequency for the operation, using the unix-cron format.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type VMStopScheduleParameters struct {

	// Specifies the frequency for the operation, using the unix-cron format.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type WeeklyScheduleInitParameters struct {

	// May contain up to seven (one for each day of the week) snapshot times.
	// Structure is documented below.
	DayOfWeeks []DayOfWeeksInitParameters `json:"dayOfWeeks,omitempty" tf:"day_of_weeks,omitempty"`
}

type WeeklyScheduleObservation struct {

	// May contain up to seven (one for each day of the week) snapshot times.
	// Structure is documented below.
	DayOfWeeks []DayOfWeeksObservation `json:"dayOfWeeks,omitempty" tf:"day_of_weeks,omitempty"`
}

type WeeklyScheduleParameters struct {

	// May contain up to seven (one for each day of the week) snapshot times.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DayOfWeeks []DayOfWeeksParameters `json:"dayOfWeeks" tf:"day_of_weeks,omitempty"`
}

// ResourcePolicySpec defines the desired state of ResourcePolicy
type ResourcePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourcePolicyParameters `json:"forProvider"`
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
	InitProvider ResourcePolicyInitParameters `json:"initProvider,omitempty"`
}

// ResourcePolicyStatus defines the observed state of ResourcePolicy.
type ResourcePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourcePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicy is the Schema for the ResourcePolicys API. A policy that can be attached to a resource to specify or schedule actions on that resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourcePolicySpec   `json:"spec"`
	Status            ResourcePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicyList contains a list of ResourcePolicys
type ResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcePolicy `json:"items"`
}

// Repository type metadata.
var (
	ResourcePolicy_Kind             = "ResourcePolicy"
	ResourcePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourcePolicy_Kind}.String()
	ResourcePolicy_KindAPIVersion   = ResourcePolicy_Kind + "." + CRDGroupVersion.String()
	ResourcePolicy_GroupVersionKind = CRDGroupVersion.WithKind(ResourcePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourcePolicy{}, &ResourcePolicyList{})
}
