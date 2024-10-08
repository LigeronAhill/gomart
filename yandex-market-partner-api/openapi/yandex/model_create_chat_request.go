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

// checks if the CreateChatRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateChatRequest{}

// CreateChatRequest Заказ, для которого нужно создать чат. 
type CreateChatRequest struct {
	// Идентификатор заказа на Маркете.
	OrderId int64 `json:"orderId"`
}

type _CreateChatRequest CreateChatRequest

// NewCreateChatRequest instantiates a new CreateChatRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateChatRequest(orderId int64) *CreateChatRequest {
	this := CreateChatRequest{}
	this.OrderId = orderId
	return &this
}

// NewCreateChatRequestWithDefaults instantiates a new CreateChatRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateChatRequestWithDefaults() *CreateChatRequest {
	this := CreateChatRequest{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *CreateChatRequest) GetOrderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *CreateChatRequest) GetOrderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *CreateChatRequest) SetOrderId(v int64) {
	o.OrderId = v
}

func (o CreateChatRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateChatRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["orderId"] = o.OrderId
	return toSerialize, nil
}

func (o *CreateChatRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orderId",
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

	varCreateChatRequest := _CreateChatRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateChatRequest)

	if err != nil {
		return err
	}

	*o = CreateChatRequest(varCreateChatRequest)

	return err
}

type NullableCreateChatRequest struct {
	value *CreateChatRequest
	isSet bool
}

func (v NullableCreateChatRequest) Get() *CreateChatRequest {
	return v.value
}

func (v *NullableCreateChatRequest) Set(val *CreateChatRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateChatRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateChatRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateChatRequest(val *CreateChatRequest) *NullableCreateChatRequest {
	return &NullableCreateChatRequest{value: val, isSet: true}
}

func (v NullableCreateChatRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateChatRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


