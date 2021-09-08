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

type S3BucketPolicyObservation struct {
}

type S3BucketPolicyParameters struct {
	Bucket string `json:"bucket" tf:"bucket"`

	Policy string `json:"policy" tf:"policy"`
}

// S3BucketPolicySpec defines the desired state of S3BucketPolicy
type S3BucketPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       S3BucketPolicyParameters `json:"forProvider"`
}

// S3BucketPolicyStatus defines the observed state of S3BucketPolicy.
type S3BucketPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          S3BucketPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketPolicy is the Schema for the S3BucketPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type S3BucketPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketPolicySpec   `json:"spec"`
	Status            S3BucketPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketPolicyList contains a list of S3BucketPolicys
type S3BucketPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketPolicy `json:"items"`
}

// Repository type metadata.
var (
	S3BucketPolicyKind             = "S3BucketPolicy"
	S3BucketPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: S3BucketPolicyKind}.String()
	S3BucketPolicyKindAPIVersion   = S3BucketPolicyKind + "." + GroupVersion.String()
	S3BucketPolicyGroupVersionKind = GroupVersion.WithKind(S3BucketPolicyKind)
)

func init() {
	SchemeBuilder.Register(&S3BucketPolicy{}, &S3BucketPolicyList{})
}