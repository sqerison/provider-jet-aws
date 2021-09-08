// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingPolicy) DeepCopyInto(out *AppautoscalingPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingPolicy.
func (in *AppautoscalingPolicy) DeepCopy() *AppautoscalingPolicy {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppautoscalingPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingPolicyList) DeepCopyInto(out *AppautoscalingPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppautoscalingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingPolicyList.
func (in *AppautoscalingPolicyList) DeepCopy() *AppautoscalingPolicyList {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppautoscalingPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingPolicyObservation) DeepCopyInto(out *AppautoscalingPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingPolicyObservation.
func (in *AppautoscalingPolicyObservation) DeepCopy() *AppautoscalingPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingPolicyParameters) DeepCopyInto(out *AppautoscalingPolicyParameters) {
	*out = *in
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
	if in.StepScalingPolicyConfiguration != nil {
		in, out := &in.StepScalingPolicyConfiguration, &out.StepScalingPolicyConfiguration
		*out = make([]StepScalingPolicyConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TargetTrackingScalingPolicyConfiguration != nil {
		in, out := &in.TargetTrackingScalingPolicyConfiguration, &out.TargetTrackingScalingPolicyConfiguration
		*out = make([]TargetTrackingScalingPolicyConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingPolicyParameters.
func (in *AppautoscalingPolicyParameters) DeepCopy() *AppautoscalingPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingPolicySpec) DeepCopyInto(out *AppautoscalingPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingPolicySpec.
func (in *AppautoscalingPolicySpec) DeepCopy() *AppautoscalingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingPolicyStatus) DeepCopyInto(out *AppautoscalingPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingPolicyStatus.
func (in *AppautoscalingPolicyStatus) DeepCopy() *AppautoscalingPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingScheduledAction) DeepCopyInto(out *AppautoscalingScheduledAction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingScheduledAction.
func (in *AppautoscalingScheduledAction) DeepCopy() *AppautoscalingScheduledAction {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingScheduledAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppautoscalingScheduledAction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingScheduledActionList) DeepCopyInto(out *AppautoscalingScheduledActionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppautoscalingScheduledAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingScheduledActionList.
func (in *AppautoscalingScheduledActionList) DeepCopy() *AppautoscalingScheduledActionList {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingScheduledActionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppautoscalingScheduledActionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingScheduledActionObservation) DeepCopyInto(out *AppautoscalingScheduledActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingScheduledActionObservation.
func (in *AppautoscalingScheduledActionObservation) DeepCopy() *AppautoscalingScheduledActionObservation {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingScheduledActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingScheduledActionParameters) DeepCopyInto(out *AppautoscalingScheduledActionParameters) {
	*out = *in
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(string)
		**out = **in
	}
	if in.ScalableTargetAction != nil {
		in, out := &in.ScalableTargetAction, &out.ScalableTargetAction
		*out = make([]ScalableTargetActionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingScheduledActionParameters.
func (in *AppautoscalingScheduledActionParameters) DeepCopy() *AppautoscalingScheduledActionParameters {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingScheduledActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingScheduledActionSpec) DeepCopyInto(out *AppautoscalingScheduledActionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingScheduledActionSpec.
func (in *AppautoscalingScheduledActionSpec) DeepCopy() *AppautoscalingScheduledActionSpec {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingScheduledActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingScheduledActionStatus) DeepCopyInto(out *AppautoscalingScheduledActionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingScheduledActionStatus.
func (in *AppautoscalingScheduledActionStatus) DeepCopy() *AppautoscalingScheduledActionStatus {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingScheduledActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingTarget) DeepCopyInto(out *AppautoscalingTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTarget.
func (in *AppautoscalingTarget) DeepCopy() *AppautoscalingTarget {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppautoscalingTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingTargetList) DeepCopyInto(out *AppautoscalingTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppautoscalingTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTargetList.
func (in *AppautoscalingTargetList) DeepCopy() *AppautoscalingTargetList {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppautoscalingTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingTargetObservation) DeepCopyInto(out *AppautoscalingTargetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTargetObservation.
func (in *AppautoscalingTargetObservation) DeepCopy() *AppautoscalingTargetObservation {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingTargetParameters) DeepCopyInto(out *AppautoscalingTargetParameters) {
	*out = *in
	if in.RoleArn != nil {
		in, out := &in.RoleArn, &out.RoleArn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTargetParameters.
func (in *AppautoscalingTargetParameters) DeepCopy() *AppautoscalingTargetParameters {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTargetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingTargetSpec) DeepCopyInto(out *AppautoscalingTargetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTargetSpec.
func (in *AppautoscalingTargetSpec) DeepCopy() *AppautoscalingTargetSpec {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppautoscalingTargetStatus) DeepCopyInto(out *AppautoscalingTargetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppautoscalingTargetStatus.
func (in *AppautoscalingTargetStatus) DeepCopy() *AppautoscalingTargetStatus {
	if in == nil {
		return nil
	}
	out := new(AppautoscalingTargetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedMetricSpecificationObservation) DeepCopyInto(out *CustomizedMetricSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedMetricSpecificationObservation.
func (in *CustomizedMetricSpecificationObservation) DeepCopy() *CustomizedMetricSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(CustomizedMetricSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedMetricSpecificationParameters) DeepCopyInto(out *CustomizedMetricSpecificationParameters) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]DimensionsParameters, len(*in))
		copy(*out, *in)
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedMetricSpecificationParameters.
func (in *CustomizedMetricSpecificationParameters) DeepCopy() *CustomizedMetricSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(CustomizedMetricSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionsObservation) DeepCopyInto(out *DimensionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionsObservation.
func (in *DimensionsObservation) DeepCopy() *DimensionsObservation {
	if in == nil {
		return nil
	}
	out := new(DimensionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DimensionsParameters) DeepCopyInto(out *DimensionsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DimensionsParameters.
func (in *DimensionsParameters) DeepCopy() *DimensionsParameters {
	if in == nil {
		return nil
	}
	out := new(DimensionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredefinedMetricSpecificationObservation) DeepCopyInto(out *PredefinedMetricSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredefinedMetricSpecificationObservation.
func (in *PredefinedMetricSpecificationObservation) DeepCopy() *PredefinedMetricSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(PredefinedMetricSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredefinedMetricSpecificationParameters) DeepCopyInto(out *PredefinedMetricSpecificationParameters) {
	*out = *in
	if in.ResourceLabel != nil {
		in, out := &in.ResourceLabel, &out.ResourceLabel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredefinedMetricSpecificationParameters.
func (in *PredefinedMetricSpecificationParameters) DeepCopy() *PredefinedMetricSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(PredefinedMetricSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetActionObservation) DeepCopyInto(out *ScalableTargetActionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetActionObservation.
func (in *ScalableTargetActionObservation) DeepCopy() *ScalableTargetActionObservation {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetActionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetActionParameters) DeepCopyInto(out *ScalableTargetActionParameters) {
	*out = *in
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(string)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetActionParameters.
func (in *ScalableTargetActionParameters) DeepCopy() *ScalableTargetActionParameters {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetActionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepAdjustmentObservation) DeepCopyInto(out *StepAdjustmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepAdjustmentObservation.
func (in *StepAdjustmentObservation) DeepCopy() *StepAdjustmentObservation {
	if in == nil {
		return nil
	}
	out := new(StepAdjustmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepAdjustmentParameters) DeepCopyInto(out *StepAdjustmentParameters) {
	*out = *in
	if in.MetricIntervalLowerBound != nil {
		in, out := &in.MetricIntervalLowerBound, &out.MetricIntervalLowerBound
		*out = new(string)
		**out = **in
	}
	if in.MetricIntervalUpperBound != nil {
		in, out := &in.MetricIntervalUpperBound, &out.MetricIntervalUpperBound
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepAdjustmentParameters.
func (in *StepAdjustmentParameters) DeepCopy() *StepAdjustmentParameters {
	if in == nil {
		return nil
	}
	out := new(StepAdjustmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepScalingPolicyConfigurationObservation) DeepCopyInto(out *StepScalingPolicyConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepScalingPolicyConfigurationObservation.
func (in *StepScalingPolicyConfigurationObservation) DeepCopy() *StepScalingPolicyConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(StepScalingPolicyConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StepScalingPolicyConfigurationParameters) DeepCopyInto(out *StepScalingPolicyConfigurationParameters) {
	*out = *in
	if in.AdjustmentType != nil {
		in, out := &in.AdjustmentType, &out.AdjustmentType
		*out = new(string)
		**out = **in
	}
	if in.Cooldown != nil {
		in, out := &in.Cooldown, &out.Cooldown
		*out = new(int64)
		**out = **in
	}
	if in.MetricAggregationType != nil {
		in, out := &in.MetricAggregationType, &out.MetricAggregationType
		*out = new(string)
		**out = **in
	}
	if in.MinAdjustmentMagnitude != nil {
		in, out := &in.MinAdjustmentMagnitude, &out.MinAdjustmentMagnitude
		*out = new(int64)
		**out = **in
	}
	if in.StepAdjustment != nil {
		in, out := &in.StepAdjustment, &out.StepAdjustment
		*out = make([]StepAdjustmentParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StepScalingPolicyConfigurationParameters.
func (in *StepScalingPolicyConfigurationParameters) DeepCopy() *StepScalingPolicyConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(StepScalingPolicyConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetTrackingScalingPolicyConfigurationObservation) DeepCopyInto(out *TargetTrackingScalingPolicyConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetTrackingScalingPolicyConfigurationObservation.
func (in *TargetTrackingScalingPolicyConfigurationObservation) DeepCopy() *TargetTrackingScalingPolicyConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(TargetTrackingScalingPolicyConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetTrackingScalingPolicyConfigurationParameters) DeepCopyInto(out *TargetTrackingScalingPolicyConfigurationParameters) {
	*out = *in
	if in.CustomizedMetricSpecification != nil {
		in, out := &in.CustomizedMetricSpecification, &out.CustomizedMetricSpecification
		*out = make([]CustomizedMetricSpecificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisableScaleIn != nil {
		in, out := &in.DisableScaleIn, &out.DisableScaleIn
		*out = new(bool)
		**out = **in
	}
	if in.PredefinedMetricSpecification != nil {
		in, out := &in.PredefinedMetricSpecification, &out.PredefinedMetricSpecification
		*out = make([]PredefinedMetricSpecificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScaleInCooldown != nil {
		in, out := &in.ScaleInCooldown, &out.ScaleInCooldown
		*out = new(int64)
		**out = **in
	}
	if in.ScaleOutCooldown != nil {
		in, out := &in.ScaleOutCooldown, &out.ScaleOutCooldown
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetTrackingScalingPolicyConfigurationParameters.
func (in *TargetTrackingScalingPolicyConfigurationParameters) DeepCopy() *TargetTrackingScalingPolicyConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(TargetTrackingScalingPolicyConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}