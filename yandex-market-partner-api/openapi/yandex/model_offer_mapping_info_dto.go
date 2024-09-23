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

// checks if the OfferMappingInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferMappingInfoDTO{}

// OfferMappingInfoDTO Информация о карточке товара.
type OfferMappingInfoDTO struct {
	Mapping *OfferMappingDTO `json:"mapping,omitempty"`
	AwaitingModerationMapping *OfferMappingDTO `json:"awaitingModerationMapping,omitempty"`
	RejectedMapping *OfferMappingDTO `json:"rejectedMapping,omitempty"`
}

// NewOfferMappingInfoDTO instantiates a new OfferMappingInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferMappingInfoDTO() *OfferMappingInfoDTO {
	this := OfferMappingInfoDTO{}
	return &this
}

// NewOfferMappingInfoDTOWithDefaults instantiates a new OfferMappingInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferMappingInfoDTOWithDefaults() *OfferMappingInfoDTO {
	this := OfferMappingInfoDTO{}
	return &this
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *OfferMappingInfoDTO) GetMapping() OfferMappingDTO {
	if o == nil || IsNil(o.Mapping) {
		var ret OfferMappingDTO
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferMappingInfoDTO) GetMappingOk() (*OfferMappingDTO, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *OfferMappingInfoDTO) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given OfferMappingDTO and assigns it to the Mapping field.
func (o *OfferMappingInfoDTO) SetMapping(v OfferMappingDTO) {
	o.Mapping = &v
}

// GetAwaitingModerationMapping returns the AwaitingModerationMapping field value if set, zero value otherwise.
func (o *OfferMappingInfoDTO) GetAwaitingModerationMapping() OfferMappingDTO {
	if o == nil || IsNil(o.AwaitingModerationMapping) {
		var ret OfferMappingDTO
		return ret
	}
	return *o.AwaitingModerationMapping
}

// GetAwaitingModerationMappingOk returns a tuple with the AwaitingModerationMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferMappingInfoDTO) GetAwaitingModerationMappingOk() (*OfferMappingDTO, bool) {
	if o == nil || IsNil(o.AwaitingModerationMapping) {
		return nil, false
	}
	return o.AwaitingModerationMapping, true
}

// HasAwaitingModerationMapping returns a boolean if a field has been set.
func (o *OfferMappingInfoDTO) HasAwaitingModerationMapping() bool {
	if o != nil && !IsNil(o.AwaitingModerationMapping) {
		return true
	}

	return false
}

// SetAwaitingModerationMapping gets a reference to the given OfferMappingDTO and assigns it to the AwaitingModerationMapping field.
func (o *OfferMappingInfoDTO) SetAwaitingModerationMapping(v OfferMappingDTO) {
	o.AwaitingModerationMapping = &v
}

// GetRejectedMapping returns the RejectedMapping field value if set, zero value otherwise.
func (o *OfferMappingInfoDTO) GetRejectedMapping() OfferMappingDTO {
	if o == nil || IsNil(o.RejectedMapping) {
		var ret OfferMappingDTO
		return ret
	}
	return *o.RejectedMapping
}

// GetRejectedMappingOk returns a tuple with the RejectedMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferMappingInfoDTO) GetRejectedMappingOk() (*OfferMappingDTO, bool) {
	if o == nil || IsNil(o.RejectedMapping) {
		return nil, false
	}
	return o.RejectedMapping, true
}

// HasRejectedMapping returns a boolean if a field has been set.
func (o *OfferMappingInfoDTO) HasRejectedMapping() bool {
	if o != nil && !IsNil(o.RejectedMapping) {
		return true
	}

	return false
}

// SetRejectedMapping gets a reference to the given OfferMappingDTO and assigns it to the RejectedMapping field.
func (o *OfferMappingInfoDTO) SetRejectedMapping(v OfferMappingDTO) {
	o.RejectedMapping = &v
}

func (o OfferMappingInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferMappingInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	if !IsNil(o.AwaitingModerationMapping) {
		toSerialize["awaitingModerationMapping"] = o.AwaitingModerationMapping
	}
	if !IsNil(o.RejectedMapping) {
		toSerialize["rejectedMapping"] = o.RejectedMapping
	}
	return toSerialize, nil
}

type NullableOfferMappingInfoDTO struct {
	value *OfferMappingInfoDTO
	isSet bool
}

func (v NullableOfferMappingInfoDTO) Get() *OfferMappingInfoDTO {
	return v.value
}

func (v *NullableOfferMappingInfoDTO) Set(val *OfferMappingInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferMappingInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferMappingInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferMappingInfoDTO(val *OfferMappingInfoDTO) *NullableOfferMappingInfoDTO {
	return &NullableOfferMappingInfoDTO{value: val, isSet: true}
}

func (v NullableOfferMappingInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferMappingInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


