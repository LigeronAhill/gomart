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

// checks if the PriceDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceDTO{}

// PriceDTO Цена с указанием скидки, валюты и времени последнего обновления.
type PriceDTO struct {
	// Цена на товар.
	Value *float32 `json:"value,omitempty"`
	// Цена на товар без скидки.  Число должно быть целым. Вы можете указать цену со скидкой от 5 до 99%.  Передавайте этот параметр при каждом обновлении цены, если предоставляете скидку на товар. 
	DiscountBase *float32 `json:"discountBase,omitempty"`
	CurrencyId *CurrencyType `json:"currencyId,omitempty"`
	// Идентификатор ставки НДС, применяемой для товара:  * `2` — 10%. * `5` — 0%. * `6` — не облагается НДС. * `7` — 20%.  Если параметр не указан, используется ставка НДС, установленная в кабинете. 
	Vat *int32 `json:"vat,omitempty"`
}

// NewPriceDTO instantiates a new PriceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceDTO() *PriceDTO {
	this := PriceDTO{}
	return &this
}

// NewPriceDTOWithDefaults instantiates a new PriceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceDTOWithDefaults() *PriceDTO {
	this := PriceDTO{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PriceDTO) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceDTO) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PriceDTO) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *PriceDTO) SetValue(v float32) {
	o.Value = &v
}

// GetDiscountBase returns the DiscountBase field value if set, zero value otherwise.
func (o *PriceDTO) GetDiscountBase() float32 {
	if o == nil || IsNil(o.DiscountBase) {
		var ret float32
		return ret
	}
	return *o.DiscountBase
}

// GetDiscountBaseOk returns a tuple with the DiscountBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceDTO) GetDiscountBaseOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountBase) {
		return nil, false
	}
	return o.DiscountBase, true
}

// HasDiscountBase returns a boolean if a field has been set.
func (o *PriceDTO) HasDiscountBase() bool {
	if o != nil && !IsNil(o.DiscountBase) {
		return true
	}

	return false
}

// SetDiscountBase gets a reference to the given float32 and assigns it to the DiscountBase field.
func (o *PriceDTO) SetDiscountBase(v float32) {
	o.DiscountBase = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *PriceDTO) GetCurrencyId() CurrencyType {
	if o == nil || IsNil(o.CurrencyId) {
		var ret CurrencyType
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceDTO) GetCurrencyIdOk() (*CurrencyType, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *PriceDTO) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given CurrencyType and assigns it to the CurrencyId field.
func (o *PriceDTO) SetCurrencyId(v CurrencyType) {
	o.CurrencyId = &v
}

// GetVat returns the Vat field value if set, zero value otherwise.
func (o *PriceDTO) GetVat() int32 {
	if o == nil || IsNil(o.Vat) {
		var ret int32
		return ret
	}
	return *o.Vat
}

// GetVatOk returns a tuple with the Vat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceDTO) GetVatOk() (*int32, bool) {
	if o == nil || IsNil(o.Vat) {
		return nil, false
	}
	return o.Vat, true
}

// HasVat returns a boolean if a field has been set.
func (o *PriceDTO) HasVat() bool {
	if o != nil && !IsNil(o.Vat) {
		return true
	}

	return false
}

// SetVat gets a reference to the given int32 and assigns it to the Vat field.
func (o *PriceDTO) SetVat(v int32) {
	o.Vat = &v
}

func (o PriceDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.DiscountBase) {
		toSerialize["discountBase"] = o.DiscountBase
	}
	if !IsNil(o.CurrencyId) {
		toSerialize["currencyId"] = o.CurrencyId
	}
	if !IsNil(o.Vat) {
		toSerialize["vat"] = o.Vat
	}
	return toSerialize, nil
}

type NullablePriceDTO struct {
	value *PriceDTO
	isSet bool
}

func (v NullablePriceDTO) Get() *PriceDTO {
	return v.value
}

func (v *NullablePriceDTO) Set(val *PriceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceDTO(val *PriceDTO) *NullablePriceDTO {
	return &NullablePriceDTO{value: val, isSet: true}
}

func (v NullablePriceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


