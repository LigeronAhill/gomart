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

// checks if the DeleteCampaignOffersDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteCampaignOffersDTO{}

// DeleteCampaignOffersDTO Список товаров, которые не удалось удалить, потому что они не найдены или хранятся на складе Маркета.
type DeleteCampaignOffersDTO struct {
	// Список SKU.
	NotDeletedOfferIds []string `json:"notDeletedOfferIds,omitempty"`
}

// NewDeleteCampaignOffersDTO instantiates a new DeleteCampaignOffersDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCampaignOffersDTO() *DeleteCampaignOffersDTO {
	this := DeleteCampaignOffersDTO{}
	return &this
}

// NewDeleteCampaignOffersDTOWithDefaults instantiates a new DeleteCampaignOffersDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCampaignOffersDTOWithDefaults() *DeleteCampaignOffersDTO {
	this := DeleteCampaignOffersDTO{}
	return &this
}

// GetNotDeletedOfferIds returns the NotDeletedOfferIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteCampaignOffersDTO) GetNotDeletedOfferIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NotDeletedOfferIds
}

// GetNotDeletedOfferIdsOk returns a tuple with the NotDeletedOfferIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteCampaignOffersDTO) GetNotDeletedOfferIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.NotDeletedOfferIds) {
		return nil, false
	}
	return o.NotDeletedOfferIds, true
}

// HasNotDeletedOfferIds returns a boolean if a field has been set.
func (o *DeleteCampaignOffersDTO) HasNotDeletedOfferIds() bool {
	if o != nil && !IsNil(o.NotDeletedOfferIds) {
		return true
	}

	return false
}

// SetNotDeletedOfferIds gets a reference to the given []string and assigns it to the NotDeletedOfferIds field.
func (o *DeleteCampaignOffersDTO) SetNotDeletedOfferIds(v []string) {
	o.NotDeletedOfferIds = v
}

func (o DeleteCampaignOffersDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteCampaignOffersDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NotDeletedOfferIds != nil {
		toSerialize["notDeletedOfferIds"] = o.NotDeletedOfferIds
	}
	return toSerialize, nil
}

type NullableDeleteCampaignOffersDTO struct {
	value *DeleteCampaignOffersDTO
	isSet bool
}

func (v NullableDeleteCampaignOffersDTO) Get() *DeleteCampaignOffersDTO {
	return v.value
}

func (v *NullableDeleteCampaignOffersDTO) Set(val *DeleteCampaignOffersDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCampaignOffersDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCampaignOffersDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCampaignOffersDTO(val *DeleteCampaignOffersDTO) *NullableDeleteCampaignOffersDTO {
	return &NullableDeleteCampaignOffersDTO{value: val, isSet: true}
}

func (v NullableDeleteCampaignOffersDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCampaignOffersDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


