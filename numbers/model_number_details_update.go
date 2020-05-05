/*
 * Numbers API
 *
 * The Numbers API enables you to manage your existing numbers and buy new virtual numbers for use with Nexmo's APIs. Further information is here: <https://developer.nexmo.com/numbers/overview>
 *
 * API version: 1.0.18
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package numbers
// NumberDetailsUpdate struct for NumberDetailsUpdate
type NumberDetailsUpdate struct {
	// The two character country code in ISO 3166-1 alpha-2 format
	Country string `json:"country"`
	// An available inbound virtual number.
	Msisdn string `json:"msisdn"`
	// The Application that will handle inbound traffic to this number.
	AppId string `json:"app_id,omitempty"`
	// An URL-encoded URI to the webhook endpoint that handles inbound messages. Your webhook endpoint must be active before you make this request. Nexmo makes a `GET` request to the endpoint and checks that it returns a `200 OK` response. Set this parameter's value to an empty string to remove the webhook.
	MoHttpUrl string `json:"moHttpUrl,omitempty"`
	// The associated system type for your SMPP client
	MoSmppSysType string `json:"moSmppSysType,omitempty"`
	// Specify whether inbound voice calls on your number are forwarded to a SIP or a telephone number.  This must be used with the `voiceCallbackValue` parameter. If set, `sip` or `tel` are prioritized over the Voice capability in your Application.  *Note: The `app` value is deprecated and will be removed in future.* 
	VoiceCallbackType string `json:"voiceCallbackType,omitempty"`
	// A SIP URI or telephone number. Must be used with the `voiceCallbackType` parameter.
	VoiceCallbackValue string `json:"voiceCallbackValue,omitempty"`
	// A webhook URI for Nexmo to send a request to when a call ends
	VoiceStatusCallback string `json:"voiceStatusCallback,omitempty"`
	// <strong>DEPRECATED</strong> - We recommend that you use `app_id` instead.  Specifies the Messages webhook type (always `app`) associated with this number and must be used with the `messagesCallbackValue` parameter. 
	MessagesCallbackType string `json:"messagesCallbackType,omitempty"`
	// <strong>DEPRECATED</strong> - We recommend that you use `app_id` instead.  Specifies the Application ID of your Messages application. It must be used with the `messagesCallbackType` parameter. 
	MessagesCallbackValue string `json:"messagesCallbackValue,omitempty"`
}
