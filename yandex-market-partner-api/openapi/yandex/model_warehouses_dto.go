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

// checks if the WarehousesDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarehousesDTO{}

// WarehousesDTO Информация о складах и группах складов.
type WarehousesDTO struct {
	// Список складов, не входящих в группы.
	Warehouses []WarehouseDTO `json:"warehouses"`
	// Список групп складов.
	WarehouseGroups []WarehouseGroupDTO `json:"warehouseGroups"`
}

type _WarehousesDTO WarehousesDTO

// NewWarehousesDTO instantiates a new WarehousesDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehousesDTO(warehouses []WarehouseDTO, warehouseGroups []WarehouseGroupDTO) *WarehousesDTO {
	this := WarehousesDTO{}
	this.Warehouses = warehouses
	this.WarehouseGroups = warehouseGroups
	return &this
}

// NewWarehousesDTOWithDefaults instantiates a new WarehousesDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehousesDTOWithDefaults() *WarehousesDTO {
	this := WarehousesDTO{}
	return &this
}

// GetWarehouses returns the Warehouses field value
func (o *WarehousesDTO) GetWarehouses() []WarehouseDTO {
	if o == nil {
		var ret []WarehouseDTO
		return ret
	}

	return o.Warehouses
}

// GetWarehousesOk returns a tuple with the Warehouses field value
// and a boolean to check if the value has been set.
func (o *WarehousesDTO) GetWarehousesOk() ([]WarehouseDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Warehouses, true
}

// SetWarehouses sets field value
func (o *WarehousesDTO) SetWarehouses(v []WarehouseDTO) {
	o.Warehouses = v
}

// GetWarehouseGroups returns the WarehouseGroups field value
func (o *WarehousesDTO) GetWarehouseGroups() []WarehouseGroupDTO {
	if o == nil {
		var ret []WarehouseGroupDTO
		return ret
	}

	return o.WarehouseGroups
}

// GetWarehouseGroupsOk returns a tuple with the WarehouseGroups field value
// and a boolean to check if the value has been set.
func (o *WarehousesDTO) GetWarehouseGroupsOk() ([]WarehouseGroupDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarehouseGroups, true
}

// SetWarehouseGroups sets field value
func (o *WarehousesDTO) SetWarehouseGroups(v []WarehouseGroupDTO) {
	o.WarehouseGroups = v
}

func (o WarehousesDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarehousesDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["warehouses"] = o.Warehouses
	toSerialize["warehouseGroups"] = o.WarehouseGroups
	return toSerialize, nil
}

func (o *WarehousesDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"warehouses",
		"warehouseGroups",
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

	varWarehousesDTO := _WarehousesDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWarehousesDTO)

	if err != nil {
		return err
	}

	*o = WarehousesDTO(varWarehousesDTO)

	return err
}

type NullableWarehousesDTO struct {
	value *WarehousesDTO
	isSet bool
}

func (v NullableWarehousesDTO) Get() *WarehousesDTO {
	return v.value
}

func (v *NullableWarehousesDTO) Set(val *WarehousesDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehousesDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehousesDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehousesDTO(val *WarehousesDTO) *NullableWarehousesDTO {
	return &NullableWarehousesDTO{value: val, isSet: true}
}

func (v NullableWarehousesDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehousesDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


