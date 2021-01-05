# {{classname}}

All URIs are relative to *https://api.hosting.ionos.com/dns*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecord**](RecordsApi.md#DeleteRecord) | **Delete** /v1/zones/{zoneId}/records/{recordId} | 
[**GetRecord**](RecordsApi.md#GetRecord) | **Get** /v1/zones/{zoneId}/records/{recordId} | 
[**UpdateRecord**](RecordsApi.md#UpdateRecord) | **Put** /v1/zones/{zoneId}/records/{recordId} | 

# **DeleteRecord**
> DeleteRecord(ctx, zoneId, recordId)


Delete a record from the customer zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**| The id of the customer zone. | 
  **recordId** | **string**| The id of the record. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRecord**
> Record GetRecord(ctx, zoneId, recordId)


Returns the record from the customer zone with the mentioned id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**| The id of the customer zone. | 
  **recordId** | **string**| The id of the record. | 

### Return type

[**Record**](record.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRecord**
> UpdateRecord(ctx, body, zoneId, recordId)


Update a record from the customer zone. Any DomainConnect templates that conflict with the update will be reverted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RecordUpdate**](RecordUpdate.md)|  | 
  **zoneId** | **string**| The id of the customer zone. | 
  **recordId** | **string**| The id of the record. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

