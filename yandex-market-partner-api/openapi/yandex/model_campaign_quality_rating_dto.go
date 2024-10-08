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

// checks if the CampaignQualityRatingDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignQualityRatingDTO{}

// CampaignQualityRatingDTO Информация об индексе качества магазина.
type CampaignQualityRatingDTO struct {
	// Идентификатор магазина.
	CampaignId int64 `json:"campaignId"`
	// Список значений индекса качества.
	Ratings []QualityRatingDTO `json:"ratings"`
}

type _CampaignQualityRatingDTO CampaignQualityRatingDTO

// NewCampaignQualityRatingDTO instantiates a new CampaignQualityRatingDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignQualityRatingDTO(campaignId int64, ratings []QualityRatingDTO) *CampaignQualityRatingDTO {
	this := CampaignQualityRatingDTO{}
	this.CampaignId = campaignId
	this.Ratings = ratings
	return &this
}

// NewCampaignQualityRatingDTOWithDefaults instantiates a new CampaignQualityRatingDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignQualityRatingDTOWithDefaults() *CampaignQualityRatingDTO {
	this := CampaignQualityRatingDTO{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *CampaignQualityRatingDTO) GetCampaignId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *CampaignQualityRatingDTO) GetCampaignIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *CampaignQualityRatingDTO) SetCampaignId(v int64) {
	o.CampaignId = v
}

// GetRatings returns the Ratings field value
func (o *CampaignQualityRatingDTO) GetRatings() []QualityRatingDTO {
	if o == nil {
		var ret []QualityRatingDTO
		return ret
	}

	return o.Ratings
}

// GetRatingsOk returns a tuple with the Ratings field value
// and a boolean to check if the value has been set.
func (o *CampaignQualityRatingDTO) GetRatingsOk() ([]QualityRatingDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ratings, true
}

// SetRatings sets field value
func (o *CampaignQualityRatingDTO) SetRatings(v []QualityRatingDTO) {
	o.Ratings = v
}

func (o CampaignQualityRatingDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignQualityRatingDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["ratings"] = o.Ratings
	return toSerialize, nil
}

func (o *CampaignQualityRatingDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"campaignId",
		"ratings",
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

	varCampaignQualityRatingDTO := _CampaignQualityRatingDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCampaignQualityRatingDTO)

	if err != nil {
		return err
	}

	*o = CampaignQualityRatingDTO(varCampaignQualityRatingDTO)

	return err
}

type NullableCampaignQualityRatingDTO struct {
	value *CampaignQualityRatingDTO
	isSet bool
}

func (v NullableCampaignQualityRatingDTO) Get() *CampaignQualityRatingDTO {
	return v.value
}

func (v *NullableCampaignQualityRatingDTO) Set(val *CampaignQualityRatingDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignQualityRatingDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignQualityRatingDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignQualityRatingDTO(val *CampaignQualityRatingDTO) *NullableCampaignQualityRatingDTO {
	return &NullableCampaignQualityRatingDTO{value: val, isSet: true}
}

func (v NullableCampaignQualityRatingDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignQualityRatingDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


