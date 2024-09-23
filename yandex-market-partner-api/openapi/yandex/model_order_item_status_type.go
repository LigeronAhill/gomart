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

// OrderItemStatusType Возвращенный или невыкупленный товар:  * `REJECTED` — невыкупленный.  * `RETURNED` — возвращенный. 
type OrderItemStatusType string

// List of OrderItemStatusType
const (
	REJECTED OrderItemStatusType = "REJECTED"
	RETURNED OrderItemStatusType = "RETURNED"
)

// All allowed values of OrderItemStatusType enum
var AllowedOrderItemStatusTypeEnumValues = []OrderItemStatusType{
	"REJECTED",
	"RETURNED",
}

func (v *OrderItemStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderItemStatusType(value)
	for _, existing := range AllowedOrderItemStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderItemStatusType", value)
}

// NewOrderItemStatusTypeFromValue returns a pointer to a valid OrderItemStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderItemStatusTypeFromValue(v string) (*OrderItemStatusType, error) {
	ev := OrderItemStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderItemStatusType: valid values are %v", v, AllowedOrderItemStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderItemStatusType) IsValid() bool {
	for _, existing := range AllowedOrderItemStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderItemStatusType value
func (v OrderItemStatusType) Ptr() *OrderItemStatusType {
	return &v
}

type NullableOrderItemStatusType struct {
	value *OrderItemStatusType
	isSet bool
}

func (v NullableOrderItemStatusType) Get() *OrderItemStatusType {
	return v.value
}

func (v *NullableOrderItemStatusType) Set(val *OrderItemStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItemStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItemStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItemStatusType(val *OrderItemStatusType) *NullableOrderItemStatusType {
	return &NullableOrderItemStatusType{value: val, isSet: true}
}

func (v NullableOrderItemStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderItemStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

