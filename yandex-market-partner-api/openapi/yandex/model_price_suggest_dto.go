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

// checks if the PriceSuggestDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceSuggestDTO{}

// PriceSuggestDTO Тип цены.
type PriceSuggestDTO struct {
	Type *PriceSuggestType `json:"type,omitempty"`
	// Цена в рублях.
	Price *float32 `json:"price,omitempty"`
}

// NewPriceSuggestDTO instantiates a new PriceSuggestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceSuggestDTO() *PriceSuggestDTO {
	this := PriceSuggestDTO{}
	return &this
}

// NewPriceSuggestDTOWithDefaults instantiates a new PriceSuggestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceSuggestDTOWithDefaults() *PriceSuggestDTO {
	this := PriceSuggestDTO{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PriceSuggestDTO) GetType() PriceSuggestType {
	if o == nil || IsNil(o.Type) {
		var ret PriceSuggestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceSuggestDTO) GetTypeOk() (*PriceSuggestType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PriceSuggestDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PriceSuggestType and assigns it to the Type field.
func (o *PriceSuggestDTO) SetType(v PriceSuggestType) {
	o.Type = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PriceSuggestDTO) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceSuggestDTO) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PriceSuggestDTO) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *PriceSuggestDTO) SetPrice(v float32) {
	o.Price = &v
}

func (o PriceSuggestDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceSuggestDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullablePriceSuggestDTO struct {
	value *PriceSuggestDTO
	isSet bool
}

func (v NullablePriceSuggestDTO) Get() *PriceSuggestDTO {
	return v.value
}

func (v *NullablePriceSuggestDTO) Set(val *PriceSuggestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceSuggestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceSuggestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceSuggestDTO(val *PriceSuggestDTO) *NullablePriceSuggestDTO {
	return &NullablePriceSuggestDTO{value: val, isSet: true}
}

func (v NullablePriceSuggestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceSuggestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


