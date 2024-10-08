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

// checks if the CampaignSettingsTimePeriodDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignSettingsTimePeriodDTO{}

// CampaignSettingsTimePeriodDTO Период, за который рассчитывается итоговый список нерабочих дней службы доставки.
type CampaignSettingsTimePeriodDTO struct {
	// Формат даты: `ДД-ММ-ГГГГ`. 
	FromDate *string `json:"fromDate,omitempty"`
	// Формат даты: `ДД-ММ-ГГГГ`. 
	ToDate *string `json:"toDate,omitempty"`
}

// NewCampaignSettingsTimePeriodDTO instantiates a new CampaignSettingsTimePeriodDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignSettingsTimePeriodDTO() *CampaignSettingsTimePeriodDTO {
	this := CampaignSettingsTimePeriodDTO{}
	return &this
}

// NewCampaignSettingsTimePeriodDTOWithDefaults instantiates a new CampaignSettingsTimePeriodDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignSettingsTimePeriodDTOWithDefaults() *CampaignSettingsTimePeriodDTO {
	this := CampaignSettingsTimePeriodDTO{}
	return &this
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *CampaignSettingsTimePeriodDTO) GetFromDate() string {
	if o == nil || IsNil(o.FromDate) {
		var ret string
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettingsTimePeriodDTO) GetFromDateOk() (*string, bool) {
	if o == nil || IsNil(o.FromDate) {
		return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *CampaignSettingsTimePeriodDTO) HasFromDate() bool {
	if o != nil && !IsNil(o.FromDate) {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given string and assigns it to the FromDate field.
func (o *CampaignSettingsTimePeriodDTO) SetFromDate(v string) {
	o.FromDate = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *CampaignSettingsTimePeriodDTO) GetToDate() string {
	if o == nil || IsNil(o.ToDate) {
		var ret string
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignSettingsTimePeriodDTO) GetToDateOk() (*string, bool) {
	if o == nil || IsNil(o.ToDate) {
		return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *CampaignSettingsTimePeriodDTO) HasToDate() bool {
	if o != nil && !IsNil(o.ToDate) {
		return true
	}

	return false
}

// SetToDate gets a reference to the given string and assigns it to the ToDate field.
func (o *CampaignSettingsTimePeriodDTO) SetToDate(v string) {
	o.ToDate = &v
}

func (o CampaignSettingsTimePeriodDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignSettingsTimePeriodDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromDate) {
		toSerialize["fromDate"] = o.FromDate
	}
	if !IsNil(o.ToDate) {
		toSerialize["toDate"] = o.ToDate
	}
	return toSerialize, nil
}

type NullableCampaignSettingsTimePeriodDTO struct {
	value *CampaignSettingsTimePeriodDTO
	isSet bool
}

func (v NullableCampaignSettingsTimePeriodDTO) Get() *CampaignSettingsTimePeriodDTO {
	return v.value
}

func (v *NullableCampaignSettingsTimePeriodDTO) Set(val *CampaignSettingsTimePeriodDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignSettingsTimePeriodDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignSettingsTimePeriodDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignSettingsTimePeriodDTO(val *CampaignSettingsTimePeriodDTO) *NullableCampaignSettingsTimePeriodDTO {
	return &NullableCampaignSettingsTimePeriodDTO{value: val, isSet: true}
}

func (v NullableCampaignSettingsTimePeriodDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignSettingsTimePeriodDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


