// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EventBusSpec defines the desired state of EventBus.
//
// An event bus receives events from a source, uses rules to evaluate them,
// applies any configured input transformation, and routes them to the appropriate
// target(s). Your account's default event bus receives events from Amazon Web
// Services services. A custom event bus can receive events from your custom
// applications and services. A partner event bus receives events from an event
// source created by an SaaS partner. These events come from the partners services
// or applications.
type EventBusSpec struct {
	DeadLetterConfig *DeadLetterConfig `json:"deadLetterConfig,omitempty"`
	// The event bus description.
	Description *string `json:"description,omitempty"`
	// If you are creating a partner event bus, this specifies the partner event
	// source that the new event bus will be matched with.
	EventSourceName *string `json:"eventSourceName,omitempty"`
	// The identifier of the KMS customer managed key for EventBridge to use, if
	// you choose to use a customer managed key to encrypt events on this event
	// bus. The identifier can be the key Amazon Resource Name (ARN), KeyId, key
	// alias, or key alias ARN.
	//
	// If you do not specify a customer managed key identifier, EventBridge uses
	// an Amazon Web Services owned key to encrypt events on the event bus.
	//
	// For more information, see Managing keys (https://docs.aws.amazon.com/kms/latest/developerguide/getting-started.html)
	// in the Key Management Service Developer Guide.
	//
	// Archives and schema discovery are not supported for event buses encrypted
	// using a customer managed key. EventBridge returns an error if:
	//
	//   - You call CreateArchive (https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_CreateArchive.html)
	//     on an event bus set to use a customer managed key for encryption.
	//
	//   - You call CreateDiscoverer (https://docs.aws.amazon.com/eventbridge/latest/schema-reference/v1-discoverers.html#CreateDiscoverer)
	//     on an event bus set to use a customer managed key for encryption.
	//
	//   - You call UpdatedEventBus (https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_UpdatedEventBus.html)
	//     to set a customer managed key on an event bus with an archives or schema
	//     discovery enabled.
	//
	// To enable archives or schema discovery on an event bus, choose to use an
	// Amazon Web Services owned key. For more information, see Data encryption
	// in EventBridge (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-encryption.html)
	// in the Amazon EventBridge User Guide.
	KMSKeyIdentifier *string `json:"kmsKeyIdentifier,omitempty"`
	// The name of the new event bus.
	//
	// Custom event bus names can't contain the / character, but you can use the
	// / character in partner event bus names. In addition, for partner event buses,
	// the name must exactly match the name of the partner event source that this
	// event bus is matched to.
	//
	// You can't use the name default for a custom event bus, as this name is already
	// used for your account's default event bus.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// Tags to associate with the event bus.
	Tags []*Tag `json:"tags,omitempty"`
}

// EventBusStatus defines the observed state of EventBus
type EventBusStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
}

// EventBus is the Schema for the EventBuses API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ARN",type=string,priority=1,JSONPath=`.status.ackResourceMetadata.arn`
// +kubebuilder:printcolumn:name="Synced",type="string",priority=0,JSONPath=".status.conditions[?(@.type==\"ACK.ResourceSynced\")].status"
// +kubebuilder:printcolumn:name="Age",type="date",priority=0,JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:shortName=eb;bus
type EventBus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventBusSpec   `json:"spec,omitempty"`
	Status            EventBusStatus `json:"status,omitempty"`
}

// EventBusList contains a list of EventBus
// +kubebuilder:object:root=true
type EventBusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventBus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventBus{}, &EventBusList{})
}
