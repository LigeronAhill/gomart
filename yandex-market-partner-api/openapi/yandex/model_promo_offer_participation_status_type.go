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

// PromoOfferParticipationStatusType Статус товара в акции:  * `AUTO` — добавлен автоматически во всех магазинах кабинета, в которых товар доступен для покупки.  * `PARTIALLY_AUTO` — добавлен автоматически у части магазинов.  * `MANUAL` — добавлен вручную.  * `NOT_PARTICIPATING` — не участвует в акции.  Об автоматическом и ручном добавлении товаров в акцию читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/marketing/promos/market/index). 
type PromoOfferParticipationStatusType string

// List of PromoOfferParticipationStatusType
const (
	AUTO PromoOfferParticipationStatusType = "AUTO"
	PARTIALLY_AUTO PromoOfferParticipationStatusType = "PARTIALLY_AUTO"
	MANUAL PromoOfferParticipationStatusType = "MANUAL"
	NOT_PARTICIPATING PromoOfferParticipationStatusType = "NOT_PARTICIPATING"
)

// All allowed values of PromoOfferParticipationStatusType enum
var AllowedPromoOfferParticipationStatusTypeEnumValues = []PromoOfferParticipationStatusType{
	"AUTO",
	"PARTIALLY_AUTO",
	"MANUAL",
	"NOT_PARTICIPATING",
}

func (v *PromoOfferParticipationStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PromoOfferParticipationStatusType(value)
	for _, existing := range AllowedPromoOfferParticipationStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PromoOfferParticipationStatusType", value)
}

// NewPromoOfferParticipationStatusTypeFromValue returns a pointer to a valid PromoOfferParticipationStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromoOfferParticipationStatusTypeFromValue(v string) (*PromoOfferParticipationStatusType, error) {
	ev := PromoOfferParticipationStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromoOfferParticipationStatusType: valid values are %v", v, AllowedPromoOfferParticipationStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromoOfferParticipationStatusType) IsValid() bool {
	for _, existing := range AllowedPromoOfferParticipationStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PromoOfferParticipationStatusType value
func (v PromoOfferParticipationStatusType) Ptr() *PromoOfferParticipationStatusType {
	return &v
}

type NullablePromoOfferParticipationStatusType struct {
	value *PromoOfferParticipationStatusType
	isSet bool
}

func (v NullablePromoOfferParticipationStatusType) Get() *PromoOfferParticipationStatusType {
	return v.value
}

func (v *NullablePromoOfferParticipationStatusType) Set(val *PromoOfferParticipationStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullablePromoOfferParticipationStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullablePromoOfferParticipationStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromoOfferParticipationStatusType(val *PromoOfferParticipationStatusType) *NullablePromoOfferParticipationStatusType {
	return &NullablePromoOfferParticipationStatusType{value: val, isSet: true}
}

func (v NullablePromoOfferParticipationStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromoOfferParticipationStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

