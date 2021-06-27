/*
 * XRELY
 *
 * API Documentation for XRELY platform
 *
 * API version: 1.0.0
 * Contact: contact@xrely.com
 */

package xrely

type DataMessage struct {
	Committed bool `json:"committed,omitempty"`
	Message string `json:"message,omitempty"`
}
