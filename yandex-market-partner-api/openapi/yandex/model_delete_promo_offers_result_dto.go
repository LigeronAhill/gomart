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

// checks if the DeletePromoOffersResultDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletePromoOffersResultDTO{}

// DeletePromoOffersResultDTO Результат удаления товаров из акции.  Возвращается, только если в запросе был передан параметр `offerIds`. 
type DeletePromoOffersResultDTO struct {
	// Товары, при удалении которых появились ошибки.  Возвращается, только если есть такие товары. 
	RejectedOffers []RejectedPromoOfferDeleteDTO `json:"rejectedOffers,omitempty"`
}

// NewDeletePromoOffersResultDTO instantiates a new DeletePromoOffersResultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePromoOffersResultDTO() *DeletePromoOffersResultDTO {
	this := DeletePromoOffersResultDTO{}
	return &this
}

// NewDeletePromoOffersResultDTOWithDefaults instantiates a new DeletePromoOffersResultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePromoOffersResultDTOWithDefaults() *DeletePromoOffersResultDTO {
	this := DeletePromoOffersResultDTO{}
	return &this
}

// GetRejectedOffers returns the RejectedOffers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeletePromoOffersResultDTO) GetRejectedOffers() []RejectedPromoOfferDeleteDTO {
	if o == nil {
		var ret []RejectedPromoOfferDeleteDTO
		return ret
	}
	return o.RejectedOffers
}

// GetRejectedOffersOk returns a tuple with the RejectedOffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeletePromoOffersResultDTO) GetRejectedOffersOk() ([]RejectedPromoOfferDeleteDTO, bool) {
	if o == nil || IsNil(o.RejectedOffers) {
		return nil, false
	}
	return o.RejectedOffers, true
}

// HasRejectedOffers returns a boolean if a field has been set.
func (o *DeletePromoOffersResultDTO) HasRejectedOffers() bool {
	if o != nil && !IsNil(o.RejectedOffers) {
		return true
	}

	return false
}

// SetRejectedOffers gets a reference to the given []RejectedPromoOfferDeleteDTO and assigns it to the RejectedOffers field.
func (o *DeletePromoOffersResultDTO) SetRejectedOffers(v []RejectedPromoOfferDeleteDTO) {
	o.RejectedOffers = v
}

func (o DeletePromoOffersResultDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletePromoOffersResultDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RejectedOffers != nil {
		toSerialize["rejectedOffers"] = o.RejectedOffers
	}
	return toSerialize, nil
}

type NullableDeletePromoOffersResultDTO struct {
	value *DeletePromoOffersResultDTO
	isSet bool
}

func (v NullableDeletePromoOffersResultDTO) Get() *DeletePromoOffersResultDTO {
	return v.value
}

func (v *NullableDeletePromoOffersResultDTO) Set(val *DeletePromoOffersResultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePromoOffersResultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePromoOffersResultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePromoOffersResultDTO(val *DeletePromoOffersResultDTO) *NullableDeletePromoOffersResultDTO {
	return &NullableDeletePromoOffersResultDTO{value: val, isSet: true}
}

func (v NullableDeletePromoOffersResultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePromoOffersResultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


