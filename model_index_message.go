/*
 * XRELY
 *
 * API Documentation for XRELY platform
 *
 * API version: 1.0.0
 * Contact: contact@xrely.com
 */

package xrely

type IndexMessage struct {
	Indexed bool `json:"indexed,omitempty"`
	Message string `json:"message,omitempty"`
}
