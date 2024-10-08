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

// EacVerificationStatusType Статус проверки кода подтверждения:  * `ACCEPTED` — код верный. * `REJECTED` — код неверный. * `NEED_UPDATE` — Маркет отправит новый код. Значение возвращается, если превышено количество попыток отправки кода. 
type EacVerificationStatusType string

// List of EacVerificationStatusType
const (
	ACCEPTED EacVerificationStatusType = "ACCEPTED"
	REJECTED EacVerificationStatusType = "REJECTED"
	NEED_UPDATE EacVerificationStatusType = "NEED_UPDATE"
)

// All allowed values of EacVerificationStatusType enum
var AllowedEacVerificationStatusTypeEnumValues = []EacVerificationStatusType{
	"ACCEPTED",
	"REJECTED",
	"NEED_UPDATE",
}

func (v *EacVerificationStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EacVerificationStatusType(value)
	for _, existing := range AllowedEacVerificationStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EacVerificationStatusType", value)
}

// NewEacVerificationStatusTypeFromValue returns a pointer to a valid EacVerificationStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEacVerificationStatusTypeFromValue(v string) (*EacVerificationStatusType, error) {
	ev := EacVerificationStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EacVerificationStatusType: valid values are %v", v, AllowedEacVerificationStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EacVerificationStatusType) IsValid() bool {
	for _, existing := range AllowedEacVerificationStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EacVerificationStatusType value
func (v EacVerificationStatusType) Ptr() *EacVerificationStatusType {
	return &v
}

type NullableEacVerificationStatusType struct {
	value *EacVerificationStatusType
	isSet bool
}

func (v NullableEacVerificationStatusType) Get() *EacVerificationStatusType {
	return v.value
}

func (v *NullableEacVerificationStatusType) Set(val *EacVerificationStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableEacVerificationStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableEacVerificationStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEacVerificationStatusType(val *EacVerificationStatusType) *NullableEacVerificationStatusType {
	return &NullableEacVerificationStatusType{value: val, isSet: true}
}

func (v NullableEacVerificationStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEacVerificationStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

