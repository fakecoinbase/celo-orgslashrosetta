/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SubNetworkIdentifier - In blockchains with sharded state, the `sub_network_identifier` is required to wholly specify a shard block. 
type SubNetworkIdentifier struct {

	SubNetwork string `json:"sub_network"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`
}