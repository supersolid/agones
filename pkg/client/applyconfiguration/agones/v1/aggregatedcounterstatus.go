// Copyright 2023 Google LLC All Rights Reserved.
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

// This code was autogenerated. Do not edit directly.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// AggregatedCounterStatusApplyConfiguration represents an declarative configuration of the AggregatedCounterStatus type for use
// with apply.
type AggregatedCounterStatusApplyConfiguration struct {
	AllocatedCount    *int64 `json:"allocatedCount,omitempty"`
	AllocatedCapacity *int64 `json:"allocatedCapacity,omitempty"`
	Count             *int64 `json:"count,omitempty"`
	Capacity          *int64 `json:"capacity,omitempty"`
}

// AggregatedCounterStatusApplyConfiguration constructs an declarative configuration of the AggregatedCounterStatus type for use with
// apply.
func AggregatedCounterStatus() *AggregatedCounterStatusApplyConfiguration {
	return &AggregatedCounterStatusApplyConfiguration{}
}

// WithAllocatedCount sets the AllocatedCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllocatedCount field is set to the value of the last call.
func (b *AggregatedCounterStatusApplyConfiguration) WithAllocatedCount(value int64) *AggregatedCounterStatusApplyConfiguration {
	b.AllocatedCount = &value
	return b
}

// WithAllocatedCapacity sets the AllocatedCapacity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllocatedCapacity field is set to the value of the last call.
func (b *AggregatedCounterStatusApplyConfiguration) WithAllocatedCapacity(value int64) *AggregatedCounterStatusApplyConfiguration {
	b.AllocatedCapacity = &value
	return b
}

// WithCount sets the Count field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Count field is set to the value of the last call.
func (b *AggregatedCounterStatusApplyConfiguration) WithCount(value int64) *AggregatedCounterStatusApplyConfiguration {
	b.Count = &value
	return b
}

// WithCapacity sets the Capacity field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Capacity field is set to the value of the last call.
func (b *AggregatedCounterStatusApplyConfiguration) WithCapacity(value int64) *AggregatedCounterStatusApplyConfiguration {
	b.Capacity = &value
	return b
}
