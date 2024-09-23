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

// AffectedOrderQualityRatingComponentType Составляющие индекса качества.  **Для модели DBS:** * `DBS_CANCELLATION_RATE` — доля отмененных товаров. * `DBS_LATE_DELIVERY_RATE` — доля заказов, доставленных после плановой даты.  **Для моделей FBS и Экспресс:** * `FBS_CANCELLATION_RATE` — доля отмененных товаров. * `FBS_LATE_SHIP_RATE` — доля не вовремя отгруженных заказов. 
type AffectedOrderQualityRatingComponentType string

// List of AffectedOrderQualityRatingComponentType
const (
	DBS_CANCELLATION_RATE AffectedOrderQualityRatingComponentType = "DBS_CANCELLATION_RATE"
	DBS_LATE_DELIVERY_RATE AffectedOrderQualityRatingComponentType = "DBS_LATE_DELIVERY_RATE"
	FBS_CANCELLATION_RATE AffectedOrderQualityRatingComponentType = "FBS_CANCELLATION_RATE"
	FBS_LATE_SHIP_RATE AffectedOrderQualityRatingComponentType = "FBS_LATE_SHIP_RATE"
)

// All allowed values of AffectedOrderQualityRatingComponentType enum
var AllowedAffectedOrderQualityRatingComponentTypeEnumValues = []AffectedOrderQualityRatingComponentType{
	"DBS_CANCELLATION_RATE",
	"DBS_LATE_DELIVERY_RATE",
	"FBS_CANCELLATION_RATE",
	"FBS_LATE_SHIP_RATE",
}

func (v *AffectedOrderQualityRatingComponentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AffectedOrderQualityRatingComponentType(value)
	for _, existing := range AllowedAffectedOrderQualityRatingComponentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AffectedOrderQualityRatingComponentType", value)
}

// NewAffectedOrderQualityRatingComponentTypeFromValue returns a pointer to a valid AffectedOrderQualityRatingComponentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAffectedOrderQualityRatingComponentTypeFromValue(v string) (*AffectedOrderQualityRatingComponentType, error) {
	ev := AffectedOrderQualityRatingComponentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AffectedOrderQualityRatingComponentType: valid values are %v", v, AllowedAffectedOrderQualityRatingComponentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AffectedOrderQualityRatingComponentType) IsValid() bool {
	for _, existing := range AllowedAffectedOrderQualityRatingComponentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AffectedOrderQualityRatingComponentType value
func (v AffectedOrderQualityRatingComponentType) Ptr() *AffectedOrderQualityRatingComponentType {
	return &v
}

type NullableAffectedOrderQualityRatingComponentType struct {
	value *AffectedOrderQualityRatingComponentType
	isSet bool
}

func (v NullableAffectedOrderQualityRatingComponentType) Get() *AffectedOrderQualityRatingComponentType {
	return v.value
}

func (v *NullableAffectedOrderQualityRatingComponentType) Set(val *AffectedOrderQualityRatingComponentType) {
	v.value = val
	v.isSet = true
}

func (v NullableAffectedOrderQualityRatingComponentType) IsSet() bool {
	return v.isSet
}

func (v *NullableAffectedOrderQualityRatingComponentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffectedOrderQualityRatingComponentType(val *AffectedOrderQualityRatingComponentType) *NullableAffectedOrderQualityRatingComponentType {
	return &NullableAffectedOrderQualityRatingComponentType{value: val, isSet: true}
}

func (v NullableAffectedOrderQualityRatingComponentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffectedOrderQualityRatingComponentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

