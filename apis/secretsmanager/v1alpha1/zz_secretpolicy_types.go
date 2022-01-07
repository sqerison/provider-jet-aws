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

type SecretPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecretPolicyParameters struct {

	// +kubebuilder:validation:Optional
	BlockPublicPolicy *bool `json:"blockPublicPolicy,omitempty" tf:"block_public_policy,omitempty"`

	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	SecretArn *string `json:"secretArn" tf:"secret_arn,omitempty"`
}

// SecretPolicySpec defines the desired state of SecretPolicy
type SecretPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretPolicyParameters `json:"forProvider"`
}

// SecretPolicyStatus defines the observed state of SecretPolicy.
type SecretPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretPolicy is the Schema for the SecretPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type SecretPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretPolicySpec   `json:"spec"`
	Status            SecretPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretPolicyList contains a list of SecretPolicys
type SecretPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretPolicy `json:"items"`
}

// Repository type metadata.
var (
	SecretPolicy_Kind             = "SecretPolicy"
	SecretPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretPolicy_Kind}.String()
	SecretPolicy_KindAPIVersion   = SecretPolicy_Kind + "." + CRDGroupVersion.String()
	SecretPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SecretPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretPolicy{}, &SecretPolicyList{})
}