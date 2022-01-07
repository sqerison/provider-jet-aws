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

type ReplicationInstanceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ReplicationInstanceArn *string `json:"replicationInstanceArn,omitempty" tf:"replication_instance_arn,omitempty"`

	ReplicationInstancePrivateIps []*string `json:"replicationInstancePrivateIps,omitempty" tf:"replication_instance_private_ips,omitempty"`

	ReplicationInstancePublicIps []*string `json:"replicationInstancePublicIps,omitempty" tf:"replication_instance_public_ips,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ReplicationInstanceParameters struct {

	// +kubebuilder:validation:Optional
	AllocatedStorage *int64 `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`

	// +kubebuilder:validation:Optional
	AllowMajorVersionUpgrade *bool `json:"allowMajorVersionUpgrade,omitempty" tf:"allow_major_version_upgrade,omitempty"`

	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	MultiAz *bool `json:"multiAz,omitempty" tf:"multi_az,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window,omitempty"`

	// +kubebuilder:validation:Optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ReplicationInstanceClass *string `json:"replicationInstanceClass" tf:"replication_instance_class,omitempty"`

	// +kubebuilder:validation:Required
	ReplicationInstanceID *string `json:"replicationInstanceId" tf:"replication_instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	ReplicationSubnetGroupID *string `json:"replicationSubnetGroupId,omitempty" tf:"replication_subnet_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupIdRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupIdSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIdRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIdSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`
}

// ReplicationInstanceSpec defines the desired state of ReplicationInstance
type ReplicationInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReplicationInstanceParameters `json:"forProvider"`
}

// ReplicationInstanceStatus defines the observed state of ReplicationInstance.
type ReplicationInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReplicationInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationInstance is the Schema for the ReplicationInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ReplicationInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicationInstanceSpec   `json:"spec"`
	Status            ReplicationInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReplicationInstanceList contains a list of ReplicationInstances
type ReplicationInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationInstance `json:"items"`
}

// Repository type metadata.
var (
	ReplicationInstance_Kind             = "ReplicationInstance"
	ReplicationInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReplicationInstance_Kind}.String()
	ReplicationInstance_KindAPIVersion   = ReplicationInstance_Kind + "." + CRDGroupVersion.String()
	ReplicationInstance_GroupVersionKind = CRDGroupVersion.WithKind(ReplicationInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&ReplicationInstance{}, &ReplicationInstanceList{})
}