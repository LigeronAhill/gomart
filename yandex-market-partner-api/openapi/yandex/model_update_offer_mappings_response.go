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

// checks if the UpdateOfferMappingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOfferMappingsResponse{}

// UpdateOfferMappingsResponse Описывает проблемы, возникшие при сохранении товара.
type UpdateOfferMappingsResponse struct {
	Status *ApiResponseStatusType `json:"status,omitempty"`
	// Ошибки и предупреждения, которые появились при обработке списка характеристик. Каждый элемент списка соответствует одному товару.  Если ошибок и предупреждений нет, поле не передается. 
	Results []UpdateOfferMappingResultDTO `json:"results,omitempty"`
}

// NewUpdateOfferMappingsResponse instantiates a new UpdateOfferMappingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOfferMappingsResponse() *UpdateOfferMappingsResponse {
	this := UpdateOfferMappingsResponse{}
	return &this
}

// NewUpdateOfferMappingsResponseWithDefaults instantiates a new UpdateOfferMappingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOfferMappingsResponseWithDefaults() *UpdateOfferMappingsResponse {
	this := UpdateOfferMappingsResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateOfferMappingsResponse) GetStatus() ApiResponseStatusType {
	if o == nil || IsNil(o.Status) {
		var ret ApiResponseStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOfferMappingsResponse) GetStatusOk() (*ApiResponseStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateOfferMappingsResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiResponseStatusType and assigns it to the Status field.
func (o *UpdateOfferMappingsResponse) SetStatus(v ApiResponseStatusType) {
	o.Status = &v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateOfferMappingsResponse) GetResults() []UpdateOfferMappingResultDTO {
	if o == nil {
		var ret []UpdateOfferMappingResultDTO
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateOfferMappingsResponse) GetResultsOk() ([]UpdateOfferMappingResultDTO, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *UpdateOfferMappingsResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []UpdateOfferMappingResultDTO and assigns it to the Results field.
func (o *UpdateOfferMappingsResponse) SetResults(v []UpdateOfferMappingResultDTO) {
	o.Results = v
}

func (o UpdateOfferMappingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOfferMappingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableUpdateOfferMappingsResponse struct {
	value *UpdateOfferMappingsResponse
	isSet bool
}

func (v NullableUpdateOfferMappingsResponse) Get() *UpdateOfferMappingsResponse {
	return v.value
}

func (v *NullableUpdateOfferMappingsResponse) Set(val *UpdateOfferMappingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOfferMappingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOfferMappingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOfferMappingsResponse(val *UpdateOfferMappingsResponse) *NullableUpdateOfferMappingsResponse {
	return &NullableUpdateOfferMappingsResponse{value: val, isSet: true}
}

func (v NullableUpdateOfferMappingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOfferMappingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


