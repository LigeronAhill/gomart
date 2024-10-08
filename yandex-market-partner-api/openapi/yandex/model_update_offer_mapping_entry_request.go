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

// checks if the UpdateOfferMappingEntryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOfferMappingEntryRequest{}

// UpdateOfferMappingEntryRequest Запрос на обновление товаров.
type UpdateOfferMappingEntryRequest struct {
	// Информация о товарах в каталоге.
	OfferMappingEntries []UpdateOfferMappingEntryDTO `json:"offerMappingEntries"`
}

type _UpdateOfferMappingEntryRequest UpdateOfferMappingEntryRequest

// NewUpdateOfferMappingEntryRequest instantiates a new UpdateOfferMappingEntryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOfferMappingEntryRequest(offerMappingEntries []UpdateOfferMappingEntryDTO) *UpdateOfferMappingEntryRequest {
	this := UpdateOfferMappingEntryRequest{}
	this.OfferMappingEntries = offerMappingEntries
	return &this
}

// NewUpdateOfferMappingEntryRequestWithDefaults instantiates a new UpdateOfferMappingEntryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOfferMappingEntryRequestWithDefaults() *UpdateOfferMappingEntryRequest {
	this := UpdateOfferMappingEntryRequest{}
	return &this
}

// GetOfferMappingEntries returns the OfferMappingEntries field value
func (o *UpdateOfferMappingEntryRequest) GetOfferMappingEntries() []UpdateOfferMappingEntryDTO {
	if o == nil {
		var ret []UpdateOfferMappingEntryDTO
		return ret
	}

	return o.OfferMappingEntries
}

// GetOfferMappingEntriesOk returns a tuple with the OfferMappingEntries field value
// and a boolean to check if the value has been set.
func (o *UpdateOfferMappingEntryRequest) GetOfferMappingEntriesOk() ([]UpdateOfferMappingEntryDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfferMappingEntries, true
}

// SetOfferMappingEntries sets field value
func (o *UpdateOfferMappingEntryRequest) SetOfferMappingEntries(v []UpdateOfferMappingEntryDTO) {
	o.OfferMappingEntries = v
}

func (o UpdateOfferMappingEntryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOfferMappingEntryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offerMappingEntries"] = o.OfferMappingEntries
	return toSerialize, nil
}

func (o *UpdateOfferMappingEntryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offerMappingEntries",
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

	varUpdateOfferMappingEntryRequest := _UpdateOfferMappingEntryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateOfferMappingEntryRequest)

	if err != nil {
		return err
	}

	*o = UpdateOfferMappingEntryRequest(varUpdateOfferMappingEntryRequest)

	return err
}

type NullableUpdateOfferMappingEntryRequest struct {
	value *UpdateOfferMappingEntryRequest
	isSet bool
}

func (v NullableUpdateOfferMappingEntryRequest) Get() *UpdateOfferMappingEntryRequest {
	return v.value
}

func (v *NullableUpdateOfferMappingEntryRequest) Set(val *UpdateOfferMappingEntryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOfferMappingEntryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOfferMappingEntryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOfferMappingEntryRequest(val *UpdateOfferMappingEntryRequest) *NullableUpdateOfferMappingEntryRequest {
	return &NullableUpdateOfferMappingEntryRequest{value: val, isSet: true}
}

func (v NullableUpdateOfferMappingEntryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOfferMappingEntryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


