# \CollectApi

All URIs are relative to *https://api.xrely.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectActivityGet**](CollectApi.md#CollectActivityGet) | **Get** /collect/activity | collects information related to user activity
[**CollectInfoGet**](CollectApi.md#CollectInfoGet) | **Get** /collect/info | collects information related to device


# **CollectActivityGet**
> ModelApiResponse CollectActivityGet(ctx, optional)
collects information related to user activity



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectActivityGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectActivityGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **k** | **optional.String**| key | 
 **uid** | **optional.String**| Unique Id | 
 **domain** | **optional.String**| Domain Name | 
 **entId** | **optional.String**| Entity Id | 
 **entName** | **optional.String**| Entity Name | 
 **keyword** | **optional.String**| Keyword | 
 **dPos** | **optional.String**| Display Position | 
 **ref** | **optional.String**| Referrer | 
 **tag** | **optional.String**| Listing Tags | 

### Return type

[**ModelApiResponse**](ApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectInfoGet**
> ModelApiResponse CollectInfoGet(ctx, optional)
collects information related to device



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectInfoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectInfoGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **optional.String**| Unique Id | 
 **k** | **optional.String**| key | 
 **aid** | **optional.String**| Account Id | 
 **aV** | **optional.String**| App Version | 
 **netwk** | **optional.String**| Network | 
 **os** | **optional.String**| Operating System | 
 **osV** | **optional.String**| Oerating System Version | 
 **url** | **optional.String**| Current url / Deep link | 
 **ref** | **optional.String**| Referrer | 
 **uagnt** | **optional.String**| User Agent / App Signature | 
 **udid** | **optional.String**| Unique Device ID | 
 **imei** | **optional.String**| Device IMEI Number | 
 **domain** | **optional.String**| Domain Name | 
 **netCrr** | **optional.String**| Network Carrier | 
 **platform** | **optional.String**| Mobile/Web | 
 **dM** | **optional.String**| Device Model | 
 **dMgf** | **optional.String**| Device Manufacturer | 
 **email** | **optional.String**| Email Id | 
 **mobile** | **optional.String**| Mobile Number | 

### Return type

[**ModelApiResponse**](ApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

