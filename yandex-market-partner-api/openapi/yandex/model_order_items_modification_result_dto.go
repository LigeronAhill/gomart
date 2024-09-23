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

// checks if the OrderItemsModificationResultDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderItemsModificationResultDTO{}

// OrderItemsModificationResultDTO Краткие сведения о промаркированных товарах. Параметр возвращается, если ответ `OK`. 
type OrderItemsModificationResultDTO struct {
	// Список позиций в заказе, подлежащих маркировке.
	Items []BriefOrderItemDTO `json:"items"`
}

type _OrderItemsModificationResultDTO OrderItemsModificationResultDTO

// NewOrderItemsModificationResultDTO instantiates a new OrderItemsModificationResultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderItemsModificationResultDTO(items []BriefOrderItemDTO) *OrderItemsModificationResultDTO {
	this := OrderItemsModificationResultDTO{}
	this.Items = items
	return &this
}

// NewOrderItemsModificationResultDTOWithDefaults instantiates a new OrderItemsModificationResultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderItemsModificationResultDTOWithDefaults() *OrderItemsModificationResultDTO {
	this := OrderItemsModificationResultDTO{}
	return &this
}

// GetItems returns the Items field value
func (o *OrderItemsModificationResultDTO) GetItems() []BriefOrderItemDTO {
	if o == nil {
		var ret []BriefOrderItemDTO
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *OrderItemsModificationResultDTO) GetItemsOk() ([]BriefOrderItemDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *OrderItemsModificationResultDTO) SetItems(v []BriefOrderItemDTO) {
	o.Items = v
}

func (o OrderItemsModificationResultDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderItemsModificationResultDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *OrderItemsModificationResultDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varOrderItemsModificationResultDTO := _OrderItemsModificationResultDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderItemsModificationResultDTO)

	if err != nil {
		return err
	}

	*o = OrderItemsModificationResultDTO(varOrderItemsModificationResultDTO)

	return err
}

type NullableOrderItemsModificationResultDTO struct {
	value *OrderItemsModificationResultDTO
	isSet bool
}

func (v NullableOrderItemsModificationResultDTO) Get() *OrderItemsModificationResultDTO {
	return v.value
}

func (v *NullableOrderItemsModificationResultDTO) Set(val *OrderItemsModificationResultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderItemsModificationResultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderItemsModificationResultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderItemsModificationResultDTO(val *OrderItemsModificationResultDTO) *NullableOrderItemsModificationResultDTO {
	return &NullableOrderItemsModificationResultDTO{value: val, isSet: true}
}

func (v NullableOrderItemsModificationResultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderItemsModificationResultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


