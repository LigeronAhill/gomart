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

// LicenseCheckStatusType Статус проверки лицензии:  * `NEW` — лицензия проверяется. * `SUCCESS` — лицензия прошла проверку. * `FAIL` — лицензия не прошла проверку. 
type LicenseCheckStatusType string

// List of LicenseCheckStatusType
const (
	NEW LicenseCheckStatusType = "NEW"
	SUCCESS LicenseCheckStatusType = "SUCCESS"
	FAIL LicenseCheckStatusType = "FAIL"
	REVOKE LicenseCheckStatusType = "REVOKE"
	DONT_WANT LicenseCheckStatusType = "DONT_WANT"
	FAIL_MANUAL LicenseCheckStatusType = "FAIL_MANUAL"
)

// All allowed values of LicenseCheckStatusType enum
var AllowedLicenseCheckStatusTypeEnumValues = []LicenseCheckStatusType{
	"NEW",
	"SUCCESS",
	"FAIL",
	"REVOKE",
	"DONT_WANT",
	"FAIL_MANUAL",
}

func (v *LicenseCheckStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LicenseCheckStatusType(value)
	for _, existing := range AllowedLicenseCheckStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LicenseCheckStatusType", value)
}

// NewLicenseCheckStatusTypeFromValue returns a pointer to a valid LicenseCheckStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLicenseCheckStatusTypeFromValue(v string) (*LicenseCheckStatusType, error) {
	ev := LicenseCheckStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LicenseCheckStatusType: valid values are %v", v, AllowedLicenseCheckStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LicenseCheckStatusType) IsValid() bool {
	for _, existing := range AllowedLicenseCheckStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LicenseCheckStatusType value
func (v LicenseCheckStatusType) Ptr() *LicenseCheckStatusType {
	return &v
}

type NullableLicenseCheckStatusType struct {
	value *LicenseCheckStatusType
	isSet bool
}

func (v NullableLicenseCheckStatusType) Get() *LicenseCheckStatusType {
	return v.value
}

func (v *NullableLicenseCheckStatusType) Set(val *LicenseCheckStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseCheckStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseCheckStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseCheckStatusType(val *LicenseCheckStatusType) *NullableLicenseCheckStatusType {
	return &NullableLicenseCheckStatusType{value: val, isSet: true}
}

func (v NullableLicenseCheckStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseCheckStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

