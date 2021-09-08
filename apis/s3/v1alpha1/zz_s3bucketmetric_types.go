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

type S3BucketMetricFilterObservation struct {
}

type S3BucketMetricFilterParameters struct {
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type S3BucketMetricObservation struct {
}

type S3BucketMetricParameters struct {
	Bucket string `json:"bucket" tf:"bucket"`

	Filter []S3BucketMetricFilterParameters `json:"filter,omitempty" tf:"filter"`

	Name string `json:"name" tf:"name"`
}

// S3BucketMetricSpec defines the desired state of S3BucketMetric
type S3BucketMetricSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       S3BucketMetricParameters `json:"forProvider"`
}

// S3BucketMetricStatus defines the observed state of S3BucketMetric.
type S3BucketMetricStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          S3BucketMetricObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketMetric is the Schema for the S3BucketMetrics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type S3BucketMetric struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketMetricSpec   `json:"spec"`
	Status            S3BucketMetricStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketMetricList contains a list of S3BucketMetrics
type S3BucketMetricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketMetric `json:"items"`
}

// Repository type metadata.
var (
	S3BucketMetricKind             = "S3BucketMetric"
	S3BucketMetricGroupKind        = schema.GroupKind{Group: Group, Kind: S3BucketMetricKind}.String()
	S3BucketMetricKindAPIVersion   = S3BucketMetricKind + "." + GroupVersion.String()
	S3BucketMetricGroupVersionKind = GroupVersion.WithKind(S3BucketMetricKind)
)

func init() {
	SchemeBuilder.Register(&S3BucketMetric{}, &S3BucketMetricList{})
}