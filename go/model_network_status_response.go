/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type NetworkStatusResponse struct {

	NetworkStatus NetworkStatus `json:"network_status"`

	// If a node supports multiple sub-networks, their statuses should be returned in this array. 
	SubNetworkStatus []SubNetworkStatus `json:"sub_network_status,omitempty"`

	Version Version `json:"version"`

	Options Options `json:"options"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`
}