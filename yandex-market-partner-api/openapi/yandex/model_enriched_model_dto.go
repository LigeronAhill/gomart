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

// checks if the EnrichedModelDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrichedModelDTO{}

// EnrichedModelDTO Модель товара.
type EnrichedModelDTO struct {
	// Идентификатор модели товара.
	Id *int64 `json:"id,omitempty"`
	// Название модели товара.
	Name *string `json:"name,omitempty"`
	Prices *ModelPriceDTO `json:"prices,omitempty"`
	// Список первых десяти предложений, расположенных на карточке модели.  В ответе на запрос возвращаются предложения различных магазинов. Если есть несколько предложений от одного магазина, в ответе отображается только одно, наиболее релевантное из них. 
	Offers []ModelOfferDTO `json:"offers,omitempty"`
	// Суммарное количество предложений в розничных магазинах в регионе. Учитываются все предложения от каждого магазина.
	OfflineOffers *int32 `json:"offlineOffers,omitempty"`
	// Суммарное количество предложений в интернет-магазинах в регионе. Учитываются все предложения от каждого магазина.
	OnlineOffers *int32 `json:"onlineOffers,omitempty"`
}

// NewEnrichedModelDTO instantiates a new EnrichedModelDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrichedModelDTO() *EnrichedModelDTO {
	this := EnrichedModelDTO{}
	return &this
}

// NewEnrichedModelDTOWithDefaults instantiates a new EnrichedModelDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrichedModelDTOWithDefaults() *EnrichedModelDTO {
	this := EnrichedModelDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnrichedModelDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrichedModelDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnrichedModelDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *EnrichedModelDTO) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnrichedModelDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrichedModelDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnrichedModelDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnrichedModelDTO) SetName(v string) {
	o.Name = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *EnrichedModelDTO) GetPrices() ModelPriceDTO {
	if o == nil || IsNil(o.Prices) {
		var ret ModelPriceDTO
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrichedModelDTO) GetPricesOk() (*ModelPriceDTO, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *EnrichedModelDTO) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given ModelPriceDTO and assigns it to the Prices field.
func (o *EnrichedModelDTO) SetPrices(v ModelPriceDTO) {
	o.Prices = &v
}

// GetOffers returns the Offers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EnrichedModelDTO) GetOffers() []ModelOfferDTO {
	if o == nil {
		var ret []ModelOfferDTO
		return ret
	}
	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EnrichedModelDTO) GetOffersOk() ([]ModelOfferDTO, bool) {
	if o == nil || IsNil(o.Offers) {
		return nil, false
	}
	return o.Offers, true
}

// HasOffers returns a boolean if a field has been set.
func (o *EnrichedModelDTO) HasOffers() bool {
	if o != nil && !IsNil(o.Offers) {
		return true
	}

	return false
}

// SetOffers gets a reference to the given []ModelOfferDTO and assigns it to the Offers field.
func (o *EnrichedModelDTO) SetOffers(v []ModelOfferDTO) {
	o.Offers = v
}

// GetOfflineOffers returns the OfflineOffers field value if set, zero value otherwise.
func (o *EnrichedModelDTO) GetOfflineOffers() int32 {
	if o == nil || IsNil(o.OfflineOffers) {
		var ret int32
		return ret
	}
	return *o.OfflineOffers
}

// GetOfflineOffersOk returns a tuple with the OfflineOffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrichedModelDTO) GetOfflineOffersOk() (*int32, bool) {
	if o == nil || IsNil(o.OfflineOffers) {
		return nil, false
	}
	return o.OfflineOffers, true
}

// HasOfflineOffers returns a boolean if a field has been set.
func (o *EnrichedModelDTO) HasOfflineOffers() bool {
	if o != nil && !IsNil(o.OfflineOffers) {
		return true
	}

	return false
}

// SetOfflineOffers gets a reference to the given int32 and assigns it to the OfflineOffers field.
func (o *EnrichedModelDTO) SetOfflineOffers(v int32) {
	o.OfflineOffers = &v
}

// GetOnlineOffers returns the OnlineOffers field value if set, zero value otherwise.
func (o *EnrichedModelDTO) GetOnlineOffers() int32 {
	if o == nil || IsNil(o.OnlineOffers) {
		var ret int32
		return ret
	}
	return *o.OnlineOffers
}

// GetOnlineOffersOk returns a tuple with the OnlineOffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrichedModelDTO) GetOnlineOffersOk() (*int32, bool) {
	if o == nil || IsNil(o.OnlineOffers) {
		return nil, false
	}
	return o.OnlineOffers, true
}

// HasOnlineOffers returns a boolean if a field has been set.
func (o *EnrichedModelDTO) HasOnlineOffers() bool {
	if o != nil && !IsNil(o.OnlineOffers) {
		return true
	}

	return false
}

// SetOnlineOffers gets a reference to the given int32 and assigns it to the OnlineOffers field.
func (o *EnrichedModelDTO) SetOnlineOffers(v int32) {
	o.OnlineOffers = &v
}

func (o EnrichedModelDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrichedModelDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	if o.Offers != nil {
		toSerialize["offers"] = o.Offers
	}
	if !IsNil(o.OfflineOffers) {
		toSerialize["offlineOffers"] = o.OfflineOffers
	}
	if !IsNil(o.OnlineOffers) {
		toSerialize["onlineOffers"] = o.OnlineOffers
	}
	return toSerialize, nil
}

type NullableEnrichedModelDTO struct {
	value *EnrichedModelDTO
	isSet bool
}

func (v NullableEnrichedModelDTO) Get() *EnrichedModelDTO {
	return v.value
}

func (v *NullableEnrichedModelDTO) Set(val *EnrichedModelDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrichedModelDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrichedModelDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrichedModelDTO(val *EnrichedModelDTO) *NullableEnrichedModelDTO {
	return &NullableEnrichedModelDTO{value: val, isSet: true}
}

func (v NullableEnrichedModelDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrichedModelDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


