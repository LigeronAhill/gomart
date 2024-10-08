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

// checks if the GetSuggestedOfferMappingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSuggestedOfferMappingsRequest{}

// GetSuggestedOfferMappingsRequest struct for GetSuggestedOfferMappingsRequest
type GetSuggestedOfferMappingsRequest struct {
	// Список товаров.
	Offers []SuggestedOfferDTO `json:"offers,omitempty"`
}

// NewGetSuggestedOfferMappingsRequest instantiates a new GetSuggestedOfferMappingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSuggestedOfferMappingsRequest() *GetSuggestedOfferMappingsRequest {
	this := GetSuggestedOfferMappingsRequest{}
	return &this
}

// NewGetSuggestedOfferMappingsRequestWithDefaults instantiates a new GetSuggestedOfferMappingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSuggestedOfferMappingsRequestWithDefaults() *GetSuggestedOfferMappingsRequest {
	this := GetSuggestedOfferMappingsRequest{}
	return &this
}

// GetOffers returns the Offers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetSuggestedOfferMappingsRequest) GetOffers() []SuggestedOfferDTO {
	if o == nil {
		var ret []SuggestedOfferDTO
		return ret
	}
	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetSuggestedOfferMappingsRequest) GetOffersOk() ([]SuggestedOfferDTO, bool) {
	if o == nil || IsNil(o.Offers) {
		return nil, false
	}
	return o.Offers, true
}

// HasOffers returns a boolean if a field has been set.
func (o *GetSuggestedOfferMappingsRequest) HasOffers() bool {
	if o != nil && !IsNil(o.Offers) {
		return true
	}

	return false
}

// SetOffers gets a reference to the given []SuggestedOfferDTO and assigns it to the Offers field.
func (o *GetSuggestedOfferMappingsRequest) SetOffers(v []SuggestedOfferDTO) {
	o.Offers = v
}

func (o GetSuggestedOfferMappingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSuggestedOfferMappingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Offers != nil {
		toSerialize["offers"] = o.Offers
	}
	return toSerialize, nil
}

type NullableGetSuggestedOfferMappingsRequest struct {
	value *GetSuggestedOfferMappingsRequest
	isSet bool
}

func (v NullableGetSuggestedOfferMappingsRequest) Get() *GetSuggestedOfferMappingsRequest {
	return v.value
}

func (v *NullableGetSuggestedOfferMappingsRequest) Set(val *GetSuggestedOfferMappingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSuggestedOfferMappingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSuggestedOfferMappingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSuggestedOfferMappingsRequest(val *GetSuggestedOfferMappingsRequest) *NullableGetSuggestedOfferMappingsRequest {
	return &NullableGetSuggestedOfferMappingsRequest{value: val, isSet: true}
}

func (v NullableGetSuggestedOfferMappingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSuggestedOfferMappingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


