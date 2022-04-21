/*
Connector Service Fleet Manager Admin APIs

Connector Service Fleet Manager Admin is a Rest API to manage connector clusters.

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

// MetaV1Condition struct for MetaV1Condition
type MetaV1Condition struct {
	Type               string `json:"type,omitempty"`
	Reason             string `json:"reason,omitempty"`
	Message            string `json:"message,omitempty"`
	Status             string `json:"status,omitempty"`
	LastTransitionTime string `json:"last_transition_time,omitempty"`
}