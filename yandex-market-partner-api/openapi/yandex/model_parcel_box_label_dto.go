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

// checks if the ParcelBoxLabelDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParcelBoxLabelDTO{}

// ParcelBoxLabelDTO Информация о ярлыке для коробки.
type ParcelBoxLabelDTO struct {
	// Соответствует URL, по которому выполняется запрос [GET campaigns/{campaignId}/orders/{orderId}/delivery/shipments/{shipmentId}/boxes/{boxId}/label](../../reference/orders/generateOrderLabel.md). 
	Url string `json:"url"`
	// Юридическое название магазина.
	SupplierName string `json:"supplierName"`
	// Юридическое название службы доставки.
	DeliveryServiceName string `json:"deliveryServiceName"`
	// Идентификатор заказа в системе Маркета.
	OrderId int64 `json:"orderId"`
	// Идентификатор заказа в информационной системе магазина.  Совпадает с `orderId`, если Маркету неизвестен номер заказа в системе магазина. 
	OrderNum string `json:"orderNum"`
	// Фамилия и инициалы получателя заказа.
	RecipientName string `json:"recipientName"`
	// Идентификатор коробки.
	BoxId int64 `json:"boxId"`
	// Идентификатор коробки в информационной системе магазина.  Возвращается в формате: `номер заказа на Маркете-номер коробки`. Например, `7206821‑1`, `7206821‑2` и т. д. 
	FulfilmentId string `json:"fulfilmentId"`
	// Номер коробки в заказе. Возвращается в формате: `номер места/общее количество мест`. 
	Place string `json:"place"`
	// {% note warning \"\" %}  Этот параметр устарел. Не используйте его.  {% endnote %}  Общая масса всех товаров в заказе. Возвращается в формате: `weight кг`. 
	Weight string `json:"weight"`
	// Идентификатор службы доставки. Информацию о службе доставки можно получить с помощью запроса [GET delivery/services](../../reference/orders/getDeliveryServices.md).
	DeliveryServiceId string `json:"deliveryServiceId"`
	// Адрес получателя.
	DeliveryAddress *string `json:"deliveryAddress,omitempty"`
	// Дата отгрузки в формате `dd.MM.yyyy`.
	ShipmentDate *string `json:"shipmentDate,omitempty"`
}

type _ParcelBoxLabelDTO ParcelBoxLabelDTO

// NewParcelBoxLabelDTO instantiates a new ParcelBoxLabelDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParcelBoxLabelDTO(url string, supplierName string, deliveryServiceName string, orderId int64, orderNum string, recipientName string, boxId int64, fulfilmentId string, place string, weight string, deliveryServiceId string) *ParcelBoxLabelDTO {
	this := ParcelBoxLabelDTO{}
	this.Url = url
	this.SupplierName = supplierName
	this.DeliveryServiceName = deliveryServiceName
	this.OrderId = orderId
	this.OrderNum = orderNum
	this.RecipientName = recipientName
	this.BoxId = boxId
	this.FulfilmentId = fulfilmentId
	this.Place = place
	this.Weight = weight
	this.DeliveryServiceId = deliveryServiceId
	return &this
}

// NewParcelBoxLabelDTOWithDefaults instantiates a new ParcelBoxLabelDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParcelBoxLabelDTOWithDefaults() *ParcelBoxLabelDTO {
	this := ParcelBoxLabelDTO{}
	return &this
}

// GetUrl returns the Url field value
func (o *ParcelBoxLabelDTO) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ParcelBoxLabelDTO) SetUrl(v string) {
	o.Url = v
}

// GetSupplierName returns the SupplierName field value
func (o *ParcelBoxLabelDTO) GetSupplierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupplierName
}

// GetSupplierNameOk returns a tuple with the SupplierName field value
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetSupplierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupplierName, true
}

// SetSupplierName sets field value
func (o *ParcelBoxLabelDTO) SetSupplierName(v string) {
	o.SupplierName = v
}

// GetDeliveryServiceName returns the DeliveryServiceName field value
func (o *ParcelBoxLabelDTO) GetDeliveryServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryServiceName
}

// GetDeliveryServiceNameOk returns a tuple with the DeliveryServiceName field value
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetDeliveryServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryServiceName, true
}

// SetDeliveryServiceName sets field value
func (o *ParcelBoxLabelDTO) SetDeliveryServiceName(v string) {
	o.DeliveryServiceName = v
}

// GetOrderId returns the OrderId field value
func (o *ParcelBoxLabelDTO) GetOrderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetOrderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *ParcelBoxLabelDTO) SetOrderId(v int64) {
	o.OrderId = v
}

// GetOrderNum returns the OrderNum field value
func (o *ParcelBoxLabelDTO) GetOrderNum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderNum
}

// GetOrderNumOk returns a tuple with the OrderNum field value
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetOrderNumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderNum, true
}

// SetOrderNum sets field value
func (o *ParcelBoxLabelDTO) SetOrderNum(v string) {
	o.OrderNum = v
}

// GetRecipientName returns the RecipientName field value
func (o *ParcelBoxLabelDTO) GetRecipientName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientName
}

// GetRecipientNameOk returns a tuple with the RecipientName field value
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetRecipientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientName, true
}

// SetRecipientName sets field value
func (o *ParcelBoxLabelDTO) SetRecipientName(v string) {
	o.RecipientName = v
}

// GetBoxId returns the BoxId field value
func (o *ParcelBoxLabelDTO) GetBoxId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BoxId
}

// GetBoxIdOk returns a tuple with the BoxId field value
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetBoxIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoxId, true
}

// SetBoxId sets field value
func (o *ParcelBoxLabelDTO) SetBoxId(v int64) {
	o.BoxId = v
}

// GetFulfilmentId returns the FulfilmentId field value
func (o *ParcelBoxLabelDTO) GetFulfilmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfilmentId
}

// GetFulfilmentIdOk returns a tuple with the FulfilmentId field value
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetFulfilmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfilmentId, true
}

// SetFulfilmentId sets field value
func (o *ParcelBoxLabelDTO) SetFulfilmentId(v string) {
	o.FulfilmentId = v
}

// GetPlace returns the Place field value
func (o *ParcelBoxLabelDTO) GetPlace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Place
}

// GetPlaceOk returns a tuple with the Place field value
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetPlaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Place, true
}

// SetPlace sets field value
func (o *ParcelBoxLabelDTO) SetPlace(v string) {
	o.Place = v
}

// GetWeight returns the Weight field value
func (o *ParcelBoxLabelDTO) GetWeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetWeightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *ParcelBoxLabelDTO) SetWeight(v string) {
	o.Weight = v
}

// GetDeliveryServiceId returns the DeliveryServiceId field value
func (o *ParcelBoxLabelDTO) GetDeliveryServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliveryServiceId
}

// GetDeliveryServiceIdOk returns a tuple with the DeliveryServiceId field value
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetDeliveryServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveryServiceId, true
}

// SetDeliveryServiceId sets field value
func (o *ParcelBoxLabelDTO) SetDeliveryServiceId(v string) {
	o.DeliveryServiceId = v
}

// GetDeliveryAddress returns the DeliveryAddress field value if set, zero value otherwise.
func (o *ParcelBoxLabelDTO) GetDeliveryAddress() string {
	if o == nil || IsNil(o.DeliveryAddress) {
		var ret string
		return ret
	}
	return *o.DeliveryAddress
}

// GetDeliveryAddressOk returns a tuple with the DeliveryAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetDeliveryAddressOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryAddress) {
		return nil, false
	}
	return o.DeliveryAddress, true
}

// HasDeliveryAddress returns a boolean if a field has been set.
func (o *ParcelBoxLabelDTO) HasDeliveryAddress() bool {
	if o != nil && !IsNil(o.DeliveryAddress) {
		return true
	}

	return false
}

// SetDeliveryAddress gets a reference to the given string and assigns it to the DeliveryAddress field.
func (o *ParcelBoxLabelDTO) SetDeliveryAddress(v string) {
	o.DeliveryAddress = &v
}

// GetShipmentDate returns the ShipmentDate field value if set, zero value otherwise.
func (o *ParcelBoxLabelDTO) GetShipmentDate() string {
	if o == nil || IsNil(o.ShipmentDate) {
		var ret string
		return ret
	}
	return *o.ShipmentDate
}

// GetShipmentDateOk returns a tuple with the ShipmentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParcelBoxLabelDTO) GetShipmentDateOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentDate) {
		return nil, false
	}
	return o.ShipmentDate, true
}

// HasShipmentDate returns a boolean if a field has been set.
func (o *ParcelBoxLabelDTO) HasShipmentDate() bool {
	if o != nil && !IsNil(o.ShipmentDate) {
		return true
	}

	return false
}

// SetShipmentDate gets a reference to the given string and assigns it to the ShipmentDate field.
func (o *ParcelBoxLabelDTO) SetShipmentDate(v string) {
	o.ShipmentDate = &v
}

func (o ParcelBoxLabelDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParcelBoxLabelDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["supplierName"] = o.SupplierName
	toSerialize["deliveryServiceName"] = o.DeliveryServiceName
	toSerialize["orderId"] = o.OrderId
	toSerialize["orderNum"] = o.OrderNum
	toSerialize["recipientName"] = o.RecipientName
	toSerialize["boxId"] = o.BoxId
	toSerialize["fulfilmentId"] = o.FulfilmentId
	toSerialize["place"] = o.Place
	toSerialize["weight"] = o.Weight
	toSerialize["deliveryServiceId"] = o.DeliveryServiceId
	if !IsNil(o.DeliveryAddress) {
		toSerialize["deliveryAddress"] = o.DeliveryAddress
	}
	if !IsNil(o.ShipmentDate) {
		toSerialize["shipmentDate"] = o.ShipmentDate
	}
	return toSerialize, nil
}

func (o *ParcelBoxLabelDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"supplierName",
		"deliveryServiceName",
		"orderId",
		"orderNum",
		"recipientName",
		"boxId",
		"fulfilmentId",
		"place",
		"weight",
		"deliveryServiceId",
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

	varParcelBoxLabelDTO := _ParcelBoxLabelDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varParcelBoxLabelDTO)

	if err != nil {
		return err
	}

	*o = ParcelBoxLabelDTO(varParcelBoxLabelDTO)

	return err
}

type NullableParcelBoxLabelDTO struct {
	value *ParcelBoxLabelDTO
	isSet bool
}

func (v NullableParcelBoxLabelDTO) Get() *ParcelBoxLabelDTO {
	return v.value
}

func (v *NullableParcelBoxLabelDTO) Set(val *ParcelBoxLabelDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableParcelBoxLabelDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableParcelBoxLabelDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParcelBoxLabelDTO(val *ParcelBoxLabelDTO) *NullableParcelBoxLabelDTO {
	return &NullableParcelBoxLabelDTO{value: val, isSet: true}
}

func (v NullableParcelBoxLabelDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParcelBoxLabelDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


