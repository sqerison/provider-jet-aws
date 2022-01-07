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

type HAProxyLayerEBSVolumeObservation struct {
}

type HAProxyLayerEBSVolumeParameters struct {

	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Required
	MountPoint *string `json:"mountPoint" tf:"mount_point,omitempty"`

	// +kubebuilder:validation:Required
	NumberOfDisks *int64 `json:"numberOfDisks" tf:"number_of_disks,omitempty"`

	// +kubebuilder:validation:Optional
	RaidLevel *string `json:"raidLevel,omitempty" tf:"raid_level,omitempty"`

	// +kubebuilder:validation:Required
	Size *int64 `json:"size" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HAProxyLayerObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type HAProxyLayerParameters struct {

	// +kubebuilder:validation:Optional
	AutoAssignElasticIps *bool `json:"autoAssignElasticIps,omitempty" tf:"auto_assign_elastic_ips,omitempty"`

	// +kubebuilder:validation:Optional
	AutoAssignPublicIps *bool `json:"autoAssignPublicIps,omitempty" tf:"auto_assign_public_ips,omitempty"`

	// +kubebuilder:validation:Optional
	AutoHealing *bool `json:"autoHealing,omitempty" tf:"auto_healing,omitempty"`

	// +kubebuilder:validation:Optional
	CustomConfigureRecipes []*string `json:"customConfigureRecipes,omitempty" tf:"custom_configure_recipes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomDeployRecipes []*string `json:"customDeployRecipes,omitempty" tf:"custom_deploy_recipes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomInstanceProfileArn *string `json:"customInstanceProfileArn,omitempty" tf:"custom_instance_profile_arn,omitempty"`

	// +kubebuilder:validation:Optional
	CustomJSON *string `json:"customJson,omitempty" tf:"custom_json,omitempty"`

	// +kubebuilder:validation:Optional
	CustomSecurityGroupIdRefs []v1.Reference `json:"customSecurityGroupIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CustomSecurityGroupIdSelector *v1.Selector `json:"customSecurityGroupIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.SecurityGroup
	// +crossplane:generate:reference:refFieldName=CustomSecurityGroupIdRefs
	// +crossplane:generate:reference:selectorFieldName=CustomSecurityGroupIdSelector
	// +kubebuilder:validation:Optional
	CustomSecurityGroupIds []*string `json:"customSecurityGroupIds,omitempty" tf:"custom_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	CustomSetupRecipes []*string `json:"customSetupRecipes,omitempty" tf:"custom_setup_recipes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomShutdownRecipes []*string `json:"customShutdownRecipes,omitempty" tf:"custom_shutdown_recipes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomUndeployRecipes []*string `json:"customUndeployRecipes,omitempty" tf:"custom_undeploy_recipes,omitempty"`

	// +kubebuilder:validation:Optional
	DrainELBOnShutdown *bool `json:"drainElbOnShutdown,omitempty" tf:"drain_elb_on_shutdown,omitempty"`

	// +kubebuilder:validation:Optional
	EBSVolume []HAProxyLayerEBSVolumeParameters `json:"ebsVolume,omitempty" tf:"ebs_volume,omitempty"`

	// +kubebuilder:validation:Optional
	ElasticLoadBalancer *string `json:"elasticLoadBalancer,omitempty" tf:"elastic_load_balancer,omitempty"`

	// +kubebuilder:validation:Optional
	HealthcheckMethod *string `json:"healthcheckMethod,omitempty" tf:"healthcheck_method,omitempty"`

	// +kubebuilder:validation:Optional
	HealthcheckURL *string `json:"healthcheckUrl,omitempty" tf:"healthcheck_url,omitempty"`

	// +kubebuilder:validation:Optional
	InstallUpdatesOnBoot *bool `json:"installUpdatesOnBoot,omitempty" tf:"install_updates_on_boot,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceShutdownTimeout *int64 `json:"instanceShutdownTimeout,omitempty" tf:"instance_shutdown_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	StackID *string `json:"stackId" tf:"stack_id,omitempty"`

	// +kubebuilder:validation:Optional
	StatsEnabled *bool `json:"statsEnabled,omitempty" tf:"stats_enabled,omitempty"`

	// +kubebuilder:validation:Required
	StatsPassword *string `json:"statsPassword" tf:"stats_password,omitempty"`

	// +kubebuilder:validation:Optional
	StatsURL *string `json:"statsUrl,omitempty" tf:"stats_url,omitempty"`

	// +kubebuilder:validation:Optional
	StatsUser *string `json:"statsUser,omitempty" tf:"stats_user,omitempty"`

	// +kubebuilder:validation:Optional
	SystemPackages []*string `json:"systemPackages,omitempty" tf:"system_packages,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	UseEBSOptimizedInstances *bool `json:"useEbsOptimizedInstances,omitempty" tf:"use_ebs_optimized_instances,omitempty"`
}

// HAProxyLayerSpec defines the desired state of HAProxyLayer
type HAProxyLayerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HAProxyLayerParameters `json:"forProvider"`
}

// HAProxyLayerStatus defines the observed state of HAProxyLayer.
type HAProxyLayerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HAProxyLayerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HAProxyLayer is the Schema for the HAProxyLayers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type HAProxyLayer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HAProxyLayerSpec   `json:"spec"`
	Status            HAProxyLayerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HAProxyLayerList contains a list of HAProxyLayers
type HAProxyLayerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HAProxyLayer `json:"items"`
}

// Repository type metadata.
var (
	HAProxyLayer_Kind             = "HAProxyLayer"
	HAProxyLayer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HAProxyLayer_Kind}.String()
	HAProxyLayer_KindAPIVersion   = HAProxyLayer_Kind + "." + CRDGroupVersion.String()
	HAProxyLayer_GroupVersionKind = CRDGroupVersion.WithKind(HAProxyLayer_Kind)
)

func init() {
	SchemeBuilder.Register(&HAProxyLayer{}, &HAProxyLayerList{})
}