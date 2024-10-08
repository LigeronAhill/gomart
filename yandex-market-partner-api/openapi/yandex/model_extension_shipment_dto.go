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

// checks if the ExtensionShipmentDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionShipmentDTO{}

// ExtensionShipmentDTO Информация об отгрузке.
type ExtensionShipmentDTO struct {
	CurrentStatus *ShipmentStatusChangeDTO `json:"currentStatus,omitempty"`
	// Доступные действия над отгрузкой.
	AvailableActions []ShipmentActionType `json:"availableActions"`
}

type _ExtensionShipmentDTO ExtensionShipmentDTO

// NewExtensionShipmentDTO instantiates a new ExtensionShipmentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionShipmentDTO(availableActions []ShipmentActionType) *ExtensionShipmentDTO {
	this := ExtensionShipmentDTO{}
	this.AvailableActions = availableActions
	return &this
}

// NewExtensionShipmentDTOWithDefaults instantiates a new ExtensionShipmentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionShipmentDTOWithDefaults() *ExtensionShipmentDTO {
	this := ExtensionShipmentDTO{}
	return &this
}

// GetCurrentStatus returns the CurrentStatus field value if set, zero value otherwise.
func (o *ExtensionShipmentDTO) GetCurrentStatus() ShipmentStatusChangeDTO {
	if o == nil || IsNil(o.CurrentStatus) {
		var ret ShipmentStatusChangeDTO
		return ret
	}
	return *o.CurrentStatus
}

// GetCurrentStatusOk returns a tuple with the CurrentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionShipmentDTO) GetCurrentStatusOk() (*ShipmentStatusChangeDTO, bool) {
	if o == nil || IsNil(o.CurrentStatus) {
		return nil, false
	}
	return o.CurrentStatus, true
}

// HasCurrentStatus returns a boolean if a field has been set.
func (o *ExtensionShipmentDTO) HasCurrentStatus() bool {
	if o != nil && !IsNil(o.CurrentStatus) {
		return true
	}

	return false
}

// SetCurrentStatus gets a reference to the given ShipmentStatusChangeDTO and assigns it to the CurrentStatus field.
func (o *ExtensionShipmentDTO) SetCurrentStatus(v ShipmentStatusChangeDTO) {
	o.CurrentStatus = &v
}

// GetAvailableActions returns the AvailableActions field value
func (o *ExtensionShipmentDTO) GetAvailableActions() []ShipmentActionType {
	if o == nil {
		var ret []ShipmentActionType
		return ret
	}

	return o.AvailableActions
}

// GetAvailableActionsOk returns a tuple with the AvailableActions field value
// and a boolean to check if the value has been set.
func (o *ExtensionShipmentDTO) GetAvailableActionsOk() ([]ShipmentActionType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableActions, true
}

// SetAvailableActions sets field value
func (o *ExtensionShipmentDTO) SetAvailableActions(v []ShipmentActionType) {
	o.AvailableActions = v
}

func (o ExtensionShipmentDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionShipmentDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentStatus) {
		toSerialize["currentStatus"] = o.CurrentStatus
	}
	toSerialize["availableActions"] = o.AvailableActions
	return toSerialize, nil
}

func (o *ExtensionShipmentDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"availableActions",
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

	varExtensionShipmentDTO := _ExtensionShipmentDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExtensionShipmentDTO)

	if err != nil {
		return err
	}

	*o = ExtensionShipmentDTO(varExtensionShipmentDTO)

	return err
}

type NullableExtensionShipmentDTO struct {
	value *ExtensionShipmentDTO
	isSet bool
}

func (v NullableExtensionShipmentDTO) Get() *ExtensionShipmentDTO {
	return v.value
}

func (v *NullableExtensionShipmentDTO) Set(val *ExtensionShipmentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionShipmentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionShipmentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionShipmentDTO(val *ExtensionShipmentDTO) *NullableExtensionShipmentDTO {
	return &NullableExtensionShipmentDTO{value: val, isSet: true}
}

func (v NullableExtensionShipmentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionShipmentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


