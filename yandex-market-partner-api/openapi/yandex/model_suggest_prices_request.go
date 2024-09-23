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

// checks if the SuggestPricesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuggestPricesRequest{}

// SuggestPricesRequest Запрос на получение списка цен для продвижения.
type SuggestPricesRequest struct {
	// Список товаров.
	Offers []SuggestOfferPriceDTO `json:"offers"`
}

type _SuggestPricesRequest SuggestPricesRequest

// NewSuggestPricesRequest instantiates a new SuggestPricesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuggestPricesRequest(offers []SuggestOfferPriceDTO) *SuggestPricesRequest {
	this := SuggestPricesRequest{}
	this.Offers = offers
	return &this
}

// NewSuggestPricesRequestWithDefaults instantiates a new SuggestPricesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuggestPricesRequestWithDefaults() *SuggestPricesRequest {
	this := SuggestPricesRequest{}
	return &this
}

// GetOffers returns the Offers field value
func (o *SuggestPricesRequest) GetOffers() []SuggestOfferPriceDTO {
	if o == nil {
		var ret []SuggestOfferPriceDTO
		return ret
	}

	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value
// and a boolean to check if the value has been set.
func (o *SuggestPricesRequest) GetOffersOk() ([]SuggestOfferPriceDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offers, true
}

// SetOffers sets field value
func (o *SuggestPricesRequest) SetOffers(v []SuggestOfferPriceDTO) {
	o.Offers = v
}

func (o SuggestPricesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuggestPricesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offers"] = o.Offers
	return toSerialize, nil
}

func (o *SuggestPricesRequest) UnmarshalJSON(data []byte) (err error) {
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

	varSuggestPricesRequest := _SuggestPricesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSuggestPricesRequest)

	if err != nil {
		return err
	}

	*o = SuggestPricesRequest(varSuggestPricesRequest)

	return err
}

type NullableSuggestPricesRequest struct {
	value *SuggestPricesRequest
	isSet bool
}

func (v NullableSuggestPricesRequest) Get() *SuggestPricesRequest {
	return v.value
}

func (v *NullableSuggestPricesRequest) Set(val *SuggestPricesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSuggestPricesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSuggestPricesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuggestPricesRequest(val *SuggestPricesRequest) *NullableSuggestPricesRequest {
	return &NullableSuggestPricesRequest{value: val, isSet: true}
}

func (v NullableSuggestPricesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuggestPricesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


