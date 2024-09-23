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

// checks if the OutletLicensesResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutletLicensesResponseDTO{}

// OutletLicensesResponseDTO Ответ на запрос информации о лицензиях для точек продаж.
type OutletLicensesResponseDTO struct {
	// Список лицензий.
	Licenses []FullOutletLicenseDTO `json:"licenses"`
}

type _OutletLicensesResponseDTO OutletLicensesResponseDTO

// NewOutletLicensesResponseDTO instantiates a new OutletLicensesResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutletLicensesResponseDTO(licenses []FullOutletLicenseDTO) *OutletLicensesResponseDTO {
	this := OutletLicensesResponseDTO{}
	this.Licenses = licenses
	return &this
}

// NewOutletLicensesResponseDTOWithDefaults instantiates a new OutletLicensesResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutletLicensesResponseDTOWithDefaults() *OutletLicensesResponseDTO {
	this := OutletLicensesResponseDTO{}
	return &this
}

// GetLicenses returns the Licenses field value
func (o *OutletLicensesResponseDTO) GetLicenses() []FullOutletLicenseDTO {
	if o == nil {
		var ret []FullOutletLicenseDTO
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *OutletLicensesResponseDTO) GetLicensesOk() ([]FullOutletLicenseDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *OutletLicensesResponseDTO) SetLicenses(v []FullOutletLicenseDTO) {
	o.Licenses = v
}

func (o OutletLicensesResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutletLicensesResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["licenses"] = o.Licenses
	return toSerialize, nil
}

func (o *OutletLicensesResponseDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"licenses",
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

	varOutletLicensesResponseDTO := _OutletLicensesResponseDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutletLicensesResponseDTO)

	if err != nil {
		return err
	}

	*o = OutletLicensesResponseDTO(varOutletLicensesResponseDTO)

	return err
}

type NullableOutletLicensesResponseDTO struct {
	value *OutletLicensesResponseDTO
	isSet bool
}

func (v NullableOutletLicensesResponseDTO) Get() *OutletLicensesResponseDTO {
	return v.value
}

func (v *NullableOutletLicensesResponseDTO) Set(val *OutletLicensesResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOutletLicensesResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOutletLicensesResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutletLicensesResponseDTO(val *OutletLicensesResponseDTO) *NullableOutletLicensesResponseDTO {
	return &NullableOutletLicensesResponseDTO{value: val, isSet: true}
}

func (v NullableOutletLicensesResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutletLicensesResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


