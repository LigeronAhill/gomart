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

// ReturnDecisionReasonType Причины возврата:  * `BAD_QUALITY` — бракованный товар (есть недостатки).  * `DO_NOT_FIT` — товар не подошел.  * `WRONG_ITEM` — привезли не тот товар.  * `DAMAGE_DELIVERY` — товар поврежден при доставке.  * `LOYALTY_FAIL` — невозможно установить виновного в браке/пересорте.  * `CONTENT_FAIL` — ошибочное описание товара по вине Маркета.  * `UNKNOWN` — причина не известна. 
type ReturnDecisionReasonType string

// List of ReturnDecisionReasonType
const (
	BAD_QUALITY ReturnDecisionReasonType = "BAD_QUALITY"
	DOES_NOT_FIT ReturnDecisionReasonType = "DOES_NOT_FIT"
	WRONG_ITEM ReturnDecisionReasonType = "WRONG_ITEM"
	DAMAGE_DELIVERY ReturnDecisionReasonType = "DAMAGE_DELIVERY"
	LOYALTY_FAIL ReturnDecisionReasonType = "LOYALTY_FAIL"
	CONTENT_FAIL ReturnDecisionReasonType = "CONTENT_FAIL"
	UNKNOWN ReturnDecisionReasonType = "UNKNOWN"
)

// All allowed values of ReturnDecisionReasonType enum
var AllowedReturnDecisionReasonTypeEnumValues = []ReturnDecisionReasonType{
	"BAD_QUALITY",
	"DOES_NOT_FIT",
	"WRONG_ITEM",
	"DAMAGE_DELIVERY",
	"LOYALTY_FAIL",
	"CONTENT_FAIL",
	"UNKNOWN",
}

func (v *ReturnDecisionReasonType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReturnDecisionReasonType(value)
	for _, existing := range AllowedReturnDecisionReasonTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReturnDecisionReasonType", value)
}

// NewReturnDecisionReasonTypeFromValue returns a pointer to a valid ReturnDecisionReasonType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReturnDecisionReasonTypeFromValue(v string) (*ReturnDecisionReasonType, error) {
	ev := ReturnDecisionReasonType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReturnDecisionReasonType: valid values are %v", v, AllowedReturnDecisionReasonTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReturnDecisionReasonType) IsValid() bool {
	for _, existing := range AllowedReturnDecisionReasonTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReturnDecisionReasonType value
func (v ReturnDecisionReasonType) Ptr() *ReturnDecisionReasonType {
	return &v
}

type NullableReturnDecisionReasonType struct {
	value *ReturnDecisionReasonType
	isSet bool
}

func (v NullableReturnDecisionReasonType) Get() *ReturnDecisionReasonType {
	return v.value
}

func (v *NullableReturnDecisionReasonType) Set(val *ReturnDecisionReasonType) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnDecisionReasonType) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnDecisionReasonType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnDecisionReasonType(val *ReturnDecisionReasonType) *NullableReturnDecisionReasonType {
	return &NullableReturnDecisionReasonType{value: val, isSet: true}
}

func (v NullableReturnDecisionReasonType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnDecisionReasonType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

