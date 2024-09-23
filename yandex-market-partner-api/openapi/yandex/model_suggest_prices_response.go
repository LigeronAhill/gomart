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

// checks if the SuggestPricesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestPricesResponse{}

// SuggestPricesResponse Ответ на запрос списка цен для продвижения.
type SuggestPricesResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *SuggestPricesResultDTO `json:"result,omitempty"`
}

// NewSuggestPricesResponse instantiates a new SuggestPricesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestPricesResponse() *SuggestPricesResponse {
	this := SuggestPricesResponse{}
	return &this
}

// NewSuggestPricesResponseWithDefaults instantiates a new SuggestPricesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestPricesResponseWithDefaults() *SuggestPricesResponse {
	this := SuggestPricesResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SuggestPricesResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestPricesResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SuggestPricesResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *SuggestPricesResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *SuggestPricesResponse) GetResult() SuggestPricesResultDTO {
	if o == nil || IsNil(o.Result) {
		var ret SuggestPricesResultDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestPricesResponse) GetResultOk() (*SuggestPricesResultDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *SuggestPricesResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given SuggestPricesResultDTO and assigns it to the Result field.
func (o *SuggestPricesResponse) SetResult(v SuggestPricesResultDTO) {
	o.Result = &v
}

func (o SuggestPricesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuggestPricesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableSuggestPricesResponse struct {
	value *SuggestPricesResponse
	isSet bool
}

func (v NullableSuggestPricesResponse) Get() *SuggestPricesResponse {
	return v.value
}

func (v *NullableSuggestPricesResponse) Set(val *SuggestPricesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestPricesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestPricesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestPricesResponse(val *SuggestPricesResponse) *NullableSuggestPricesResponse {
	return &NullableSuggestPricesResponse{value: val, isSet: true}
}

func (v NullableSuggestPricesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestPricesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


