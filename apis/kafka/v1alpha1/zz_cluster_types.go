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

type BrokerLogsObservation struct {
}

type BrokerLogsParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchLogs []CloudwatchLogsParameters `json:"cloudwatchLogs,omitempty" tf:"cloudwatch_logs,omitempty"`

	// +kubebuilder:validation:Optional
	Firehose []FirehoseParameters `json:"firehose,omitempty" tf:"firehose,omitempty"`

	// +kubebuilder:validation:Optional
	S3 []S3Parameters `json:"s3,omitempty" tf:"s3,omitempty"`
}

type BrokerNodeGroupInfoObservation struct {
}

type BrokerNodeGroupInfoParameters struct {

	// +kubebuilder:validation:Optional
	AzDistribution *string `json:"azDistribution,omitempty" tf:"az_distribution,omitempty"`

	// +kubebuilder:validation:Required
	ClientSubnets []*string `json:"clientSubnets" tf:"client_subnets,omitempty"`

	// +kubebuilder:validation:Required
	EBSVolumeSize *int64 `json:"ebsVolumeSize" tf:"ebs_volume_size,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType *string `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Required
	SecurityGroups []*string `json:"securityGroups" tf:"security_groups,omitempty"`
}

type ClientAuthenticationObservation struct {
}

type ClientAuthenticationParameters struct {

	// +kubebuilder:validation:Optional
	Sasl []SaslParameters `json:"sasl,omitempty" tf:"sasl,omitempty"`

	// +kubebuilder:validation:Optional
	TLS []TLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`
}

type CloudwatchLogsObservation struct {
}

type CloudwatchLogsParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LogGroup *string `json:"logGroup,omitempty" tf:"log_group,omitempty"`
}

type ClusterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	BootstrapBrokers *string `json:"bootstrapBrokers,omitempty" tf:"bootstrap_brokers,omitempty"`

	BootstrapBrokersSaslIAM *string `json:"bootstrapBrokersSaslIam,omitempty" tf:"bootstrap_brokers_sasl_iam,omitempty"`

	BootstrapBrokersSaslScram *string `json:"bootstrapBrokersSaslScram,omitempty" tf:"bootstrap_brokers_sasl_scram,omitempty"`

	BootstrapBrokersTLS *string `json:"bootstrapBrokersTls,omitempty" tf:"bootstrap_brokers_tls,omitempty"`

	CurrentVersion *string `json:"currentVersion,omitempty" tf:"current_version,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	ZookeeperConnectString *string `json:"zookeeperConnectString,omitempty" tf:"zookeeper_connect_string,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Required
	BrokerNodeGroupInfo []BrokerNodeGroupInfoParameters `json:"brokerNodeGroupInfo" tf:"broker_node_group_info,omitempty"`

	// +kubebuilder:validation:Optional
	ClientAuthentication []ClientAuthenticationParameters `json:"clientAuthentication,omitempty" tf:"client_authentication,omitempty"`

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Optional
	ConfigurationInfo []ConfigurationInfoParameters `json:"configurationInfo,omitempty" tf:"configuration_info,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionInfo []EncryptionInfoParameters `json:"encryptionInfo,omitempty" tf:"encryption_info,omitempty"`

	// +kubebuilder:validation:Optional
	EnhancedMonitoring *string `json:"enhancedMonitoring,omitempty" tf:"enhanced_monitoring,omitempty"`

	// +kubebuilder:validation:Required
	KafkaVersion *string `json:"kafkaVersion" tf:"kafka_version,omitempty"`

	// +kubebuilder:validation:Optional
	LoggingInfo []LoggingInfoParameters `json:"loggingInfo,omitempty" tf:"logging_info,omitempty"`

	// +kubebuilder:validation:Required
	NumberOfBrokerNodes *int64 `json:"numberOfBrokerNodes" tf:"number_of_broker_nodes,omitempty"`

	// +kubebuilder:validation:Optional
	OpenMonitoring []OpenMonitoringParameters `json:"openMonitoring,omitempty" tf:"open_monitoring,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ConfigurationInfoObservation struct {
}

type ConfigurationInfoParameters struct {

	// +kubebuilder:validation:Required
	Arn *string `json:"arn" tf:"arn,omitempty"`

	// +kubebuilder:validation:Required
	Revision *int64 `json:"revision" tf:"revision,omitempty"`
}

type EncryptionInTransitObservation struct {
}

type EncryptionInTransitParameters struct {

	// +kubebuilder:validation:Optional
	ClientBroker *string `json:"clientBroker,omitempty" tf:"client_broker,omitempty"`

	// +kubebuilder:validation:Optional
	InCluster *bool `json:"inCluster,omitempty" tf:"in_cluster,omitempty"`
}

type EncryptionInfoObservation struct {
}

type EncryptionInfoParameters struct {

	// +kubebuilder:validation:Optional
	EncryptionAtRestKMSKeyArn *string `json:"encryptionAtRestKmsKeyArn,omitempty" tf:"encryption_at_rest_kms_key_arn,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionInTransit []EncryptionInTransitParameters `json:"encryptionInTransit,omitempty" tf:"encryption_in_transit,omitempty"`
}

type FirehoseObservation struct {
}

type FirehoseParameters struct {

	// +kubebuilder:validation:Optional
	DeliveryStream *string `json:"deliveryStream,omitempty" tf:"delivery_stream,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`
}

type JmxExporterObservation struct {
}

type JmxExporterParameters struct {

	// +kubebuilder:validation:Required
	EnabledInBroker *bool `json:"enabledInBroker" tf:"enabled_in_broker,omitempty"`
}

type LoggingInfoObservation struct {
}

type LoggingInfoParameters struct {

	// +kubebuilder:validation:Required
	BrokerLogs []BrokerLogsParameters `json:"brokerLogs" tf:"broker_logs,omitempty"`
}

type NodeExporterObservation struct {
}

type NodeExporterParameters struct {

	// +kubebuilder:validation:Required
	EnabledInBroker *bool `json:"enabledInBroker" tf:"enabled_in_broker,omitempty"`
}

type OpenMonitoringObservation struct {
}

type OpenMonitoringParameters struct {

	// +kubebuilder:validation:Required
	Prometheus []PrometheusParameters `json:"prometheus" tf:"prometheus,omitempty"`
}

type PrometheusObservation struct {
}

type PrometheusParameters struct {

	// +kubebuilder:validation:Optional
	JmxExporter []JmxExporterParameters `json:"jmxExporter,omitempty" tf:"jmx_exporter,omitempty"`

	// +kubebuilder:validation:Optional
	NodeExporter []NodeExporterParameters `json:"nodeExporter,omitempty" tf:"node_exporter,omitempty"`
}

type S3Observation struct {
}

type S3Parameters struct {

	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type SaslObservation struct {
}

type SaslParameters struct {

	// +kubebuilder:validation:Optional
	IAM *bool `json:"iam,omitempty" tf:"iam,omitempty"`

	// +kubebuilder:validation:Optional
	Scram *bool `json:"scram,omitempty" tf:"scram,omitempty"`
}

type TLSObservation struct {
}

type TLSParameters struct {

	// +kubebuilder:validation:Optional
	CertificateAuthorityArns []*string `json:"certificateAuthorityArns,omitempty" tf:"certificate_authority_arns,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}