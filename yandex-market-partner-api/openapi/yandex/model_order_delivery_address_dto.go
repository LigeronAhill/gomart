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

// checks if the OrderDeliveryAddressDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDeliveryAddressDTO{}

// OrderDeliveryAddressDTO Адрес доставки.  Указывается, если `type=DELIVERY`, `type=POST` или `type=PICKUP` (адрес пункта выдачи). 
type OrderDeliveryAddressDTO struct {
	// Страна.  Обязательный параметр. 
	Country *string `json:"country,omitempty"`
	// Почтовый индекс.  Указывается, если выбрана доставка почтой (`delivery type=POST`). 
	Postcode *string `json:"postcode,omitempty"`
	// Город или населенный пункт.  Обязательный параметр. 
	City *string `json:"city,omitempty"`
	// Район.
	District *string `json:"district,omitempty"`
	// Станция метро.
	Subway *string `json:"subway,omitempty"`
	// Улица.  Обязательный параметр. 
	Street *string `json:"street,omitempty"`
	// Дом или владение.  Обязательный параметр. 
	House *string `json:"house,omitempty"`
	// Корпус или строение.
	Block *string `json:"block,omitempty"`
	// Подъезд.
	Entrance *string `json:"entrance,omitempty"`
	// Код домофона.
	Entryphone *string `json:"entryphone,omitempty"`
	// Этаж.
	Floor *string `json:"floor,omitempty"`
	// Квартира или офис.
	Apartment *string `json:"apartment,omitempty"`
	// Телефон получателя заказа.  Обязательный параметр. 
	Phone *string `json:"phone,omitempty"`
	// Фамилия, имя и отчество получателя заказа.  Обязательный параметр. 
	Recipient *string `json:"recipient,omitempty"`
	Gps *GpsDTO `json:"gps,omitempty"`
}

// NewOrderDeliveryAddressDTO instantiates a new OrderDeliveryAddressDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDeliveryAddressDTO() *OrderDeliveryAddressDTO {
	this := OrderDeliveryAddressDTO{}
	return &this
}

// NewOrderDeliveryAddressDTOWithDefaults instantiates a new OrderDeliveryAddressDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDeliveryAddressDTOWithDefaults() *OrderDeliveryAddressDTO {
	this := OrderDeliveryAddressDTO{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *OrderDeliveryAddressDTO) SetCountry(v string) {
	o.Country = &v
}

// GetPostcode returns the Postcode field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetPostcode() string {
	if o == nil || IsNil(o.Postcode) {
		var ret string
		return ret
	}
	return *o.Postcode
}

// GetPostcodeOk returns a tuple with the Postcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetPostcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Postcode) {
		return nil, false
	}
	return o.Postcode, true
}

// HasPostcode returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasPostcode() bool {
	if o != nil && !IsNil(o.Postcode) {
		return true
	}

	return false
}

// SetPostcode gets a reference to the given string and assigns it to the Postcode field.
func (o *OrderDeliveryAddressDTO) SetPostcode(v string) {
	o.Postcode = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OrderDeliveryAddressDTO) SetCity(v string) {
	o.City = &v
}

// GetDistrict returns the District field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetDistrict() string {
	if o == nil || IsNil(o.District) {
		var ret string
		return ret
	}
	return *o.District
}

// GetDistrictOk returns a tuple with the District field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetDistrictOk() (*string, bool) {
	if o == nil || IsNil(o.District) {
		return nil, false
	}
	return o.District, true
}

// HasDistrict returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasDistrict() bool {
	if o != nil && !IsNil(o.District) {
		return true
	}

	return false
}

// SetDistrict gets a reference to the given string and assigns it to the District field.
func (o *OrderDeliveryAddressDTO) SetDistrict(v string) {
	o.District = &v
}

// GetSubway returns the Subway field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetSubway() string {
	if o == nil || IsNil(o.Subway) {
		var ret string
		return ret
	}
	return *o.Subway
}

// GetSubwayOk returns a tuple with the Subway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetSubwayOk() (*string, bool) {
	if o == nil || IsNil(o.Subway) {
		return nil, false
	}
	return o.Subway, true
}

// HasSubway returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasSubway() bool {
	if o != nil && !IsNil(o.Subway) {
		return true
	}

	return false
}

// SetSubway gets a reference to the given string and assigns it to the Subway field.
func (o *OrderDeliveryAddressDTO) SetSubway(v string) {
	o.Subway = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetStreet() string {
	if o == nil || IsNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetStreetOk() (*string, bool) {
	if o == nil || IsNil(o.Street) {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasStreet() bool {
	if o != nil && !IsNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *OrderDeliveryAddressDTO) SetStreet(v string) {
	o.Street = &v
}

// GetHouse returns the House field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetHouse() string {
	if o == nil || IsNil(o.House) {
		var ret string
		return ret
	}
	return *o.House
}

// GetHouseOk returns a tuple with the House field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetHouseOk() (*string, bool) {
	if o == nil || IsNil(o.House) {
		return nil, false
	}
	return o.House, true
}

// HasHouse returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasHouse() bool {
	if o != nil && !IsNil(o.House) {
		return true
	}

	return false
}

// SetHouse gets a reference to the given string and assigns it to the House field.
func (o *OrderDeliveryAddressDTO) SetHouse(v string) {
	o.House = &v
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetBlock() string {
	if o == nil || IsNil(o.Block) {
		var ret string
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetBlockOk() (*string, bool) {
	if o == nil || IsNil(o.Block) {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasBlock() bool {
	if o != nil && !IsNil(o.Block) {
		return true
	}

	return false
}

// SetBlock gets a reference to the given string and assigns it to the Block field.
func (o *OrderDeliveryAddressDTO) SetBlock(v string) {
	o.Block = &v
}

// GetEntrance returns the Entrance field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetEntrance() string {
	if o == nil || IsNil(o.Entrance) {
		var ret string
		return ret
	}
	return *o.Entrance
}

// GetEntranceOk returns a tuple with the Entrance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetEntranceOk() (*string, bool) {
	if o == nil || IsNil(o.Entrance) {
		return nil, false
	}
	return o.Entrance, true
}

// HasEntrance returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasEntrance() bool {
	if o != nil && !IsNil(o.Entrance) {
		return true
	}

	return false
}

// SetEntrance gets a reference to the given string and assigns it to the Entrance field.
func (o *OrderDeliveryAddressDTO) SetEntrance(v string) {
	o.Entrance = &v
}

// GetEntryphone returns the Entryphone field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetEntryphone() string {
	if o == nil || IsNil(o.Entryphone) {
		var ret string
		return ret
	}
	return *o.Entryphone
}

// GetEntryphoneOk returns a tuple with the Entryphone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetEntryphoneOk() (*string, bool) {
	if o == nil || IsNil(o.Entryphone) {
		return nil, false
	}
	return o.Entryphone, true
}

// HasEntryphone returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasEntryphone() bool {
	if o != nil && !IsNil(o.Entryphone) {
		return true
	}

	return false
}

// SetEntryphone gets a reference to the given string and assigns it to the Entryphone field.
func (o *OrderDeliveryAddressDTO) SetEntryphone(v string) {
	o.Entryphone = &v
}

// GetFloor returns the Floor field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetFloor() string {
	if o == nil || IsNil(o.Floor) {
		var ret string
		return ret
	}
	return *o.Floor
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetFloorOk() (*string, bool) {
	if o == nil || IsNil(o.Floor) {
		return nil, false
	}
	return o.Floor, true
}

// HasFloor returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasFloor() bool {
	if o != nil && !IsNil(o.Floor) {
		return true
	}

	return false
}

// SetFloor gets a reference to the given string and assigns it to the Floor field.
func (o *OrderDeliveryAddressDTO) SetFloor(v string) {
	o.Floor = &v
}

// GetApartment returns the Apartment field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetApartment() string {
	if o == nil || IsNil(o.Apartment) {
		var ret string
		return ret
	}
	return *o.Apartment
}

// GetApartmentOk returns a tuple with the Apartment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetApartmentOk() (*string, bool) {
	if o == nil || IsNil(o.Apartment) {
		return nil, false
	}
	return o.Apartment, true
}

// HasApartment returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasApartment() bool {
	if o != nil && !IsNil(o.Apartment) {
		return true
	}

	return false
}

// SetApartment gets a reference to the given string and assigns it to the Apartment field.
func (o *OrderDeliveryAddressDTO) SetApartment(v string) {
	o.Apartment = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *OrderDeliveryAddressDTO) SetPhone(v string) {
	o.Phone = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetRecipient() string {
	if o == nil || IsNil(o.Recipient) {
		var ret string
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetRecipientOk() (*string, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given string and assigns it to the Recipient field.
func (o *OrderDeliveryAddressDTO) SetRecipient(v string) {
	o.Recipient = &v
}

// GetGps returns the Gps field value if set, zero value otherwise.
func (o *OrderDeliveryAddressDTO) GetGps() GpsDTO {
	if o == nil || IsNil(o.Gps) {
		var ret GpsDTO
		return ret
	}
	return *o.Gps
}

// GetGpsOk returns a tuple with the Gps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDeliveryAddressDTO) GetGpsOk() (*GpsDTO, bool) {
	if o == nil || IsNil(o.Gps) {
		return nil, false
	}
	return o.Gps, true
}

// HasGps returns a boolean if a field has been set.
func (o *OrderDeliveryAddressDTO) HasGps() bool {
	if o != nil && !IsNil(o.Gps) {
		return true
	}

	return false
}

// SetGps gets a reference to the given GpsDTO and assigns it to the Gps field.
func (o *OrderDeliveryAddressDTO) SetGps(v GpsDTO) {
	o.Gps = &v
}

func (o OrderDeliveryAddressDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDeliveryAddressDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Postcode) {
		toSerialize["postcode"] = o.Postcode
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.District) {
		toSerialize["district"] = o.District
	}
	if !IsNil(o.Subway) {
		toSerialize["subway"] = o.Subway
	}
	if !IsNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	if !IsNil(o.House) {
		toSerialize["house"] = o.House
	}
	if !IsNil(o.Block) {
		toSerialize["block"] = o.Block
	}
	if !IsNil(o.Entrance) {
		toSerialize["entrance"] = o.Entrance
	}
	if !IsNil(o.Entryphone) {
		toSerialize["entryphone"] = o.Entryphone
	}
	if !IsNil(o.Floor) {
		toSerialize["floor"] = o.Floor
	}
	if !IsNil(o.Apartment) {
		toSerialize["apartment"] = o.Apartment
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	if !IsNil(o.Gps) {
		toSerialize["gps"] = o.Gps
	}
	return toSerialize, nil
}

type NullableOrderDeliveryAddressDTO struct {
	value *OrderDeliveryAddressDTO
	isSet bool
}

func (v NullableOrderDeliveryAddressDTO) Get() *OrderDeliveryAddressDTO {
	return v.value
}

func (v *NullableOrderDeliveryAddressDTO) Set(val *OrderDeliveryAddressDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDeliveryAddressDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDeliveryAddressDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDeliveryAddressDTO(val *OrderDeliveryAddressDTO) *NullableOrderDeliveryAddressDTO {
	return &NullableOrderDeliveryAddressDTO{value: val, isSet: true}
}

func (v NullableOrderDeliveryAddressDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDeliveryAddressDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


