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

type CognitoIdentityProviderObservation struct {
}

type CognitoIdentityProviderParameters struct {
	AttributeMapping map[string]string `json:"attributeMapping,omitempty" tf:"attribute_mapping"`

	IdpIdentifiers []string `json:"idpIdentifiers,omitempty" tf:"idp_identifiers"`

	ProviderDetails map[string]string `json:"providerDetails" tf:"provider_details"`

	ProviderName string `json:"providerName" tf:"provider_name"`

	ProviderType string `json:"providerType" tf:"provider_type"`

	UserPoolId string `json:"userPoolId" tf:"user_pool_id"`
}

// CognitoIdentityProviderSpec defines the desired state of CognitoIdentityProvider
type CognitoIdentityProviderSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CognitoIdentityProviderParameters `json:"forProvider"`
}

// CognitoIdentityProviderStatus defines the observed state of CognitoIdentityProvider.
type CognitoIdentityProviderStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CognitoIdentityProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoIdentityProvider is the Schema for the CognitoIdentityProviders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CognitoIdentityProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoIdentityProviderSpec   `json:"spec"`
	Status            CognitoIdentityProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitoIdentityProviderList contains a list of CognitoIdentityProviders
type CognitoIdentityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitoIdentityProvider `json:"items"`
}

// Repository type metadata.
var (
	CognitoIdentityProviderKind             = "CognitoIdentityProvider"
	CognitoIdentityProviderGroupKind        = schema.GroupKind{Group: Group, Kind: CognitoIdentityProviderKind}.String()
	CognitoIdentityProviderKindAPIVersion   = CognitoIdentityProviderKind + "." + GroupVersion.String()
	CognitoIdentityProviderGroupVersionKind = GroupVersion.WithKind(CognitoIdentityProviderKind)
)

func init() {
	SchemeBuilder.Register(&CognitoIdentityProvider{}, &CognitoIdentityProviderList{})
}