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

type HostedPublicVirtualInterfaceAccepterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type HostedPublicVirtualInterfaceAccepterParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VirtualInterfaceID *string `json:"virtualInterfaceId" tf:"virtual_interface_id,omitempty"`
}

// HostedPublicVirtualInterfaceAccepterSpec defines the desired state of HostedPublicVirtualInterfaceAccepter
type HostedPublicVirtualInterfaceAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostedPublicVirtualInterfaceAccepterParameters `json:"forProvider"`
}

// HostedPublicVirtualInterfaceAccepterStatus defines the observed state of HostedPublicVirtualInterfaceAccepter.
type HostedPublicVirtualInterfaceAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostedPublicVirtualInterfaceAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HostedPublicVirtualInterfaceAccepter is the Schema for the HostedPublicVirtualInterfaceAccepters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type HostedPublicVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostedPublicVirtualInterfaceAccepterSpec   `json:"spec"`
	Status            HostedPublicVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedPublicVirtualInterfaceAccepterList contains a list of HostedPublicVirtualInterfaceAccepters
type HostedPublicVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedPublicVirtualInterfaceAccepter `json:"items"`
}

// Repository type metadata.
var (
	HostedPublicVirtualInterfaceAccepter_Kind             = "HostedPublicVirtualInterfaceAccepter"
	HostedPublicVirtualInterfaceAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostedPublicVirtualInterfaceAccepter_Kind}.String()
	HostedPublicVirtualInterfaceAccepter_KindAPIVersion   = HostedPublicVirtualInterfaceAccepter_Kind + "." + CRDGroupVersion.String()
	HostedPublicVirtualInterfaceAccepter_GroupVersionKind = CRDGroupVersion.WithKind(HostedPublicVirtualInterfaceAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&HostedPublicVirtualInterfaceAccepter{}, &HostedPublicVirtualInterfaceAccepterList{})
}