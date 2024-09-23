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

// checks if the ParameterValueOptionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterValueOptionDTO{}

// ParameterValueOptionDTO Значение характеристики.
type ParameterValueOptionDTO struct {
	// Идентификатор значения.
	Id int64 `json:"id"`
	// Значение.
	Value string `json:"value"`
	// Описание значения.
	Description *string `json:"description,omitempty"`
}

type _ParameterValueOptionDTO ParameterValueOptionDTO

// NewParameterValueOptionDTO instantiates a new ParameterValueOptionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterValueOptionDTO(id int64, value string) *ParameterValueOptionDTO {
	this := ParameterValueOptionDTO{}
	this.Id = id
	this.Value = value
	return &this
}

// NewParameterValueOptionDTOWithDefaults instantiates a new ParameterValueOptionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterValueOptionDTOWithDefaults() *ParameterValueOptionDTO {
	this := ParameterValueOptionDTO{}
	return &this
}

// GetId returns the Id field value
func (o *ParameterValueOptionDTO) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ParameterValueOptionDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ParameterValueOptionDTO) SetId(v int64) {
	o.Id = v
}

// GetValue returns the Value field value
func (o *ParameterValueOptionDTO) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ParameterValueOptionDTO) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ParameterValueOptionDTO) SetValue(v string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ParameterValueOptionDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterValueOptionDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ParameterValueOptionDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ParameterValueOptionDTO) SetDescription(v string) {
	o.Description = &v
}

func (o ParameterValueOptionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterValueOptionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["value"] = o.Value
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *ParameterValueOptionDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"value",
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

	varParameterValueOptionDTO := _ParameterValueOptionDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParameterValueOptionDTO)

	if err != nil {
		return err
	}

	*o = ParameterValueOptionDTO(varParameterValueOptionDTO)

	return err
}

type NullableParameterValueOptionDTO struct {
	value *ParameterValueOptionDTO
	isSet bool
}

func (v NullableParameterValueOptionDTO) Get() *ParameterValueOptionDTO {
	return v.value
}

func (v *NullableParameterValueOptionDTO) Set(val *ParameterValueOptionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterValueOptionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterValueOptionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterValueOptionDTO(val *ParameterValueOptionDTO) *NullableParameterValueOptionDTO {
	return &NullableParameterValueOptionDTO{value: val, isSet: true}
}

func (v NullableParameterValueOptionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterValueOptionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


