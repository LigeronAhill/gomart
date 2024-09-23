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

// checks if the PartnerShipmentWarehouseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartnerShipmentWarehouseDTO{}

// PartnerShipmentWarehouseDTO Данные о складе отправления.
type PartnerShipmentWarehouseDTO struct {
	// Идентификатор склада отправления.
	Id *int64 `json:"id,omitempty"`
	// Наименование склада отправления.
	Name *string `json:"name,omitempty"`
	// Адрес склада отправления.
	Address *string `json:"address,omitempty"`
}

// NewPartnerShipmentWarehouseDTO instantiates a new PartnerShipmentWarehouseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartnerShipmentWarehouseDTO() *PartnerShipmentWarehouseDTO {
	this := PartnerShipmentWarehouseDTO{}
	return &this
}

// NewPartnerShipmentWarehouseDTOWithDefaults instantiates a new PartnerShipmentWarehouseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartnerShipmentWarehouseDTOWithDefaults() *PartnerShipmentWarehouseDTO {
	this := PartnerShipmentWarehouseDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PartnerShipmentWarehouseDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerShipmentWarehouseDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PartnerShipmentWarehouseDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PartnerShipmentWarehouseDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartnerShipmentWarehouseDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerShipmentWarehouseDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartnerShipmentWarehouseDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartnerShipmentWarehouseDTO) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PartnerShipmentWarehouseDTO) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartnerShipmentWarehouseDTO) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PartnerShipmentWarehouseDTO) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *PartnerShipmentWarehouseDTO) SetAddress(v string) {
	o.Address = &v
}

func (o PartnerShipmentWarehouseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartnerShipmentWarehouseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullablePartnerShipmentWarehouseDTO struct {
	value *PartnerShipmentWarehouseDTO
	isSet bool
}

func (v NullablePartnerShipmentWarehouseDTO) Get() *PartnerShipmentWarehouseDTO {
	return v.value
}

func (v *NullablePartnerShipmentWarehouseDTO) Set(val *PartnerShipmentWarehouseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerShipmentWarehouseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerShipmentWarehouseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerShipmentWarehouseDTO(val *PartnerShipmentWarehouseDTO) *NullablePartnerShipmentWarehouseDTO {
	return &NullablePartnerShipmentWarehouseDTO{value: val, isSet: true}
}

func (v NullablePartnerShipmentWarehouseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerShipmentWarehouseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


