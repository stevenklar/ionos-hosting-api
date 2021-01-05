# {{classname}}

All URIs are relative to *https://api.hosting.ionos.com/ssl*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificate**](CertificatesApi.md#CreateCertificate) | **Post** /v1/certificates | Create a new certificate
[**GetCertificateDetails**](CertificatesApi.md#GetCertificateDetails) | **Get** /v1/certificates/{id} | Get certificate details
[**GetCertificates**](CertificatesApi.md#GetCertificates) | **Get** /v1/certificates | Get certificates
[**UnassignCertificate**](CertificatesApi.md#UnassignCertificate) | **Delete** /v1/certificates/{id} | Unassign

# **CreateCertificate**
> Certificate CreateCertificate(ctx, body)
Create a new certificate

Request a new certificate. This requires that there is an unused SSL certificate of the respective type already purchased within the customer account. Creating a certificate is an asynchronous operation. Currently, the supported method for getting the completed state is by polling operations (see GET /v1/certificates/{id}).  The DV authentication method is automatically selected based on the provided common name. `FILE` and `DNS` methods are preferred. When `EMAIL` authentication is selected,  Digicert sends an email to the email address of the hostmaster of your domain (i.e. hostmaster@example.com). You must confirm the request for the certificate via the link contained in the email.  The CSR must be supplied when the certificate installation is managed by you.  The CSR must use RSA with key size 2048 bits. On Linux, you can generate the CSR using OpenSSL:      openssl req -new -newkey rsa:2048 -nodes -keyout example.key -out example.csr -subj \"/C=DE/CN=example.com\"  The generated CSR will be formatted like in the following example:      -----BEGIN CERTIFICATE REQUEST-----     MIICaDCCAVACAQAwIzELMAkGA1UEBhMCREUxFDASBgNVBAMMC2V4YW1wbGUuY29t     MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1UWfeVKGB0ekYy3Yv2R6     tNJrwEQSIPVUkd0M+3LqEOYY3O2Hq3QQ7cbpsQEwPNXasTKmE/IbG89fDmiESfsp     v9goxkAs0CAD6G8/e7jNe8KMrYBPBV+tFx8DV5UubmBAqd7ahLK0+dDhFbmQAnbs     CkshE9e8yuaWKijLjr2BqiIWKR+pjCnWiBr4ApfT38vYt56ImhfNfQT/HKgxsYDO     x1AAorauzQ6Vko5VEKtaeLFemh7AKboCza6JuDDgSqj0c8TsdStu/ftOE05fErzf     HnJaQFfkQY+C5bo72tmtexrBXdZsGDQgvOpsE8497fmYoJ/A9D+C9udGyN9U64bM     QwIDAQABoAAwDQYJKoZIhvcNAQELBQADggEBAEHBXVLQBb5FqdEf1b7VDFeq7ZFN     zl61jbzefL5JcMA9wlPfJiT4nhxai3ueEHcH4mhyEvONzFBbB12TeS7GtEO9T21h     ZGXaok4UMy/XZrJXxw733BsJyEzacuzzBNbG9BP1xPfVrTWKfmsgnurCutlNXzuj     a0bjETgKV6QHJdB8hNWMNrpW4N6GyODBHZm0hPsQkEf4wDXnKMeO7OhQo4zOeZTn     ZR9WVxzcWfgQ7TQ5Od+mVJHc39WhMQlFzaDsHGmipJZYN/iIeP4LBymTnnM9N84K     z1N44/dwlIP2HxW9qjUMy+Ks6eymEAgHMeF3aym2qnPBWcd92Py7OFXZhxQ=     -----END CERTIFICATE REQUEST-----  For more details, you can consult the following guide from Digicert:  [How to create a CSR](https://www.digicert.com/csr-creation.htm).  When a CSR is not provided, the certificate is automatically installed on your IONOS managed website. An error will be raised if the common name is not associated with an website managed by IONOS.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EnrollRequest**](EnrollRequest.md)|  | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

[X-API-Key](../README.md#X-API-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertificateDetails**
> Certificate GetCertificateDetails(ctx, id)
Get certificate details

Retrieve the details for the certificate having the specified id. Use this operation to get the status of the certificate or the certificate itself after the certificate was issued.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCertificates**
> []Certificate GetCertificates(ctx, optional)
Get certificates

Retrieve the details for all the certificates in the current customer account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertificatesApiGetCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificatesApiGetCertificatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**|  | [default to 1]
 **size** | **optional.Int32**|  | [default to 50]

### Return type

[**[]Certificate**](Certificate.md)

### Authorization

[X-API-Key](../README.md#X-API-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnassignCertificate**
> UnassignCertificate(ctx, id)
Unassign

Unassign the certificate having the specified id. After the operation completes the slot becomes available for creating a new certificate.  Note: The same restrictions as in Control Panel also apply for the API operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[X-API-Key](../README.md#X-API-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

