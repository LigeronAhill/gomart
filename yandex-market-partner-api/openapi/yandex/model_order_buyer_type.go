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

// OrderBuyerType Тип покупателя:  * `PERSON` — физическое лицо.  * `BUSINESS` — организация. 
type OrderBuyerType string

// List of OrderBuyerType
const (
	PERSON OrderBuyerType = "PERSON"
	BUSINESS OrderBuyerType = "BUSINESS"
)

// All allowed values of OrderBuyerType enum
var AllowedOrderBuyerTypeEnumValues = []OrderBuyerType{
	"PERSON",
	"BUSINESS",
}

func (v *OrderBuyerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderBuyerType(value)
	for _, existing := range AllowedOrderBuyerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderBuyerType", value)
}

// NewOrderBuyerTypeFromValue returns a pointer to a valid OrderBuyerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderBuyerTypeFromValue(v string) (*OrderBuyerType, error) {
	ev := OrderBuyerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderBuyerType: valid values are %v", v, AllowedOrderBuyerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderBuyerType) IsValid() bool {
	for _, existing := range AllowedOrderBuyerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderBuyerType value
func (v OrderBuyerType) Ptr() *OrderBuyerType {
	return &v
}

type NullableOrderBuyerType struct {
	value *OrderBuyerType
	isSet bool
}

func (v NullableOrderBuyerType) Get() *OrderBuyerType {
	return v.value
}

func (v *NullableOrderBuyerType) Set(val *OrderBuyerType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBuyerType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBuyerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderBuyerType(val *OrderBuyerType) *NullableOrderBuyerType {
	return &NullableOrderBuyerType{value: val, isSet: true}
}

func (v NullableOrderBuyerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBuyerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

