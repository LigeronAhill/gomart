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

// checks if the OfferRecommendationsResultDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferRecommendationsResultDTO{}

// OfferRecommendationsResultDTO Список товаров с рекомендациями.
type OfferRecommendationsResultDTO struct {
	Paging *ScrollingPagerDTO `json:"paging,omitempty"`
	// Страница списка товаров.
	OfferRecommendations []OfferRecommendationDTO `json:"offerRecommendations"`
}

type _OfferRecommendationsResultDTO OfferRecommendationsResultDTO

// NewOfferRecommendationsResultDTO instantiates a new OfferRecommendationsResultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferRecommendationsResultDTO(offerRecommendations []OfferRecommendationDTO) *OfferRecommendationsResultDTO {
	this := OfferRecommendationsResultDTO{}
	this.OfferRecommendations = offerRecommendations
	return &this
}

// NewOfferRecommendationsResultDTOWithDefaults instantiates a new OfferRecommendationsResultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferRecommendationsResultDTOWithDefaults() *OfferRecommendationsResultDTO {
	this := OfferRecommendationsResultDTO{}
	return &this
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *OfferRecommendationsResultDTO) GetPaging() ScrollingPagerDTO {
	if o == nil || IsNil(o.Paging) {
		var ret ScrollingPagerDTO
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferRecommendationsResultDTO) GetPagingOk() (*ScrollingPagerDTO, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *OfferRecommendationsResultDTO) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ScrollingPagerDTO and assigns it to the Paging field.
func (o *OfferRecommendationsResultDTO) SetPaging(v ScrollingPagerDTO) {
	o.Paging = &v
}

// GetOfferRecommendations returns the OfferRecommendations field value
func (o *OfferRecommendationsResultDTO) GetOfferRecommendations() []OfferRecommendationDTO {
	if o == nil {
		var ret []OfferRecommendationDTO
		return ret
	}

	return o.OfferRecommendations
}

// GetOfferRecommendationsOk returns a tuple with the OfferRecommendations field value
// and a boolean to check if the value has been set.
func (o *OfferRecommendationsResultDTO) GetOfferRecommendationsOk() ([]OfferRecommendationDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfferRecommendations, true
}

// SetOfferRecommendations sets field value
func (o *OfferRecommendationsResultDTO) SetOfferRecommendations(v []OfferRecommendationDTO) {
	o.OfferRecommendations = v
}

func (o OfferRecommendationsResultDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferRecommendationsResultDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	toSerialize["offerRecommendations"] = o.OfferRecommendations
	return toSerialize, nil
}

func (o *OfferRecommendationsResultDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offerRecommendations",
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

	varOfferRecommendationsResultDTO := _OfferRecommendationsResultDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOfferRecommendationsResultDTO)

	if err != nil {
		return err
	}

	*o = OfferRecommendationsResultDTO(varOfferRecommendationsResultDTO)

	return err
}

type NullableOfferRecommendationsResultDTO struct {
	value *OfferRecommendationsResultDTO
	isSet bool
}

func (v NullableOfferRecommendationsResultDTO) Get() *OfferRecommendationsResultDTO {
	return v.value
}

func (v *NullableOfferRecommendationsResultDTO) Set(val *OfferRecommendationsResultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferRecommendationsResultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferRecommendationsResultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferRecommendationsResultDTO(val *OfferRecommendationsResultDTO) *NullableOfferRecommendationsResultDTO {
	return &NullableOfferRecommendationsResultDTO{value: val, isSet: true}
}

func (v NullableOfferRecommendationsResultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferRecommendationsResultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


