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

// FeedIndexLogsErrorType Тип ошибки индексации прайс-листа.  Возможные значения:  * `DOWNLOAD_ERROR` — ошибка загрузки прайс-листа. Например, проблема с DNS-сервером или обрыв интернет-соединения.    Проблема описана в параметре `description`.  * `DOWNLOAD_HTTP_ERROR` — Маркет передал запрос на получение прайс-листа и получил в ответ HTTP-код, отличный от 2xx.    HTTP-код выведен в параметре `httpStatusCode`.  * `PARSE_ERROR` — ошибка при проверке прайс-листа, не связанная с форматом YML. Например, прайс-лист пустой или его не удалось разархивировать.  * `PARSE_XML_ERROR` — несоответствие техническим требованиям формата YML. Например, элементы и их значения описаны некорректно.  * `TOO_MANY_REJECTED_OFFERS` — более чем в половине предложений из прайс-листа найдены ошибки. Все предложения из прайс-листа не будут опубликованы на Маркете. 
type FeedIndexLogsErrorType string

// List of FeedIndexLogsErrorType
const (
	DOWNLOAD_ERROR FeedIndexLogsErrorType = "DOWNLOAD_ERROR"
	DOWNLOAD_HTTP_ERROR FeedIndexLogsErrorType = "DOWNLOAD_HTTP_ERROR"
	PARSE_ERROR FeedIndexLogsErrorType = "PARSE_ERROR"
	PARSE_XML_ERROR FeedIndexLogsErrorType = "PARSE_XML_ERROR"
	TOO_MANY_REJECTED_OFFERS FeedIndexLogsErrorType = "TOO_MANY_REJECTED_OFFERS"
	NOT_INDEXED FeedIndexLogsErrorType = "NOT_INDEXED"
)

// All allowed values of FeedIndexLogsErrorType enum
var AllowedFeedIndexLogsErrorTypeEnumValues = []FeedIndexLogsErrorType{
	"DOWNLOAD_ERROR",
	"DOWNLOAD_HTTP_ERROR",
	"PARSE_ERROR",
	"PARSE_XML_ERROR",
	"TOO_MANY_REJECTED_OFFERS",
	"NOT_INDEXED",
}

func (v *FeedIndexLogsErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeedIndexLogsErrorType(value)
	for _, existing := range AllowedFeedIndexLogsErrorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeedIndexLogsErrorType", value)
}

// NewFeedIndexLogsErrorTypeFromValue returns a pointer to a valid FeedIndexLogsErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeedIndexLogsErrorTypeFromValue(v string) (*FeedIndexLogsErrorType, error) {
	ev := FeedIndexLogsErrorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeedIndexLogsErrorType: valid values are %v", v, AllowedFeedIndexLogsErrorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeedIndexLogsErrorType) IsValid() bool {
	for _, existing := range AllowedFeedIndexLogsErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeedIndexLogsErrorType value
func (v FeedIndexLogsErrorType) Ptr() *FeedIndexLogsErrorType {
	return &v
}

type NullableFeedIndexLogsErrorType struct {
	value *FeedIndexLogsErrorType
	isSet bool
}

func (v NullableFeedIndexLogsErrorType) Get() *FeedIndexLogsErrorType {
	return v.value
}

func (v *NullableFeedIndexLogsErrorType) Set(val *FeedIndexLogsErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedIndexLogsErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedIndexLogsErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedIndexLogsErrorType(val *FeedIndexLogsErrorType) *NullableFeedIndexLogsErrorType {
	return &NullableFeedIndexLogsErrorType{value: val, isSet: true}
}

func (v NullableFeedIndexLogsErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedIndexLogsErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

