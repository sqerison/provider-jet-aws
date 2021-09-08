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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type FsxLustreFileSystemObservation struct {
	Arn string `json:"arn" tf:"arn"`

	DnsName string `json:"dnsName" tf:"dns_name"`

	MountName string `json:"mountName" tf:"mount_name"`

	NetworkInterfaceIds []string `json:"networkInterfaceIds" tf:"network_interface_ids"`

	OwnerId string `json:"ownerId" tf:"owner_id"`

	VpcId string `json:"vpcId" tf:"vpc_id"`
}

type FsxLustreFileSystemParameters struct {
	AutoImportPolicy *string `json:"autoImportPolicy,omitempty" tf:"auto_import_policy"`

	AutomaticBackupRetentionDays *int64 `json:"automaticBackupRetentionDays,omitempty" tf:"automatic_backup_retention_days"`

	CopyTagsToBackups *bool `json:"copyTagsToBackups,omitempty" tf:"copy_tags_to_backups"`

	DailyAutomaticBackupStartTime *string `json:"dailyAutomaticBackupStartTime,omitempty" tf:"daily_automatic_backup_start_time"`

	DataCompressionType *string `json:"dataCompressionType,omitempty" tf:"data_compression_type"`

	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type"`

	DriveCacheType *string `json:"driveCacheType,omitempty" tf:"drive_cache_type"`

	ExportPath *string `json:"exportPath,omitempty" tf:"export_path"`

	ImportPath *string `json:"importPath,omitempty" tf:"import_path"`

	ImportedFileChunkSize *int64 `json:"importedFileChunkSize,omitempty" tf:"imported_file_chunk_size"`

	KmsKeyId *string `json:"kmsKeyId,omitempty" tf:"kms_key_id"`

	PerUnitStorageThroughput *int64 `json:"perUnitStorageThroughput,omitempty" tf:"per_unit_storage_throughput"`

	SecurityGroupIds []string `json:"securityGroupIds,omitempty" tf:"security_group_ids"`

	StorageCapacity int64 `json:"storageCapacity" tf:"storage_capacity"`

	StorageType *string `json:"storageType,omitempty" tf:"storage_type"`

	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	WeeklyMaintenanceStartTime *string `json:"weeklyMaintenanceStartTime,omitempty" tf:"weekly_maintenance_start_time"`
}

// FsxLustreFileSystemSpec defines the desired state of FsxLustreFileSystem
type FsxLustreFileSystemSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FsxLustreFileSystemParameters `json:"forProvider"`
}

// FsxLustreFileSystemStatus defines the observed state of FsxLustreFileSystem.
type FsxLustreFileSystemStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FsxLustreFileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FsxLustreFileSystem is the Schema for the FsxLustreFileSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type FsxLustreFileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FsxLustreFileSystemSpec   `json:"spec"`
	Status            FsxLustreFileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FsxLustreFileSystemList contains a list of FsxLustreFileSystems
type FsxLustreFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FsxLustreFileSystem `json:"items"`
}

// Repository type metadata.
var (
	FsxLustreFileSystemKind             = "FsxLustreFileSystem"
	FsxLustreFileSystemGroupKind        = schema.GroupKind{Group: Group, Kind: FsxLustreFileSystemKind}.String()
	FsxLustreFileSystemKindAPIVersion   = FsxLustreFileSystemKind + "." + GroupVersion.String()
	FsxLustreFileSystemGroupVersionKind = GroupVersion.WithKind(FsxLustreFileSystemKind)
)

func init() {
	SchemeBuilder.Register(&FsxLustreFileSystem{}, &FsxLustreFileSystemList{})
}