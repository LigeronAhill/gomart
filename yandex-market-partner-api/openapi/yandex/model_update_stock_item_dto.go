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

// checks if the UpdateStockItemDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStockItemDTO{}

// UpdateStockItemDTO Информация об остатках товара. 
type UpdateStockItemDTO struct {
	// Количество доступного товара. 
	Count int64 `json:"count"`
	// Дата и время последнего обновления информации об остатках. <br><br> Если вы не передали параметр `updatedAt`, используется текущее время. <br><br> Формат даты и времени: ISO 8601 со смещением относительно UTC. Например, `2017-11-21T00:42:42+03:00`. 
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type _UpdateStockItemDTO UpdateStockItemDTO

// NewUpdateStockItemDTO instantiates a new UpdateStockItemDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStockItemDTO(count int64) *UpdateStockItemDTO {
	this := UpdateStockItemDTO{}
	this.Count = count
	return &this
}

// NewUpdateStockItemDTOWithDefaults instantiates a new UpdateStockItemDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStockItemDTOWithDefaults() *UpdateStockItemDTO {
	this := UpdateStockItemDTO{}
	return &this
}

// GetCount returns the Count field value
func (o *UpdateStockItemDTO) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *UpdateStockItemDTO) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *UpdateStockItemDTO) SetCount(v int64) {
	o.Count = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UpdateStockItemDTO) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStockItemDTO) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UpdateStockItemDTO) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *UpdateStockItemDTO) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o UpdateStockItemDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStockItemDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *UpdateStockItemDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
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

	varUpdateStockItemDTO := _UpdateStockItemDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateStockItemDTO)

	if err != nil {
		return err
	}

	*o = UpdateStockItemDTO(varUpdateStockItemDTO)

	return err
}

type NullableUpdateStockItemDTO struct {
	value *UpdateStockItemDTO
	isSet bool
}

func (v NullableUpdateStockItemDTO) Get() *UpdateStockItemDTO {
	return v.value
}

func (v *NullableUpdateStockItemDTO) Set(val *UpdateStockItemDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStockItemDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStockItemDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStockItemDTO(val *UpdateStockItemDTO) *NullableUpdateStockItemDTO {
	return &NullableUpdateStockItemDTO{value: val, isSet: true}
}

func (v NullableUpdateStockItemDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStockItemDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


