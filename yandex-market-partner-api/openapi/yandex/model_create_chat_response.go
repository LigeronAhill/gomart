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

// checks if the CreateChatResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChatResponse{}

// CreateChatResponse Результат создания чата.
type CreateChatResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *CreateChatResultDTO `json:"result,omitempty"`
}

// NewCreateChatResponse instantiates a new CreateChatResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChatResponse() *CreateChatResponse {
	this := CreateChatResponse{}
	return &this
}

// NewCreateChatResponseWithDefaults instantiates a new CreateChatResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChatResponseWithDefaults() *CreateChatResponse {
	this := CreateChatResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateChatResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChatResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateChatResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *CreateChatResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateChatResponse) GetResult() CreateChatResultDTO {
	if o == nil || IsNil(o.Result) {
		var ret CreateChatResultDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateChatResponse) GetResultOk() (*CreateChatResultDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateChatResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given CreateChatResultDTO and assigns it to the Result field.
func (o *CreateChatResponse) SetResult(v CreateChatResultDTO) {
	o.Result = &v
}

func (o CreateChatResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChatResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableCreateChatResponse struct {
	value *CreateChatResponse
	isSet bool
}

func (v NullableCreateChatResponse) Get() *CreateChatResponse {
	return v.value
}

func (v *NullableCreateChatResponse) Set(val *CreateChatResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChatResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChatResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChatResponse(val *CreateChatResponse) *NullableCreateChatResponse {
	return &NullableCreateChatResponse{value: val, isSet: true}
}

func (v NullableCreateChatResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChatResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


