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

// checks if the UpdateStockDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStockDTO{}

// UpdateStockDTO Информация об остатках одного товара на одном из складов.
type UpdateStockDTO struct {
	// Ваш SKU — идентификатор товара в вашей системе.  Разрешена любая последовательность длиной до 255 знаков.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * SKU товара нельзя менять — можно только удалить товар и добавить заново с новым SKU.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields) 
	Sku string `json:"sku" validate:"regexp=^[^\\\\x00-\\\\x08\\\\x0A-\\\\x1f\\\\x7f]{1,255}$"`
	// Информация об остатках товара. 
	Items []UpdateStockItemDTO `json:"items"`
}

type _UpdateStockDTO UpdateStockDTO

// NewUpdateStockDTO instantiates a new UpdateStockDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStockDTO(sku string, items []UpdateStockItemDTO) *UpdateStockDTO {
	this := UpdateStockDTO{}
	this.Sku = sku
	this.Items = items
	return &this
}

// NewUpdateStockDTOWithDefaults instantiates a new UpdateStockDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStockDTOWithDefaults() *UpdateStockDTO {
	this := UpdateStockDTO{}
	return &this
}

// GetSku returns the Sku field value
func (o *UpdateStockDTO) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *UpdateStockDTO) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *UpdateStockDTO) SetSku(v string) {
	o.Sku = v
}

// GetItems returns the Items field value
func (o *UpdateStockDTO) GetItems() []UpdateStockItemDTO {
	if o == nil {
		var ret []UpdateStockItemDTO
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *UpdateStockDTO) GetItemsOk() ([]UpdateStockItemDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *UpdateStockDTO) SetItems(v []UpdateStockItemDTO) {
	o.Items = v
}

func (o UpdateStockDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStockDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sku"] = o.Sku
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *UpdateStockDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sku",
		"items",
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

	varUpdateStockDTO := _UpdateStockDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateStockDTO)

	if err != nil {
		return err
	}

	*o = UpdateStockDTO(varUpdateStockDTO)

	return err
}

type NullableUpdateStockDTO struct {
	value *UpdateStockDTO
	isSet bool
}

func (v NullableUpdateStockDTO) Get() *UpdateStockDTO {
	return v.value
}

func (v *NullableUpdateStockDTO) Set(val *UpdateStockDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStockDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStockDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStockDTO(val *UpdateStockDTO) *NullableUpdateStockDTO {
	return &NullableUpdateStockDTO{value: val, isSet: true}
}

func (v NullableUpdateStockDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStockDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


