/*
 * XRELY
 *
 * API Documentation for XRELY platform
 *
 * API version: 1.0.0
 * Contact: contact@xrely.com
 */

package xrely

type SearchRequest struct {
	// Account Key
	Ak string `json:"ak,omitempty"`
	// Query Term
	Q string `json:"q,omitempty"`
	// Number of results
	Size string `json:"size,omitempty"`
	AggField []AggrigationField `json:"aggField,omitempty"`
	FilterField []FilterField `json:"filterField,omitempty"`
}
