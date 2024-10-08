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

// checks if the GetOrderBuyerInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrderBuyerInfoResponse{}

// GetOrderBuyerInfoResponse struct for GetOrderBuyerInfoResponse
type GetOrderBuyerInfoResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *OrderBuyerInfoDTO `json:"result,omitempty"`
}

// NewGetOrderBuyerInfoResponse instantiates a new GetOrderBuyerInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderBuyerInfoResponse() *GetOrderBuyerInfoResponse {
	this := GetOrderBuyerInfoResponse{}
	return &this
}

// NewGetOrderBuyerInfoResponseWithDefaults instantiates a new GetOrderBuyerInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderBuyerInfoResponseWithDefaults() *GetOrderBuyerInfoResponse {
	this := GetOrderBuyerInfoResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrderBuyerInfoResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderBuyerInfoResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrderBuyerInfoResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetOrderBuyerInfoResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetOrderBuyerInfoResponse) GetResult() OrderBuyerInfoDTO {
	if o == nil || IsNil(o.Result) {
		var ret OrderBuyerInfoDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderBuyerInfoResponse) GetResultOk() (*OrderBuyerInfoDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetOrderBuyerInfoResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OrderBuyerInfoDTO and assigns it to the Result field.
func (o *GetOrderBuyerInfoResponse) SetResult(v OrderBuyerInfoDTO) {
	o.Result = &v
}

func (o GetOrderBuyerInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderBuyerInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetOrderBuyerInfoResponse struct {
	value *GetOrderBuyerInfoResponse
	isSet bool
}

func (v NullableGetOrderBuyerInfoResponse) Get() *GetOrderBuyerInfoResponse {
	return v.value
}

func (v *NullableGetOrderBuyerInfoResponse) Set(val *GetOrderBuyerInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderBuyerInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderBuyerInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderBuyerInfoResponse(val *GetOrderBuyerInfoResponse) *NullableGetOrderBuyerInfoResponse {
	return &NullableGetOrderBuyerInfoResponse{value: val, isSet: true}
}

func (v NullableGetOrderBuyerInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderBuyerInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


