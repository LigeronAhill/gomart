/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BusinessSettingsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BusinessSettingsDTO{}

// BusinessSettingsDTO Настройки кабинета.
type BusinessSettingsDTO struct {
	// Можно ли установить только [базовую цену](*rule): * `false` — можно задать и базовую цену, и цены в конкретных магазинах. * `true` — можно задать только базовую цену. 
	OnlyDefaultPrice *bool `json:"onlyDefaultPrice,omitempty"`
	Currency *CurrencyType `json:"currency,omitempty"`
}

// NewBusinessSettingsDTO instantiates a new BusinessSettingsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBusinessSettingsDTO() *BusinessSettingsDTO {
	this := BusinessSettingsDTO{}
	return &this
}

// NewBusinessSettingsDTOWithDefaults instantiates a new BusinessSettingsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBusinessSettingsDTOWithDefaults() *BusinessSettingsDTO {
	this := BusinessSettingsDTO{}
	return &this
}

// GetOnlyDefaultPrice returns the OnlyDefaultPrice field value if set, zero value otherwise.
func (o *BusinessSettingsDTO) GetOnlyDefaultPrice() bool {
	if o == nil || IsNil(o.OnlyDefaultPrice) {
		var ret bool
		return ret
	}
	return *o.OnlyDefaultPrice
}

// GetOnlyDefaultPriceOk returns a tuple with the OnlyDefaultPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessSettingsDTO) GetOnlyDefaultPriceOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyDefaultPrice) {
		return nil, false
	}
	return o.OnlyDefaultPrice, true
}

// HasOnlyDefaultPrice returns a boolean if a field has been set.
func (o *BusinessSettingsDTO) HasOnlyDefaultPrice() bool {
	if o != nil && !IsNil(o.OnlyDefaultPrice) {
		return true
	}

	return false
}

// SetOnlyDefaultPrice gets a reference to the given bool and assigns it to the OnlyDefaultPrice field.
func (o *BusinessSettingsDTO) SetOnlyDefaultPrice(v bool) {
	o.OnlyDefaultPrice = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *BusinessSettingsDTO) GetCurrency() CurrencyType {
	if o == nil || IsNil(o.Currency) {
		var ret CurrencyType
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BusinessSettingsDTO) GetCurrencyOk() (*CurrencyType, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *BusinessSettingsDTO) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given CurrencyType and assigns it to the Currency field.
func (o *BusinessSettingsDTO) SetCurrency(v CurrencyType) {
	o.Currency = &v
}

func (o BusinessSettingsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BusinessSettingsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnlyDefaultPrice) {
		toSerialize["onlyDefaultPrice"] = o.OnlyDefaultPrice
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableBusinessSettingsDTO struct {
	value *BusinessSettingsDTO
	isSet bool
}

func (v NullableBusinessSettingsDTO) Get() *BusinessSettingsDTO {
	return v.value
}

func (v *NullableBusinessSettingsDTO) Set(val *BusinessSettingsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableBusinessSettingsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableBusinessSettingsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBusinessSettingsDTO(val *BusinessSettingsDTO) *NullableBusinessSettingsDTO {
	return &NullableBusinessSettingsDTO{value: val, isSet: true}
}

func (v NullableBusinessSettingsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBusinessSettingsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


