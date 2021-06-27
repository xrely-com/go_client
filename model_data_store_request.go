/*
 * XRELY
 *
 * API Documentation for XRELY platform
 *
 * API version: 1.0.0
 * Contact: contact@xrely.com
 */

package xrely

type DataStoreRequest struct {
	// private key of the account
	SecretKey string `json:"secretKey,omitempty"`
	Docs []DocStoreItem `json:"docs,omitempty"`
}
