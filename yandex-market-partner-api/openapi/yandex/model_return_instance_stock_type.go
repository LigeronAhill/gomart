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

// ReturnInstanceStockType Тип остатка на складе.
type ReturnInstanceStockType string

// List of ReturnInstanceStockType
const (
	FIT ReturnInstanceStockType = "FIT"
	DEFECT ReturnInstanceStockType = "DEFECT"
	ANOMALY ReturnInstanceStockType = "ANOMALY"
	SURPLUS ReturnInstanceStockType = "SURPLUS"
	EXPIRED ReturnInstanceStockType = "EXPIRED"
	MISGRADING ReturnInstanceStockType = "MISGRADING"
	UNDEFINED ReturnInstanceStockType = "UNDEFINED"
	INCORRECT_IMEI ReturnInstanceStockType = "INCORRECT_IMEI"
	INCORRECT_SERIAL_NUMBER ReturnInstanceStockType = "INCORRECT_SERIAL_NUMBER"
	INCORRECT_CIS ReturnInstanceStockType = "INCORRECT_CIS"
	PART_MISSING ReturnInstanceStockType = "PART_MISSING"
	NON_COMPLIENT ReturnInstanceStockType = "NON_COMPLIENT"
	NOT_ACCEPTABLE ReturnInstanceStockType = "NOT_ACCEPTABLE"
	SERVICE ReturnInstanceStockType = "SERVICE"
	MARKDOWN ReturnInstanceStockType = "MARKDOWN"
	DEMO ReturnInstanceStockType = "DEMO"
	REPAIR ReturnInstanceStockType = "REPAIR"
	FIRMWARE ReturnInstanceStockType = "FIRMWARE"
	UNKNOWN ReturnInstanceStockType = "UNKNOWN"
)

// All allowed values of ReturnInstanceStockType enum
var AllowedReturnInstanceStockTypeEnumValues = []ReturnInstanceStockType{
	"FIT",
	"DEFECT",
	"ANOMALY",
	"SURPLUS",
	"EXPIRED",
	"MISGRADING",
	"UNDEFINED",
	"INCORRECT_IMEI",
	"INCORRECT_SERIAL_NUMBER",
	"INCORRECT_CIS",
	"PART_MISSING",
	"NON_COMPLIENT",
	"NOT_ACCEPTABLE",
	"SERVICE",
	"MARKDOWN",
	"DEMO",
	"REPAIR",
	"FIRMWARE",
	"UNKNOWN",
}

func (v *ReturnInstanceStockType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReturnInstanceStockType(value)
	for _, existing := range AllowedReturnInstanceStockTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReturnInstanceStockType", value)
}

// NewReturnInstanceStockTypeFromValue returns a pointer to a valid ReturnInstanceStockType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReturnInstanceStockTypeFromValue(v string) (*ReturnInstanceStockType, error) {
	ev := ReturnInstanceStockType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReturnInstanceStockType: valid values are %v", v, AllowedReturnInstanceStockTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReturnInstanceStockType) IsValid() bool {
	for _, existing := range AllowedReturnInstanceStockTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReturnInstanceStockType value
func (v ReturnInstanceStockType) Ptr() *ReturnInstanceStockType {
	return &v
}

type NullableReturnInstanceStockType struct {
	value *ReturnInstanceStockType
	isSet bool
}

func (v NullableReturnInstanceStockType) Get() *ReturnInstanceStockType {
	return v.value
}

func (v *NullableReturnInstanceStockType) Set(val *ReturnInstanceStockType) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnInstanceStockType) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnInstanceStockType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnInstanceStockType(val *ReturnInstanceStockType) *NullableReturnInstanceStockType {
	return &NullableReturnInstanceStockType{value: val, isSet: true}
}

func (v NullableReturnInstanceStockType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnInstanceStockType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

