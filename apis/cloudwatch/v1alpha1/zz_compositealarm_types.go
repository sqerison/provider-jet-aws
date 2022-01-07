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

type CompositeAlarmObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CompositeAlarmParameters struct {

	// +kubebuilder:validation:Optional
	ActionsEnabled *bool `json:"actionsEnabled,omitempty" tf:"actions_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	AlarmActions []*string `json:"alarmActions,omitempty" tf:"alarm_actions,omitempty"`

	// +kubebuilder:validation:Optional
	AlarmDescription *string `json:"alarmDescription,omitempty" tf:"alarm_description,omitempty"`

	// +kubebuilder:validation:Required
	AlarmName *string `json:"alarmName" tf:"alarm_name,omitempty"`

	// +kubebuilder:validation:Required
	AlarmRule *string `json:"alarmRule" tf:"alarm_rule,omitempty"`

	// +kubebuilder:validation:Optional
	InsufficientDataActions []*string `json:"insufficientDataActions,omitempty" tf:"insufficient_data_actions,omitempty"`

	// +kubebuilder:validation:Optional
	OkActions []*string `json:"okActions,omitempty" tf:"ok_actions,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// CompositeAlarmSpec defines the desired state of CompositeAlarm
type CompositeAlarmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CompositeAlarmParameters `json:"forProvider"`
}

// CompositeAlarmStatus defines the observed state of CompositeAlarm.
type CompositeAlarmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CompositeAlarmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CompositeAlarm is the Schema for the CompositeAlarms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type CompositeAlarm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CompositeAlarmSpec   `json:"spec"`
	Status            CompositeAlarmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CompositeAlarmList contains a list of CompositeAlarms
type CompositeAlarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CompositeAlarm `json:"items"`
}

// Repository type metadata.
var (
	CompositeAlarm_Kind             = "CompositeAlarm"
	CompositeAlarm_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CompositeAlarm_Kind}.String()
	CompositeAlarm_KindAPIVersion   = CompositeAlarm_Kind + "." + CRDGroupVersion.String()
	CompositeAlarm_GroupVersionKind = CRDGroupVersion.WithKind(CompositeAlarm_Kind)
)

func init() {
	SchemeBuilder.Register(&CompositeAlarm{}, &CompositeAlarmList{})
}