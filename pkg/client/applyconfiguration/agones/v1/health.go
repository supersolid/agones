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

// HealthApplyConfiguration represents an declarative configuration of the Health type for use
// with apply.
type HealthApplyConfiguration struct {
	Disabled            *bool  `json:"disabled,omitempty"`
	PeriodSeconds       *int32 `json:"periodSeconds,omitempty"`
	FailureThreshold    *int32 `json:"failureThreshold,omitempty"`
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty"`
}

// HealthApplyConfiguration constructs an declarative configuration of the Health type for use with
// apply.
func Health() *HealthApplyConfiguration {
	return &HealthApplyConfiguration{}
}

// WithDisabled sets the Disabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Disabled field is set to the value of the last call.
func (b *HealthApplyConfiguration) WithDisabled(value bool) *HealthApplyConfiguration {
	b.Disabled = &value
	return b
}

// WithPeriodSeconds sets the PeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PeriodSeconds field is set to the value of the last call.
func (b *HealthApplyConfiguration) WithPeriodSeconds(value int32) *HealthApplyConfiguration {
	b.PeriodSeconds = &value
	return b
}

// WithFailureThreshold sets the FailureThreshold field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailureThreshold field is set to the value of the last call.
func (b *HealthApplyConfiguration) WithFailureThreshold(value int32) *HealthApplyConfiguration {
	b.FailureThreshold = &value
	return b
}

// WithInitialDelaySeconds sets the InitialDelaySeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InitialDelaySeconds field is set to the value of the last call.
func (b *HealthApplyConfiguration) WithInitialDelaySeconds(value int32) *HealthApplyConfiguration {
	b.InitialDelaySeconds = &value
	return b
}
