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

type AudioCodecOptionsObservation struct {
}

type AudioCodecOptionsParameters struct {
	BitDepth *string `json:"bitDepth,omitempty" tf:"bit_depth"`

	BitOrder *string `json:"bitOrder,omitempty" tf:"bit_order"`

	Profile *string `json:"profile,omitempty" tf:"profile"`

	Signed *string `json:"signed,omitempty" tf:"signed"`
}

type AudioObservation struct {
}

type AudioParameters struct {
	AudioPackingMode *string `json:"audioPackingMode,omitempty" tf:"audio_packing_mode"`

	BitRate *string `json:"bitRate,omitempty" tf:"bit_rate"`

	Channels *string `json:"channels,omitempty" tf:"channels"`

	Codec *string `json:"codec,omitempty" tf:"codec"`

	SampleRate *string `json:"sampleRate,omitempty" tf:"sample_rate"`
}

type ElastictranscoderPresetObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type ElastictranscoderPresetParameters struct {
	Audio []AudioParameters `json:"audio,omitempty" tf:"audio"`

	AudioCodecOptions []AudioCodecOptionsParameters `json:"audioCodecOptions,omitempty" tf:"audio_codec_options"`

	Container string `json:"container" tf:"container"`

	Description *string `json:"description,omitempty" tf:"description"`

	Name *string `json:"name,omitempty" tf:"name"`

	Thumbnails []ThumbnailsParameters `json:"thumbnails,omitempty" tf:"thumbnails"`

	Type *string `json:"type,omitempty" tf:"type"`

	Video []VideoParameters `json:"video,omitempty" tf:"video"`

	VideoCodecOptions map[string]string `json:"videoCodecOptions,omitempty" tf:"video_codec_options"`

	VideoWatermarks []VideoWatermarksParameters `json:"videoWatermarks,omitempty" tf:"video_watermarks"`
}

type ThumbnailsObservation struct {
}

type ThumbnailsParameters struct {
	AspectRatio *string `json:"aspectRatio,omitempty" tf:"aspect_ratio"`

	Format *string `json:"format,omitempty" tf:"format"`

	Interval *string `json:"interval,omitempty" tf:"interval"`

	MaxHeight *string `json:"maxHeight,omitempty" tf:"max_height"`

	MaxWidth *string `json:"maxWidth,omitempty" tf:"max_width"`

	PaddingPolicy *string `json:"paddingPolicy,omitempty" tf:"padding_policy"`

	Resolution *string `json:"resolution,omitempty" tf:"resolution"`

	SizingPolicy *string `json:"sizingPolicy,omitempty" tf:"sizing_policy"`
}

type VideoObservation struct {
}

type VideoParameters struct {
	AspectRatio *string `json:"aspectRatio,omitempty" tf:"aspect_ratio"`

	BitRate *string `json:"bitRate,omitempty" tf:"bit_rate"`

	Codec *string `json:"codec,omitempty" tf:"codec"`

	DisplayAspectRatio *string `json:"displayAspectRatio,omitempty" tf:"display_aspect_ratio"`

	FixedGop *string `json:"fixedGop,omitempty" tf:"fixed_gop"`

	FrameRate *string `json:"frameRate,omitempty" tf:"frame_rate"`

	KeyframesMaxDist *string `json:"keyframesMaxDist,omitempty" tf:"keyframes_max_dist"`

	MaxFrameRate *string `json:"maxFrameRate,omitempty" tf:"max_frame_rate"`

	MaxHeight *string `json:"maxHeight,omitempty" tf:"max_height"`

	MaxWidth *string `json:"maxWidth,omitempty" tf:"max_width"`

	PaddingPolicy *string `json:"paddingPolicy,omitempty" tf:"padding_policy"`

	Resolution *string `json:"resolution,omitempty" tf:"resolution"`

	SizingPolicy *string `json:"sizingPolicy,omitempty" tf:"sizing_policy"`
}

type VideoWatermarksObservation struct {
}

type VideoWatermarksParameters struct {
	HorizontalAlign *string `json:"horizontalAlign,omitempty" tf:"horizontal_align"`

	HorizontalOffset *string `json:"horizontalOffset,omitempty" tf:"horizontal_offset"`

	Id *string `json:"id,omitempty" tf:"id"`

	MaxHeight *string `json:"maxHeight,omitempty" tf:"max_height"`

	MaxWidth *string `json:"maxWidth,omitempty" tf:"max_width"`

	Opacity *string `json:"opacity,omitempty" tf:"opacity"`

	SizingPolicy *string `json:"sizingPolicy,omitempty" tf:"sizing_policy"`

	Target *string `json:"target,omitempty" tf:"target"`

	VerticalAlign *string `json:"verticalAlign,omitempty" tf:"vertical_align"`

	VerticalOffset *string `json:"verticalOffset,omitempty" tf:"vertical_offset"`
}

// ElastictranscoderPresetSpec defines the desired state of ElastictranscoderPreset
type ElastictranscoderPresetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElastictranscoderPresetParameters `json:"forProvider"`
}

// ElastictranscoderPresetStatus defines the observed state of ElastictranscoderPreset.
type ElastictranscoderPresetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElastictranscoderPresetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElastictranscoderPreset is the Schema for the ElastictranscoderPresets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ElastictranscoderPreset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElastictranscoderPresetSpec   `json:"spec"`
	Status            ElastictranscoderPresetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElastictranscoderPresetList contains a list of ElastictranscoderPresets
type ElastictranscoderPresetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElastictranscoderPreset `json:"items"`
}

// Repository type metadata.
var (
	ElastictranscoderPresetKind             = "ElastictranscoderPreset"
	ElastictranscoderPresetGroupKind        = schema.GroupKind{Group: Group, Kind: ElastictranscoderPresetKind}.String()
	ElastictranscoderPresetKindAPIVersion   = ElastictranscoderPresetKind + "." + GroupVersion.String()
	ElastictranscoderPresetGroupVersionKind = GroupVersion.WithKind(ElastictranscoderPresetKind)
)

func init() {
	SchemeBuilder.Register(&ElastictranscoderPreset{}, &ElastictranscoderPresetList{})
}