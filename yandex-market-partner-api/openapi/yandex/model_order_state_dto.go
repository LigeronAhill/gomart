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

// checks if the OrderStateDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderStateDTO{}

// OrderStateDTO Информация по заказу.
type OrderStateDTO struct {
	// Идентификатор заказа.
	Id int64 `json:"id"`
	Status OrderStatusType `json:"status"`
	Substatus *OrderSubstatusType `json:"substatus,omitempty"`
}

type _OrderStateDTO OrderStateDTO

// NewOrderStateDTO instantiates a new OrderStateDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStateDTO(id int64, status OrderStatusType) *OrderStateDTO {
	this := OrderStateDTO{}
	this.Id = id
	this.Status = status
	return &this
}

// NewOrderStateDTOWithDefaults instantiates a new OrderStateDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStateDTOWithDefaults() *OrderStateDTO {
	this := OrderStateDTO{}
	return &this
}

// GetId returns the Id field value
func (o *OrderStateDTO) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderStateDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderStateDTO) SetId(v int64) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *OrderStateDTO) GetStatus() OrderStatusType {
	if o == nil {
		var ret OrderStatusType
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrderStateDTO) GetStatusOk() (*OrderStatusType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrderStateDTO) SetStatus(v OrderStatusType) {
	o.Status = v
}

// GetSubstatus returns the Substatus field value if set, zero value otherwise.
func (o *OrderStateDTO) GetSubstatus() OrderSubstatusType {
	if o == nil || IsNil(o.Substatus) {
		var ret OrderSubstatusType
		return ret
	}
	return *o.Substatus
}

// GetSubstatusOk returns a tuple with the Substatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStateDTO) GetSubstatusOk() (*OrderSubstatusType, bool) {
	if o == nil || IsNil(o.Substatus) {
		return nil, false
	}
	return o.Substatus, true
}

// HasSubstatus returns a boolean if a field has been set.
func (o *OrderStateDTO) HasSubstatus() bool {
	if o != nil && !IsNil(o.Substatus) {
		return true
	}

	return false
}

// SetSubstatus gets a reference to the given OrderSubstatusType and assigns it to the Substatus field.
func (o *OrderStateDTO) SetSubstatus(v OrderSubstatusType) {
	o.Substatus = &v
}

func (o OrderStateDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderStateDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	if !IsNil(o.Substatus) {
		toSerialize["substatus"] = o.Substatus
	}
	return toSerialize, nil
}

func (o *OrderStateDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
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

	varOrderStateDTO := _OrderStateDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderStateDTO)

	if err != nil {
		return err
	}

	*o = OrderStateDTO(varOrderStateDTO)

	return err
}

type NullableOrderStateDTO struct {
	value *OrderStateDTO
	isSet bool
}

func (v NullableOrderStateDTO) Get() *OrderStateDTO {
	return v.value
}

func (v *NullableOrderStateDTO) Set(val *OrderStateDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStateDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStateDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStateDTO(val *OrderStateDTO) *NullableOrderStateDTO {
	return &NullableOrderStateDTO{value: val, isSet: true}
}

func (v NullableOrderStateDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStateDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


