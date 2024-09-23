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

// checks if the GetChatsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetChatsResponse{}

// GetChatsResponse struct for GetChatsResponse
type GetChatsResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *GetChatsInfoDTO `json:"result,omitempty"`
}

// NewGetChatsResponse instantiates a new GetChatsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetChatsResponse() *GetChatsResponse {
	this := GetChatsResponse{}
	return &this
}

// NewGetChatsResponseWithDefaults instantiates a new GetChatsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetChatsResponseWithDefaults() *GetChatsResponse {
	this := GetChatsResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetChatsResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatsResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetChatsResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetChatsResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetChatsResponse) GetResult() GetChatsInfoDTO {
	if o == nil || IsNil(o.Result) {
		var ret GetChatsInfoDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChatsResponse) GetResultOk() (*GetChatsInfoDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetChatsResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GetChatsInfoDTO and assigns it to the Result field.
func (o *GetChatsResponse) SetResult(v GetChatsInfoDTO) {
	o.Result = &v
}

func (o GetChatsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetChatsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetChatsResponse struct {
	value *GetChatsResponse
	isSet bool
}

func (v NullableGetChatsResponse) Get() *GetChatsResponse {
	return v.value
}

func (v *NullableGetChatsResponse) Set(val *GetChatsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChatsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChatsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChatsResponse(val *GetChatsResponse) *NullableGetChatsResponse {
	return &NullableGetChatsResponse{value: val, isSet: true}
}

func (v NullableGetChatsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChatsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


