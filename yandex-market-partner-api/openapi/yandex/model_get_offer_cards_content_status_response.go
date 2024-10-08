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

// checks if the GetOfferCardsContentStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOfferCardsContentStatusResponse{}

// GetOfferCardsContentStatusResponse Ответ со списком состояний товаров и пагинацией.
type GetOfferCardsContentStatusResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	Result *OfferCardsContentStatusDTO `json:"result,omitempty"`
}

// NewGetOfferCardsContentStatusResponse instantiates a new GetOfferCardsContentStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOfferCardsContentStatusResponse() *GetOfferCardsContentStatusResponse {
	this := GetOfferCardsContentStatusResponse{}
	return &this
}

// NewGetOfferCardsContentStatusResponseWithDefaults instantiates a new GetOfferCardsContentStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOfferCardsContentStatusResponseWithDefaults() *GetOfferCardsContentStatusResponse {
	this := GetOfferCardsContentStatusResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOfferCardsContentStatusResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOfferCardsContentStatusResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOfferCardsContentStatusResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *GetOfferCardsContentStatusResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetOfferCardsContentStatusResponse) GetResult() OfferCardsContentStatusDTO {
	if o == nil || IsNil(o.Result) {
		var ret OfferCardsContentStatusDTO
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOfferCardsContentStatusResponse) GetResultOk() (*OfferCardsContentStatusDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetOfferCardsContentStatusResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OfferCardsContentStatusDTO and assigns it to the Result field.
func (o *GetOfferCardsContentStatusResponse) SetResult(v OfferCardsContentStatusDTO) {
	o.Result = &v
}

func (o GetOfferCardsContentStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOfferCardsContentStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetOfferCardsContentStatusResponse struct {
	value *GetOfferCardsContentStatusResponse
	isSet bool
}

func (v NullableGetOfferCardsContentStatusResponse) Get() *GetOfferCardsContentStatusResponse {
	return v.value
}

func (v *NullableGetOfferCardsContentStatusResponse) Set(val *GetOfferCardsContentStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOfferCardsContentStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOfferCardsContentStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOfferCardsContentStatusResponse(val *GetOfferCardsContentStatusResponse) *NullableGetOfferCardsContentStatusResponse {
	return &NullableGetOfferCardsContentStatusResponse{value: val, isSet: true}
}

func (v NullableGetOfferCardsContentStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOfferCardsContentStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


