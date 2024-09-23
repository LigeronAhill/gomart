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

// checks if the FeedbackFactorDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedbackFactorDTO{}

// FeedbackFactorDTO Информация об оценках по параметрам, указанных в отзыве.  При создании отзыва автору предлагается поставить оценки магазину по нескольким параметрам: например, за скорость обработки заказа или удобство самовывоза. Набор параметров зависит от того, какой способ покупки (параметр `delivery`) указал автор. 
type FeedbackFactorDTO struct {
	// Идентификатор параметра.
	Id *int64 `json:"id,omitempty"`
	// Название параметра. Например, `Скорость обработки заказа`.
	Title *string `json:"title,omitempty"`
	// Описание параметра. Например, `Как быстро с вами связались для подтверждения заказа?`.
	Description *string `json:"description,omitempty"`
	// Оценка по параметру, указанная в отзыве: от `1` (низшая оценка) до `5` (высшая оценка). 
	Value *int32 `json:"value,omitempty"`
}

// NewFeedbackFactorDTO instantiates a new FeedbackFactorDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackFactorDTO() *FeedbackFactorDTO {
	this := FeedbackFactorDTO{}
	return &this
}

// NewFeedbackFactorDTOWithDefaults instantiates a new FeedbackFactorDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackFactorDTOWithDefaults() *FeedbackFactorDTO {
	this := FeedbackFactorDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FeedbackFactorDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFactorDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FeedbackFactorDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *FeedbackFactorDTO) SetId(v int64) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *FeedbackFactorDTO) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFactorDTO) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *FeedbackFactorDTO) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *FeedbackFactorDTO) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FeedbackFactorDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFactorDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FeedbackFactorDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FeedbackFactorDTO) SetDescription(v string) {
	o.Description = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FeedbackFactorDTO) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFactorDTO) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FeedbackFactorDTO) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *FeedbackFactorDTO) SetValue(v int32) {
	o.Value = &v
}

func (o FeedbackFactorDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedbackFactorDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFeedbackFactorDTO struct {
	value *FeedbackFactorDTO
	isSet bool
}

func (v NullableFeedbackFactorDTO) Get() *FeedbackFactorDTO {
	return v.value
}

func (v *NullableFeedbackFactorDTO) Set(val *FeedbackFactorDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackFactorDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackFactorDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackFactorDTO(val *FeedbackFactorDTO) *NullableFeedbackFactorDTO {
	return &NullableFeedbackFactorDTO{value: val, isSet: true}
}

func (v NullableFeedbackFactorDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackFactorDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


