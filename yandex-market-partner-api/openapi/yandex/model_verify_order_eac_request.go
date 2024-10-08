/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the VerifyOrderEacRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifyOrderEacRequest{}

// VerifyOrderEacRequest struct for VerifyOrderEacRequest
type VerifyOrderEacRequest struct {
	// Код для подтверждения ЭАПП.
	Code *string `json:"code,omitempty"`
}

// NewVerifyOrderEacRequest instantiates a new VerifyOrderEacRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyOrderEacRequest() *VerifyOrderEacRequest {
	this := VerifyOrderEacRequest{}
	return &this
}

// NewVerifyOrderEacRequestWithDefaults instantiates a new VerifyOrderEacRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyOrderEacRequestWithDefaults() *VerifyOrderEacRequest {
	this := VerifyOrderEacRequest{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *VerifyOrderEacRequest) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyOrderEacRequest) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *VerifyOrderEacRequest) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *VerifyOrderEacRequest) SetCode(v string) {
	o.Code = &v
}

func (o VerifyOrderEacRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifyOrderEacRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableVerifyOrderEacRequest struct {
	value *VerifyOrderEacRequest
	isSet bool
}

func (v NullableVerifyOrderEacRequest) Get() *VerifyOrderEacRequest {
	return v.value
}

func (v *NullableVerifyOrderEacRequest) Set(val *VerifyOrderEacRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyOrderEacRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyOrderEacRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyOrderEacRequest(val *VerifyOrderEacRequest) *NullableVerifyOrderEacRequest {
	return &NullableVerifyOrderEacRequest{value: val, isSet: true}
}

func (v NullableVerifyOrderEacRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyOrderEacRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


