/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureVMSchedulerStopNHSpec defines the desired state of AzureVMSchedulerStopNH
type AzureVMSchedulerStopNHSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of AzureVMSchedulerStopNH. Edit azurevmschedulerstopnh_types.go to remove/update
	AzureVMName string `json:"azureVMName"`
	AzureRGName string `json:"azureRGName"`
	Command     string `json:"command"`

	// Schedule period for the CronJob.
	// This spec allows you to setup the start schedule for VM
	StopSchedule string `json:"stopSchedule"`
	Image        string `json:"image"`
}

// AzureVMSchedulerStopNHStatus defines the observed state of AzureVMSchedulerStopNH
type AzureVMSchedulerStopNHStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	VMStopStatus string `json:"vmStopStatus"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AzureVMSchedulerStopNH is the Schema for the azurevmschedulerstopnhs API
type AzureVMSchedulerStopNH struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureVMSchedulerStopNHSpec   `json:"spec,omitempty"`
	Status AzureVMSchedulerStopNHStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AzureVMSchedulerStopNHList contains a list of AzureVMSchedulerStopNH
type AzureVMSchedulerStopNHList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureVMSchedulerStopNH `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureVMSchedulerStopNH{}, &AzureVMSchedulerStopNHList{})
}
