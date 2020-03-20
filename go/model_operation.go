/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Operation struct {

	OperationIdentifier OperationIdentifier `json:"operation_identifier"`

	// Restrict referenced `related_operations` to identifier indexes `<` the current `operation_identifier.index`. This ensures there exists a clear DAG-structure of relations.  Since `operations` are one-sided, one could imagine relating operations in a single transfer or linking `operations` in a call tree. 
	RelatedOperations []OperationIdentifier `json:"related_operations,omitempty"`

	// The network-specific type of the operation. Ensure that any type that can be returned here is also specified in the `NetowrkStatus`. This can be very useful to downstream consumers that parse all block data. 
	Type string `json:"type"`

	// The network-specific status of the operation. Status is not defined on the transaction object because blockchains with smart contracts may have transactions that partially apply.  Blockchains with atomic transactions (all operations succeed or all operations fail) will have the same `status` for each operation. 
	Status string `json:"status"`

	Account AccountIdentifier `json:"account,omitempty"`

	Amount Amount `json:"amount,omitempty"`

	Metadata map[string]interface{} `json:"metadata,omitempty"`
}