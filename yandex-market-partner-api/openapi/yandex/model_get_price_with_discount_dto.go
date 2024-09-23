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

// checks if the GetPriceWithDiscountDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPriceWithDiscountDTO{}

// GetPriceWithDiscountDTO Цена с указанием скидки и времени последнего обновления.
type GetPriceWithDiscountDTO struct {
	// Значение.
	Value float32 `json:"value"`
	CurrencyId CurrencyType `json:"currencyId"`
	// Цена до скидки.  Число должно быть целым. Вы можете указать цену со скидкой от 5 до 99%.  Передавайте этот параметр при каждом обновлении цены, если предоставляете скидку на товар. 
	DiscountBase *float32 `json:"discountBase,omitempty"`
	// Время последнего обновления.
	UpdatedAt time.Time `json:"updatedAt"`
}

type _GetPriceWithDiscountDTO GetPriceWithDiscountDTO

// NewGetPriceWithDiscountDTO instantiates a new GetPriceWithDiscountDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPriceWithDiscountDTO(value float32, currencyId CurrencyType, updatedAt time.Time) *GetPriceWithDiscountDTO {
	this := GetPriceWithDiscountDTO{}
	this.Value = value
	this.CurrencyId = currencyId
	this.UpdatedAt = updatedAt
	return &this
}

// NewGetPriceWithDiscountDTOWithDefaults instantiates a new GetPriceWithDiscountDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPriceWithDiscountDTOWithDefaults() *GetPriceWithDiscountDTO {
	this := GetPriceWithDiscountDTO{}
	return &this
}

// GetValue returns the Value field value
func (o *GetPriceWithDiscountDTO) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *GetPriceWithDiscountDTO) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *GetPriceWithDiscountDTO) SetValue(v float32) {
	o.Value = v
}

// GetCurrencyId returns the CurrencyId field value
func (o *GetPriceWithDiscountDTO) GetCurrencyId() CurrencyType {
	if o == nil {
		var ret CurrencyType
		return ret
	}

	return o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value
// and a boolean to check if the value has been set.
func (o *GetPriceWithDiscountDTO) GetCurrencyIdOk() (*CurrencyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyId, true
}

// SetCurrencyId sets field value
func (o *GetPriceWithDiscountDTO) SetCurrencyId(v CurrencyType) {
	o.CurrencyId = v
}

// GetDiscountBase returns the DiscountBase field value if set, zero value otherwise.
func (o *GetPriceWithDiscountDTO) GetDiscountBase() float32 {
	if o == nil || IsNil(o.DiscountBase) {
		var ret float32
		return ret
	}
	return *o.DiscountBase
}

// GetDiscountBaseOk returns a tuple with the DiscountBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPriceWithDiscountDTO) GetDiscountBaseOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountBase) {
		return nil, false
	}
	return o.DiscountBase, true
}

// HasDiscountBase returns a boolean if a field has been set.
func (o *GetPriceWithDiscountDTO) HasDiscountBase() bool {
	if o != nil && !IsNil(o.DiscountBase) {
		return true
	}

	return false
}

// SetDiscountBase gets a reference to the given float32 and assigns it to the DiscountBase field.
func (o *GetPriceWithDiscountDTO) SetDiscountBase(v float32) {
	o.DiscountBase = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetPriceWithDiscountDTO) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetPriceWithDiscountDTO) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetPriceWithDiscountDTO) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o GetPriceWithDiscountDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPriceWithDiscountDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["currencyId"] = o.CurrencyId
	if !IsNil(o.DiscountBase) {
		toSerialize["discountBase"] = o.DiscountBase
	}
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *GetPriceWithDiscountDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"currencyId",
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

	varGetPriceWithDiscountDTO := _GetPriceWithDiscountDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPriceWithDiscountDTO)

	if err != nil {
		return err
	}

	*o = GetPriceWithDiscountDTO(varGetPriceWithDiscountDTO)

	return err
}

type NullableGetPriceWithDiscountDTO struct {
	value *GetPriceWithDiscountDTO
	isSet bool
}

func (v NullableGetPriceWithDiscountDTO) Get() *GetPriceWithDiscountDTO {
	return v.value
}

func (v *NullableGetPriceWithDiscountDTO) Set(val *GetPriceWithDiscountDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPriceWithDiscountDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPriceWithDiscountDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPriceWithDiscountDTO(val *GetPriceWithDiscountDTO) *NullableGetPriceWithDiscountDTO {
	return &NullableGetPriceWithDiscountDTO{value: val, isSet: true}
}

func (v NullableGetPriceWithDiscountDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPriceWithDiscountDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


