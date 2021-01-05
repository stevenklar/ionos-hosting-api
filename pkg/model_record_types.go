/*
 * DNS API
 *
 * ## Working with the API Every endpoint uses the `X-API-Key` header for authorization, to obtain the key please see the [Official Documentation](/docs/getstarted).  Please note that any zone or record updates might conflict with active services. In such cases, the DNS records belonging to the conflicting services will be deactivated.  ## Support Support questions may be posted in English: <a href='/docs/getstarted#support'>API Beta Support</a>.  Please note that in the Beta phase we offer support in the business Hours Mo-Fri 9:00-17:00 EET. <h2> <details>   <summary>Release notes</summary>    Vesion 1.0.0   Exposed CRUD operations for customer zone. </details> </h2> 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pkg
// RecordTypes : Holds supported dns record types.
type RecordTypes string

// List of recordTypes
const (
	A_RecordTypes RecordTypes = "A"
	AAAA_RecordTypes RecordTypes = "AAAA"
	CNAME_RecordTypes RecordTypes = "CNAME"
	MX_RecordTypes RecordTypes = "MX"
	NS_RecordTypes RecordTypes = "NS"
	SOA_RecordTypes RecordTypes = "SOA"
	SRV_RecordTypes RecordTypes = "SRV"
	TXT_RecordTypes RecordTypes = "TXT"
	CAA_RecordTypes RecordTypes = "CAA"
)
