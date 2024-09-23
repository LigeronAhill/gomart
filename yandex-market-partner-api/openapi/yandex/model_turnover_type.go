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

// TurnoverType Оценка оборачиваемости.  |enum|Диапазон оборачиваемости|Комментарий| |-|-|-| |`LOW`|`turnoverDays` ≥ 120|| |`ALMOST_LOW`|100 ≤ `turnoverDays` < 120|| |`HIGH`|45 ≤ `turnoverDays` < 100|| |`VERY_HIGH`|0 ≤ `turnoverDays` < 45|| |`NO_SALES`|—|Продаж нет.| |`FREE_STORE`|Любое значение.|Товары этой категории сейчас хранятся бесплатно.| 
type TurnoverType string

// List of TurnoverType
const (
	LOW TurnoverType = "LOW"
	ALMOST_LOW TurnoverType = "ALMOST_LOW"
	HIGH TurnoverType = "HIGH"
	VERY_HIGH TurnoverType = "VERY_HIGH"
	NO_SALES TurnoverType = "NO_SALES"
	FREE_STORE TurnoverType = "FREE_STORE"
)

// All allowed values of TurnoverType enum
var AllowedTurnoverTypeEnumValues = []TurnoverType{
	"LOW",
	"ALMOST_LOW",
	"HIGH",
	"VERY_HIGH",
	"NO_SALES",
	"FREE_STORE",
}

func (v *TurnoverType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TurnoverType(value)
	for _, existing := range AllowedTurnoverTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TurnoverType", value)
}

// NewTurnoverTypeFromValue returns a pointer to a valid TurnoverType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTurnoverTypeFromValue(v string) (*TurnoverType, error) {
	ev := TurnoverType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TurnoverType: valid values are %v", v, AllowedTurnoverTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TurnoverType) IsValid() bool {
	for _, existing := range AllowedTurnoverTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TurnoverType value
func (v TurnoverType) Ptr() *TurnoverType {
	return &v
}

type NullableTurnoverType struct {
	value *TurnoverType
	isSet bool
}

func (v NullableTurnoverType) Get() *TurnoverType {
	return v.value
}

func (v *NullableTurnoverType) Set(val *TurnoverType) {
	v.value = val
	v.isSet = true
}

func (v NullableTurnoverType) IsSet() bool {
	return v.isSet
}

func (v *NullableTurnoverType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTurnoverType(val *TurnoverType) *NullableTurnoverType {
	return &NullableTurnoverType{value: val, isSet: true}
}

func (v NullableTurnoverType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTurnoverType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

