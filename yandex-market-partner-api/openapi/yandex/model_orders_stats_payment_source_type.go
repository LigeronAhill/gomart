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

// OrdersStatsPaymentSourceType Способ денежного перевода: - `BUYER` — оплата или возврат деньгами. - `CASHBACK` — оплата или возврат баллами Плюса. - `MARKETPLACE` — оплата или возврат купонами. - `SPLIT` — оплата картой по частям (Сплит). 
type OrdersStatsPaymentSourceType string

// List of OrdersStatsPaymentSourceType
const (
	BUYER OrdersStatsPaymentSourceType = "BUYER"
	CASHBACK OrdersStatsPaymentSourceType = "CASHBACK"
	MARKETPLACE OrdersStatsPaymentSourceType = "MARKETPLACE"
	SPLIT OrdersStatsPaymentSourceType = "SPLIT"
)

// All allowed values of OrdersStatsPaymentSourceType enum
var AllowedOrdersStatsPaymentSourceTypeEnumValues = []OrdersStatsPaymentSourceType{
	"BUYER",
	"CASHBACK",
	"MARKETPLACE",
	"SPLIT",
}

func (v *OrdersStatsPaymentSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrdersStatsPaymentSourceType(value)
	for _, existing := range AllowedOrdersStatsPaymentSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrdersStatsPaymentSourceType", value)
}

// NewOrdersStatsPaymentSourceTypeFromValue returns a pointer to a valid OrdersStatsPaymentSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrdersStatsPaymentSourceTypeFromValue(v string) (*OrdersStatsPaymentSourceType, error) {
	ev := OrdersStatsPaymentSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrdersStatsPaymentSourceType: valid values are %v", v, AllowedOrdersStatsPaymentSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrdersStatsPaymentSourceType) IsValid() bool {
	for _, existing := range AllowedOrdersStatsPaymentSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrdersStatsPaymentSourceType value
func (v OrdersStatsPaymentSourceType) Ptr() *OrdersStatsPaymentSourceType {
	return &v
}

type NullableOrdersStatsPaymentSourceType struct {
	value *OrdersStatsPaymentSourceType
	isSet bool
}

func (v NullableOrdersStatsPaymentSourceType) Get() *OrdersStatsPaymentSourceType {
	return v.value
}

func (v *NullableOrdersStatsPaymentSourceType) Set(val *OrdersStatsPaymentSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersStatsPaymentSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersStatsPaymentSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersStatsPaymentSourceType(val *OrdersStatsPaymentSourceType) *NullableOrdersStatsPaymentSourceType {
	return &NullableOrdersStatsPaymentSourceType{value: val, isSet: true}
}

func (v NullableOrdersStatsPaymentSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersStatsPaymentSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

