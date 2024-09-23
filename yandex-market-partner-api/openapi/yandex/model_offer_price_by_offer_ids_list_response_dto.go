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

// checks if the OfferPriceByOfferIdsListResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferPriceByOfferIdsListResponseDTO{}

// OfferPriceByOfferIdsListResponseDTO Список цен.
type OfferPriceByOfferIdsListResponseDTO struct {
	// Страница списка цен.
	Offers []OfferPriceByOfferIdsResponseDTO `json:"offers"`
	Paging *ScrollingPagerDTO `json:"paging,omitempty"`
}

type _OfferPriceByOfferIdsListResponseDTO OfferPriceByOfferIdsListResponseDTO

// NewOfferPriceByOfferIdsListResponseDTO instantiates a new OfferPriceByOfferIdsListResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferPriceByOfferIdsListResponseDTO(offers []OfferPriceByOfferIdsResponseDTO) *OfferPriceByOfferIdsListResponseDTO {
	this := OfferPriceByOfferIdsListResponseDTO{}
	this.Offers = offers
	return &this
}

// NewOfferPriceByOfferIdsListResponseDTOWithDefaults instantiates a new OfferPriceByOfferIdsListResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferPriceByOfferIdsListResponseDTOWithDefaults() *OfferPriceByOfferIdsListResponseDTO {
	this := OfferPriceByOfferIdsListResponseDTO{}
	return &this
}

// GetOffers returns the Offers field value
func (o *OfferPriceByOfferIdsListResponseDTO) GetOffers() []OfferPriceByOfferIdsResponseDTO {
	if o == nil {
		var ret []OfferPriceByOfferIdsResponseDTO
		return ret
	}

	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value
// and a boolean to check if the value has been set.
func (o *OfferPriceByOfferIdsListResponseDTO) GetOffersOk() ([]OfferPriceByOfferIdsResponseDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offers, true
}

// SetOffers sets field value
func (o *OfferPriceByOfferIdsListResponseDTO) SetOffers(v []OfferPriceByOfferIdsResponseDTO) {
	o.Offers = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *OfferPriceByOfferIdsListResponseDTO) GetPaging() ScrollingPagerDTO {
	if o == nil || IsNil(o.Paging) {
		var ret ScrollingPagerDTO
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferPriceByOfferIdsListResponseDTO) GetPagingOk() (*ScrollingPagerDTO, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *OfferPriceByOfferIdsListResponseDTO) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ScrollingPagerDTO and assigns it to the Paging field.
func (o *OfferPriceByOfferIdsListResponseDTO) SetPaging(v ScrollingPagerDTO) {
	o.Paging = &v
}

func (o OfferPriceByOfferIdsListResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferPriceByOfferIdsListResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offers"] = o.Offers
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

func (o *OfferPriceByOfferIdsListResponseDTO) UnmarshalJSON(data []byte) (err error) {
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

	varOfferPriceByOfferIdsListResponseDTO := _OfferPriceByOfferIdsListResponseDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOfferPriceByOfferIdsListResponseDTO)

	if err != nil {
		return err
	}

	*o = OfferPriceByOfferIdsListResponseDTO(varOfferPriceByOfferIdsListResponseDTO)

	return err
}

type NullableOfferPriceByOfferIdsListResponseDTO struct {
	value *OfferPriceByOfferIdsListResponseDTO
	isSet bool
}

func (v NullableOfferPriceByOfferIdsListResponseDTO) Get() *OfferPriceByOfferIdsListResponseDTO {
	return v.value
}

func (v *NullableOfferPriceByOfferIdsListResponseDTO) Set(val *OfferPriceByOfferIdsListResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferPriceByOfferIdsListResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferPriceByOfferIdsListResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferPriceByOfferIdsListResponseDTO(val *OfferPriceByOfferIdsListResponseDTO) *NullableOfferPriceByOfferIdsListResponseDTO {
	return &NullableOfferPriceByOfferIdsListResponseDTO{value: val, isSet: true}
}

func (v NullableOfferPriceByOfferIdsListResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferPriceByOfferIdsListResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


