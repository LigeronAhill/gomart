/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ShipmentType Способ отгрузки заказов:  * `IMPORT` — вы самостоятельно привозите заказы в выбранный сортировочный центр или пункт приема заказов. * `WITHDRAW` — вы отгружаете заказы со своего склада курьерам Яндекс Маркета. 
type ShipmentType string

// List of ShipmentType
const (
	IMPORT ShipmentType = "IMPORT"
	WITHDRAW ShipmentType = "WITHDRAW"
)

// All allowed values of ShipmentType enum
var AllowedShipmentTypeEnumValues = []ShipmentType{
	"IMPORT",
	"WITHDRAW",
}

func (v *ShipmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShipmentType(value)
	for _, existing := range AllowedShipmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShipmentType", value)
}

// NewShipmentTypeFromValue returns a pointer to a valid ShipmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShipmentTypeFromValue(v string) (*ShipmentType, error) {
	ev := ShipmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShipmentType: valid values are %v", v, AllowedShipmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShipmentType) IsValid() bool {
	for _, existing := range AllowedShipmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShipmentType value
func (v ShipmentType) Ptr() *ShipmentType {
	return &v
}

type NullableShipmentType struct {
	value *ShipmentType
	isSet bool
}

func (v NullableShipmentType) Get() *ShipmentType {
	return v.value
}

func (v *NullableShipmentType) Set(val *ShipmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentType(val *ShipmentType) *NullableShipmentType {
	return &NullableShipmentType{value: val, isSet: true}
}

func (v NullableShipmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

