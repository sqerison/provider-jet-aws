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
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/terrajet/pkg/conversion"
	"github.com/crossplane-contrib/terrajet/pkg/json"
)

// GetTerraformResourceType returns Terraform resource type for this RdsClusterEndpoint
func (mg *RdsClusterEndpoint) GetTerraformResourceType() string {
	return "aws_rds_cluster_endpoint"
}

// GetTerraformResourceIdField returns Terraform identifier field for this RdsClusterEndpoint
func (tr *RdsClusterEndpoint) GetTerraformResourceIdField() string {
	return "id"
}

// GetObservation of this RdsClusterEndpoint
func (tr *RdsClusterEndpoint) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RdsClusterEndpoint
func (tr *RdsClusterEndpoint) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetParameters of this RdsClusterEndpoint
func (tr *RdsClusterEndpoint) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RdsClusterEndpoint
func (tr *RdsClusterEndpoint) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this RdsClusterEndpoint using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RdsClusterEndpoint) LateInitialize(attrs []byte) (bool, error) {
	params := &RdsClusterEndpointParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	li := conversion.NewLateInitializer(conversion.WithZeroValueJSONOmitEmptyFilter(conversion.CNameWildcard),
		conversion.WithZeroElemPtrFilter(conversion.CNameWildcard))
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}
