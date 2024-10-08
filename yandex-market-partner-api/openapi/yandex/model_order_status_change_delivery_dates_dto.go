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

// checks if the OrderStatusChangeDeliveryDatesDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderStatusChangeDeliveryDatesDTO{}

// OrderStatusChangeDeliveryDatesDTO Диапазон дат доставки.
type OrderStatusChangeDeliveryDatesDTO struct {
	// **Только для модели DBS**  Фактическая дата доставки. <br><br> Когда передавать параметр `realDeliveryDate`:  * Не передавайте параметр, если:   * переводите заказ в любой статус, кроме `PICKUP` или `DELIVERED`;   * меняете статус заказа на `PICKUP` или `DELIVERED` в день доставки — будет указана дата выполнения запроса. * Передавайте дату доставки, если переводите заказ в статус `PICKUP` или `DELIVERED` не в день доставки. Нельзя указывать дату доставки в будущем.    {% note warning \"Индекс качества\" %}    Передача статуса после установленного срока снижает индекс качества. О сроках читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/quality/tech#dbs).    {% endnote %}      
	RealDeliveryDate *string `json:"realDeliveryDate,omitempty"`
}

// NewOrderStatusChangeDeliveryDatesDTO instantiates a new OrderStatusChangeDeliveryDatesDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStatusChangeDeliveryDatesDTO() *OrderStatusChangeDeliveryDatesDTO {
	this := OrderStatusChangeDeliveryDatesDTO{}
	return &this
}

// NewOrderStatusChangeDeliveryDatesDTOWithDefaults instantiates a new OrderStatusChangeDeliveryDatesDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStatusChangeDeliveryDatesDTOWithDefaults() *OrderStatusChangeDeliveryDatesDTO {
	this := OrderStatusChangeDeliveryDatesDTO{}
	return &this
}

// GetRealDeliveryDate returns the RealDeliveryDate field value if set, zero value otherwise.
func (o *OrderStatusChangeDeliveryDatesDTO) GetRealDeliveryDate() string {
	if o == nil || IsNil(o.RealDeliveryDate) {
		var ret string
		return ret
	}
	return *o.RealDeliveryDate
}

// GetRealDeliveryDateOk returns a tuple with the RealDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusChangeDeliveryDatesDTO) GetRealDeliveryDateOk() (*string, bool) {
	if o == nil || IsNil(o.RealDeliveryDate) {
		return nil, false
	}
	return o.RealDeliveryDate, true
}

// HasRealDeliveryDate returns a boolean if a field has been set.
func (o *OrderStatusChangeDeliveryDatesDTO) HasRealDeliveryDate() bool {
	if o != nil && !IsNil(o.RealDeliveryDate) {
		return true
	}

	return false
}

// SetRealDeliveryDate gets a reference to the given string and assigns it to the RealDeliveryDate field.
func (o *OrderStatusChangeDeliveryDatesDTO) SetRealDeliveryDate(v string) {
	o.RealDeliveryDate = &v
}

func (o OrderStatusChangeDeliveryDatesDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderStatusChangeDeliveryDatesDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RealDeliveryDate) {
		toSerialize["realDeliveryDate"] = o.RealDeliveryDate
	}
	return toSerialize, nil
}

type NullableOrderStatusChangeDeliveryDatesDTO struct {
	value *OrderStatusChangeDeliveryDatesDTO
	isSet bool
}

func (v NullableOrderStatusChangeDeliveryDatesDTO) Get() *OrderStatusChangeDeliveryDatesDTO {
	return v.value
}

func (v *NullableOrderStatusChangeDeliveryDatesDTO) Set(val *OrderStatusChangeDeliveryDatesDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatusChangeDeliveryDatesDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatusChangeDeliveryDatesDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatusChangeDeliveryDatesDTO(val *OrderStatusChangeDeliveryDatesDTO) *NullableOrderStatusChangeDeliveryDatesDTO {
	return &NullableOrderStatusChangeDeliveryDatesDTO{value: val, isSet: true}
}

func (v NullableOrderStatusChangeDeliveryDatesDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatusChangeDeliveryDatesDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


