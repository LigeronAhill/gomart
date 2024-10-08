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

// checks if the UpdateCampaignOffersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCampaignOffersRequest{}

// UpdateCampaignOffersRequest Запрос на обновление предложений товаров магазина.
type UpdateCampaignOffersRequest struct {
	// Параметры размещения товаров в заданном магазине.
	Offers []UpdateCampaignOfferDTO `json:"offers"`
}

type _UpdateCampaignOffersRequest UpdateCampaignOffersRequest

// NewUpdateCampaignOffersRequest instantiates a new UpdateCampaignOffersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCampaignOffersRequest(offers []UpdateCampaignOfferDTO) *UpdateCampaignOffersRequest {
	this := UpdateCampaignOffersRequest{}
	this.Offers = offers
	return &this
}

// NewUpdateCampaignOffersRequestWithDefaults instantiates a new UpdateCampaignOffersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCampaignOffersRequestWithDefaults() *UpdateCampaignOffersRequest {
	this := UpdateCampaignOffersRequest{}
	return &this
}

// GetOffers returns the Offers field value
func (o *UpdateCampaignOffersRequest) GetOffers() []UpdateCampaignOfferDTO {
	if o == nil {
		var ret []UpdateCampaignOfferDTO
		return ret
	}

	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value
// and a boolean to check if the value has been set.
func (o *UpdateCampaignOffersRequest) GetOffersOk() ([]UpdateCampaignOfferDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offers, true
}

// SetOffers sets field value
func (o *UpdateCampaignOffersRequest) SetOffers(v []UpdateCampaignOfferDTO) {
	o.Offers = v
}

func (o UpdateCampaignOffersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCampaignOffersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offers"] = o.Offers
	return toSerialize, nil
}

func (o *UpdateCampaignOffersRequest) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateCampaignOffersRequest := _UpdateCampaignOffersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateCampaignOffersRequest)

	if err != nil {
		return err
	}

	*o = UpdateCampaignOffersRequest(varUpdateCampaignOffersRequest)

	return err
}

type NullableUpdateCampaignOffersRequest struct {
	value *UpdateCampaignOffersRequest
	isSet bool
}

func (v NullableUpdateCampaignOffersRequest) Get() *UpdateCampaignOffersRequest {
	return v.value
}

func (v *NullableUpdateCampaignOffersRequest) Set(val *UpdateCampaignOffersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCampaignOffersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCampaignOffersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCampaignOffersRequest(val *UpdateCampaignOffersRequest) *NullableUpdateCampaignOffersRequest {
	return &NullableUpdateCampaignOffersRequest{value: val, isSet: true}
}

func (v NullableUpdateCampaignOffersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCampaignOffersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


