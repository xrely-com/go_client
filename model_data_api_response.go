/*
 * XRELY
 *
 * API Documentation for XRELY platform
 *
 * API version: 1.0.0
 * Contact: contact@xrely.com
 */

package xrely

type DataApiResponse struct {
	Code int32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data *DataMessage `json:"data,omitempty"`
}
