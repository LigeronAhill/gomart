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

// checks if the GetPromoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPromoDTO{}

// GetPromoDTO Информация об акции.
type GetPromoDTO struct {
	// Идентификатор акции.
	Id string `json:"id"`
	// Название акции.
	Name string `json:"name"`
	Period PromoPeriodDTO `json:"period"`
	// Участвует или участвовал ли продавец в этой акции.  Для текущих и будущих акций возвращается со значением `true`, если в акции есть товары, которые были добавлены вручную. Если товары не участвуют в акции или добавлены в нее автоматически, параметр возвращается со значением `false`.  Для прошедших акций всегда возвращается со значением `true`.  Об автоматическом и ручном добавлении товаров в акцию читайте [в Справке Маркета для продавцов](https://yandex.ru/support2/marketplace/ru/marketing/promos/market/index). 
	Participating bool `json:"participating"`
	AssortmentInfo GetPromoAssortmentInfoDTO `json:"assortmentInfo"`
	MechanicsInfo GetPromoMechanicsInfoDTO `json:"mechanicsInfo"`
	BestsellerInfo GetPromoBestsellerInfoDTO `json:"bestsellerInfo"`
	// Список каналов продвижения товаров.
	Channels []ChannelType `json:"channels,omitempty"`
	Constraints *GetPromoConstraintsDTO `json:"constraints,omitempty"`
}

type _GetPromoDTO GetPromoDTO

// NewGetPromoDTO instantiates a new GetPromoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPromoDTO(id string, name string, period PromoPeriodDTO, participating bool, assortmentInfo GetPromoAssortmentInfoDTO, mechanicsInfo GetPromoMechanicsInfoDTO, bestsellerInfo GetPromoBestsellerInfoDTO) *GetPromoDTO {
	this := GetPromoDTO{}
	this.Id = id
	this.Name = name
	this.Period = period
	this.Participating = participating
	this.AssortmentInfo = assortmentInfo
	this.MechanicsInfo = mechanicsInfo
	this.BestsellerInfo = bestsellerInfo
	return &this
}

// NewGetPromoDTOWithDefaults instantiates a new GetPromoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPromoDTOWithDefaults() *GetPromoDTO {
	this := GetPromoDTO{}
	return &this
}

// GetId returns the Id field value
func (o *GetPromoDTO) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetPromoDTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetPromoDTO) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GetPromoDTO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetPromoDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetPromoDTO) SetName(v string) {
	o.Name = v
}

// GetPeriod returns the Period field value
func (o *GetPromoDTO) GetPeriod() PromoPeriodDTO {
	if o == nil {
		var ret PromoPeriodDTO
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *GetPromoDTO) GetPeriodOk() (*PromoPeriodDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *GetPromoDTO) SetPeriod(v PromoPeriodDTO) {
	o.Period = v
}

// GetParticipating returns the Participating field value
func (o *GetPromoDTO) GetParticipating() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Participating
}

// GetParticipatingOk returns a tuple with the Participating field value
// and a boolean to check if the value has been set.
func (o *GetPromoDTO) GetParticipatingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Participating, true
}

// SetParticipating sets field value
func (o *GetPromoDTO) SetParticipating(v bool) {
	o.Participating = v
}

// GetAssortmentInfo returns the AssortmentInfo field value
func (o *GetPromoDTO) GetAssortmentInfo() GetPromoAssortmentInfoDTO {
	if o == nil {
		var ret GetPromoAssortmentInfoDTO
		return ret
	}

	return o.AssortmentInfo
}

// GetAssortmentInfoOk returns a tuple with the AssortmentInfo field value
// and a boolean to check if the value has been set.
func (o *GetPromoDTO) GetAssortmentInfoOk() (*GetPromoAssortmentInfoDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssortmentInfo, true
}

// SetAssortmentInfo sets field value
func (o *GetPromoDTO) SetAssortmentInfo(v GetPromoAssortmentInfoDTO) {
	o.AssortmentInfo = v
}

// GetMechanicsInfo returns the MechanicsInfo field value
func (o *GetPromoDTO) GetMechanicsInfo() GetPromoMechanicsInfoDTO {
	if o == nil {
		var ret GetPromoMechanicsInfoDTO
		return ret
	}

	return o.MechanicsInfo
}

// GetMechanicsInfoOk returns a tuple with the MechanicsInfo field value
// and a boolean to check if the value has been set.
func (o *GetPromoDTO) GetMechanicsInfoOk() (*GetPromoMechanicsInfoDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MechanicsInfo, true
}

// SetMechanicsInfo sets field value
func (o *GetPromoDTO) SetMechanicsInfo(v GetPromoMechanicsInfoDTO) {
	o.MechanicsInfo = v
}

// GetBestsellerInfo returns the BestsellerInfo field value
func (o *GetPromoDTO) GetBestsellerInfo() GetPromoBestsellerInfoDTO {
	if o == nil {
		var ret GetPromoBestsellerInfoDTO
		return ret
	}

	return o.BestsellerInfo
}

// GetBestsellerInfoOk returns a tuple with the BestsellerInfo field value
// and a boolean to check if the value has been set.
func (o *GetPromoDTO) GetBestsellerInfoOk() (*GetPromoBestsellerInfoDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BestsellerInfo, true
}

// SetBestsellerInfo sets field value
func (o *GetPromoDTO) SetBestsellerInfo(v GetPromoBestsellerInfoDTO) {
	o.BestsellerInfo = v
}

// GetChannels returns the Channels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetPromoDTO) GetChannels() []ChannelType {
	if o == nil {
		var ret []ChannelType
		return ret
	}
	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetPromoDTO) GetChannelsOk() ([]ChannelType, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *GetPromoDTO) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given []ChannelType and assigns it to the Channels field.
func (o *GetPromoDTO) SetChannels(v []ChannelType) {
	o.Channels = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *GetPromoDTO) GetConstraints() GetPromoConstraintsDTO {
	if o == nil || IsNil(o.Constraints) {
		var ret GetPromoConstraintsDTO
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPromoDTO) GetConstraintsOk() (*GetPromoConstraintsDTO, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *GetPromoDTO) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given GetPromoConstraintsDTO and assigns it to the Constraints field.
func (o *GetPromoDTO) SetConstraints(v GetPromoConstraintsDTO) {
	o.Constraints = &v
}

func (o GetPromoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPromoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["period"] = o.Period
	toSerialize["participating"] = o.Participating
	toSerialize["assortmentInfo"] = o.AssortmentInfo
	toSerialize["mechanicsInfo"] = o.MechanicsInfo
	toSerialize["bestsellerInfo"] = o.BestsellerInfo
	if o.Channels != nil {
		toSerialize["channels"] = o.Channels
	}
	if !IsNil(o.Constraints) {
		toSerialize["constraints"] = o.Constraints
	}
	return toSerialize, nil
}

func (o *GetPromoDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"period",
		"participating",
		"assortmentInfo",
		"mechanicsInfo",
		"bestsellerInfo",
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

	varGetPromoDTO := _GetPromoDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPromoDTO)

	if err != nil {
		return err
	}

	*o = GetPromoDTO(varGetPromoDTO)

	return err
}

type NullableGetPromoDTO struct {
	value *GetPromoDTO
	isSet bool
}

func (v NullableGetPromoDTO) Get() *GetPromoDTO {
	return v.value
}

func (v *NullableGetPromoDTO) Set(val *GetPromoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPromoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPromoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPromoDTO(val *GetPromoDTO) *NullableGetPromoDTO {
	return &NullableGetPromoDTO{value: val, isSet: true}
}

func (v NullableGetPromoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPromoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


