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
func (in *ApplicationSourceObservation) DeepCopyInto(out *ApplicationSourceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSourceObservation.
func (in *ApplicationSourceObservation) DeepCopy() *ApplicationSourceObservation {
	if in == nil {
		return nil
	}
	out := new(ApplicationSourceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSourceParameters) DeepCopyInto(out *ApplicationSourceParameters) {
	*out = *in
	if in.CloudFormationStackArn != nil {
		in, out := &in.CloudFormationStackArn, &out.CloudFormationStackArn
		*out = new(string)
		**out = **in
	}
	if in.TagFilter != nil {
		in, out := &in.TagFilter, &out.TagFilter
		*out = make([]TagFilterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSourceParameters.
func (in *ApplicationSourceParameters) DeepCopy() *ApplicationSourceParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationSourceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedLoadMetricSpecificationObservation) DeepCopyInto(out *CustomizedLoadMetricSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedLoadMetricSpecificationObservation.
func (in *CustomizedLoadMetricSpecificationObservation) DeepCopy() *CustomizedLoadMetricSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(CustomizedLoadMetricSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedLoadMetricSpecificationParameters) DeepCopyInto(out *CustomizedLoadMetricSpecificationParameters) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedLoadMetricSpecificationParameters.
func (in *CustomizedLoadMetricSpecificationParameters) DeepCopy() *CustomizedLoadMetricSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(CustomizedLoadMetricSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedScalingMetricSpecificationObservation) DeepCopyInto(out *CustomizedScalingMetricSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedScalingMetricSpecificationObservation.
func (in *CustomizedScalingMetricSpecificationObservation) DeepCopy() *CustomizedScalingMetricSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(CustomizedScalingMetricSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomizedScalingMetricSpecificationParameters) DeepCopyInto(out *CustomizedScalingMetricSpecificationParameters) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomizedScalingMetricSpecificationParameters.
func (in *CustomizedScalingMetricSpecificationParameters) DeepCopy() *CustomizedScalingMetricSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(CustomizedScalingMetricSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredefinedLoadMetricSpecificationObservation) DeepCopyInto(out *PredefinedLoadMetricSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredefinedLoadMetricSpecificationObservation.
func (in *PredefinedLoadMetricSpecificationObservation) DeepCopy() *PredefinedLoadMetricSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(PredefinedLoadMetricSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredefinedLoadMetricSpecificationParameters) DeepCopyInto(out *PredefinedLoadMetricSpecificationParameters) {
	*out = *in
	if in.PredefinedLoadMetricType != nil {
		in, out := &in.PredefinedLoadMetricType, &out.PredefinedLoadMetricType
		*out = new(string)
		**out = **in
	}
	if in.ResourceLabel != nil {
		in, out := &in.ResourceLabel, &out.ResourceLabel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredefinedLoadMetricSpecificationParameters.
func (in *PredefinedLoadMetricSpecificationParameters) DeepCopy() *PredefinedLoadMetricSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(PredefinedLoadMetricSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredefinedScalingMetricSpecificationObservation) DeepCopyInto(out *PredefinedScalingMetricSpecificationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredefinedScalingMetricSpecificationObservation.
func (in *PredefinedScalingMetricSpecificationObservation) DeepCopy() *PredefinedScalingMetricSpecificationObservation {
	if in == nil {
		return nil
	}
	out := new(PredefinedScalingMetricSpecificationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredefinedScalingMetricSpecificationParameters) DeepCopyInto(out *PredefinedScalingMetricSpecificationParameters) {
	*out = *in
	if in.PredefinedScalingMetricType != nil {
		in, out := &in.PredefinedScalingMetricType, &out.PredefinedScalingMetricType
		*out = new(string)
		**out = **in
	}
	if in.ResourceLabel != nil {
		in, out := &in.ResourceLabel, &out.ResourceLabel
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredefinedScalingMetricSpecificationParameters.
func (in *PredefinedScalingMetricSpecificationParameters) DeepCopy() *PredefinedScalingMetricSpecificationParameters {
	if in == nil {
		return nil
	}
	out := new(PredefinedScalingMetricSpecificationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingInstructionObservation) DeepCopyInto(out *ScalingInstructionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingInstructionObservation.
func (in *ScalingInstructionObservation) DeepCopy() *ScalingInstructionObservation {
	if in == nil {
		return nil
	}
	out := new(ScalingInstructionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingInstructionParameters) DeepCopyInto(out *ScalingInstructionParameters) {
	*out = *in
	if in.CustomizedLoadMetricSpecification != nil {
		in, out := &in.CustomizedLoadMetricSpecification, &out.CustomizedLoadMetricSpecification
		*out = make([]CustomizedLoadMetricSpecificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisableDynamicScaling != nil {
		in, out := &in.DisableDynamicScaling, &out.DisableDynamicScaling
		*out = new(bool)
		**out = **in
	}
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(int64)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(int64)
		**out = **in
	}
	if in.PredefinedLoadMetricSpecification != nil {
		in, out := &in.PredefinedLoadMetricSpecification, &out.PredefinedLoadMetricSpecification
		*out = make([]PredefinedLoadMetricSpecificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PredictiveScalingMaxCapacityBehavior != nil {
		in, out := &in.PredictiveScalingMaxCapacityBehavior, &out.PredictiveScalingMaxCapacityBehavior
		*out = new(string)
		**out = **in
	}
	if in.PredictiveScalingMaxCapacityBuffer != nil {
		in, out := &in.PredictiveScalingMaxCapacityBuffer, &out.PredictiveScalingMaxCapacityBuffer
		*out = new(int64)
		**out = **in
	}
	if in.PredictiveScalingMode != nil {
		in, out := &in.PredictiveScalingMode, &out.PredictiveScalingMode
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ScalableDimension != nil {
		in, out := &in.ScalableDimension, &out.ScalableDimension
		*out = new(string)
		**out = **in
	}
	if in.ScalingPolicyUpdateBehavior != nil {
		in, out := &in.ScalingPolicyUpdateBehavior, &out.ScalingPolicyUpdateBehavior
		*out = new(string)
		**out = **in
	}
	if in.ScheduledActionBufferTime != nil {
		in, out := &in.ScheduledActionBufferTime, &out.ScheduledActionBufferTime
		*out = new(int64)
		**out = **in
	}
	if in.ServiceNamespace != nil {
		in, out := &in.ServiceNamespace, &out.ServiceNamespace
		*out = new(string)
		**out = **in
	}
	if in.TargetTrackingConfiguration != nil {
		in, out := &in.TargetTrackingConfiguration, &out.TargetTrackingConfiguration
		*out = make([]TargetTrackingConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingInstructionParameters.
func (in *ScalingInstructionParameters) DeepCopy() *ScalingInstructionParameters {
	if in == nil {
		return nil
	}
	out := new(ScalingInstructionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlan) DeepCopyInto(out *ScalingPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlan.
func (in *ScalingPlan) DeepCopy() *ScalingPlan {
	if in == nil {
		return nil
	}
	out := new(ScalingPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanList) DeepCopyInto(out *ScalingPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScalingPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanList.
func (in *ScalingPlanList) DeepCopy() *ScalingPlanList {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanObservation) DeepCopyInto(out *ScalingPlanObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ScalingPlanVersion != nil {
		in, out := &in.ScalingPlanVersion, &out.ScalingPlanVersion
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanObservation.
func (in *ScalingPlanObservation) DeepCopy() *ScalingPlanObservation {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanParameters) DeepCopyInto(out *ScalingPlanParameters) {
	*out = *in
	if in.ApplicationSource != nil {
		in, out := &in.ApplicationSource, &out.ApplicationSource
		*out = make([]ApplicationSourceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ScalingInstruction != nil {
		in, out := &in.ScalingInstruction, &out.ScalingInstruction
		*out = make([]ScalingInstructionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanParameters.
func (in *ScalingPlanParameters) DeepCopy() *ScalingPlanParameters {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpec) DeepCopyInto(out *ScalingPlanSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpec.
func (in *ScalingPlanSpec) DeepCopy() *ScalingPlanSpec {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanStatus) DeepCopyInto(out *ScalingPlanStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanStatus.
func (in *ScalingPlanStatus) DeepCopy() *ScalingPlanStatus {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagFilterObservation) DeepCopyInto(out *TagFilterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagFilterObservation.
func (in *TagFilterObservation) DeepCopy() *TagFilterObservation {
	if in == nil {
		return nil
	}
	out := new(TagFilterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagFilterParameters) DeepCopyInto(out *TagFilterParameters) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagFilterParameters.
func (in *TagFilterParameters) DeepCopy() *TagFilterParameters {
	if in == nil {
		return nil
	}
	out := new(TagFilterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetTrackingConfigurationObservation) DeepCopyInto(out *TargetTrackingConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetTrackingConfigurationObservation.
func (in *TargetTrackingConfigurationObservation) DeepCopy() *TargetTrackingConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(TargetTrackingConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetTrackingConfigurationParameters) DeepCopyInto(out *TargetTrackingConfigurationParameters) {
	*out = *in
	if in.CustomizedScalingMetricSpecification != nil {
		in, out := &in.CustomizedScalingMetricSpecification, &out.CustomizedScalingMetricSpecification
		*out = make([]CustomizedScalingMetricSpecificationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisableScaleIn != nil {
		in, out := &in.DisableScaleIn, &out.DisableScaleIn
		*out = new(bool)
		**out = **in
	}
	if in.EstimatedInstanceWarmup != nil {
		in, out := &in.EstimatedInstanceWarmup, &out.EstimatedInstanceWarmup
		*out = new(int64)
		**out = **in
	}
	if in.PredefinedScalingMetricSpecification != nil {
		in, out := &in.PredefinedScalingMetricSpecification, &out.PredefinedScalingMetricSpecification
		*out = make([]PredefinedScalingMetricSpecificationParameters, len(*in))
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
	if in.TargetValue != nil {
		in, out := &in.TargetValue, &out.TargetValue
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetTrackingConfigurationParameters.
func (in *TargetTrackingConfigurationParameters) DeepCopy() *TargetTrackingConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(TargetTrackingConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}