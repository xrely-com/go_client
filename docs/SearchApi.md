# \SearchApi

All URIs are relative to *https://api.xrely.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchGet**](SearchApi.md#SearchGet) | **Get** /search | Get autocomplete phrase or keywords for your query
[**SearchPost**](SearchApi.md#SearchPost) | **Post** /search | Provides relevant result for your query


# **SearchGet**
> ModelApiResponse SearchGet(ctx, q, ak)
Get autocomplete phrase or keywords for your query



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| Search Term or Keyword | 
  **ak** | **string**| Account Key | 

### Return type

[**ModelApiResponse**](ApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchPost**
> ModelApiResponse SearchPost(ctx, body)
Provides relevant result for your query



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchRequest**](SearchRequest.md)| Search query term or keyword | 

### Return type

[**ModelApiResponse**](ApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

