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

// checks if the OffersDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OffersDTO{}

// OffersDTO Найденные предложения магазина.
type OffersDTO struct {
	// Список предложений магазина.
	Offers []OfferDTO `json:"offers"`
}

type _OffersDTO OffersDTO

// NewOffersDTO instantiates a new OffersDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffersDTO(offers []OfferDTO) *OffersDTO {
	this := OffersDTO{}
	this.Offers = offers
	return &this
}

// NewOffersDTOWithDefaults instantiates a new OffersDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffersDTOWithDefaults() *OffersDTO {
	this := OffersDTO{}
	return &this
}

// GetOffers returns the Offers field value
func (o *OffersDTO) GetOffers() []OfferDTO {
	if o == nil {
		var ret []OfferDTO
		return ret
	}

	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value
// and a boolean to check if the value has been set.
func (o *OffersDTO) GetOffersOk() ([]OfferDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offers, true
}

// SetOffers sets field value
func (o *OffersDTO) SetOffers(v []OfferDTO) {
	o.Offers = v
}

func (o OffersDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OffersDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offers"] = o.Offers
	return toSerialize, nil
}

func (o *OffersDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offers",
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

	varOffersDTO := _OffersDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOffersDTO)

	if err != nil {
		return err
	}

	*o = OffersDTO(varOffersDTO)

	return err
}

type NullableOffersDTO struct {
	value *OffersDTO
	isSet bool
}

func (v NullableOffersDTO) Get() *OffersDTO {
	return v.value
}

func (v *NullableOffersDTO) Set(val *OffersDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOffersDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOffersDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffersDTO(val *OffersDTO) *NullableOffersDTO {
	return &NullableOffersDTO{value: val, isSet: true}
}

func (v NullableOffersDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffersDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


