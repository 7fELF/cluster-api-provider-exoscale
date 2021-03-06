/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha1

import (
	"github.com/exoscale/egoscale"
	yaml "github.com/ghodss/yaml"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExoscaleClusterProviderStatus is the Schema for the exoscaleclusterproviderstatuses API
// +k8s:openapi-gen=true
type ExoscaleClusterProviderStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	MasterSecurityGroupID *egoscale.UUID `json:"masterSecurityGroupID"`
	NodeSecurityGroupID   *egoscale.UUID `json:"nodeSecurityGroupID"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

func init() {
	SchemeBuilder.Register(&ExoscaleClusterProviderStatus{})
}

// ClusterStatusFromProviderStatus return cluster provider specs from cluster provider custom resources (/config/crds)
func ClusterStatusFromProviderStatus(providerStatus *runtime.RawExtension) (*ExoscaleClusterProviderStatus, error) {
	config := new(ExoscaleClusterProviderStatus)
	if providerStatus != nil {
		if err := yaml.Unmarshal(providerStatus.Raw, config); err != nil {
			return nil, err
		}
	}
	return config, nil
}
