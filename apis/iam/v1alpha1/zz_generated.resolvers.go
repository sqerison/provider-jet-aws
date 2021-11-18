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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	common "github.com/crossplane-contrib/provider-jet-aws/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AccessKey.
func (mg *AccessKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.User),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserRef,
		Selector:     mg.Spec.ForProvider.UserSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.User")
	}
	mg.Spec.ForProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupPolicy.
func (mg *GroupPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Group),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupRef,
		Selector:     mg.Spec.ForProvider.GroupSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Group")
	}
	mg.Spec.ForProvider.Group = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GroupPolicyAttachment.
func (mg *GroupPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Group),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupRef,
		Selector:     mg.Spec.ForProvider.GroupSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Group")
	}
	mg.Spec.ForProvider.Group = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.PolicyArnRef,
		Selector:     mg.Spec.ForProvider.PolicyArnSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyArn")
	}
	mg.Spec.ForProvider.PolicyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this InstanceProfile.
func (mg *InstanceProfile) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RoleRef,
		Selector:     mg.Spec.ForProvider.RoleSelector,
		To: reference.To{
			List:    &RoleList{},
			Managed: &Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PolicyAttachment.
func (mg *PolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Groups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.GroupRefs,
		Selector:      mg.Spec.ForProvider.GroupSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Groups")
	}
	mg.Spec.ForProvider.Groups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.GroupRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.PolicyArnRef,
		Selector:     mg.Spec.ForProvider.PolicyArnSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyArn")
	}
	mg.Spec.ForProvider.PolicyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyArnRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Roles),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.RoleRefs,
		Selector:      mg.Spec.ForProvider.RoleSelector,
		To: reference.To{
			List:    &RoleList{},
			Managed: &Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Roles")
	}
	mg.Spec.ForProvider.Roles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.RoleRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Users),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.UserRefs,
		Selector:      mg.Spec.ForProvider.UserSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Users")
	}
	mg.Spec.ForProvider.Users = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.UserRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this RolePolicyAttachment.
func (mg *RolePolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.PolicyArnRef,
		Selector:     mg.Spec.ForProvider.PolicyArnSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyArn")
	}
	mg.Spec.ForProvider.PolicyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Role),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RoleRef,
		Selector:     mg.Spec.ForProvider.RoleSelector,
		To: reference.To{
			List:    &RoleList{},
			Managed: &Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Role")
	}
	mg.Spec.ForProvider.Role = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserGroupMembership.
func (mg *UserGroupMembership) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Groups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.GroupRefs,
		Selector:      mg.Spec.ForProvider.GroupSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Groups")
	}
	mg.Spec.ForProvider.Groups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.GroupRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.User),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserRef,
		Selector:     mg.Spec.ForProvider.UserSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.User")
	}
	mg.Spec.ForProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPolicy.
func (mg *UserPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.User),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserRef,
		Selector:     mg.Spec.ForProvider.UserSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.User")
	}
	mg.Spec.ForProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this UserPolicyAttachment.
func (mg *UserPolicyAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.PolicyArnRef,
		Selector:     mg.Spec.ForProvider.PolicyArnSelector,
		To: reference.To{
			List:    &PolicyList{},
			Managed: &Policy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyArn")
	}
	mg.Spec.ForProvider.PolicyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.User),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.UserRef,
		Selector:     mg.Spec.ForProvider.UserSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.User")
	}
	mg.Spec.ForProvider.User = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.UserRef = rsp.ResolvedReference

	return nil
}
