/*
 * SSL API
 *
 * # Working with the API Every endpoint uses the `X-API-Key` header for authorization, to obtain the key please see the [Official Documentation](/docs/getstarted).  This SSL API version only supports DV certificates. SSL Unlimited (Flat) is not supported in the current beta.  # Support Support questions may be posted in English: <a href='/docs/getstarted#support'>API Beta Support</a>.  Please note that in the Beta phase we offer support in the business Hours Mo-Fri 9:00-17:00 EET.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pkg
import (
	"time"
)

type Certificate struct {
	Id string `json:"id,omitempty"`
	CertificateType string `json:"certificateType,omitempty"`
	AuthenticationMethod string `json:"authenticationMethod,omitempty"`
	Status string `json:"status,omitempty"`
	CommonName string `json:"commonName,omitempty"`
	AlternativeNames []string `json:"alternativeNames,omitempty"`
	ValidFrom time.Time `json:"validFrom,omitempty"`
	ValidUntil time.Time `json:"validUntil,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
	CaOrderId string `json:"caOrderId,omitempty"`
	ReplacedCertificates []ReplacedCertificate `json:"replacedCertificates,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	CaCertificates []CaCertificate `json:"caCertificates,omitempty"`
	DeploymentStatus string `json:"deploymentStatus,omitempty"`
	Links []Link `json:"links,omitempty"`
}