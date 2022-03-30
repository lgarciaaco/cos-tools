/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public
// ConnectorNamespaceStatus struct for ConnectorNamespaceStatus
type ConnectorNamespaceStatus struct {
	State ConnectorNamespaceState `json:"state"`
	Version string `json:"version,omitempty"`
	ConnectorsDeployed int32 `json:"connectors_deployed"`
	Error string `json:"error,omitempty"`
}