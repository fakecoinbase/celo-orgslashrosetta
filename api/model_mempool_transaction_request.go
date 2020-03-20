/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type MempoolTransactionRequest struct {

	NetworkIdentifier NetworkIdentifier `json:"network_identifier"`

	TransactionIdentifier TransactionIdentifier `json:"transaction_identifier"`
}