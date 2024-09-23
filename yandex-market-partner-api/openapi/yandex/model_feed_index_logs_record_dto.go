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
)

// checks if the FeedIndexLogsRecordDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedIndexLogsRecordDTO{}

// FeedIndexLogsRecordDTO Список отчетов по индексации прайс-листа.
type FeedIndexLogsRecordDTO struct {
	// Дата и время загрузки прайс-листа.  Формат даты: ISO 8601 со смещением относительно UTC. Например, `2017-11-21T00:42:42+03:00`. 
	DownloadTime *time.Time `json:"downloadTime,omitempty"`
	// Дата и время, которые магазин указал в прайс-листе.  Формат даты: ISO 8601 со смещением относительно UTC. Например, `2017-11-21T00:42:42+03:00`. 
	FileTime *time.Time `json:"fileTime,omitempty"`
	// Идентификатор индексации.
	GenerationId *int64 `json:"generationId,omitempty"`
	IndexType *FeedIndexLogsIndexType `json:"indexType,omitempty"`
	// Дата и время публикации предложений из прайс-листа на Яндекс Маркете.  Формат даты: ISO 8601 со смещением относительно UTC. Например, `2017-11-21T00:42:42+03:00`. 
	PublishedTime *time.Time `json:"publishedTime,omitempty"`
	Status *FeedIndexLogsStatusType `json:"status,omitempty"`
	Error *FeedIndexLogsErrorDTO `json:"error,omitempty"`
	Offers *FeedIndexLogsOffersDTO `json:"offers,omitempty"`
}

// NewFeedIndexLogsRecordDTO instantiates a new FeedIndexLogsRecordDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedIndexLogsRecordDTO() *FeedIndexLogsRecordDTO {
	this := FeedIndexLogsRecordDTO{}
	return &this
}

// NewFeedIndexLogsRecordDTOWithDefaults instantiates a new FeedIndexLogsRecordDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedIndexLogsRecordDTOWithDefaults() *FeedIndexLogsRecordDTO {
	this := FeedIndexLogsRecordDTO{}
	return &this
}

// GetDownloadTime returns the DownloadTime field value if set, zero value otherwise.
func (o *FeedIndexLogsRecordDTO) GetDownloadTime() time.Time {
	if o == nil || IsNil(o.DownloadTime) {
		var ret time.Time
		return ret
	}
	return *o.DownloadTime
}

// GetDownloadTimeOk returns a tuple with the DownloadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndexLogsRecordDTO) GetDownloadTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DownloadTime) {
		return nil, false
	}
	return o.DownloadTime, true
}

// HasDownloadTime returns a boolean if a field has been set.
func (o *FeedIndexLogsRecordDTO) HasDownloadTime() bool {
	if o != nil && !IsNil(o.DownloadTime) {
		return true
	}

	return false
}

// SetDownloadTime gets a reference to the given time.Time and assigns it to the DownloadTime field.
func (o *FeedIndexLogsRecordDTO) SetDownloadTime(v time.Time) {
	o.DownloadTime = &v
}

// GetFileTime returns the FileTime field value if set, zero value otherwise.
func (o *FeedIndexLogsRecordDTO) GetFileTime() time.Time {
	if o == nil || IsNil(o.FileTime) {
		var ret time.Time
		return ret
	}
	return *o.FileTime
}

// GetFileTimeOk returns a tuple with the FileTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndexLogsRecordDTO) GetFileTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FileTime) {
		return nil, false
	}
	return o.FileTime, true
}

// HasFileTime returns a boolean if a field has been set.
func (o *FeedIndexLogsRecordDTO) HasFileTime() bool {
	if o != nil && !IsNil(o.FileTime) {
		return true
	}

	return false
}

// SetFileTime gets a reference to the given time.Time and assigns it to the FileTime field.
func (o *FeedIndexLogsRecordDTO) SetFileTime(v time.Time) {
	o.FileTime = &v
}

// GetGenerationId returns the GenerationId field value if set, zero value otherwise.
func (o *FeedIndexLogsRecordDTO) GetGenerationId() int64 {
	if o == nil || IsNil(o.GenerationId) {
		var ret int64
		return ret
	}
	return *o.GenerationId
}

// GetGenerationIdOk returns a tuple with the GenerationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndexLogsRecordDTO) GetGenerationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GenerationId) {
		return nil, false
	}
	return o.GenerationId, true
}

// HasGenerationId returns a boolean if a field has been set.
func (o *FeedIndexLogsRecordDTO) HasGenerationId() bool {
	if o != nil && !IsNil(o.GenerationId) {
		return true
	}

	return false
}

// SetGenerationId gets a reference to the given int64 and assigns it to the GenerationId field.
func (o *FeedIndexLogsRecordDTO) SetGenerationId(v int64) {
	o.GenerationId = &v
}

// GetIndexType returns the IndexType field value if set, zero value otherwise.
func (o *FeedIndexLogsRecordDTO) GetIndexType() FeedIndexLogsIndexType {
	if o == nil || IsNil(o.IndexType) {
		var ret FeedIndexLogsIndexType
		return ret
	}
	return *o.IndexType
}

// GetIndexTypeOk returns a tuple with the IndexType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndexLogsRecordDTO) GetIndexTypeOk() (*FeedIndexLogsIndexType, bool) {
	if o == nil || IsNil(o.IndexType) {
		return nil, false
	}
	return o.IndexType, true
}

// HasIndexType returns a boolean if a field has been set.
func (o *FeedIndexLogsRecordDTO) HasIndexType() bool {
	if o != nil && !IsNil(o.IndexType) {
		return true
	}

	return false
}

// SetIndexType gets a reference to the given FeedIndexLogsIndexType and assigns it to the IndexType field.
func (o *FeedIndexLogsRecordDTO) SetIndexType(v FeedIndexLogsIndexType) {
	o.IndexType = &v
}

// GetPublishedTime returns the PublishedTime field value if set, zero value otherwise.
func (o *FeedIndexLogsRecordDTO) GetPublishedTime() time.Time {
	if o == nil || IsNil(o.PublishedTime) {
		var ret time.Time
		return ret
	}
	return *o.PublishedTime
}

// GetPublishedTimeOk returns a tuple with the PublishedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndexLogsRecordDTO) GetPublishedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishedTime) {
		return nil, false
	}
	return o.PublishedTime, true
}

// HasPublishedTime returns a boolean if a field has been set.
func (o *FeedIndexLogsRecordDTO) HasPublishedTime() bool {
	if o != nil && !IsNil(o.PublishedTime) {
		return true
	}

	return false
}

// SetPublishedTime gets a reference to the given time.Time and assigns it to the PublishedTime field.
func (o *FeedIndexLogsRecordDTO) SetPublishedTime(v time.Time) {
	o.PublishedTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FeedIndexLogsRecordDTO) GetStatus() FeedIndexLogsStatusType {
	if o == nil || IsNil(o.Status) {
		var ret FeedIndexLogsStatusType
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndexLogsRecordDTO) GetStatusOk() (*FeedIndexLogsStatusType, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FeedIndexLogsRecordDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given FeedIndexLogsStatusType and assigns it to the Status field.
func (o *FeedIndexLogsRecordDTO) SetStatus(v FeedIndexLogsStatusType) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *FeedIndexLogsRecordDTO) GetError() FeedIndexLogsErrorDTO {
	if o == nil || IsNil(o.Error) {
		var ret FeedIndexLogsErrorDTO
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndexLogsRecordDTO) GetErrorOk() (*FeedIndexLogsErrorDTO, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *FeedIndexLogsRecordDTO) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given FeedIndexLogsErrorDTO and assigns it to the Error field.
func (o *FeedIndexLogsRecordDTO) SetError(v FeedIndexLogsErrorDTO) {
	o.Error = &v
}

// GetOffers returns the Offers field value if set, zero value otherwise.
func (o *FeedIndexLogsRecordDTO) GetOffers() FeedIndexLogsOffersDTO {
	if o == nil || IsNil(o.Offers) {
		var ret FeedIndexLogsOffersDTO
		return ret
	}
	return *o.Offers
}

// GetOffersOk returns a tuple with the Offers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedIndexLogsRecordDTO) GetOffersOk() (*FeedIndexLogsOffersDTO, bool) {
	if o == nil || IsNil(o.Offers) {
		return nil, false
	}
	return o.Offers, true
}

// HasOffers returns a boolean if a field has been set.
func (o *FeedIndexLogsRecordDTO) HasOffers() bool {
	if o != nil && !IsNil(o.Offers) {
		return true
	}

	return false
}

// SetOffers gets a reference to the given FeedIndexLogsOffersDTO and assigns it to the Offers field.
func (o *FeedIndexLogsRecordDTO) SetOffers(v FeedIndexLogsOffersDTO) {
	o.Offers = &v
}

func (o FeedIndexLogsRecordDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedIndexLogsRecordDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownloadTime) {
		toSerialize["downloadTime"] = o.DownloadTime
	}
	if !IsNil(o.FileTime) {
		toSerialize["fileTime"] = o.FileTime
	}
	if !IsNil(o.GenerationId) {
		toSerialize["generationId"] = o.GenerationId
	}
	if !IsNil(o.IndexType) {
		toSerialize["indexType"] = o.IndexType
	}
	if !IsNil(o.PublishedTime) {
		toSerialize["publishedTime"] = o.PublishedTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Offers) {
		toSerialize["offers"] = o.Offers
	}
	return toSerialize, nil
}

type NullableFeedIndexLogsRecordDTO struct {
	value *FeedIndexLogsRecordDTO
	isSet bool
}

func (v NullableFeedIndexLogsRecordDTO) Get() *FeedIndexLogsRecordDTO {
	return v.value
}

func (v *NullableFeedIndexLogsRecordDTO) Set(val *FeedIndexLogsRecordDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedIndexLogsRecordDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedIndexLogsRecordDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedIndexLogsRecordDTO(val *FeedIndexLogsRecordDTO) *NullableFeedIndexLogsRecordDTO {
	return &NullableFeedIndexLogsRecordDTO{value: val, isSet: true}
}

func (v NullableFeedIndexLogsRecordDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedIndexLogsRecordDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


