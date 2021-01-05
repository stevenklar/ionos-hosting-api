# {{classname}}

All URIs are relative to *https://api.hosting.ionos.com/dns*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZone**](ZonesApi.md#GetZone) | **Get** /v1/zones/{zoneId} | 
[**GetZones**](ZonesApi.md#GetZones) | **Get** /v1/zones | 
[**PatchZone**](ZonesApi.md#PatchZone) | **Patch** /v1/zones/{zoneId} | 
[**UpdateZone**](ZonesApi.md#UpdateZone) | **Put** /v1/zones/{zoneId} | 

# **GetZone**
> CustomerZone GetZone(ctx, zoneId, optional)


Returns a customer zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**| The id of the customer zone. | 
 **optional** | ***ZonesApiGetZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ZonesApiGetZoneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suffix** | **optional.String**| The FQDN used to filter all the record names that end with it. | 
 **recordName** | **optional.String**| The record names that should be included (same as name field of Record) | 
 **recordType** | **optional.String**| A comma-separated list of record types that should be included | 

### Return type

[**CustomerZone**](customer-zone.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetZones**
> []Zone GetZones(ctx, )


Returns list of customer zones.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Zone**](zone.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchZone**
> PatchZone(ctx, body, zoneId)


Replaces all records of the same name and type with the ones provided.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]Record**](record.md)| records | 
  **zoneId** | **string**| The id of the customer zone. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateZone**
> UpdateZone(ctx, body, zoneId)


Replaces all records in the zone with the ones provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]Record**](record.md)| records | 
  **zoneId** | **string**| The id of the customer zone. | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

