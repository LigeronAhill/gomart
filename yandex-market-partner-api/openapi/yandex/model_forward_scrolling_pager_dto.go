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

// checks if the ForwardScrollingPagerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForwardScrollingPagerDTO{}

// ForwardScrollingPagerDTO Ссылка на следующую страницу. 
type ForwardScrollingPagerDTO struct {
	// Идентификатор следующей страницы результатов.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

// NewForwardScrollingPagerDTO instantiates a new ForwardScrollingPagerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForwardScrollingPagerDTO() *ForwardScrollingPagerDTO {
	this := ForwardScrollingPagerDTO{}
	return &this
}

// NewForwardScrollingPagerDTOWithDefaults instantiates a new ForwardScrollingPagerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForwardScrollingPagerDTOWithDefaults() *ForwardScrollingPagerDTO {
	this := ForwardScrollingPagerDTO{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ForwardScrollingPagerDTO) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardScrollingPagerDTO) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ForwardScrollingPagerDTO) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ForwardScrollingPagerDTO) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ForwardScrollingPagerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForwardScrollingPagerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}
	return toSerialize, nil
}

type NullableForwardScrollingPagerDTO struct {
	value *ForwardScrollingPagerDTO
	isSet bool
}

func (v NullableForwardScrollingPagerDTO) Get() *ForwardScrollingPagerDTO {
	return v.value
}

func (v *NullableForwardScrollingPagerDTO) Set(val *ForwardScrollingPagerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableForwardScrollingPagerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableForwardScrollingPagerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForwardScrollingPagerDTO(val *ForwardScrollingPagerDTO) *NullableForwardScrollingPagerDTO {
	return &NullableForwardScrollingPagerDTO{value: val, isSet: true}
}

func (v NullableForwardScrollingPagerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForwardScrollingPagerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


