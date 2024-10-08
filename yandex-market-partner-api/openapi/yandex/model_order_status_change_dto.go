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

// checks if the OrderStatusChangeDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderStatusChangeDTO{}

// OrderStatusChangeDTO Заказ.
type OrderStatusChangeDTO struct {
	Status OrderStatusType `json:"status"`
	Substatus *OrderSubstatusType `json:"substatus,omitempty"`
	Delivery *OrderStatusChangeDeliveryDTO `json:"delivery,omitempty"`
}

type _OrderStatusChangeDTO OrderStatusChangeDTO

// NewOrderStatusChangeDTO instantiates a new OrderStatusChangeDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStatusChangeDTO(status OrderStatusType) *OrderStatusChangeDTO {
	this := OrderStatusChangeDTO{}
	this.Status = status
	return &this
}

// NewOrderStatusChangeDTOWithDefaults instantiates a new OrderStatusChangeDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStatusChangeDTOWithDefaults() *OrderStatusChangeDTO {
	this := OrderStatusChangeDTO{}
	return &this
}

// GetStatus returns the Status field value
func (o *OrderStatusChangeDTO) GetStatus() OrderStatusType {
	if o == nil {
		var ret OrderStatusType
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrderStatusChangeDTO) GetStatusOk() (*OrderStatusType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrderStatusChangeDTO) SetStatus(v OrderStatusType) {
	o.Status = v
}

// GetSubstatus returns the Substatus field value if set, zero value otherwise.
func (o *OrderStatusChangeDTO) GetSubstatus() OrderSubstatusType {
	if o == nil || IsNil(o.Substatus) {
		var ret OrderSubstatusType
		return ret
	}
	return *o.Substatus
}

// GetSubstatusOk returns a tuple with the Substatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusChangeDTO) GetSubstatusOk() (*OrderSubstatusType, bool) {
	if o == nil || IsNil(o.Substatus) {
		return nil, false
	}
	return o.Substatus, true
}

// HasSubstatus returns a boolean if a field has been set.
func (o *OrderStatusChangeDTO) HasSubstatus() bool {
	if o != nil && !IsNil(o.Substatus) {
		return true
	}

	return false
}

// SetSubstatus gets a reference to the given OrderSubstatusType and assigns it to the Substatus field.
func (o *OrderStatusChangeDTO) SetSubstatus(v OrderSubstatusType) {
	o.Substatus = &v
}

// GetDelivery returns the Delivery field value if set, zero value otherwise.
func (o *OrderStatusChangeDTO) GetDelivery() OrderStatusChangeDeliveryDTO {
	if o == nil || IsNil(o.Delivery) {
		var ret OrderStatusChangeDeliveryDTO
		return ret
	}
	return *o.Delivery
}

// GetDeliveryOk returns a tuple with the Delivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusChangeDTO) GetDeliveryOk() (*OrderStatusChangeDeliveryDTO, bool) {
	if o == nil || IsNil(o.Delivery) {
		return nil, false
	}
	return o.Delivery, true
}

// HasDelivery returns a boolean if a field has been set.
func (o *OrderStatusChangeDTO) HasDelivery() bool {
	if o != nil && !IsNil(o.Delivery) {
		return true
	}

	return false
}

// SetDelivery gets a reference to the given OrderStatusChangeDeliveryDTO and assigns it to the Delivery field.
func (o *OrderStatusChangeDTO) SetDelivery(v OrderStatusChangeDeliveryDTO) {
	o.Delivery = &v
}

func (o OrderStatusChangeDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderStatusChangeDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.Substatus) {
		toSerialize["substatus"] = o.Substatus
	}
	if !IsNil(o.Delivery) {
		toSerialize["delivery"] = o.Delivery
	}
	return toSerialize, nil
}

func (o *OrderStatusChangeDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varOrderStatusChangeDTO := _OrderStatusChangeDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderStatusChangeDTO)

	if err != nil {
		return err
	}

	*o = OrderStatusChangeDTO(varOrderStatusChangeDTO)

	return err
}

type NullableOrderStatusChangeDTO struct {
	value *OrderStatusChangeDTO
	isSet bool
}

func (v NullableOrderStatusChangeDTO) Get() *OrderStatusChangeDTO {
	return v.value
}

func (v *NullableOrderStatusChangeDTO) Set(val *OrderStatusChangeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatusChangeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatusChangeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatusChangeDTO(val *OrderStatusChangeDTO) *NullableOrderStatusChangeDTO {
	return &NullableOrderStatusChangeDTO{value: val, isSet: true}
}

func (v NullableOrderStatusChangeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatusChangeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


