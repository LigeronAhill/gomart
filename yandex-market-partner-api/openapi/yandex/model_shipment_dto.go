/*
Партнерский API Маркета

API Яндекс Маркета помогает продавцам автоматизировать и упростить работу с маркетплейсом.  В числе возможностей интеграции:  * управление каталогом товаров и витриной,  * обработка заказов,  * изменение настроек магазина,  * получение отчетов. 

API version: LATEST
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the ShipmentDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDTO{}

// ShipmentDTO Информация об отгрузке.
type ShipmentDTO struct {
	// Идентификатор отгрузки.
	Id *int64 `json:"id,omitempty"`
	// Начало планового интервала отгрузки.
	PlanIntervalFrom *time.Time `json:"planIntervalFrom,omitempty"`
	// Конец планового интервала отгрузки.
	PlanIntervalTo *time.Time `json:"planIntervalTo,omitempty"`
	ShipmentType *ShipmentType `json:"shipmentType,omitempty"`
	Warehouse *PartnerShipmentWarehouseDTO `json:"warehouse,omitempty"`
	WarehouseTo *PartnerShipmentWarehouseDTO `json:"warehouseTo,omitempty"`
	// Идентификатор отгрузки в вашей системе. Если вы еще не передавали идентификатор, вернется идентификатор из параметра `id`.
	ExternalId *string `json:"externalId,omitempty"`
	DeliveryService *DeliveryServiceDTO `json:"deliveryService,omitempty"`
	PalletsCount *PalletsCountDTO `json:"palletsCount,omitempty"`
	// Идентификаторы заказов в отгрузке.
	OrderIds []int64 `json:"orderIds"`
	// Количество заказов, которое Маркет запланировал к отгрузке.
	DraftCount *int32 `json:"draftCount,omitempty"`
	// Количество заказов, которое Маркет подтвердил к отгрузке.
	PlannedCount *int32 `json:"plannedCount,omitempty"`
	// Количество заказов, принятых в сортировочном центре или пункте приема.
	FactCount *int32 `json:"factCount,omitempty"`
	CurrentStatus *ShipmentStatusChangeDTO `json:"currentStatus,omitempty"`
	// Доступные действия над отгрузкой.
	AvailableActions []ShipmentActionType `json:"availableActions"`
}

type _ShipmentDTO ShipmentDTO

// NewShipmentDTO instantiates a new ShipmentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDTO(orderIds []int64, availableActions []ShipmentActionType) *ShipmentDTO {
	this := ShipmentDTO{}
	this.OrderIds = orderIds
	this.AvailableActions = availableActions
	return &this
}

// NewShipmentDTOWithDefaults instantiates a new ShipmentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDTOWithDefaults() *ShipmentDTO {
	this := ShipmentDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ShipmentDTO) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ShipmentDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ShipmentDTO) SetId(v int64) {
	o.Id = &v
}

// GetPlanIntervalFrom returns the PlanIntervalFrom field value if set, zero value otherwise.
func (o *ShipmentDTO) GetPlanIntervalFrom() time.Time {
	if o == nil || IsNil(o.PlanIntervalFrom) {
		var ret time.Time
		return ret
	}
	return *o.PlanIntervalFrom
}

// GetPlanIntervalFromOk returns a tuple with the PlanIntervalFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetPlanIntervalFromOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PlanIntervalFrom) {
		return nil, false
	}
	return o.PlanIntervalFrom, true
}

// HasPlanIntervalFrom returns a boolean if a field has been set.
func (o *ShipmentDTO) HasPlanIntervalFrom() bool {
	if o != nil && !IsNil(o.PlanIntervalFrom) {
		return true
	}

	return false
}

// SetPlanIntervalFrom gets a reference to the given time.Time and assigns it to the PlanIntervalFrom field.
func (o *ShipmentDTO) SetPlanIntervalFrom(v time.Time) {
	o.PlanIntervalFrom = &v
}

// GetPlanIntervalTo returns the PlanIntervalTo field value if set, zero value otherwise.
func (o *ShipmentDTO) GetPlanIntervalTo() time.Time {
	if o == nil || IsNil(o.PlanIntervalTo) {
		var ret time.Time
		return ret
	}
	return *o.PlanIntervalTo
}

// GetPlanIntervalToOk returns a tuple with the PlanIntervalTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetPlanIntervalToOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PlanIntervalTo) {
		return nil, false
	}
	return o.PlanIntervalTo, true
}

// HasPlanIntervalTo returns a boolean if a field has been set.
func (o *ShipmentDTO) HasPlanIntervalTo() bool {
	if o != nil && !IsNil(o.PlanIntervalTo) {
		return true
	}

	return false
}

// SetPlanIntervalTo gets a reference to the given time.Time and assigns it to the PlanIntervalTo field.
func (o *ShipmentDTO) SetPlanIntervalTo(v time.Time) {
	o.PlanIntervalTo = &v
}

// GetShipmentType returns the ShipmentType field value if set, zero value otherwise.
func (o *ShipmentDTO) GetShipmentType() ShipmentType {
	if o == nil || IsNil(o.ShipmentType) {
		var ret ShipmentType
		return ret
	}
	return *o.ShipmentType
}

// GetShipmentTypeOk returns a tuple with the ShipmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetShipmentTypeOk() (*ShipmentType, bool) {
	if o == nil || IsNil(o.ShipmentType) {
		return nil, false
	}
	return o.ShipmentType, true
}

// HasShipmentType returns a boolean if a field has been set.
func (o *ShipmentDTO) HasShipmentType() bool {
	if o != nil && !IsNil(o.ShipmentType) {
		return true
	}

	return false
}

// SetShipmentType gets a reference to the given ShipmentType and assigns it to the ShipmentType field.
func (o *ShipmentDTO) SetShipmentType(v ShipmentType) {
	o.ShipmentType = &v
}

// GetWarehouse returns the Warehouse field value if set, zero value otherwise.
func (o *ShipmentDTO) GetWarehouse() PartnerShipmentWarehouseDTO {
	if o == nil || IsNil(o.Warehouse) {
		var ret PartnerShipmentWarehouseDTO
		return ret
	}
	return *o.Warehouse
}

// GetWarehouseOk returns a tuple with the Warehouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetWarehouseOk() (*PartnerShipmentWarehouseDTO, bool) {
	if o == nil || IsNil(o.Warehouse) {
		return nil, false
	}
	return o.Warehouse, true
}

// HasWarehouse returns a boolean if a field has been set.
func (o *ShipmentDTO) HasWarehouse() bool {
	if o != nil && !IsNil(o.Warehouse) {
		return true
	}

	return false
}

// SetWarehouse gets a reference to the given PartnerShipmentWarehouseDTO and assigns it to the Warehouse field.
func (o *ShipmentDTO) SetWarehouse(v PartnerShipmentWarehouseDTO) {
	o.Warehouse = &v
}

// GetWarehouseTo returns the WarehouseTo field value if set, zero value otherwise.
func (o *ShipmentDTO) GetWarehouseTo() PartnerShipmentWarehouseDTO {
	if o == nil || IsNil(o.WarehouseTo) {
		var ret PartnerShipmentWarehouseDTO
		return ret
	}
	return *o.WarehouseTo
}

// GetWarehouseToOk returns a tuple with the WarehouseTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetWarehouseToOk() (*PartnerShipmentWarehouseDTO, bool) {
	if o == nil || IsNil(o.WarehouseTo) {
		return nil, false
	}
	return o.WarehouseTo, true
}

// HasWarehouseTo returns a boolean if a field has been set.
func (o *ShipmentDTO) HasWarehouseTo() bool {
	if o != nil && !IsNil(o.WarehouseTo) {
		return true
	}

	return false
}

// SetWarehouseTo gets a reference to the given PartnerShipmentWarehouseDTO and assigns it to the WarehouseTo field.
func (o *ShipmentDTO) SetWarehouseTo(v PartnerShipmentWarehouseDTO) {
	o.WarehouseTo = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ShipmentDTO) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ShipmentDTO) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ShipmentDTO) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetDeliveryService returns the DeliveryService field value if set, zero value otherwise.
func (o *ShipmentDTO) GetDeliveryService() DeliveryServiceDTO {
	if o == nil || IsNil(o.DeliveryService) {
		var ret DeliveryServiceDTO
		return ret
	}
	return *o.DeliveryService
}

// GetDeliveryServiceOk returns a tuple with the DeliveryService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetDeliveryServiceOk() (*DeliveryServiceDTO, bool) {
	if o == nil || IsNil(o.DeliveryService) {
		return nil, false
	}
	return o.DeliveryService, true
}

// HasDeliveryService returns a boolean if a field has been set.
func (o *ShipmentDTO) HasDeliveryService() bool {
	if o != nil && !IsNil(o.DeliveryService) {
		return true
	}

	return false
}

// SetDeliveryService gets a reference to the given DeliveryServiceDTO and assigns it to the DeliveryService field.
func (o *ShipmentDTO) SetDeliveryService(v DeliveryServiceDTO) {
	o.DeliveryService = &v
}

// GetPalletsCount returns the PalletsCount field value if set, zero value otherwise.
func (o *ShipmentDTO) GetPalletsCount() PalletsCountDTO {
	if o == nil || IsNil(o.PalletsCount) {
		var ret PalletsCountDTO
		return ret
	}
	return *o.PalletsCount
}

// GetPalletsCountOk returns a tuple with the PalletsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetPalletsCountOk() (*PalletsCountDTO, bool) {
	if o == nil || IsNil(o.PalletsCount) {
		return nil, false
	}
	return o.PalletsCount, true
}

// HasPalletsCount returns a boolean if a field has been set.
func (o *ShipmentDTO) HasPalletsCount() bool {
	if o != nil && !IsNil(o.PalletsCount) {
		return true
	}

	return false
}

// SetPalletsCount gets a reference to the given PalletsCountDTO and assigns it to the PalletsCount field.
func (o *ShipmentDTO) SetPalletsCount(v PalletsCountDTO) {
	o.PalletsCount = &v
}

// GetOrderIds returns the OrderIds field value
func (o *ShipmentDTO) GetOrderIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.OrderIds
}

// GetOrderIdsOk returns a tuple with the OrderIds field value
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetOrderIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderIds, true
}

// SetOrderIds sets field value
func (o *ShipmentDTO) SetOrderIds(v []int64) {
	o.OrderIds = v
}

// GetDraftCount returns the DraftCount field value if set, zero value otherwise.
func (o *ShipmentDTO) GetDraftCount() int32 {
	if o == nil || IsNil(o.DraftCount) {
		var ret int32
		return ret
	}
	return *o.DraftCount
}

// GetDraftCountOk returns a tuple with the DraftCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetDraftCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DraftCount) {
		return nil, false
	}
	return o.DraftCount, true
}

// HasDraftCount returns a boolean if a field has been set.
func (o *ShipmentDTO) HasDraftCount() bool {
	if o != nil && !IsNil(o.DraftCount) {
		return true
	}

	return false
}

// SetDraftCount gets a reference to the given int32 and assigns it to the DraftCount field.
func (o *ShipmentDTO) SetDraftCount(v int32) {
	o.DraftCount = &v
}

// GetPlannedCount returns the PlannedCount field value if set, zero value otherwise.
func (o *ShipmentDTO) GetPlannedCount() int32 {
	if o == nil || IsNil(o.PlannedCount) {
		var ret int32
		return ret
	}
	return *o.PlannedCount
}

// GetPlannedCountOk returns a tuple with the PlannedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetPlannedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PlannedCount) {
		return nil, false
	}
	return o.PlannedCount, true
}

// HasPlannedCount returns a boolean if a field has been set.
func (o *ShipmentDTO) HasPlannedCount() bool {
	if o != nil && !IsNil(o.PlannedCount) {
		return true
	}

	return false
}

// SetPlannedCount gets a reference to the given int32 and assigns it to the PlannedCount field.
func (o *ShipmentDTO) SetPlannedCount(v int32) {
	o.PlannedCount = &v
}

// GetFactCount returns the FactCount field value if set, zero value otherwise.
func (o *ShipmentDTO) GetFactCount() int32 {
	if o == nil || IsNil(o.FactCount) {
		var ret int32
		return ret
	}
	return *o.FactCount
}

// GetFactCountOk returns a tuple with the FactCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetFactCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FactCount) {
		return nil, false
	}
	return o.FactCount, true
}

// HasFactCount returns a boolean if a field has been set.
func (o *ShipmentDTO) HasFactCount() bool {
	if o != nil && !IsNil(o.FactCount) {
		return true
	}

	return false
}

// SetFactCount gets a reference to the given int32 and assigns it to the FactCount field.
func (o *ShipmentDTO) SetFactCount(v int32) {
	o.FactCount = &v
}

// GetCurrentStatus returns the CurrentStatus field value if set, zero value otherwise.
func (o *ShipmentDTO) GetCurrentStatus() ShipmentStatusChangeDTO {
	if o == nil || IsNil(o.CurrentStatus) {
		var ret ShipmentStatusChangeDTO
		return ret
	}
	return *o.CurrentStatus
}

// GetCurrentStatusOk returns a tuple with the CurrentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetCurrentStatusOk() (*ShipmentStatusChangeDTO, bool) {
	if o == nil || IsNil(o.CurrentStatus) {
		return nil, false
	}
	return o.CurrentStatus, true
}

// HasCurrentStatus returns a boolean if a field has been set.
func (o *ShipmentDTO) HasCurrentStatus() bool {
	if o != nil && !IsNil(o.CurrentStatus) {
		return true
	}

	return false
}

// SetCurrentStatus gets a reference to the given ShipmentStatusChangeDTO and assigns it to the CurrentStatus field.
func (o *ShipmentDTO) SetCurrentStatus(v ShipmentStatusChangeDTO) {
	o.CurrentStatus = &v
}

// GetAvailableActions returns the AvailableActions field value
func (o *ShipmentDTO) GetAvailableActions() []ShipmentActionType {
	if o == nil {
		var ret []ShipmentActionType
		return ret
	}

	return o.AvailableActions
}

// GetAvailableActionsOk returns a tuple with the AvailableActions field value
// and a boolean to check if the value has been set.
func (o *ShipmentDTO) GetAvailableActionsOk() ([]ShipmentActionType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableActions, true
}

// SetAvailableActions sets field value
func (o *ShipmentDTO) SetAvailableActions(v []ShipmentActionType) {
	o.AvailableActions = v
}

func (o ShipmentDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PlanIntervalFrom) {
		toSerialize["planIntervalFrom"] = o.PlanIntervalFrom
	}
	if !IsNil(o.PlanIntervalTo) {
		toSerialize["planIntervalTo"] = o.PlanIntervalTo
	}
	if !IsNil(o.ShipmentType) {
		toSerialize["shipmentType"] = o.ShipmentType
	}
	if !IsNil(o.Warehouse) {
		toSerialize["warehouse"] = o.Warehouse
	}
	if !IsNil(o.WarehouseTo) {
		toSerialize["warehouseTo"] = o.WarehouseTo
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.DeliveryService) {
		toSerialize["deliveryService"] = o.DeliveryService
	}
	if !IsNil(o.PalletsCount) {
		toSerialize["palletsCount"] = o.PalletsCount
	}
	toSerialize["orderIds"] = o.OrderIds
	if !IsNil(o.DraftCount) {
		toSerialize["draftCount"] = o.DraftCount
	}
	if !IsNil(o.PlannedCount) {
		toSerialize["plannedCount"] = o.PlannedCount
	}
	if !IsNil(o.FactCount) {
		toSerialize["factCount"] = o.FactCount
	}
	if !IsNil(o.CurrentStatus) {
		toSerialize["currentStatus"] = o.CurrentStatus
	}
	toSerialize["availableActions"] = o.AvailableActions
	return toSerialize, nil
}

func (o *ShipmentDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"orderIds",
		"availableActions",
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

	varShipmentDTO := _ShipmentDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipmentDTO)

	if err != nil {
		return err
	}

	*o = ShipmentDTO(varShipmentDTO)

	return err
}

type NullableShipmentDTO struct {
	value *ShipmentDTO
	isSet bool
}

func (v NullableShipmentDTO) Get() *ShipmentDTO {
	return v.value
}

func (v *NullableShipmentDTO) Set(val *ShipmentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDTO(val *ShipmentDTO) *NullableShipmentDTO {
	return &NullableShipmentDTO{value: val, isSet: true}
}

func (v NullableShipmentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


