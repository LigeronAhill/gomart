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

// checks if the OrderDeliveryDateDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDeliveryDateDTO{}

// OrderDeliveryDateDTO Информация о новой дате доставки заказа.
type OrderDeliveryDateDTO struct {
	// Новая дата доставки заказа.  Формат даты: `ГГГГ-ММ-ДД`. 
	ToDate string `json:"toDate"`
}

type _OrderDeliveryDateDTO OrderDeliveryDateDTO

// NewOrderDeliveryDateDTO instantiates a new OrderDeliveryDateDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDeliveryDateDTO(toDate string) *OrderDeliveryDateDTO {
	this := OrderDeliveryDateDTO{}
	this.ToDate = toDate
	return &this
}

// NewOrderDeliveryDateDTOWithDefaults instantiates a new OrderDeliveryDateDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDeliveryDateDTOWithDefaults() *OrderDeliveryDateDTO {
	this := OrderDeliveryDateDTO{}
	return &this
}

// GetToDate returns the ToDate field value
func (o *OrderDeliveryDateDTO) GetToDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value
// and a boolean to check if the value has been set.
func (o *OrderDeliveryDateDTO) GetToDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToDate, true
}

// SetToDate sets field value
func (o *OrderDeliveryDateDTO) SetToDate(v string) {
	o.ToDate = v
}

func (o OrderDeliveryDateDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDeliveryDateDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["toDate"] = o.ToDate
	return toSerialize, nil
}

func (o *OrderDeliveryDateDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"toDate",
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

	varOrderDeliveryDateDTO := _OrderDeliveryDateDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderDeliveryDateDTO)

	if err != nil {
		return err
	}

	*o = OrderDeliveryDateDTO(varOrderDeliveryDateDTO)

	return err
}

type NullableOrderDeliveryDateDTO struct {
	value *OrderDeliveryDateDTO
	isSet bool
}

func (v NullableOrderDeliveryDateDTO) Get() *OrderDeliveryDateDTO {
	return v.value
}

func (v *NullableOrderDeliveryDateDTO) Set(val *OrderDeliveryDateDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDeliveryDateDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDeliveryDateDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDeliveryDateDTO(val *OrderDeliveryDateDTO) *NullableOrderDeliveryDateDTO {
	return &NullableOrderDeliveryDateDTO{value: val, isSet: true}
}

func (v NullableOrderDeliveryDateDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDeliveryDateDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


