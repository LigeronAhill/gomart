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

// checks if the ReportInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportInfoDTO{}

// ReportInfoDTO Статус генерации и ссылка на готовый отчет.
type ReportInfoDTO struct {
	Status ReportStatusType `json:"status"`
	SubStatus *ReportSubStatusType `json:"subStatus,omitempty"`
	// Дата и время запроса на генерацию.
	GenerationRequestedAt time.Time `json:"generationRequestedAt"`
	// Дата и время завершения генерации.
	GenerationFinishedAt *time.Time `json:"generationFinishedAt,omitempty"`
	// Ссылка на готовый отчет.
	File *string `json:"file,omitempty"`
	// Ожидаемая продолжительность генерации в миллисекундах.
	EstimatedGenerationTime *int64 `json:"estimatedGenerationTime,omitempty"`
}

type _ReportInfoDTO ReportInfoDTO

// NewReportInfoDTO instantiates a new ReportInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportInfoDTO(status ReportStatusType, generationRequestedAt time.Time) *ReportInfoDTO {
	this := ReportInfoDTO{}
	this.Status = status
	this.GenerationRequestedAt = generationRequestedAt
	return &this
}

// NewReportInfoDTOWithDefaults instantiates a new ReportInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportInfoDTOWithDefaults() *ReportInfoDTO {
	this := ReportInfoDTO{}
	return &this
}

// GetStatus returns the Status field value
func (o *ReportInfoDTO) GetStatus() ReportStatusType {
	if o == nil {
		var ret ReportStatusType
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReportInfoDTO) GetStatusOk() (*ReportStatusType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReportInfoDTO) SetStatus(v ReportStatusType) {
	o.Status = v
}

// GetSubStatus returns the SubStatus field value if set, zero value otherwise.
func (o *ReportInfoDTO) GetSubStatus() ReportSubStatusType {
	if o == nil || IsNil(o.SubStatus) {
		var ret ReportSubStatusType
		return ret
	}
	return *o.SubStatus
}

// GetSubStatusOk returns a tuple with the SubStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfoDTO) GetSubStatusOk() (*ReportSubStatusType, bool) {
	if o == nil || IsNil(o.SubStatus) {
		return nil, false
	}
	return o.SubStatus, true
}

// HasSubStatus returns a boolean if a field has been set.
func (o *ReportInfoDTO) HasSubStatus() bool {
	if o != nil && !IsNil(o.SubStatus) {
		return true
	}

	return false
}

// SetSubStatus gets a reference to the given ReportSubStatusType and assigns it to the SubStatus field.
func (o *ReportInfoDTO) SetSubStatus(v ReportSubStatusType) {
	o.SubStatus = &v
}

// GetGenerationRequestedAt returns the GenerationRequestedAt field value
func (o *ReportInfoDTO) GetGenerationRequestedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.GenerationRequestedAt
}

// GetGenerationRequestedAtOk returns a tuple with the GenerationRequestedAt field value
// and a boolean to check if the value has been set.
func (o *ReportInfoDTO) GetGenerationRequestedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GenerationRequestedAt, true
}

// SetGenerationRequestedAt sets field value
func (o *ReportInfoDTO) SetGenerationRequestedAt(v time.Time) {
	o.GenerationRequestedAt = v
}

// GetGenerationFinishedAt returns the GenerationFinishedAt field value if set, zero value otherwise.
func (o *ReportInfoDTO) GetGenerationFinishedAt() time.Time {
	if o == nil || IsNil(o.GenerationFinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.GenerationFinishedAt
}

// GetGenerationFinishedAtOk returns a tuple with the GenerationFinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfoDTO) GetGenerationFinishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.GenerationFinishedAt) {
		return nil, false
	}
	return o.GenerationFinishedAt, true
}

// HasGenerationFinishedAt returns a boolean if a field has been set.
func (o *ReportInfoDTO) HasGenerationFinishedAt() bool {
	if o != nil && !IsNil(o.GenerationFinishedAt) {
		return true
	}

	return false
}

// SetGenerationFinishedAt gets a reference to the given time.Time and assigns it to the GenerationFinishedAt field.
func (o *ReportInfoDTO) SetGenerationFinishedAt(v time.Time) {
	o.GenerationFinishedAt = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *ReportInfoDTO) GetFile() string {
	if o == nil || IsNil(o.File) {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfoDTO) GetFileOk() (*string, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *ReportInfoDTO) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *ReportInfoDTO) SetFile(v string) {
	o.File = &v
}

// GetEstimatedGenerationTime returns the EstimatedGenerationTime field value if set, zero value otherwise.
func (o *ReportInfoDTO) GetEstimatedGenerationTime() int64 {
	if o == nil || IsNil(o.EstimatedGenerationTime) {
		var ret int64
		return ret
	}
	return *o.EstimatedGenerationTime
}

// GetEstimatedGenerationTimeOk returns a tuple with the EstimatedGenerationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportInfoDTO) GetEstimatedGenerationTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.EstimatedGenerationTime) {
		return nil, false
	}
	return o.EstimatedGenerationTime, true
}

// HasEstimatedGenerationTime returns a boolean if a field has been set.
func (o *ReportInfoDTO) HasEstimatedGenerationTime() bool {
	if o != nil && !IsNil(o.EstimatedGenerationTime) {
		return true
	}

	return false
}

// SetEstimatedGenerationTime gets a reference to the given int64 and assigns it to the EstimatedGenerationTime field.
func (o *ReportInfoDTO) SetEstimatedGenerationTime(v int64) {
	o.EstimatedGenerationTime = &v
}

func (o ReportInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.SubStatus) {
		toSerialize["subStatus"] = o.SubStatus
	}
	toSerialize["generationRequestedAt"] = o.GenerationRequestedAt
	if !IsNil(o.GenerationFinishedAt) {
		toSerialize["generationFinishedAt"] = o.GenerationFinishedAt
	}
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.EstimatedGenerationTime) {
		toSerialize["estimatedGenerationTime"] = o.EstimatedGenerationTime
	}
	return toSerialize, nil
}

func (o *ReportInfoDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"generationRequestedAt",
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

	varReportInfoDTO := _ReportInfoDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportInfoDTO)

	if err != nil {
		return err
	}

	*o = ReportInfoDTO(varReportInfoDTO)

	return err
}

type NullableReportInfoDTO struct {
	value *ReportInfoDTO
	isSet bool
}

func (v NullableReportInfoDTO) Get() *ReportInfoDTO {
	return v.value
}

func (v *NullableReportInfoDTO) Set(val *ReportInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableReportInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableReportInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportInfoDTO(val *ReportInfoDTO) *NullableReportInfoDTO {
	return &NullableReportInfoDTO{value: val, isSet: true}
}

func (v NullableReportInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


