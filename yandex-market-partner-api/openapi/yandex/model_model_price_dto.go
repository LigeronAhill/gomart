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

// checks if the ModelPriceDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPriceDTO{}

// ModelPriceDTO Информация о ценах на модель товара.
type ModelPriceDTO struct {
	// Средняя цена предложения для модели в регионе.
	Avg *float32 `json:"avg,omitempty"`
	// Максимальная цена предложения для модели в регионе.
	Max *float32 `json:"max,omitempty"`
	// Минимальная цена предложения для модели в регионе.
	Min *float32 `json:"min,omitempty"`
}

// NewModelPriceDTO instantiates a new ModelPriceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPriceDTO() *ModelPriceDTO {
	this := ModelPriceDTO{}
	return &this
}

// NewModelPriceDTOWithDefaults instantiates a new ModelPriceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPriceDTOWithDefaults() *ModelPriceDTO {
	this := ModelPriceDTO{}
	return &this
}

// GetAvg returns the Avg field value if set, zero value otherwise.
func (o *ModelPriceDTO) GetAvg() float32 {
	if o == nil || IsNil(o.Avg) {
		var ret float32
		return ret
	}
	return *o.Avg
}

// GetAvgOk returns a tuple with the Avg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPriceDTO) GetAvgOk() (*float32, bool) {
	if o == nil || IsNil(o.Avg) {
		return nil, false
	}
	return o.Avg, true
}

// HasAvg returns a boolean if a field has been set.
func (o *ModelPriceDTO) HasAvg() bool {
	if o != nil && !IsNil(o.Avg) {
		return true
	}

	return false
}

// SetAvg gets a reference to the given float32 and assigns it to the Avg field.
func (o *ModelPriceDTO) SetAvg(v float32) {
	o.Avg = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ModelPriceDTO) GetMax() float32 {
	if o == nil || IsNil(o.Max) {
		var ret float32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPriceDTO) GetMaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ModelPriceDTO) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float32 and assigns it to the Max field.
func (o *ModelPriceDTO) SetMax(v float32) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ModelPriceDTO) GetMin() float32 {
	if o == nil || IsNil(o.Min) {
		var ret float32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPriceDTO) GetMinOk() (*float32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ModelPriceDTO) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float32 and assigns it to the Min field.
func (o *ModelPriceDTO) SetMin(v float32) {
	o.Min = &v
}

func (o ModelPriceDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPriceDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Avg) {
		toSerialize["avg"] = o.Avg
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	return toSerialize, nil
}

type NullableModelPriceDTO struct {
	value *ModelPriceDTO
	isSet bool
}

func (v NullableModelPriceDTO) Get() *ModelPriceDTO {
	return v.value
}

func (v *NullableModelPriceDTO) Set(val *ModelPriceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPriceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPriceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPriceDTO(val *ModelPriceDTO) *NullableModelPriceDTO {
	return &NullableModelPriceDTO{value: val, isSet: true}
}

func (v NullableModelPriceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPriceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


