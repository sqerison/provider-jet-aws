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

type AppmeshVirtualNodeObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreatedDate string `json:"createdDate" tf:"created_date"`

	LastUpdatedDate string `json:"lastUpdatedDate" tf:"last_updated_date"`

	ResourceOwner string `json:"resourceOwner" tf:"resource_owner"`
}

type AppmeshVirtualNodeParameters struct {
	MeshName string `json:"meshName" tf:"mesh_name"`

	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner"`

	Name string `json:"name" tf:"name"`

	Spec []AppmeshVirtualNodeSpecParameters `json:"spec" tf:"spec"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type AppmeshVirtualNodeSpecObservation struct {
}

type AppmeshVirtualNodeSpecParameters struct {
	Backend []BackendParameters `json:"backend,omitempty" tf:"backend"`

	BackendDefaults []SpecBackendDefaultsParameters `json:"backendDefaults,omitempty" tf:"backend_defaults"`

	Listener []SpecListenerParameters `json:"listener,omitempty" tf:"listener"`

	Logging []SpecLoggingParameters `json:"logging,omitempty" tf:"logging"`

	ServiceDiscovery []ServiceDiscoveryParameters `json:"serviceDiscovery,omitempty" tf:"service_discovery"`
}

type AwsCloudMapObservation struct {
}

type AwsCloudMapParameters struct {
	Attributes map[string]string `json:"attributes,omitempty" tf:"attributes"`

	NamespaceName string `json:"namespaceName" tf:"namespace_name"`

	ServiceName string `json:"serviceName" tf:"service_name"`
}

type BackendDefaultsClientPolicyObservation struct {
}

type BackendDefaultsClientPolicyParameters struct {
	Tls []BackendDefaultsClientPolicyTlsParameters `json:"tls,omitempty" tf:"tls"`
}

type BackendDefaultsClientPolicyTlsCertificateObservation struct {
}

type BackendDefaultsClientPolicyTlsCertificateParameters struct {
	File []ClientPolicyTlsCertificateFileParameters `json:"file,omitempty" tf:"file"`

	Sds []ClientPolicyTlsCertificateSdsParameters `json:"sds,omitempty" tf:"sds"`
}

type BackendDefaultsClientPolicyTlsObservation struct {
}

type BackendDefaultsClientPolicyTlsParameters struct {
	Certificate []BackendDefaultsClientPolicyTlsCertificateParameters `json:"certificate,omitempty" tf:"certificate"`

	Enforce *bool `json:"enforce,omitempty" tf:"enforce"`

	Ports []int64 `json:"ports,omitempty" tf:"ports"`

	Validation []BackendDefaultsClientPolicyTlsValidationParameters `json:"validation" tf:"validation"`
}

type BackendDefaultsClientPolicyTlsValidationObservation struct {
}

type BackendDefaultsClientPolicyTlsValidationParameters struct {
	SubjectAlternativeNames []ClientPolicyTlsValidationSubjectAlternativeNamesParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names"`

	Trust []ClientPolicyTlsValidationTrustParameters `json:"trust" tf:"trust"`
}

type BackendObservation struct {
}

type BackendParameters struct {
	VirtualService []BackendVirtualServiceParameters `json:"virtualService" tf:"virtual_service"`
}

type BackendVirtualServiceObservation struct {
}

type BackendVirtualServiceParameters struct {
	ClientPolicy []VirtualServiceClientPolicyParameters `json:"clientPolicy,omitempty" tf:"client_policy"`

	VirtualServiceName string `json:"virtualServiceName" tf:"virtual_service_name"`
}

type BaseEjectionDurationObservation struct {
}

type BaseEjectionDurationParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type ClientPolicyTlsCertificateFileObservation struct {
}

type ClientPolicyTlsCertificateFileParameters struct {
	CertificateChain string `json:"certificateChain" tf:"certificate_chain"`

	PrivateKey string `json:"privateKey" tf:"private_key"`
}

type ClientPolicyTlsCertificateObservation struct {
}

type ClientPolicyTlsCertificateParameters struct {
	File []TlsCertificateFileParameters `json:"file,omitempty" tf:"file"`

	Sds []TlsCertificateSdsParameters `json:"sds,omitempty" tf:"sds"`
}

type ClientPolicyTlsCertificateSdsObservation struct {
}

type ClientPolicyTlsCertificateSdsParameters struct {
	SecretName string `json:"secretName" tf:"secret_name"`
}

type ClientPolicyTlsObservation struct {
}

type ClientPolicyTlsParameters struct {
	Certificate []ClientPolicyTlsCertificateParameters `json:"certificate,omitempty" tf:"certificate"`

	Enforce *bool `json:"enforce,omitempty" tf:"enforce"`

	Ports []int64 `json:"ports,omitempty" tf:"ports"`

	Validation []ClientPolicyTlsValidationParameters `json:"validation" tf:"validation"`
}

type ClientPolicyTlsValidationObservation struct {
}

type ClientPolicyTlsValidationParameters struct {
	SubjectAlternativeNames []TlsValidationSubjectAlternativeNamesParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names"`

	Trust []TlsValidationTrustParameters `json:"trust" tf:"trust"`
}

type ClientPolicyTlsValidationSubjectAlternativeNamesMatchObservation struct {
}

type ClientPolicyTlsValidationSubjectAlternativeNamesMatchParameters struct {
	Exact []string `json:"exact" tf:"exact"`
}

type ClientPolicyTlsValidationSubjectAlternativeNamesObservation struct {
}

type ClientPolicyTlsValidationSubjectAlternativeNamesParameters struct {
	Match []ClientPolicyTlsValidationSubjectAlternativeNamesMatchParameters `json:"match" tf:"match"`
}

type ClientPolicyTlsValidationTrustFileObservation struct {
}

type ClientPolicyTlsValidationTrustFileParameters struct {
	CertificateChain string `json:"certificateChain" tf:"certificate_chain"`
}

type ClientPolicyTlsValidationTrustObservation struct {
}

type ClientPolicyTlsValidationTrustParameters struct {
	Acm []ValidationTrustAcmParameters `json:"acm,omitempty" tf:"acm"`

	File []ClientPolicyTlsValidationTrustFileParameters `json:"file,omitempty" tf:"file"`

	Sds []ClientPolicyTlsValidationTrustSdsParameters `json:"sds,omitempty" tf:"sds"`
}

type ClientPolicyTlsValidationTrustSdsObservation struct {
}

type ClientPolicyTlsValidationTrustSdsParameters struct {
	SecretName string `json:"secretName" tf:"secret_name"`
}

type ConnectionPoolGrpcObservation struct {
}

type ConnectionPoolGrpcParameters struct {
	MaxRequests int64 `json:"maxRequests" tf:"max_requests"`
}

type ConnectionPoolHttp2Observation struct {
}

type ConnectionPoolHttp2Parameters struct {
	MaxRequests int64 `json:"maxRequests" tf:"max_requests"`
}

type ConnectionPoolHttpObservation struct {
}

type ConnectionPoolHttpParameters struct {
	MaxConnections int64 `json:"maxConnections" tf:"max_connections"`

	MaxPendingRequests *int64 `json:"maxPendingRequests,omitempty" tf:"max_pending_requests"`
}

type DnsObservation struct {
}

type DnsParameters struct {
	Hostname string `json:"hostname" tf:"hostname"`
}

type GrpcIdleObservation struct {
}

type GrpcIdleParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type GrpcPerRequestObservation struct {
}

type GrpcPerRequestParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type Http2IdleObservation struct {
}

type Http2IdleParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type Http2PerRequestObservation struct {
}

type Http2PerRequestParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type HttpIdleObservation struct {
}

type HttpIdleParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type HttpPerRequestObservation struct {
}

type HttpPerRequestParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type IntervalObservation struct {
}

type IntervalParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type ListenerConnectionPoolObservation struct {
}

type ListenerConnectionPoolParameters struct {
	Grpc []ConnectionPoolGrpcParameters `json:"grpc,omitempty" tf:"grpc"`

	Http []ConnectionPoolHttpParameters `json:"http,omitempty" tf:"http"`

	Http2 []ConnectionPoolHttp2Parameters `json:"http2,omitempty" tf:"http2"`

	Tcp []TcpParameters `json:"tcp,omitempty" tf:"tcp"`
}

type ListenerHealthCheckObservation struct {
}

type ListenerHealthCheckParameters struct {
	HealthyThreshold int64 `json:"healthyThreshold" tf:"healthy_threshold"`

	IntervalMillis int64 `json:"intervalMillis" tf:"interval_millis"`

	Path *string `json:"path,omitempty" tf:"path"`

	Port *int64 `json:"port,omitempty" tf:"port"`

	Protocol string `json:"protocol" tf:"protocol"`

	TimeoutMillis int64 `json:"timeoutMillis" tf:"timeout_millis"`

	UnhealthyThreshold int64 `json:"unhealthyThreshold" tf:"unhealthy_threshold"`
}

type ListenerPortMappingObservation struct {
}

type ListenerPortMappingParameters struct {
	Port int64 `json:"port" tf:"port"`

	Protocol string `json:"protocol" tf:"protocol"`
}

type ListenerTimeoutObservation struct {
}

type ListenerTimeoutParameters struct {
	Grpc []TimeoutGrpcParameters `json:"grpc,omitempty" tf:"grpc"`

	Http []TimeoutHttpParameters `json:"http,omitempty" tf:"http"`

	Http2 []TimeoutHttp2Parameters `json:"http2,omitempty" tf:"http2"`

	Tcp []TimeoutTcpParameters `json:"tcp,omitempty" tf:"tcp"`
}

type ListenerTlsCertificateFileObservation struct {
}

type ListenerTlsCertificateFileParameters struct {
	CertificateChain string `json:"certificateChain" tf:"certificate_chain"`

	PrivateKey string `json:"privateKey" tf:"private_key"`
}

type ListenerTlsCertificateObservation struct {
}

type ListenerTlsCertificateParameters struct {
	Acm []TlsCertificateAcmParameters `json:"acm,omitempty" tf:"acm"`

	File []ListenerTlsCertificateFileParameters `json:"file,omitempty" tf:"file"`

	Sds []ListenerTlsCertificateSdsParameters `json:"sds,omitempty" tf:"sds"`
}

type ListenerTlsCertificateSdsObservation struct {
}

type ListenerTlsCertificateSdsParameters struct {
	SecretName string `json:"secretName" tf:"secret_name"`
}

type ListenerTlsValidationObservation struct {
}

type ListenerTlsValidationParameters struct {
	SubjectAlternativeNames []ListenerTlsValidationSubjectAlternativeNamesParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names"`

	Trust []ListenerTlsValidationTrustParameters `json:"trust" tf:"trust"`
}

type ListenerTlsValidationSubjectAlternativeNamesMatchObservation struct {
}

type ListenerTlsValidationSubjectAlternativeNamesMatchParameters struct {
	Exact []string `json:"exact" tf:"exact"`
}

type ListenerTlsValidationSubjectAlternativeNamesObservation struct {
}

type ListenerTlsValidationSubjectAlternativeNamesParameters struct {
	Match []ListenerTlsValidationSubjectAlternativeNamesMatchParameters `json:"match" tf:"match"`
}

type ListenerTlsValidationTrustFileObservation struct {
}

type ListenerTlsValidationTrustFileParameters struct {
	CertificateChain string `json:"certificateChain" tf:"certificate_chain"`
}

type ListenerTlsValidationTrustObservation struct {
}

type ListenerTlsValidationTrustParameters struct {
	File []ListenerTlsValidationTrustFileParameters `json:"file,omitempty" tf:"file"`

	Sds []ListenerTlsValidationTrustSdsParameters `json:"sds,omitempty" tf:"sds"`
}

type ListenerTlsValidationTrustSdsObservation struct {
}

type ListenerTlsValidationTrustSdsParameters struct {
	SecretName string `json:"secretName" tf:"secret_name"`
}

type LoggingAccessLogFileObservation struct {
}

type LoggingAccessLogFileParameters struct {
	Path string `json:"path" tf:"path"`
}

type LoggingAccessLogObservation struct {
}

type LoggingAccessLogParameters struct {
	File []LoggingAccessLogFileParameters `json:"file,omitempty" tf:"file"`
}

type OutlierDetectionObservation struct {
}

type OutlierDetectionParameters struct {
	BaseEjectionDuration []BaseEjectionDurationParameters `json:"baseEjectionDuration" tf:"base_ejection_duration"`

	Interval []IntervalParameters `json:"interval" tf:"interval"`

	MaxEjectionPercent int64 `json:"maxEjectionPercent" tf:"max_ejection_percent"`

	MaxServerErrors int64 `json:"maxServerErrors" tf:"max_server_errors"`
}

type ServiceDiscoveryObservation struct {
}

type ServiceDiscoveryParameters struct {
	AwsCloudMap []AwsCloudMapParameters `json:"awsCloudMap,omitempty" tf:"aws_cloud_map"`

	Dns []DnsParameters `json:"dns,omitempty" tf:"dns"`
}

type SpecBackendDefaultsObservation struct {
}

type SpecBackendDefaultsParameters struct {
	ClientPolicy []BackendDefaultsClientPolicyParameters `json:"clientPolicy,omitempty" tf:"client_policy"`
}

type SpecListenerObservation struct {
}

type SpecListenerParameters struct {
	ConnectionPool []ListenerConnectionPoolParameters `json:"connectionPool,omitempty" tf:"connection_pool"`

	HealthCheck []ListenerHealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check"`

	OutlierDetection []OutlierDetectionParameters `json:"outlierDetection,omitempty" tf:"outlier_detection"`

	PortMapping []ListenerPortMappingParameters `json:"portMapping" tf:"port_mapping"`

	Timeout []ListenerTimeoutParameters `json:"timeout,omitempty" tf:"timeout"`

	Tls []SpecListenerTlsParameters `json:"tls,omitempty" tf:"tls"`
}

type SpecListenerTlsObservation struct {
}

type SpecListenerTlsParameters struct {
	Certificate []ListenerTlsCertificateParameters `json:"certificate" tf:"certificate"`

	Mode string `json:"mode" tf:"mode"`

	Validation []ListenerTlsValidationParameters `json:"validation,omitempty" tf:"validation"`
}

type SpecLoggingObservation struct {
}

type SpecLoggingParameters struct {
	AccessLog []LoggingAccessLogParameters `json:"accessLog,omitempty" tf:"access_log"`
}

type TcpIdleObservation struct {
}

type TcpIdleParameters struct {
	Unit string `json:"unit" tf:"unit"`

	Value int64 `json:"value" tf:"value"`
}

type TcpObservation struct {
}

type TcpParameters struct {
	MaxConnections int64 `json:"maxConnections" tf:"max_connections"`
}

type TimeoutGrpcObservation struct {
}

type TimeoutGrpcParameters struct {
	Idle []GrpcIdleParameters `json:"idle,omitempty" tf:"idle"`

	PerRequest []GrpcPerRequestParameters `json:"perRequest,omitempty" tf:"per_request"`
}

type TimeoutHttp2Observation struct {
}

type TimeoutHttp2Parameters struct {
	Idle []Http2IdleParameters `json:"idle,omitempty" tf:"idle"`

	PerRequest []Http2PerRequestParameters `json:"perRequest,omitempty" tf:"per_request"`
}

type TimeoutHttpObservation struct {
}

type TimeoutHttpParameters struct {
	Idle []HttpIdleParameters `json:"idle,omitempty" tf:"idle"`

	PerRequest []HttpPerRequestParameters `json:"perRequest,omitempty" tf:"per_request"`
}

type TimeoutTcpObservation struct {
}

type TimeoutTcpParameters struct {
	Idle []TcpIdleParameters `json:"idle,omitempty" tf:"idle"`
}

type TlsCertificateAcmObservation struct {
}

type TlsCertificateAcmParameters struct {
	CertificateArn string `json:"certificateArn" tf:"certificate_arn"`
}

type TlsCertificateFileObservation struct {
}

type TlsCertificateFileParameters struct {
	CertificateChain string `json:"certificateChain" tf:"certificate_chain"`

	PrivateKey string `json:"privateKey" tf:"private_key"`
}

type TlsCertificateSdsObservation struct {
}

type TlsCertificateSdsParameters struct {
	SecretName string `json:"secretName" tf:"secret_name"`
}

type TlsValidationSubjectAlternativeNamesMatchObservation struct {
}

type TlsValidationSubjectAlternativeNamesMatchParameters struct {
	Exact []string `json:"exact" tf:"exact"`
}

type TlsValidationSubjectAlternativeNamesObservation struct {
}

type TlsValidationSubjectAlternativeNamesParameters struct {
	Match []TlsValidationSubjectAlternativeNamesMatchParameters `json:"match" tf:"match"`
}

type TlsValidationTrustFileObservation struct {
}

type TlsValidationTrustFileParameters struct {
	CertificateChain string `json:"certificateChain" tf:"certificate_chain"`
}

type TlsValidationTrustObservation struct {
}

type TlsValidationTrustParameters struct {
	Acm []TrustAcmParameters `json:"acm,omitempty" tf:"acm"`

	File []TlsValidationTrustFileParameters `json:"file,omitempty" tf:"file"`

	Sds []TlsValidationTrustSdsParameters `json:"sds,omitempty" tf:"sds"`
}

type TlsValidationTrustSdsObservation struct {
}

type TlsValidationTrustSdsParameters struct {
	SecretName string `json:"secretName" tf:"secret_name"`
}

type TrustAcmObservation struct {
}

type TrustAcmParameters struct {
	CertificateAuthorityArns []string `json:"certificateAuthorityArns" tf:"certificate_authority_arns"`
}

type ValidationTrustAcmObservation struct {
}

type ValidationTrustAcmParameters struct {
	CertificateAuthorityArns []string `json:"certificateAuthorityArns" tf:"certificate_authority_arns"`
}

type VirtualServiceClientPolicyObservation struct {
}

type VirtualServiceClientPolicyParameters struct {
	Tls []ClientPolicyTlsParameters `json:"tls,omitempty" tf:"tls"`
}

// AppmeshVirtualNodeSpec defines the desired state of AppmeshVirtualNode
type AppmeshVirtualNodeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppmeshVirtualNodeParameters `json:"forProvider"`
}

// AppmeshVirtualNodeStatus defines the observed state of AppmeshVirtualNode.
type AppmeshVirtualNodeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppmeshVirtualNodeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshVirtualNode is the Schema for the AppmeshVirtualNodes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AppmeshVirtualNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshVirtualNodeSpec   `json:"spec"`
	Status            AppmeshVirtualNodeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshVirtualNodeList contains a list of AppmeshVirtualNodes
type AppmeshVirtualNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshVirtualNode `json:"items"`
}

// Repository type metadata.
var (
	AppmeshVirtualNodeKind             = "AppmeshVirtualNode"
	AppmeshVirtualNodeGroupKind        = schema.GroupKind{Group: Group, Kind: AppmeshVirtualNodeKind}.String()
	AppmeshVirtualNodeKindAPIVersion   = AppmeshVirtualNodeKind + "." + GroupVersion.String()
	AppmeshVirtualNodeGroupVersionKind = GroupVersion.WithKind(AppmeshVirtualNodeKind)
)

func init() {
	SchemeBuilder.Register(&AppmeshVirtualNode{}, &AppmeshVirtualNodeList{})
}