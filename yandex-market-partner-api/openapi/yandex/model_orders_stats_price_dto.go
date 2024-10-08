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

// checks if the OrdersStatsPriceDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrdersStatsPriceDTO{}

// OrdersStatsPriceDTO Цена или скидки на товар.
type OrdersStatsPriceDTO struct {
	Type *OrdersStatsPriceType `json:"type,omitempty"`
	// Цена или скидка на единицу товара в заказе. Указана в рублях. Точность — два знака после запятой. 
	CostPerItem *float32 `json:"costPerItem,omitempty"`
	// Суммарная цена или скидка на все единицы товара в заказе. Указана в рублях. Точность — два знака после запятой. 
	Total *float32 `json:"total,omitempty"`
}

// NewOrdersStatsPriceDTO instantiates a new OrdersStatsPriceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrdersStatsPriceDTO() *OrdersStatsPriceDTO {
	this := OrdersStatsPriceDTO{}
	return &this
}

// NewOrdersStatsPriceDTOWithDefaults instantiates a new OrdersStatsPriceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrdersStatsPriceDTOWithDefaults() *OrdersStatsPriceDTO {
	this := OrdersStatsPriceDTO{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrdersStatsPriceDTO) GetType() OrdersStatsPriceType {
	if o == nil || IsNil(o.Type) {
		var ret OrdersStatsPriceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsPriceDTO) GetTypeOk() (*OrdersStatsPriceType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrdersStatsPriceDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OrdersStatsPriceType and assigns it to the Type field.
func (o *OrdersStatsPriceDTO) SetType(v OrdersStatsPriceType) {
	o.Type = &v
}

// GetCostPerItem returns the CostPerItem field value if set, zero value otherwise.
func (o *OrdersStatsPriceDTO) GetCostPerItem() float32 {
	if o == nil || IsNil(o.CostPerItem) {
		var ret float32
		return ret
	}
	return *o.CostPerItem
}

// GetCostPerItemOk returns a tuple with the CostPerItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsPriceDTO) GetCostPerItemOk() (*float32, bool) {
	if o == nil || IsNil(o.CostPerItem) {
		return nil, false
	}
	return o.CostPerItem, true
}

// HasCostPerItem returns a boolean if a field has been set.
func (o *OrdersStatsPriceDTO) HasCostPerItem() bool {
	if o != nil && !IsNil(o.CostPerItem) {
		return true
	}

	return false
}

// SetCostPerItem gets a reference to the given float32 and assigns it to the CostPerItem field.
func (o *OrdersStatsPriceDTO) SetCostPerItem(v float32) {
	o.CostPerItem = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *OrdersStatsPriceDTO) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrdersStatsPriceDTO) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *OrdersStatsPriceDTO) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *OrdersStatsPriceDTO) SetTotal(v float32) {
	o.Total = &v
}

func (o OrdersStatsPriceDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrdersStatsPriceDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.CostPerItem) {
		toSerialize["costPerItem"] = o.CostPerItem
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableOrdersStatsPriceDTO struct {
	value *OrdersStatsPriceDTO
	isSet bool
}

func (v NullableOrdersStatsPriceDTO) Get() *OrdersStatsPriceDTO {
	return v.value
}

func (v *NullableOrdersStatsPriceDTO) Set(val *OrdersStatsPriceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersStatsPriceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersStatsPriceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersStatsPriceDTO(val *OrdersStatsPriceDTO) *NullableOrdersStatsPriceDTO {
	return &NullableOrdersStatsPriceDTO{value: val, isSet: true}
}

func (v NullableOrdersStatsPriceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersStatsPriceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


