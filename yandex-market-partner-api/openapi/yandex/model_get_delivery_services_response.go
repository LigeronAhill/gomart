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

// checks if the GetDeliveryServicesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeliveryServicesResponse{}

// GetDeliveryServicesResponse Ответ на запрос списка служб доставки.
type GetDeliveryServicesResponse struct {
	Result *DeliveryServicesDTO `json:"result,omitempty"`
}

// NewGetDeliveryServicesResponse instantiates a new GetDeliveryServicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeliveryServicesResponse() *GetDeliveryServicesResponse {
	this := GetDeliveryServicesResponse{}
	return &this
}

// NewGetDeliveryServicesResponseWithDefaults instantiates a new GetDeliveryServicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeliveryServicesResponseWithDefaults() *GetDeliveryServicesResponse {
	this := GetDeliveryServicesResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetDeliveryServicesResponse) GetResult() DeliveryServicesDTO {
	if o == nil || IsNil(o.Result) {
		var ret DeliveryServicesDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeliveryServicesResponse) GetResultOk() (*DeliveryServicesDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetDeliveryServicesResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DeliveryServicesDTO and assigns it to the Result field.
func (o *GetDeliveryServicesResponse) SetResult(v DeliveryServicesDTO) {
	o.Result = &v
}

func (o GetDeliveryServicesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeliveryServicesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetDeliveryServicesResponse struct {
	value *GetDeliveryServicesResponse
	isSet bool
}

func (v NullableGetDeliveryServicesResponse) Get() *GetDeliveryServicesResponse {
	return v.value
}

func (v *NullableGetDeliveryServicesResponse) Set(val *GetDeliveryServicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeliveryServicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeliveryServicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeliveryServicesResponse(val *GetDeliveryServicesResponse) *NullableGetDeliveryServicesResponse {
	return &NullableGetDeliveryServicesResponse{value: val, isSet: true}
}

func (v NullableGetDeliveryServicesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeliveryServicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


