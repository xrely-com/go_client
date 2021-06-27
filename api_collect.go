/*
 * XRELY
 *
 * API Documentation for XRELY platform
 *
 * API version: 1.0.0
 * Contact: contact@xrely.com
 */

package xrely

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type CollectApiService service

/* 
CollectApiService collects information related to user activity

 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CollectActivityGetOpts - Optional Parameters:
     * @param "K" (optional.String) -  key
     * @param "Uid" (optional.String) -  Unique Id
     * @param "Domain" (optional.String) -  Domain Name
     * @param "EntId" (optional.String) -  Entity Id
     * @param "EntName" (optional.String) -  Entity Name
     * @param "Keyword" (optional.String) -  Keyword
     * @param "DPos" (optional.String) -  Display Position
     * @param "Ref" (optional.String) -  Referrer
     * @param "Tag" (optional.String) -  Listing Tags

@return ModelApiResponse
*/

type CollectActivityGetOpts struct { 
	K optional.String
	Uid optional.String
	Domain optional.String
	EntId optional.String
	EntName optional.String
	Keyword optional.String
	DPos optional.String
	Ref optional.String
	Tag optional.String
}

func (a *CollectApiService) CollectActivityGet(ctx context.Context, localVarOptionals *CollectActivityGetOpts) (ModelApiResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ModelApiResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/collect/activity"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.K.IsSet() {
		localVarQueryParams.Add("k", parameterToString(localVarOptionals.K.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Uid.IsSet() {
		localVarQueryParams.Add("uid", parameterToString(localVarOptionals.Uid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Domain.IsSet() {
		localVarQueryParams.Add("domain", parameterToString(localVarOptionals.Domain.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EntId.IsSet() {
		localVarQueryParams.Add("entId", parameterToString(localVarOptionals.EntId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EntName.IsSet() {
		localVarQueryParams.Add("entName", parameterToString(localVarOptionals.EntName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Keyword.IsSet() {
		localVarQueryParams.Add("keyword", parameterToString(localVarOptionals.Keyword.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DPos.IsSet() {
		localVarQueryParams.Add("dPos", parameterToString(localVarOptionals.DPos.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ref.IsSet() {
		localVarQueryParams.Add("ref", parameterToString(localVarOptionals.Ref.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Tag.IsSet() {
		localVarQueryParams.Add("tag", parameterToString(localVarOptionals.Tag.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
CollectApiService collects information related to device

 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *CollectInfoGetOpts - Optional Parameters:
     * @param "Uid" (optional.String) -  Unique Id
     * @param "K" (optional.String) -  key
     * @param "Aid" (optional.String) -  Account Id
     * @param "AV" (optional.String) -  App Version
     * @param "Netwk" (optional.String) -  Network
     * @param "Os" (optional.String) -  Operating System
     * @param "OsV" (optional.String) -  Oerating System Version
     * @param "Url" (optional.String) -  Current url / Deep link
     * @param "Ref" (optional.String) -  Referrer
     * @param "Uagnt" (optional.String) -  User Agent / App Signature
     * @param "Udid" (optional.String) -  Unique Device ID
     * @param "Imei" (optional.String) -  Device IMEI Number
     * @param "Domain" (optional.String) -  Domain Name
     * @param "NetCrr" (optional.String) -  Network Carrier
     * @param "Platform" (optional.String) -  Mobile/Web
     * @param "DM" (optional.String) -  Device Model
     * @param "DMgf" (optional.String) -  Device Manufacturer
     * @param "Email" (optional.String) -  Email Id
     * @param "Mobile" (optional.String) -  Mobile Number

@return ModelApiResponse
*/

type CollectInfoGetOpts struct { 
	Uid optional.String
	K optional.String
	Aid optional.String
	AV optional.String
	Netwk optional.String
	Os optional.String
	OsV optional.String
	Url optional.String
	Ref optional.String
	Uagnt optional.String
	Udid optional.String
	Imei optional.String
	Domain optional.String
	NetCrr optional.String
	Platform optional.String
	DM optional.String
	DMgf optional.String
	Email optional.String
	Mobile optional.String
}

func (a *CollectApiService) CollectInfoGet(ctx context.Context, localVarOptionals *CollectInfoGetOpts) (ModelApiResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ModelApiResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/collect/info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Uid.IsSet() {
		localVarQueryParams.Add("uid", parameterToString(localVarOptionals.Uid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.K.IsSet() {
		localVarQueryParams.Add("k", parameterToString(localVarOptionals.K.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Aid.IsSet() {
		localVarQueryParams.Add("aid", parameterToString(localVarOptionals.Aid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AV.IsSet() {
		localVarQueryParams.Add("aV", parameterToString(localVarOptionals.AV.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Netwk.IsSet() {
		localVarQueryParams.Add("netwk", parameterToString(localVarOptionals.Netwk.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Os.IsSet() {
		localVarQueryParams.Add("os", parameterToString(localVarOptionals.Os.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OsV.IsSet() {
		localVarQueryParams.Add("osV", parameterToString(localVarOptionals.OsV.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Url.IsSet() {
		localVarQueryParams.Add("url", parameterToString(localVarOptionals.Url.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Ref.IsSet() {
		localVarQueryParams.Add("ref", parameterToString(localVarOptionals.Ref.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Uagnt.IsSet() {
		localVarQueryParams.Add("uagnt", parameterToString(localVarOptionals.Uagnt.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Udid.IsSet() {
		localVarQueryParams.Add("udid", parameterToString(localVarOptionals.Udid.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Imei.IsSet() {
		localVarQueryParams.Add("imei", parameterToString(localVarOptionals.Imei.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Domain.IsSet() {
		localVarQueryParams.Add("domain", parameterToString(localVarOptionals.Domain.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NetCrr.IsSet() {
		localVarQueryParams.Add("netCrr", parameterToString(localVarOptionals.NetCrr.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Platform.IsSet() {
		localVarQueryParams.Add("platform", parameterToString(localVarOptionals.Platform.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DM.IsSet() {
		localVarQueryParams.Add("dM", parameterToString(localVarOptionals.DM.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DMgf.IsSet() {
		localVarQueryParams.Add("dMgf", parameterToString(localVarOptionals.DMgf.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Email.IsSet() {
		localVarQueryParams.Add("email", parameterToString(localVarOptionals.Email.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Mobile.IsSet() {
		localVarQueryParams.Add("mobile", parameterToString(localVarOptionals.Mobile.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
