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

// ReportSubStatusType Подстатус генерации отчета: * `NO_DATA` — для такого отчета нет данных. * `TOO_LARGE` — отчет превысил допустимый размер — укажите меньший период времени или уточните условия запроса. * `RESOURCE_NOT_FOUND` — для такого отчета не удалось найти часть сущностей. 
type ReportSubStatusType string

// List of ReportSubStatusType
const (
	NO_DATA ReportSubStatusType = "NO_DATA"
	TOO_LARGE ReportSubStatusType = "TOO_LARGE"
	RESOURCE_NOT_FOUND ReportSubStatusType = "RESOURCE_NOT_FOUND"
)

// All allowed values of ReportSubStatusType enum
var AllowedReportSubStatusTypeEnumValues = []ReportSubStatusType{
	"NO_DATA",
	"TOO_LARGE",
	"RESOURCE_NOT_FOUND",
}

func (v *ReportSubStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportSubStatusType(value)
	for _, existing := range AllowedReportSubStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportSubStatusType", value)
}

// NewReportSubStatusTypeFromValue returns a pointer to a valid ReportSubStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportSubStatusTypeFromValue(v string) (*ReportSubStatusType, error) {
	ev := ReportSubStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportSubStatusType: valid values are %v", v, AllowedReportSubStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportSubStatusType) IsValid() bool {
	for _, existing := range AllowedReportSubStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportSubStatusType value
func (v ReportSubStatusType) Ptr() *ReportSubStatusType {
	return &v
}

type NullableReportSubStatusType struct {
	value *ReportSubStatusType
	isSet bool
}

func (v NullableReportSubStatusType) Get() *ReportSubStatusType {
	return v.value
}

func (v *NullableReportSubStatusType) Set(val *ReportSubStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableReportSubStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableReportSubStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportSubStatusType(val *ReportSubStatusType) *NullableReportSubStatusType {
	return &NullableReportSubStatusType{value: val, isSet: true}
}

func (v NullableReportSubStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportSubStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

