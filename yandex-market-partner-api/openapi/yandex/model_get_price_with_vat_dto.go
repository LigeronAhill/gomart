/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the GetPriceWithVatDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPriceWithVatDTO{}

// GetPriceWithVatDTO Цена с указанием ставки НДС и времени последнего обновления.
type GetPriceWithVatDTO struct {
	// Цена на товар.
	Value *float32 `json:"value,omitempty"`
	// Цена на товар без скидки.  Число должно быть целым. Вы можете указать цену со скидкой от 5 до 99%.  Передавайте этот параметр при каждом обновлении цены, если предоставляете скидку на товар. 
	DiscountBase *float32 `json:"discountBase,omitempty"`
	CurrencyId *CurrencyType `json:"currencyId,omitempty"`
	// Идентификатор ставки НДС, применяемой для товара:  * `2` — 10%. * `5` — 0%. * `6` — не облагается НДС. * `7` — 20%.  Если параметр не указан, используется ставка НДС, установленная в кабинете. 
	Vat *int32 `json:"vat,omitempty"`
	// Время последнего обновления.
	UpdatedAt time.Time `json:"updatedAt"`
}

type _GetPriceWithVatDTO GetPriceWithVatDTO

// NewGetPriceWithVatDTO instantiates a new GetPriceWithVatDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPriceWithVatDTO(updatedAt time.Time) *GetPriceWithVatDTO {
	this := GetPriceWithVatDTO{}
	this.UpdatedAt = updatedAt
	return &this
}

// NewGetPriceWithVatDTOWithDefaults instantiates a new GetPriceWithVatDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPriceWithVatDTOWithDefaults() *GetPriceWithVatDTO {
	this := GetPriceWithVatDTO{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetPriceWithVatDTO) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceWithVatDTO) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetPriceWithVatDTO) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *GetPriceWithVatDTO) SetValue(v float32) {
	o.Value = &v
}

// GetDiscountBase returns the DiscountBase field value if set, zero value otherwise.
func (o *GetPriceWithVatDTO) GetDiscountBase() float32 {
	if o == nil || IsNil(o.DiscountBase) {
		var ret float32
		return ret
	}
	return *o.DiscountBase
}

// GetDiscountBaseOk returns a tuple with the DiscountBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceWithVatDTO) GetDiscountBaseOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountBase) {
		return nil, false
	}
	return o.DiscountBase, true
}

// HasDiscountBase returns a boolean if a field has been set.
func (o *GetPriceWithVatDTO) HasDiscountBase() bool {
	if o != nil && !IsNil(o.DiscountBase) {
		return true
	}

	return false
}

// SetDiscountBase gets a reference to the given float32 and assigns it to the DiscountBase field.
func (o *GetPriceWithVatDTO) SetDiscountBase(v float32) {
	o.DiscountBase = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *GetPriceWithVatDTO) GetCurrencyId() CurrencyType {
	if o == nil || IsNil(o.CurrencyId) {
		var ret CurrencyType
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceWithVatDTO) GetCurrencyIdOk() (*CurrencyType, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *GetPriceWithVatDTO) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given CurrencyType and assigns it to the CurrencyId field.
func (o *GetPriceWithVatDTO) SetCurrencyId(v CurrencyType) {
	o.CurrencyId = &v
}

// GetVat returns the Vat field value if set, zero value otherwise.
func (o *GetPriceWithVatDTO) GetVat() int32 {
	if o == nil || IsNil(o.Vat) {
		var ret int32
		return ret
	}
	return *o.Vat
}

// GetVatOk returns a tuple with the Vat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceWithVatDTO) GetVatOk() (*int32, bool) {
	if o == nil || IsNil(o.Vat) {
		return nil, false
	}
	return o.Vat, true
}

// HasVat returns a boolean if a field has been set.
func (o *GetPriceWithVatDTO) HasVat() bool {
	if o != nil && !IsNil(o.Vat) {
		return true
	}

	return false
}

// SetVat gets a reference to the given int32 and assigns it to the Vat field.
func (o *GetPriceWithVatDTO) SetVat(v int32) {
	o.Vat = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetPriceWithVatDTO) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetPriceWithVatDTO) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetPriceWithVatDTO) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o GetPriceWithVatDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPriceWithVatDTO) ToMap() (map[string]interface{}, error) {
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
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *GetPriceWithVatDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"updatedAt",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetPriceWithVatDTO := _GetPriceWithVatDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPriceWithVatDTO)

	if err != nil {
		return err
	}

	*o = GetPriceWithVatDTO(varGetPriceWithVatDTO)

	return err
}

type NullableGetPriceWithVatDTO struct {
	value *GetPriceWithVatDTO
	isSet bool
}

func (v NullableGetPriceWithVatDTO) Get() *GetPriceWithVatDTO {
	return v.value
}

func (v *NullableGetPriceWithVatDTO) Set(val *GetPriceWithVatDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPriceWithVatDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPriceWithVatDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPriceWithVatDTO(val *GetPriceWithVatDTO) *NullableGetPriceWithVatDTO {
	return &NullableGetPriceWithVatDTO{value: val, isSet: true}
}

func (v NullableGetPriceWithVatDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPriceWithVatDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


