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

type EmrInstanceFleetInstanceTypeConfigsConfigurationsObservation struct {
}

type EmrInstanceFleetInstanceTypeConfigsConfigurationsParameters struct {
	Classification *string `json:"classification,omitempty" tf:"classification"`

	Properties map[string]string `json:"properties,omitempty" tf:"properties"`
}

type EmrInstanceFleetInstanceTypeConfigsEbsConfigObservation struct {
}

type EmrInstanceFleetInstanceTypeConfigsEbsConfigParameters struct {
	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	Size int64 `json:"size" tf:"size"`

	Type string `json:"type" tf:"type"`

	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance"`
}

type EmrInstanceFleetInstanceTypeConfigsObservation struct {
}

type EmrInstanceFleetInstanceTypeConfigsParameters struct {
	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price"`

	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice,omitempty" tf:"bid_price_as_percentage_of_on_demand_price"`

	Configurations []EmrInstanceFleetInstanceTypeConfigsConfigurationsParameters `json:"configurations,omitempty" tf:"configurations"`

	EbsConfig []EmrInstanceFleetInstanceTypeConfigsEbsConfigParameters `json:"ebsConfig,omitempty" tf:"ebs_config"`

	InstanceType string `json:"instanceType" tf:"instance_type"`

	WeightedCapacity *int64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity"`
}

type EmrInstanceFleetLaunchSpecificationsObservation struct {
}

type EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationObservation struct {
}

type EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationParameters struct {
	AllocationStrategy string `json:"allocationStrategy" tf:"allocation_strategy"`
}

type EmrInstanceFleetLaunchSpecificationsParameters struct {
	OnDemandSpecification []EmrInstanceFleetLaunchSpecificationsOnDemandSpecificationParameters `json:"onDemandSpecification,omitempty" tf:"on_demand_specification"`

	SpotSpecification []EmrInstanceFleetLaunchSpecificationsSpotSpecificationParameters `json:"spotSpecification,omitempty" tf:"spot_specification"`
}

type EmrInstanceFleetLaunchSpecificationsSpotSpecificationObservation struct {
}

type EmrInstanceFleetLaunchSpecificationsSpotSpecificationParameters struct {
	AllocationStrategy string `json:"allocationStrategy" tf:"allocation_strategy"`

	BlockDurationMinutes *int64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes"`

	TimeoutAction string `json:"timeoutAction" tf:"timeout_action"`

	TimeoutDurationMinutes int64 `json:"timeoutDurationMinutes" tf:"timeout_duration_minutes"`
}

type EmrInstanceFleetObservation struct {
	ProvisionedOnDemandCapacity int64 `json:"provisionedOnDemandCapacity" tf:"provisioned_on_demand_capacity"`

	ProvisionedSpotCapacity int64 `json:"provisionedSpotCapacity" tf:"provisioned_spot_capacity"`
}

type EmrInstanceFleetParameters struct {
	ClusterId string `json:"clusterId" tf:"cluster_id"`

	InstanceTypeConfigs []EmrInstanceFleetInstanceTypeConfigsParameters `json:"instanceTypeConfigs,omitempty" tf:"instance_type_configs"`

	LaunchSpecifications []EmrInstanceFleetLaunchSpecificationsParameters `json:"launchSpecifications,omitempty" tf:"launch_specifications"`

	Name *string `json:"name,omitempty" tf:"name"`

	TargetOnDemandCapacity *int64 `json:"targetOnDemandCapacity,omitempty" tf:"target_on_demand_capacity"`

	TargetSpotCapacity *int64 `json:"targetSpotCapacity,omitempty" tf:"target_spot_capacity"`
}

// EmrInstanceFleetSpec defines the desired state of EmrInstanceFleet
type EmrInstanceFleetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EmrInstanceFleetParameters `json:"forProvider"`
}

// EmrInstanceFleetStatus defines the observed state of EmrInstanceFleet.
type EmrInstanceFleetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EmrInstanceFleetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmrInstanceFleet is the Schema for the EmrInstanceFleets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EmrInstanceFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrInstanceFleetSpec   `json:"spec"`
	Status            EmrInstanceFleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmrInstanceFleetList contains a list of EmrInstanceFleets
type EmrInstanceFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmrInstanceFleet `json:"items"`
}

// Repository type metadata.
var (
	EmrInstanceFleetKind             = "EmrInstanceFleet"
	EmrInstanceFleetGroupKind        = schema.GroupKind{Group: Group, Kind: EmrInstanceFleetKind}.String()
	EmrInstanceFleetKindAPIVersion   = EmrInstanceFleetKind + "." + GroupVersion.String()
	EmrInstanceFleetGroupVersionKind = GroupVersion.WithKind(EmrInstanceFleetKind)
)

func init() {
	SchemeBuilder.Register(&EmrInstanceFleet{}, &EmrInstanceFleetList{})
}