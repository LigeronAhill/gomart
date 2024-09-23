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

// checks if the OfferSellingProgramDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferSellingProgramDTO{}

// OfferSellingProgramDTO Информация о том, по каким моделям можно продавать товар, а по каким нельзя.
type OfferSellingProgramDTO struct {
	SellingProgram SellingProgramType `json:"sellingProgram"`
	Status OfferSellingProgramStatusType `json:"status"`
}

type _OfferSellingProgramDTO OfferSellingProgramDTO

// NewOfferSellingProgramDTO instantiates a new OfferSellingProgramDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferSellingProgramDTO(sellingProgram SellingProgramType, status OfferSellingProgramStatusType) *OfferSellingProgramDTO {
	this := OfferSellingProgramDTO{}
	this.SellingProgram = sellingProgram
	this.Status = status
	return &this
}

// NewOfferSellingProgramDTOWithDefaults instantiates a new OfferSellingProgramDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferSellingProgramDTOWithDefaults() *OfferSellingProgramDTO {
	this := OfferSellingProgramDTO{}
	return &this
}

// GetSellingProgram returns the SellingProgram field value
func (o *OfferSellingProgramDTO) GetSellingProgram() SellingProgramType {
	if o == nil {
		var ret SellingProgramType
		return ret
	}

	return o.SellingProgram
}

// GetSellingProgramOk returns a tuple with the SellingProgram field value
// and a boolean to check if the value has been set.
func (o *OfferSellingProgramDTO) GetSellingProgramOk() (*SellingProgramType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SellingProgram, true
}

// SetSellingProgram sets field value
func (o *OfferSellingProgramDTO) SetSellingProgram(v SellingProgramType) {
	o.SellingProgram = v
}

// GetStatus returns the Status field value
func (o *OfferSellingProgramDTO) GetStatus() OfferSellingProgramStatusType {
	if o == nil {
		var ret OfferSellingProgramStatusType
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OfferSellingProgramDTO) GetStatusOk() (*OfferSellingProgramStatusType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OfferSellingProgramDTO) SetStatus(v OfferSellingProgramStatusType) {
	o.Status = v
}

func (o OfferSellingProgramDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferSellingProgramDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sellingProgram"] = o.SellingProgram
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *OfferSellingProgramDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sellingProgram",
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

	varOfferSellingProgramDTO := _OfferSellingProgramDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOfferSellingProgramDTO)

	if err != nil {
		return err
	}

	*o = OfferSellingProgramDTO(varOfferSellingProgramDTO)

	return err
}

type NullableOfferSellingProgramDTO struct {
	value *OfferSellingProgramDTO
	isSet bool
}

func (v NullableOfferSellingProgramDTO) Get() *OfferSellingProgramDTO {
	return v.value
}

func (v *NullableOfferSellingProgramDTO) Set(val *OfferSellingProgramDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferSellingProgramDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferSellingProgramDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferSellingProgramDTO(val *OfferSellingProgramDTO) *NullableOfferSellingProgramDTO {
	return &NullableOfferSellingProgramDTO{value: val, isSet: true}
}

func (v NullableOfferSellingProgramDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferSellingProgramDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


