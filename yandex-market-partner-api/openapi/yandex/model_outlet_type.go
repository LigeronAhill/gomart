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

// OutletType Тип точки продаж.  Возможные значения:  * `DEPOT` — пункт выдачи заказов. * `MIXED` — смешанный тип точки продаж (торговый зал и пункт выдачи заказов). * `RETAIL` — розничная точка продаж (торговый зал). 
type OutletType string

// List of OutletType
const (
	DEPOT OutletType = "DEPOT"
	MIXED OutletType = "MIXED"
	RETAIL OutletType = "RETAIL"
	NOT_DEFINED OutletType = "NOT_DEFINED"
)

// All allowed values of OutletType enum
var AllowedOutletTypeEnumValues = []OutletType{
	"DEPOT",
	"MIXED",
	"RETAIL",
	"NOT_DEFINED",
}

func (v *OutletType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OutletType(value)
	for _, existing := range AllowedOutletTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OutletType", value)
}

// NewOutletTypeFromValue returns a pointer to a valid OutletType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOutletTypeFromValue(v string) (*OutletType, error) {
	ev := OutletType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OutletType: valid values are %v", v, AllowedOutletTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OutletType) IsValid() bool {
	for _, existing := range AllowedOutletTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OutletType value
func (v OutletType) Ptr() *OutletType {
	return &v
}

type NullableOutletType struct {
	value *OutletType
	isSet bool
}

func (v NullableOutletType) Get() *OutletType {
	return v.value
}

func (v *NullableOutletType) Set(val *OutletType) {
	v.value = val
	v.isSet = true
}

func (v NullableOutletType) IsSet() bool {
	return v.isSet
}

func (v *NullableOutletType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutletType(val *OutletType) *NullableOutletType {
	return &NullableOutletType{value: val, isSet: true}
}

func (v NullableOutletType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutletType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

