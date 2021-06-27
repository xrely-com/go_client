/*
 * XRELY
 *
 * API Documentation for XRELY platform
 *
 * API version: 1.0.0
 * Contact: contact@xrely.com
 */

package xrely

type DocStoreItem struct {
	// keyword or phrase
	Keyword string `json:"keyword,omitempty"`
	// Related url with the keyword
	Url string `json:"url,omitempty"`
	Data *MapObject `json:"data,omitempty"`
}
