// Copyright © 2021 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cluster is the Schema for the InfraMetadata API
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

type ClusterStatus struct {
}

type SSH struct {
	User     string `json:"user,omitempty"`
	Passwd   string `json:"passwd,omitempty"`
	Pk       string `json:"pk,omitempty"`
	PkPasswd string `json:"pkPasswd,omitempty"`
	Port     uint16 `json:"port,omitempty"`
}

type Host struct {
	IPS   []string `json:"ips,omitempty"`
	Roles []string `json:"roles,omitempty"`
	//overwrite env
	Env []string `json:"env,omitempty"`
}

type ImageList []string

// ClusterSpec defines the desired state of InfraMetadata
type ClusterSpec struct {
	// desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Foo is an example field of Cluster. Edit Cluster_types.go to remove/update
	Image ImageList `json:"image,omitempty"`
	SSH   SSH       `json:"ssh"`
	Hosts []Host    `json:"hosts,omitempty"`
	// Why env not using map[string]string
	// Because some argument is list, like: CertSANS=127.0.0.1 CertSANS=localhost, if ENV is map, will merge those two values
	// but user want to InfraMetadata a list, using array we can convert it to {CertSANS:[127.0.0.1, localhost]}
	Env []string `json:"env,omitempty"`
	// Entrypoint array. Not executed within a shell.
	// The docker image's ENTRYPOINT is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
	// to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
	// produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
	// of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	Command []string `json:"command,omitempty"`
}
