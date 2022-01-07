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

type AutomationParametersObservation struct {
}

type AutomationParametersParameterObservation struct {
}

type AutomationParametersParameterParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type AutomationParametersParameters struct {

	// +kubebuilder:validation:Optional
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	// +kubebuilder:validation:Optional
	Parameter []AutomationParametersParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type CloudwatchConfigObservation struct {
}

type CloudwatchConfigParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchLogGroupName *string `json:"cloudwatchLogGroupName,omitempty" tf:"cloudwatch_log_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	CloudwatchOutputEnabled *bool `json:"cloudwatchOutputEnabled,omitempty" tf:"cloudwatch_output_enabled,omitempty"`
}

type LambdaParametersObservation struct {
}

type LambdaParametersParameters struct {

	// +kubebuilder:validation:Optional
	ClientContext *string `json:"clientContext,omitempty" tf:"client_context,omitempty"`

	// +kubebuilder:validation:Optional
	PayloadSecretRef *v1.SecretKeySelector `json:"payloadSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Qualifier *string `json:"qualifier,omitempty" tf:"qualifier,omitempty"`
}

type MaintenanceWindowTaskObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MaintenanceWindowTaskParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	MaxConcurrency *string `json:"maxConcurrency" tf:"max_concurrency,omitempty"`

	// +kubebuilder:validation:Required
	MaxErrors *string `json:"maxErrors" tf:"max_errors,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *int64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2.Role
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceRoleArnRef *v1.Reference `json:"serviceRoleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceRoleArnSelector *v1.Selector `json:"serviceRoleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Targets []MaintenanceWindowTaskTargetsParameters `json:"targets,omitempty" tf:"targets,omitempty"`

	// +kubebuilder:validation:Required
	TaskArn *string `json:"taskArn" tf:"task_arn,omitempty"`

	// +kubebuilder:validation:Optional
	TaskInvocationParameters []TaskInvocationParametersParameters `json:"taskInvocationParameters,omitempty" tf:"task_invocation_parameters,omitempty"`

	// +kubebuilder:validation:Required
	TaskType *string `json:"taskType" tf:"task_type,omitempty"`

	// +kubebuilder:validation:Required
	WindowID *string `json:"windowId" tf:"window_id,omitempty"`
}

type MaintenanceWindowTaskTargetsObservation struct {
}

type MaintenanceWindowTaskTargetsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type NotificationConfigObservation struct {
}

type NotificationConfigParameters struct {

	// +kubebuilder:validation:Optional
	NotificationArn *string `json:"notificationArn,omitempty" tf:"notification_arn,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationEvents []*string `json:"notificationEvents,omitempty" tf:"notification_events,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationType *string `json:"notificationType,omitempty" tf:"notification_type,omitempty"`
}

type RunCommandParametersObservation struct {
}

type RunCommandParametersParameterObservation struct {
}

type RunCommandParametersParameterParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type RunCommandParametersParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchConfig []CloudwatchConfigParameters `json:"cloudwatchConfig,omitempty" tf:"cloudwatch_config,omitempty"`

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Optional
	DocumentHash *string `json:"documentHash,omitempty" tf:"document_hash,omitempty"`

	// +kubebuilder:validation:Optional
	DocumentHashType *string `json:"documentHashType,omitempty" tf:"document_hash_type,omitempty"`

	// +kubebuilder:validation:Optional
	DocumentVersion *string `json:"documentVersion,omitempty" tf:"document_version,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationConfig []NotificationConfigParameters `json:"notificationConfig,omitempty" tf:"notification_config,omitempty"`

	// +kubebuilder:validation:Optional
	OutputS3Bucket *string `json:"outputS3Bucket,omitempty" tf:"output_s3_bucket,omitempty"`

	// +kubebuilder:validation:Optional
	OutputS3KeyPrefix *string `json:"outputS3KeyPrefix,omitempty" tf:"output_s3_key_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Parameter []RunCommandParametersParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`
}

type StepFunctionsParametersObservation struct {
}

type StepFunctionsParametersParameters struct {

	// +kubebuilder:validation:Optional
	InputSecretRef *v1.SecretKeySelector `json:"inputSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TaskInvocationParametersObservation struct {
}

type TaskInvocationParametersParameters struct {

	// +kubebuilder:validation:Optional
	AutomationParameters []AutomationParametersParameters `json:"automationParameters,omitempty" tf:"automation_parameters,omitempty"`

	// +kubebuilder:validation:Optional
	LambdaParameters []LambdaParametersParameters `json:"lambdaParameters,omitempty" tf:"lambda_parameters,omitempty"`

	// +kubebuilder:validation:Optional
	RunCommandParameters []RunCommandParametersParameters `json:"runCommandParameters,omitempty" tf:"run_command_parameters,omitempty"`

	// +kubebuilder:validation:Optional
	StepFunctionsParameters []StepFunctionsParametersParameters `json:"stepFunctionsParameters,omitempty" tf:"step_functions_parameters,omitempty"`
}

// MaintenanceWindowTaskSpec defines the desired state of MaintenanceWindowTask
type MaintenanceWindowTaskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MaintenanceWindowTaskParameters `json:"forProvider"`
}

// MaintenanceWindowTaskStatus defines the observed state of MaintenanceWindowTask.
type MaintenanceWindowTaskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MaintenanceWindowTaskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceWindowTask is the Schema for the MaintenanceWindowTasks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type MaintenanceWindowTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MaintenanceWindowTaskSpec   `json:"spec"`
	Status            MaintenanceWindowTaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceWindowTaskList contains a list of MaintenanceWindowTasks
type MaintenanceWindowTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MaintenanceWindowTask `json:"items"`
}

// Repository type metadata.
var (
	MaintenanceWindowTask_Kind             = "MaintenanceWindowTask"
	MaintenanceWindowTask_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MaintenanceWindowTask_Kind}.String()
	MaintenanceWindowTask_KindAPIVersion   = MaintenanceWindowTask_Kind + "." + CRDGroupVersion.String()
	MaintenanceWindowTask_GroupVersionKind = CRDGroupVersion.WithKind(MaintenanceWindowTask_Kind)
)

func init() {
	SchemeBuilder.Register(&MaintenanceWindowTask{}, &MaintenanceWindowTaskList{})
}