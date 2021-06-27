# \DataApi

All URIs are relative to *https://www.xrely.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataStoreDelete**](DataApi.md#DataStoreDelete) | **Delete** /data/store | Delete data from your account bucket
[**DataStorePost**](DataApi.md#DataStorePost) | **Post** /data/store | Insert data to your account bucket


# **DataStoreDelete**
> DataApiResponse DataStoreDelete(ctx, optional)
Delete data from your account bucket



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataStoreDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataStoreDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DataStoreRequest**](DataStoreRequest.md)|  | 

### Return type

[**DataApiResponse**](DataApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataStorePost**
> DataApiResponse DataStorePost(ctx, optional)
Insert data to your account bucket



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DataStorePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataStorePostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DataStoreRequest**](DataStoreRequest.md)|  | 

### Return type

[**DataApiResponse**](DataApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

