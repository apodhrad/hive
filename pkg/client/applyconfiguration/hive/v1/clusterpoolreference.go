// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterPoolReferenceApplyConfiguration represents an declarative configuration of the ClusterPoolReference type for use
// with apply.
type ClusterPoolReferenceApplyConfiguration struct {
	Namespace        *string                      `json:"namespace,omitempty"`
	PoolName         *string                      `json:"poolName,omitempty"`
	ClaimName        *string                      `json:"claimName,omitempty"`
	ClaimedTimestamp *v1.Time                     `json:"claimedTimestamp,omitempty"`
	CustomizationRef *corev1.LocalObjectReference `json:"clusterDeploymentCustomization,omitempty"`
}

// ClusterPoolReferenceApplyConfiguration constructs an declarative configuration of the ClusterPoolReference type for use with
// apply.
func ClusterPoolReference() *ClusterPoolReferenceApplyConfiguration {
	return &ClusterPoolReferenceApplyConfiguration{}
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ClusterPoolReferenceApplyConfiguration) WithNamespace(value string) *ClusterPoolReferenceApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithPoolName sets the PoolName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PoolName field is set to the value of the last call.
func (b *ClusterPoolReferenceApplyConfiguration) WithPoolName(value string) *ClusterPoolReferenceApplyConfiguration {
	b.PoolName = &value
	return b
}

// WithClaimName sets the ClaimName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClaimName field is set to the value of the last call.
func (b *ClusterPoolReferenceApplyConfiguration) WithClaimName(value string) *ClusterPoolReferenceApplyConfiguration {
	b.ClaimName = &value
	return b
}

// WithClaimedTimestamp sets the ClaimedTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClaimedTimestamp field is set to the value of the last call.
func (b *ClusterPoolReferenceApplyConfiguration) WithClaimedTimestamp(value v1.Time) *ClusterPoolReferenceApplyConfiguration {
	b.ClaimedTimestamp = &value
	return b
}

// WithCustomizationRef sets the CustomizationRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CustomizationRef field is set to the value of the last call.
func (b *ClusterPoolReferenceApplyConfiguration) WithCustomizationRef(value corev1.LocalObjectReference) *ClusterPoolReferenceApplyConfiguration {
	b.CustomizationRef = &value
	return b
}