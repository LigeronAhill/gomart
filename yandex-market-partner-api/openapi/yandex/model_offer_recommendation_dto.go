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

// checks if the OfferRecommendationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferRecommendationDTO{}

// OfferRecommendationDTO Информация о состоянии цен и рекомендации. 
type OfferRecommendationDTO struct {
	Offer *OfferForRecommendationDTO `json:"offer,omitempty"`
	Recommendation *OfferRecommendationInfoDTO `json:"recommendation,omitempty"`
}

// NewOfferRecommendationDTO instantiates a new OfferRecommendationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferRecommendationDTO() *OfferRecommendationDTO {
	this := OfferRecommendationDTO{}
	return &this
}

// NewOfferRecommendationDTOWithDefaults instantiates a new OfferRecommendationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferRecommendationDTOWithDefaults() *OfferRecommendationDTO {
	this := OfferRecommendationDTO{}
	return &this
}

// GetOffer returns the Offer field value if set, zero value otherwise.
func (o *OfferRecommendationDTO) GetOffer() OfferForRecommendationDTO {
	if o == nil || IsNil(o.Offer) {
		var ret OfferForRecommendationDTO
		return ret
	}
	return *o.Offer
}

// GetOfferOk returns a tuple with the Offer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferRecommendationDTO) GetOfferOk() (*OfferForRecommendationDTO, bool) {
	if o == nil || IsNil(o.Offer) {
		return nil, false
	}
	return o.Offer, true
}

// HasOffer returns a boolean if a field has been set.
func (o *OfferRecommendationDTO) HasOffer() bool {
	if o != nil && !IsNil(o.Offer) {
		return true
	}

	return false
}

// SetOffer gets a reference to the given OfferForRecommendationDTO and assigns it to the Offer field.
func (o *OfferRecommendationDTO) SetOffer(v OfferForRecommendationDTO) {
	o.Offer = &v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise.
func (o *OfferRecommendationDTO) GetRecommendation() OfferRecommendationInfoDTO {
	if o == nil || IsNil(o.Recommendation) {
		var ret OfferRecommendationInfoDTO
		return ret
	}
	return *o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferRecommendationDTO) GetRecommendationOk() (*OfferRecommendationInfoDTO, bool) {
	if o == nil || IsNil(o.Recommendation) {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *OfferRecommendationDTO) HasRecommendation() bool {
	if o != nil && !IsNil(o.Recommendation) {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given OfferRecommendationInfoDTO and assigns it to the Recommendation field.
func (o *OfferRecommendationDTO) SetRecommendation(v OfferRecommendationInfoDTO) {
	o.Recommendation = &v
}

func (o OfferRecommendationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferRecommendationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Offer) {
		toSerialize["offer"] = o.Offer
	}
	if !IsNil(o.Recommendation) {
		toSerialize["recommendation"] = o.Recommendation
	}
	return toSerialize, nil
}

type NullableOfferRecommendationDTO struct {
	value *OfferRecommendationDTO
	isSet bool
}

func (v NullableOfferRecommendationDTO) Get() *OfferRecommendationDTO {
	return v.value
}

func (v *NullableOfferRecommendationDTO) Set(val *OfferRecommendationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferRecommendationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferRecommendationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferRecommendationDTO(val *OfferRecommendationDTO) *NullableOfferRecommendationDTO {
	return &NullableOfferRecommendationDTO{value: val, isSet: true}
}

func (v NullableOfferRecommendationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferRecommendationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


