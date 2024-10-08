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

// checks if the AddOffersToArchiveResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOffersToArchiveResponse{}

// AddOffersToArchiveResponse Результат архивации товаров.
type AddOffersToArchiveResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *AddOffersToArchiveDTO `json:"result,omitempty"`
}

// NewAddOffersToArchiveResponse instantiates a new AddOffersToArchiveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOffersToArchiveResponse() *AddOffersToArchiveResponse {
	this := AddOffersToArchiveResponse{}
	return &this
}

// NewAddOffersToArchiveResponseWithDefaults instantiates a new AddOffersToArchiveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOffersToArchiveResponseWithDefaults() *AddOffersToArchiveResponse {
	this := AddOffersToArchiveResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AddOffersToArchiveResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOffersToArchiveResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AddOffersToArchiveResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *AddOffersToArchiveResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AddOffersToArchiveResponse) GetResult() AddOffersToArchiveDTO {
	if o == nil || IsNil(o.Result) {
		var ret AddOffersToArchiveDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOffersToArchiveResponse) GetResultOk() (*AddOffersToArchiveDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AddOffersToArchiveResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given AddOffersToArchiveDTO and assigns it to the Result field.
func (o *AddOffersToArchiveResponse) SetResult(v AddOffersToArchiveDTO) {
	o.Result = &v
}

func (o AddOffersToArchiveResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOffersToArchiveResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableAddOffersToArchiveResponse struct {
	value *AddOffersToArchiveResponse
	isSet bool
}

func (v NullableAddOffersToArchiveResponse) Get() *AddOffersToArchiveResponse {
	return v.value
}

func (v *NullableAddOffersToArchiveResponse) Set(val *AddOffersToArchiveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOffersToArchiveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOffersToArchiveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOffersToArchiveResponse(val *AddOffersToArchiveResponse) *NullableAddOffersToArchiveResponse {
	return &NullableAddOffersToArchiveResponse{value: val, isSet: true}
}

func (v NullableAddOffersToArchiveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOffersToArchiveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


