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

// FeedDownloadErrorType Тип ошибки загрузки прайс-листа.  Возможные значения:  * `DOWNLOAD_ERROR` — ошибка загрузки прайс-листа. Например, проблема с DNS-сервером или обрыв интернет-соединения.   Проблема описана в параметре `description`.  * `DOWNLOAD_HTTP_ERROR` — Яндекс Маркет передал запрос на получение прайс-листа и получил в ответ HTTP-код, отличный от 2xx.  HTTP-код выведен в параметре `httpStatusCode`. 
type FeedDownloadErrorType string

// List of FeedDownloadErrorType
const (
	ERROR FeedDownloadErrorType = "DOWNLOAD_ERROR"
	HTTP_ERROR FeedDownloadErrorType = "DOWNLOAD_HTTP_ERROR"
)

// All allowed values of FeedDownloadErrorType enum
var AllowedFeedDownloadErrorTypeEnumValues = []FeedDownloadErrorType{
	"DOWNLOAD_ERROR",
	"DOWNLOAD_HTTP_ERROR",
}

func (v *FeedDownloadErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeedDownloadErrorType(value)
	for _, existing := range AllowedFeedDownloadErrorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeedDownloadErrorType", value)
}

// NewFeedDownloadErrorTypeFromValue returns a pointer to a valid FeedDownloadErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeedDownloadErrorTypeFromValue(v string) (*FeedDownloadErrorType, error) {
	ev := FeedDownloadErrorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeedDownloadErrorType: valid values are %v", v, AllowedFeedDownloadErrorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeedDownloadErrorType) IsValid() bool {
	for _, existing := range AllowedFeedDownloadErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeedDownloadErrorType value
func (v FeedDownloadErrorType) Ptr() *FeedDownloadErrorType {
	return &v
}

type NullableFeedDownloadErrorType struct {
	value *FeedDownloadErrorType
	isSet bool
}

func (v NullableFeedDownloadErrorType) Get() *FeedDownloadErrorType {
	return v.value
}

func (v *NullableFeedDownloadErrorType) Set(val *FeedDownloadErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedDownloadErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedDownloadErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedDownloadErrorType(val *FeedDownloadErrorType) *NullableFeedDownloadErrorType {
	return &NullableFeedDownloadErrorType{value: val, isSet: true}
}

func (v NullableFeedDownloadErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedDownloadErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

