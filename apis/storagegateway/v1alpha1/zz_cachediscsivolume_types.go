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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CachedISCSIVolumeObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ChapEnabled *bool `json:"chapEnabled,omitempty" tf:"chap_enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LunNumber *int64 `json:"lunNumber,omitempty" tf:"lun_number,omitempty"`

	NetworkInterfacePort *int64 `json:"networkInterfacePort,omitempty" tf:"network_interface_port,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	TargetArn *string `json:"targetArn,omitempty" tf:"target_arn,omitempty"`

	VolumeArn *string `json:"volumeArn,omitempty" tf:"volume_arn,omitempty"`

	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type CachedISCSIVolumeParameters struct {

	// +kubebuilder:validation:Required
	GatewayArn *string `json:"gatewayArn" tf:"gateway_arn,omitempty"`

	// +kubebuilder:validation:Optional
	KMSEncrypted *bool `json:"kmsEncrypted,omitempty" tf:"kms_encrypted,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2.Key
	// +kubebuilder:validation:Optional
	KMSKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyRef *v1.Reference `json:"kmsKeyRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeySelector *v1.Selector `json:"kmsKeySelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	NetworkInterfaceID *string `json:"networkInterfaceId" tf:"network_interface_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// +kubebuilder:validation:Optional
	SourceVolumeArn *string `json:"sourceVolumeArn,omitempty" tf:"source_volume_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TargetName *string `json:"targetName" tf:"target_name,omitempty"`

	// +kubebuilder:validation:Required
	VolumeSizeInBytes *int64 `json:"volumeSizeInBytes" tf:"volume_size_in_bytes,omitempty"`
}

// CachedISCSIVolumeSpec defines the desired state of CachedISCSIVolume
type CachedISCSIVolumeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CachedISCSIVolumeParameters `json:"forProvider"`
}

// CachedISCSIVolumeStatus defines the observed state of CachedISCSIVolume.
type CachedISCSIVolumeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CachedISCSIVolumeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CachedISCSIVolume is the Schema for the CachedISCSIVolumes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type CachedISCSIVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CachedISCSIVolumeSpec   `json:"spec"`
	Status            CachedISCSIVolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CachedISCSIVolumeList contains a list of CachedISCSIVolumes
type CachedISCSIVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CachedISCSIVolume `json:"items"`
}

// Repository type metadata.
var (
	CachedISCSIVolume_Kind             = "CachedISCSIVolume"
	CachedISCSIVolume_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CachedISCSIVolume_Kind}.String()
	CachedISCSIVolume_KindAPIVersion   = CachedISCSIVolume_Kind + "." + CRDGroupVersion.String()
	CachedISCSIVolume_GroupVersionKind = CRDGroupVersion.WithKind(CachedISCSIVolume_Kind)
)

func init() {
	SchemeBuilder.Register(&CachedISCSIVolume{}, &CachedISCSIVolumeList{})
}