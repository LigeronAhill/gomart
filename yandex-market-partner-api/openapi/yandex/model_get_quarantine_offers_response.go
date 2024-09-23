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

// checks if the GetQuarantineOffersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetQuarantineOffersResponse{}

// GetQuarantineOffersResponse Ответ на запрос списка товаров в карантине.
type GetQuarantineOffersResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *GetQuarantineOffersResultDTO `json:"result,omitempty"`
}

// NewGetQuarantineOffersResponse instantiates a new GetQuarantineOffersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQuarantineOffersResponse() *GetQuarantineOffersResponse {
	this := GetQuarantineOffersResponse{}
	return &this
}

// NewGetQuarantineOffersResponseWithDefaults instantiates a new GetQuarantineOffersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQuarantineOffersResponseWithDefaults() *GetQuarantineOffersResponse {
	this := GetQuarantineOffersResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetQuarantineOffersResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuarantineOffersResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetQuarantineOffersResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetQuarantineOffersResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetQuarantineOffersResponse) GetResult() GetQuarantineOffersResultDTO {
	if o == nil || IsNil(o.Result) {
		var ret GetQuarantineOffersResultDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuarantineOffersResponse) GetResultOk() (*GetQuarantineOffersResultDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetQuarantineOffersResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GetQuarantineOffersResultDTO and assigns it to the Result field.
func (o *GetQuarantineOffersResponse) SetResult(v GetQuarantineOffersResultDTO) {
	o.Result = &v
}

func (o GetQuarantineOffersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetQuarantineOffersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetQuarantineOffersResponse struct {
	value *GetQuarantineOffersResponse
	isSet bool
}

func (v NullableGetQuarantineOffersResponse) Get() *GetQuarantineOffersResponse {
	return v.value
}

func (v *NullableGetQuarantineOffersResponse) Set(val *GetQuarantineOffersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuarantineOffersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuarantineOffersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuarantineOffersResponse(val *GetQuarantineOffersResponse) *NullableGetQuarantineOffersResponse {
	return &NullableGetQuarantineOffersResponse{value: val, isSet: true}
}

func (v NullableGetQuarantineOffersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuarantineOffersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


