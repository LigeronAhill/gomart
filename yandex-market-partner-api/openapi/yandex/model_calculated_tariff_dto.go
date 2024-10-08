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

// checks if the CalculatedTariffDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalculatedTariffDTO{}

// CalculatedTariffDTO Информация об услугах Маркета.
type CalculatedTariffDTO struct {
	Type CalculatedTariffType `json:"type"`
	// Стоимость услуги в рублях.
	Amount *float32 `json:"amount,omitempty"`
	// Параметры расчета тарифа.
	Parameters []TariffParameterDTO `json:"parameters"`
}

type _CalculatedTariffDTO CalculatedTariffDTO

// NewCalculatedTariffDTO instantiates a new CalculatedTariffDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalculatedTariffDTO(type_ CalculatedTariffType, parameters []TariffParameterDTO) *CalculatedTariffDTO {
	this := CalculatedTariffDTO{}
	this.Type = type_
	this.Parameters = parameters
	return &this
}

// NewCalculatedTariffDTOWithDefaults instantiates a new CalculatedTariffDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalculatedTariffDTOWithDefaults() *CalculatedTariffDTO {
	this := CalculatedTariffDTO{}
	return &this
}

// GetType returns the Type field value
func (o *CalculatedTariffDTO) GetType() CalculatedTariffType {
	if o == nil {
		var ret CalculatedTariffType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CalculatedTariffDTO) GetTypeOk() (*CalculatedTariffType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CalculatedTariffDTO) SetType(v CalculatedTariffType) {
	o.Type = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CalculatedTariffDTO) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalculatedTariffDTO) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CalculatedTariffDTO) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *CalculatedTariffDTO) SetAmount(v float32) {
	o.Amount = &v
}

// GetParameters returns the Parameters field value
func (o *CalculatedTariffDTO) GetParameters() []TariffParameterDTO {
	if o == nil {
		var ret []TariffParameterDTO
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *CalculatedTariffDTO) GetParametersOk() ([]TariffParameterDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *CalculatedTariffDTO) SetParameters(v []TariffParameterDTO) {
	o.Parameters = v
}

func (o CalculatedTariffDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalculatedTariffDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["parameters"] = o.Parameters
	return toSerialize, nil
}

func (o *CalculatedTariffDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"parameters",
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

	varCalculatedTariffDTO := _CalculatedTariffDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCalculatedTariffDTO)

	if err != nil {
		return err
	}

	*o = CalculatedTariffDTO(varCalculatedTariffDTO)

	return err
}

type NullableCalculatedTariffDTO struct {
	value *CalculatedTariffDTO
	isSet bool
}

func (v NullableCalculatedTariffDTO) Get() *CalculatedTariffDTO {
	return v.value
}

func (v *NullableCalculatedTariffDTO) Set(val *CalculatedTariffDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCalculatedTariffDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCalculatedTariffDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalculatedTariffDTO(val *CalculatedTariffDTO) *NullableCalculatedTariffDTO {
	return &NullableCalculatedTariffDTO{value: val, isSet: true}
}

func (v NullableCalculatedTariffDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalculatedTariffDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


