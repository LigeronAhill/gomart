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

// checks if the ApiNotFoundErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiNotFoundErrorResponse{}

// ApiNotFoundErrorResponse Запрашиваемый ресурс не найден.
type ApiNotFoundErrorResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	// Список ошибок.
	Errors []ApiErrorDTO `json:"errors,omitempty"`
}

// NewApiNotFoundErrorResponse instantiates a new ApiNotFoundErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNotFoundErrorResponse() *ApiNotFoundErrorResponse {
	this := ApiNotFoundErrorResponse{}
	return &this
}

// NewApiNotFoundErrorResponseWithDefaults instantiates a new ApiNotFoundErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNotFoundErrorResponseWithDefaults() *ApiNotFoundErrorResponse {
	this := ApiNotFoundErrorResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiNotFoundErrorResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNotFoundErrorResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiNotFoundErrorResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *ApiNotFoundErrorResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiNotFoundErrorResponse) GetErrors() []ApiErrorDTO {
	if o == nil {
		var ret []ApiErrorDTO
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiNotFoundErrorResponse) GetErrorsOk() ([]ApiErrorDTO, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ApiNotFoundErrorResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ApiErrorDTO and assigns it to the Errors field.
func (o *ApiNotFoundErrorResponse) SetErrors(v []ApiErrorDTO) {
	o.Errors = v
}

func (o ApiNotFoundErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiNotFoundErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableApiNotFoundErrorResponse struct {
	value *ApiNotFoundErrorResponse
	isSet bool
}

func (v NullableApiNotFoundErrorResponse) Get() *ApiNotFoundErrorResponse {
	return v.value
}

func (v *NullableApiNotFoundErrorResponse) Set(val *ApiNotFoundErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNotFoundErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNotFoundErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNotFoundErrorResponse(val *ApiNotFoundErrorResponse) *NullableApiNotFoundErrorResponse {
	return &NullableApiNotFoundErrorResponse{value: val, isSet: true}
}

func (v NullableApiNotFoundErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNotFoundErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


