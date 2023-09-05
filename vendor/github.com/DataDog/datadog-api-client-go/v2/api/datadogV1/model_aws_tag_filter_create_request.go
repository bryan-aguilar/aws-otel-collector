// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSTagFilterCreateRequest The objects used to set an AWS tag filter.
type AWSTagFilterCreateRequest struct {
	// Your AWS Account ID without dashes.
	AccountId *string `json:"account_id,omitempty"`
	// The namespace associated with the tag filter entry.
	Namespace *AWSNamespace `json:"namespace,omitempty"`
	// The tag filter string.
	TagFilterStr *string `json:"tag_filter_str,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSTagFilterCreateRequest instantiates a new AWSTagFilterCreateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSTagFilterCreateRequest() *AWSTagFilterCreateRequest {
	this := AWSTagFilterCreateRequest{}
	return &this
}

// NewAWSTagFilterCreateRequestWithDefaults instantiates a new AWSTagFilterCreateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSTagFilterCreateRequestWithDefaults() *AWSTagFilterCreateRequest {
	this := AWSTagFilterCreateRequest{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AWSTagFilterCreateRequest) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSTagFilterCreateRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AWSTagFilterCreateRequest) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AWSTagFilterCreateRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AWSTagFilterCreateRequest) GetNamespace() AWSNamespace {
	if o == nil || o.Namespace == nil {
		var ret AWSNamespace
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSTagFilterCreateRequest) GetNamespaceOk() (*AWSNamespace, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AWSTagFilterCreateRequest) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given AWSNamespace and assigns it to the Namespace field.
func (o *AWSTagFilterCreateRequest) SetNamespace(v AWSNamespace) {
	o.Namespace = &v
}

// GetTagFilterStr returns the TagFilterStr field value if set, zero value otherwise.
func (o *AWSTagFilterCreateRequest) GetTagFilterStr() string {
	if o == nil || o.TagFilterStr == nil {
		var ret string
		return ret
	}
	return *o.TagFilterStr
}

// GetTagFilterStrOk returns a tuple with the TagFilterStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSTagFilterCreateRequest) GetTagFilterStrOk() (*string, bool) {
	if o == nil || o.TagFilterStr == nil {
		return nil, false
	}
	return o.TagFilterStr, true
}

// HasTagFilterStr returns a boolean if a field has been set.
func (o *AWSTagFilterCreateRequest) HasTagFilterStr() bool {
	return o != nil && o.TagFilterStr != nil
}

// SetTagFilterStr gets a reference to the given string and assigns it to the TagFilterStr field.
func (o *AWSTagFilterCreateRequest) SetTagFilterStr(v string) {
	o.TagFilterStr = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSTagFilterCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.TagFilterStr != nil {
		toSerialize["tag_filter_str"] = o.TagFilterStr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSTagFilterCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId    *string       `json:"account_id,omitempty"`
		Namespace    *AWSNamespace `json:"namespace,omitempty"`
		TagFilterStr *string       `json:"tag_filter_str,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "namespace", "tag_filter_str"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = all.AccountId
	if all.Namespace != nil && !all.Namespace.IsValid() {
		hasInvalidField = true
	} else {
		o.Namespace = all.Namespace
	}
	o.TagFilterStr = all.TagFilterStr

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
