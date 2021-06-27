/*
 * XRELY
 *
 * API Documentation for XRELY platform
 *
 * API version: 1.0.0
 * Contact: contact@xrely.com
 */

package xrely

type ModelApiResponse struct {
	Code int32 `json:"code,omitempty"`
	Data string `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}
