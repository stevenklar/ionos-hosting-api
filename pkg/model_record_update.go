/*
 * DNS API
 *
 * ## Working with the API Every endpoint uses the `X-API-Key` header for authorization, to obtain the key please see the [Official Documentation](/docs/getstarted).  Please note that any zone or record updates might conflict with active services. In such cases, the DNS records belonging to the conflicting services will be deactivated.  ## Support Support questions may be posted in English: <a href='/docs/getstarted#support'>API Beta Support</a>.  Please note that in the Beta phase we offer support in the business Hours Mo-Fri 9:00-17:00 EET. <h2> <details>   <summary>Release notes</summary>    Vesion 1.0.0   Exposed CRUD operations for customer zone. </details> </h2> 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pkg

type RecordUpdate struct {
	// When is true, the record is not visible for lookup.
	Disabled bool `json:"disabled,omitempty"`
	Content string `json:"content,omitempty"`
	// Time to live for the record, recommended 3600.
	Ttl int32 `json:"ttl,omitempty"`
	Prio int32 `json:"prio,omitempty"`
}
