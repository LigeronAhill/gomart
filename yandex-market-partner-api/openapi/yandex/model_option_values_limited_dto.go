/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OptionValuesLimitedDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionValuesLimitedDTO{}

// OptionValuesLimitedDTO Значение ограничивающей характеристики и список допустимых значений ограничиваемой характеристики.
type OptionValuesLimitedDTO struct {
	// Идентификатор значения ограничивающей характеристики.
	LimitingOptionValueId int64 `json:"limitingOptionValueId"`
	// Идентификаторы допустимых значений ограничиваемой характеристики. 
	OptionValueIds []int64 `json:"optionValueIds"`
}

type _OptionValuesLimitedDTO OptionValuesLimitedDTO

// NewOptionValuesLimitedDTO instantiates a new OptionValuesLimitedDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionValuesLimitedDTO(limitingOptionValueId int64, optionValueIds []int64) *OptionValuesLimitedDTO {
	this := OptionValuesLimitedDTO{}
	this.LimitingOptionValueId = limitingOptionValueId
	this.OptionValueIds = optionValueIds
	return &this
}

// NewOptionValuesLimitedDTOWithDefaults instantiates a new OptionValuesLimitedDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionValuesLimitedDTOWithDefaults() *OptionValuesLimitedDTO {
	this := OptionValuesLimitedDTO{}
	return &this
}

// GetLimitingOptionValueId returns the LimitingOptionValueId field value
func (o *OptionValuesLimitedDTO) GetLimitingOptionValueId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LimitingOptionValueId
}

// GetLimitingOptionValueIdOk returns a tuple with the LimitingOptionValueId field value
// and a boolean to check if the value has been set.
func (o *OptionValuesLimitedDTO) GetLimitingOptionValueIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitingOptionValueId, true
}

// SetLimitingOptionValueId sets field value
func (o *OptionValuesLimitedDTO) SetLimitingOptionValueId(v int64) {
	o.LimitingOptionValueId = v
}

// GetOptionValueIds returns the OptionValueIds field value
func (o *OptionValuesLimitedDTO) GetOptionValueIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.OptionValueIds
}

// GetOptionValueIdsOk returns a tuple with the OptionValueIds field value
// and a boolean to check if the value has been set.
func (o *OptionValuesLimitedDTO) GetOptionValueIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptionValueIds, true
}

// SetOptionValueIds sets field value
func (o *OptionValuesLimitedDTO) SetOptionValueIds(v []int64) {
	o.OptionValueIds = v
}

func (o OptionValuesLimitedDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionValuesLimitedDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limitingOptionValueId"] = o.LimitingOptionValueId
	toSerialize["optionValueIds"] = o.OptionValueIds
	return toSerialize, nil
}

func (o *OptionValuesLimitedDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"limitingOptionValueId",
		"optionValueIds",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOptionValuesLimitedDTO := _OptionValuesLimitedDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOptionValuesLimitedDTO)

	if err != nil {
		return err
	}

	*o = OptionValuesLimitedDTO(varOptionValuesLimitedDTO)

	return err
}

type NullableOptionValuesLimitedDTO struct {
	value *OptionValuesLimitedDTO
	isSet bool
}

func (v NullableOptionValuesLimitedDTO) Get() *OptionValuesLimitedDTO {
	return v.value
}

func (v *NullableOptionValuesLimitedDTO) Set(val *OptionValuesLimitedDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionValuesLimitedDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionValuesLimitedDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionValuesLimitedDTO(val *OptionValuesLimitedDTO) *NullableOptionValuesLimitedDTO {
	return &NullableOptionValuesLimitedDTO{value: val, isSet: true}
}

func (v NullableOptionValuesLimitedDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionValuesLimitedDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


