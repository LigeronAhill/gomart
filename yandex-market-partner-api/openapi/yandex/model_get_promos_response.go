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

// checks if the GetPromosResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPromosResponse{}

// GetPromosResponse struct for GetPromosResponse
type GetPromosResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *GetPromosResultDTO `json:"result,omitempty"`
}

// NewGetPromosResponse instantiates a new GetPromosResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPromosResponse() *GetPromosResponse {
	this := GetPromosResponse{}
	return &this
}

// NewGetPromosResponseWithDefaults instantiates a new GetPromosResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPromosResponseWithDefaults() *GetPromosResponse {
	this := GetPromosResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetPromosResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPromosResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetPromosResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetPromosResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetPromosResponse) GetResult() GetPromosResultDTO {
	if o == nil || IsNil(o.Result) {
		var ret GetPromosResultDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPromosResponse) GetResultOk() (*GetPromosResultDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetPromosResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GetPromosResultDTO and assigns it to the Result field.
func (o *GetPromosResponse) SetResult(v GetPromosResultDTO) {
	o.Result = &v
}

func (o GetPromosResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPromosResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetPromosResponse struct {
	value *GetPromosResponse
	isSet bool
}

func (v NullableGetPromosResponse) Get() *GetPromosResponse {
	return v.value
}

func (v *NullableGetPromosResponse) Set(val *GetPromosResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPromosResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPromosResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPromosResponse(val *GetPromosResponse) *NullableGetPromosResponse {
	return &NullableGetPromosResponse{value: val, isSet: true}
}

func (v NullableGetPromosResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPromosResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


