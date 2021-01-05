# {{classname}}

All URIs are relative to *https://api.hosting.ionos.com/dns*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateDynDns**](DynamicDNSApi.md#ActivateDynDns) | **Post** /v1/dyndns | 
[**DisableDynDns**](DynamicDNSApi.md#DisableDynDns) | **Delete** /v1/dyndns | 
[**V1DyndnsBulkIdDelete**](DynamicDNSApi.md#V1DyndnsBulkIdDelete) | **Delete** /v1/dyndns/{bulkId} | 
[**V1DyndnsBulkIdPut**](DynamicDNSApi.md#V1DyndnsBulkIdPut) | **Put** /v1/dyndns/{bulkId} | 

# **ActivateDynDns**
> DynamicDns ActivateDynDns(ctx, body)


Activate Dynamic Dns for a bundle of (sub)domains. The url from response will be used to update the ips of the (sub)domains.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DynDnsRequest**](DynDnsRequest.md)| Dynamic Dns configuration. | 

### Return type

[**DynamicDns**](dynamic-dns.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableDynDns**
> DisableDynDns(ctx, )


Disable Dynamic Dns

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DyndnsBulkIdDelete**
> V1DyndnsBulkIdDelete(ctx, bulkId)


Disable Dynamic Dns for bulk id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bulkId** | **string**| Dynamic Dns configuration identifier. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1DyndnsBulkIdPut**
> V1DyndnsBulkIdPut(ctx, body, bulkId)


Update Dynamic Dns for bulk id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]DynDnsRequest**](dyn-dns-request.md)| Dynamic Dns configuration. | 
  **bulkId** | **string**| Dynamic Dns configuration identifier. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

