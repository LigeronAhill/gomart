/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SkuBidItemDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkuBidItemDTO{}

// SkuBidItemDTO Список товаров и ставок на них.
type SkuBidItemDTO struct {
	// Ваш SKU — идентификатор товара в вашей системе.  Разрешена любая последовательность длиной до 255 знаков.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * SKU товара нельзя менять — можно только удалить товар и добавить заново с новым SKU.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields) 
	Sku string `json:"sku" validate:"regexp=^[^\\\\x00-\\\\x08\\\\x0A-\\\\x1f\\\\x7f]{1,255}$"`
	// Значение ставки.
	Bid int32 `json:"bid"`
}

type _SkuBidItemDTO SkuBidItemDTO

// NewSkuBidItemDTO instantiates a new SkuBidItemDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuBidItemDTO(sku string, bid int32) *SkuBidItemDTO {
	this := SkuBidItemDTO{}
	this.Sku = sku
	this.Bid = bid
	return &this
}

// NewSkuBidItemDTOWithDefaults instantiates a new SkuBidItemDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuBidItemDTOWithDefaults() *SkuBidItemDTO {
	this := SkuBidItemDTO{}
	return &this
}

// GetSku returns the Sku field value
func (o *SkuBidItemDTO) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *SkuBidItemDTO) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *SkuBidItemDTO) SetSku(v string) {
	o.Sku = v
}

// GetBid returns the Bid field value
func (o *SkuBidItemDTO) GetBid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bid
}

// GetBidOk returns a tuple with the Bid field value
// and a boolean to check if the value has been set.
func (o *SkuBidItemDTO) GetBidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bid, true
}

// SetBid sets field value
func (o *SkuBidItemDTO) SetBid(v int32) {
	o.Bid = v
}

func (o SkuBidItemDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkuBidItemDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sku"] = o.Sku
	toSerialize["bid"] = o.Bid
	return toSerialize, nil
}

func (o *SkuBidItemDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sku",
		"bid",
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

	varSkuBidItemDTO := _SkuBidItemDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkuBidItemDTO)

	if err != nil {
		return err
	}

	*o = SkuBidItemDTO(varSkuBidItemDTO)

	return err
}

type NullableSkuBidItemDTO struct {
	value *SkuBidItemDTO
	isSet bool
}

func (v NullableSkuBidItemDTO) Get() *SkuBidItemDTO {
	return v.value
}

func (v *NullableSkuBidItemDTO) Set(val *SkuBidItemDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuBidItemDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuBidItemDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuBidItemDTO(val *SkuBidItemDTO) *NullableSkuBidItemDTO {
	return &NullableSkuBidItemDTO{value: val, isSet: true}
}

func (v NullableSkuBidItemDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuBidItemDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


