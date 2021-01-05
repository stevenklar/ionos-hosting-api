# Go API client for IONOS hosting api

### Working with the API Every endpoint uses the `X-API-Key` header for authorization, to obtain the key please see the Official Documentation.

## Overview
- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Documentation for API Endpoints

All URIs are relative to *https://api.hosting.ionos.com/ssl*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DynamicDNSApi* | [**ActivateDynDns**](docs/DynamicDNSApi.md#activatedyndns) | **Post** /v1/dyndns | 
*DynamicDNSApi* | [**DisableDynDns**](docs/DynamicDNSApi.md#disabledyndns) | **Delete** /v1/dyndns | 
*DynamicDNSApi* | [**V1DyndnsBulkIdDelete**](docs/DynamicDNSApi.md#v1dyndnsbulkiddelete) | **Delete** /v1/dyndns/{bulkId} | 
*DynamicDNSApi* | [**V1DyndnsBulkIdPut**](docs/DynamicDNSApi.md#v1dyndnsbulkidput) | **Put** /v1/dyndns/{bulkId} | 
*RecordsApi* | [**DeleteRecord**](docs/RecordsApi.md#deleterecord) | **Delete** /v1/zones/{zoneId}/records/{recordId} | 
*RecordsApi* | [**GetRecord**](docs/RecordsApi.md#getrecord) | **Get** /v1/zones/{zoneId}/records/{recordId} | 
*RecordsApi* | [**UpdateRecord**](docs/RecordsApi.md#updaterecord) | **Put** /v1/zones/{zoneId}/records/{recordId} | 
*ZonesApi* | [**GetZone**](docs/ZonesApi.md#getzone) | **Get** /v1/zones/{zoneId} | 
*ZonesApi* | [**GetZones**](docs/ZonesApi.md#getzones) | **Get** /v1/zones | 
*ZonesApi* | [**PatchZone**](docs/ZonesApi.md#patchzone) | **Patch** /v1/zones/{zoneId} | 
*ZonesApi* | [**UpdateZone**](docs/ZonesApi.md#updatezone) | **Put** /v1/zones/{zoneId} | 
*CertificatesApi* | [**CreateCertificate**](docs/CertificatesApi.md#createcertificate) | **Post** /v1/certificates | Create a new certificate
*CertificatesApi* | [**GetCertificateDetails**](docs/CertificatesApi.md#getcertificatedetails) | **Get** /v1/certificates/{id} | Get certificate details
*CertificatesApi* | [**GetCertificates**](docs/CertificatesApi.md#getcertificates) | **Get** /v1/certificates | Get certificates
*CertificatesApi* | [**UnassignCertificate**](docs/CertificatesApi.md#unassigncertificate) | **Delete** /v1/certificates/{id} | Unassign

## Documentation For Models

 - [CustomerZone](docs/CustomerZone.md)
 - [DynDnsRequest](docs/DynDnsRequest.md)
 - [DynamicDns](docs/DynamicDns.md)
 - [ErrorsInner](docs/ErrorsInner.md)
 - [ModelError](docs/ModelError.md)
 - [Record](docs/Record.md)
 - [RecordTypes](docs/RecordTypes.md)
 - [RecordUpdate](docs/RecordUpdate.md)
 - [Zone](docs/Zone.md)
 - [ZoneTypes](docs/ZoneTypes.md)
 - [CaCertificate](docs/CaCertificate.md)
 - [Certificate](docs/Certificate.md)
 - [EnrollRequest](docs/EnrollRequest.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [Link](docs/Link.md)
 - [ReplacedCertificate](docs/ReplacedCertificate.md)

## Documentation For Authorization

## X-API-Key
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```
