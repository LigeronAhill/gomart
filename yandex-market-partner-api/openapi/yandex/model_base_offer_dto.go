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

// checks if the BaseOfferDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseOfferDTO{}

// BaseOfferDTO Основные параметры товара.
type BaseOfferDTO struct {
	// Ваш SKU — идентификатор товара в вашей системе.  Разрешена любая последовательность длиной до 255 знаков.  Правила использования SKU:  * У каждого товара SKU должен быть свой.  * SKU товара нельзя менять — можно только удалить товар и добавить заново с новым SKU.  * Уже заданный SKU нельзя освободить и использовать заново для другого товара. Каждый товар должен получать новый идентификатор, до того никогда не использовавшийся в вашем каталоге.  [Что такое SKU и как его назначать](https://yandex.ru/support/marketplace/assortment/add/index.html#fields) 
	OfferId string `json:"offerId" validate:"regexp=^[^\\\\x00-\\\\x08\\\\x0A-\\\\x1f\\\\x7f]{1,255}$"`
	// Составляйте название по схеме: тип + бренд или производитель + модель + особенности, если есть (например, цвет, размер или вес) и количество в упаковке.  Не включайте в название условия продажи (например, «скидка», «бесплатная доставка» и т. д.), эмоциональные характеристики («хит», «супер» и т. д.). Не пишите слова большими буквами — кроме устоявшихся названий брендов и моделей.  Оптимальная длина — 50–60 символов, максимальная — 256.  [Рекомендации и правила](https://yandex.ru/support/marketplace/assortment/fields/title.html) 
	Name *string `json:"name,omitempty"`
	// Идентификатор категории на Маркете, к которой вы относите свой товар.  Если не указать `marketCategoryId`, то маркетная категория будет определена автоматически.  При изменении информации о товаре передавайте тот же идентификатор категории. Если вы укажете другой, категория товара не поменяется. Изменить ее можно только в кабинете продавца на Маркете.  Список категорий Маркета можно получить с помощью запроса  [POST categories/tree](../../reference/categories/getCategoriesTree.md). 
	MarketCategoryId *int64 `json:"marketCategoryId,omitempty"`
	// Категория товара в вашем магазине. Значение будет использовано для определения категории товара на Маркете в случае, если вы не передали категорию в параметре `marketCategoryId`.  Указывайте конкретные категории — например, набор ножей лучше отнести к категории **Столовые приборы**, а не просто **Посуда**.  Выбирайте категории, которые описывают товар, а не абстрактный признак — например, **Духи**, а не **Подарки**.  Значение будет использовано для определения категории товара на Маркете в случае, если вы не передали категорию в параметре `marketCategoryId`. 
	Category *string `json:"category,omitempty"`
	// Ссылки на изображения товара. Изображение по первой ссылке считается основным, остальные дополнительными.  **Требования к ссылкам**  * Ссылок может быть до 30. * Указывайте ссылку целиком, включая протокол http или https. * Максимальная длина — 512 символов. * Русские буквы в URL можно. * Можно использовать прямые ссылки на изображения и на Яндекс Диск. Ссылки на Яндекс Диске нужно копировать с помощью функции **Поделиться**. Относительные ссылки и ссылки на другие облачные хранилища — не работают.  ✅ `https://example-shop.ru/images/sku12345.jpg`  ✅ `https://yadi.sk/i/NaBoRsimVOLov`  ❌ `/images/sku12345.jpg`  ❌ `https://www.dropbox.com/s/818f/tovar.jpg`  Ссылки на изображение должны быть постоянными. Нельзя использовать динамические ссылки, меняющиеся от выгрузки к выгрузке.  Если нужно заменить изображение, выложите новое изображение по новой ссылке, а ссылку на старое удалите. Если просто заменить изображение по старой ссылке, оно не обновится.  [Требования к изображениям](https://yandex.ru/support/marketplace/assortment/fields/images.html) 
	Pictures []string `json:"pictures,omitempty"`
	// Ссылка (URL) на видео товара.  Максимальное количество ссылок — 6.  **Требования к ссылке**  * Указывайте ссылку целиком, включая протокол http или https. * Максимальная длина — 512 символов. * Русские буквы в URL можно. * Можно использовать прямые ссылки на видео и на Яндекс Диск. Ссылки на Яндекс Диске нужно копировать с помощью функции **Поделиться**. Относительные ссылки и ссылки на другие облачные хранилища — не работают.  ✅ `https://example-shop.ru/video/sku12345.avi`  ✅ `https://yadi.sk/i/NaBoRsimVOLov`  ❌ `/video/sku12345.avi`  ❌ `https://www.dropbox.com/s/818f/super-tovar.avi`  Ссылки на видео должны быть постоянными. Нельзя использовать динамические ссылки, меняющиеся от выгрузки к выгрузке.  Если нужно заменить видео, выложите новое видео по новой ссылке, а ссылку на старое удалите. Если просто заменить видео по старой ссылке, оно не обновится.  [Требования к видео](https://yandex.ru/support/marketplace/assortment/fields/video.html) 
	Videos []string `json:"videos,omitempty"`
	// Список инструкций по использованию товара.  Максимальное количество инструкций — 6.  Если вы передадите пустое поле `manuals`, загруженные ранее инструкции удалятся. 
	Manuals []OfferManualDTO `json:"manuals,omitempty"`
	// Название бренда или производителя. Должно быть записано так, как его пишет сам бренд.
	Vendor *string `json:"vendor,omitempty"`
	// Указывайте в виде последовательности цифр. Подойдут коды EAN-13, EAN-8, UPC-A, UPC-E или Code 128.  Для книг указывайте ISBN.  Для товаров [определенных категорий и торговых марок](https://yastatic.net/s3/doc-binary/src/support/market/ru/yandex-market-list-for-gtin.xlsx) штрихкод должен быть действительным кодом GTIN. Обратите внимание: внутренние штрихкоды, начинающиеся на 2 или 02, и коды формата Code 128 не являются GTIN.  [Что такое GTIN](*gtin)  
	Barcodes []string `json:"barcodes,omitempty"`
	// Подробное описание товара: например, его преимущества и особенности.  Не давайте в описании инструкций по установке и сборке. Не используйте слова «скидка», «распродажа», «дешевый», «подарок» (кроме подарочных категорий), «бесплатно», «акция», «специальная цена», «новинка», «new», «аналог», «заказ», «хит». Не указывайте никакой контактной информации и не давайте ссылок.  Можно использовать теги:  * \\<h>, \\<h1>, \\<h2> и так далее — для заголовков; * \\<br> и \\<p> — для переноса строки; * \\<ol> — для нумерованного списка; * \\<ul> — для маркированного списка; * \\<li> — для создания элементов списка (должен находиться внутри \\<ol> или \\<ul>); * \\<div> — поддерживается, но не влияет на отображение текста.  Оптимальная длина — 400–600 символов, максимальная — 6000.  [Рекомендации и правила](https://yandex.ru/support/marketplace/assortment/fields/description.html) 
	Description *string `json:"description,omitempty"`
	// Страна, где был произведен товар.  Записывайте названия стран так, как они записаны в [списке](https://yastatic.net/s3/doc-binary/src/support/market/ru/countries.xlsx). 
	ManufacturerCountries []string `json:"manufacturerCountries,omitempty"`
	WeightDimensions *OfferWeightDimensionsDTO `json:"weightDimensions,omitempty"`
	// Артикул товара от производителя.
	VendorCode *string `json:"vendorCode,omitempty"`
	// Метки товара, используемые магазином. Покупателям теги не видны. По тегам можно группировать и фильтровать разные товары в каталоге — например, товары одной серии, коллекции или линейки.  Максимальная длина тега 20 символов. У одного товара может быть максимум 10 тегов. Всего можно создать не больше 50 разных тегов. 
	Tags []string `json:"tags,omitempty"`
	ShelfLife *TimePeriodDTO `json:"shelfLife,omitempty"`
	LifeTime *TimePeriodDTO `json:"lifeTime,omitempty"`
	GuaranteePeriod *TimePeriodDTO `json:"guaranteePeriod,omitempty"`
	// Код товара в единой Товарной номенклатуре внешнеэкономической деятельности (ТН ВЭД) — 10 или 14 цифр без пробелов.  Обязательно укажите, если он есть. 
	CustomsCommodityCode *string `json:"customsCommodityCode,omitempty"`
	// Номера документов на товар: сертификата, декларации соответствия и т. п.  Передавать можно только номера документов, сканы которого загружены в кабинете продавца по [инструкции](https://yandex.ru/support/marketplace/assortment/restrictions/certificates.html). 
	Certificates []string `json:"certificates,omitempty"`
	// Количество грузовых мест.  Параметр используется, если товар представляет собой несколько коробок, упаковок и так далее. Например, кондиционер занимает два места — внешний и внутренний блоки в двух коробках.  Для товаров, занимающих одно место, не передавайте этот параметр. 
	BoxCount *int32 `json:"boxCount,omitempty"`
	Condition *OfferConditionDTO `json:"condition,omitempty"`
	Type *OfferType `json:"type,omitempty"`
	// Признак цифрового товара. Укажите `true`, если товар доставляется по электронной почте.  [Как работать с цифровыми товарами](../../step-by-step/digital.md) 
	Downloadable *bool `json:"downloadable,omitempty"`
	// Параметр включает для товара пометку 18+. Устанавливайте ее только для товаров, которые относятся к удовлетворению сексуальных потребностей. 
	Adult *bool `json:"adult,omitempty"`
	Age *AgeDTO `json:"age,omitempty"`
	// {% note warning \"\" %}  Этот параметр устарел. При передаче характеристик используйте `parameterValues`.  {% endnote %}  Характеристики, которые есть только у товаров конкретной категории — например, диаметр колес велосипеда или материал подошвы обуви. 
	// Deprecated
	Params []OfferParamDTO `json:"params,omitempty"`
}

type _BaseOfferDTO BaseOfferDTO

// NewBaseOfferDTO instantiates a new BaseOfferDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseOfferDTO(offerId string) *BaseOfferDTO {
	this := BaseOfferDTO{}
	this.OfferId = offerId
	return &this
}

// NewBaseOfferDTOWithDefaults instantiates a new BaseOfferDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseOfferDTOWithDefaults() *BaseOfferDTO {
	this := BaseOfferDTO{}
	return &this
}

// GetOfferId returns the OfferId field value
func (o *BaseOfferDTO) GetOfferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetOfferIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferId, true
}

// SetOfferId sets field value
func (o *BaseOfferDTO) SetOfferId(v string) {
	o.OfferId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseOfferDTO) SetName(v string) {
	o.Name = &v
}

// GetMarketCategoryId returns the MarketCategoryId field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetMarketCategoryId() int64 {
	if o == nil || IsNil(o.MarketCategoryId) {
		var ret int64
		return ret
	}
	return *o.MarketCategoryId
}

// GetMarketCategoryIdOk returns a tuple with the MarketCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetMarketCategoryIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MarketCategoryId) {
		return nil, false
	}
	return o.MarketCategoryId, true
}

// HasMarketCategoryId returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasMarketCategoryId() bool {
	if o != nil && !IsNil(o.MarketCategoryId) {
		return true
	}

	return false
}

// SetMarketCategoryId gets a reference to the given int64 and assigns it to the MarketCategoryId field.
func (o *BaseOfferDTO) SetMarketCategoryId(v int64) {
	o.MarketCategoryId = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *BaseOfferDTO) SetCategory(v string) {
	o.Category = &v
}

// GetPictures returns the Pictures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseOfferDTO) GetPictures() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Pictures
}

// GetPicturesOk returns a tuple with the Pictures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseOfferDTO) GetPicturesOk() ([]string, bool) {
	if o == nil || IsNil(o.Pictures) {
		return nil, false
	}
	return o.Pictures, true
}

// HasPictures returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasPictures() bool {
	if o != nil && !IsNil(o.Pictures) {
		return true
	}

	return false
}

// SetPictures gets a reference to the given []string and assigns it to the Pictures field.
func (o *BaseOfferDTO) SetPictures(v []string) {
	o.Pictures = v
}

// GetVideos returns the Videos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseOfferDTO) GetVideos() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Videos
}

// GetVideosOk returns a tuple with the Videos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseOfferDTO) GetVideosOk() ([]string, bool) {
	if o == nil || IsNil(o.Videos) {
		return nil, false
	}
	return o.Videos, true
}

// HasVideos returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasVideos() bool {
	if o != nil && !IsNil(o.Videos) {
		return true
	}

	return false
}

// SetVideos gets a reference to the given []string and assigns it to the Videos field.
func (o *BaseOfferDTO) SetVideos(v []string) {
	o.Videos = v
}

// GetManuals returns the Manuals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseOfferDTO) GetManuals() []OfferManualDTO {
	if o == nil {
		var ret []OfferManualDTO
		return ret
	}
	return o.Manuals
}

// GetManualsOk returns a tuple with the Manuals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseOfferDTO) GetManualsOk() ([]OfferManualDTO, bool) {
	if o == nil || IsNil(o.Manuals) {
		return nil, false
	}
	return o.Manuals, true
}

// HasManuals returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasManuals() bool {
	if o != nil && !IsNil(o.Manuals) {
		return true
	}

	return false
}

// SetManuals gets a reference to the given []OfferManualDTO and assigns it to the Manuals field.
func (o *BaseOfferDTO) SetManuals(v []OfferManualDTO) {
	o.Manuals = v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *BaseOfferDTO) SetVendor(v string) {
	o.Vendor = &v
}

// GetBarcodes returns the Barcodes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseOfferDTO) GetBarcodes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Barcodes
}

// GetBarcodesOk returns a tuple with the Barcodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseOfferDTO) GetBarcodesOk() ([]string, bool) {
	if o == nil || IsNil(o.Barcodes) {
		return nil, false
	}
	return o.Barcodes, true
}

// HasBarcodes returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasBarcodes() bool {
	if o != nil && !IsNil(o.Barcodes) {
		return true
	}

	return false
}

// SetBarcodes gets a reference to the given []string and assigns it to the Barcodes field.
func (o *BaseOfferDTO) SetBarcodes(v []string) {
	o.Barcodes = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseOfferDTO) SetDescription(v string) {
	o.Description = &v
}

// GetManufacturerCountries returns the ManufacturerCountries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseOfferDTO) GetManufacturerCountries() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ManufacturerCountries
}

// GetManufacturerCountriesOk returns a tuple with the ManufacturerCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseOfferDTO) GetManufacturerCountriesOk() ([]string, bool) {
	if o == nil || IsNil(o.ManufacturerCountries) {
		return nil, false
	}
	return o.ManufacturerCountries, true
}

// HasManufacturerCountries returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasManufacturerCountries() bool {
	if o != nil && !IsNil(o.ManufacturerCountries) {
		return true
	}

	return false
}

// SetManufacturerCountries gets a reference to the given []string and assigns it to the ManufacturerCountries field.
func (o *BaseOfferDTO) SetManufacturerCountries(v []string) {
	o.ManufacturerCountries = v
}

// GetWeightDimensions returns the WeightDimensions field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetWeightDimensions() OfferWeightDimensionsDTO {
	if o == nil || IsNil(o.WeightDimensions) {
		var ret OfferWeightDimensionsDTO
		return ret
	}
	return *o.WeightDimensions
}

// GetWeightDimensionsOk returns a tuple with the WeightDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetWeightDimensionsOk() (*OfferWeightDimensionsDTO, bool) {
	if o == nil || IsNil(o.WeightDimensions) {
		return nil, false
	}
	return o.WeightDimensions, true
}

// HasWeightDimensions returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasWeightDimensions() bool {
	if o != nil && !IsNil(o.WeightDimensions) {
		return true
	}

	return false
}

// SetWeightDimensions gets a reference to the given OfferWeightDimensionsDTO and assigns it to the WeightDimensions field.
func (o *BaseOfferDTO) SetWeightDimensions(v OfferWeightDimensionsDTO) {
	o.WeightDimensions = &v
}

// GetVendorCode returns the VendorCode field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetVendorCode() string {
	if o == nil || IsNil(o.VendorCode) {
		var ret string
		return ret
	}
	return *o.VendorCode
}

// GetVendorCodeOk returns a tuple with the VendorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetVendorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VendorCode) {
		return nil, false
	}
	return o.VendorCode, true
}

// HasVendorCode returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasVendorCode() bool {
	if o != nil && !IsNil(o.VendorCode) {
		return true
	}

	return false
}

// SetVendorCode gets a reference to the given string and assigns it to the VendorCode field.
func (o *BaseOfferDTO) SetVendorCode(v string) {
	o.VendorCode = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseOfferDTO) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseOfferDTO) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *BaseOfferDTO) SetTags(v []string) {
	o.Tags = v
}

// GetShelfLife returns the ShelfLife field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetShelfLife() TimePeriodDTO {
	if o == nil || IsNil(o.ShelfLife) {
		var ret TimePeriodDTO
		return ret
	}
	return *o.ShelfLife
}

// GetShelfLifeOk returns a tuple with the ShelfLife field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetShelfLifeOk() (*TimePeriodDTO, bool) {
	if o == nil || IsNil(o.ShelfLife) {
		return nil, false
	}
	return o.ShelfLife, true
}

// HasShelfLife returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasShelfLife() bool {
	if o != nil && !IsNil(o.ShelfLife) {
		return true
	}

	return false
}

// SetShelfLife gets a reference to the given TimePeriodDTO and assigns it to the ShelfLife field.
func (o *BaseOfferDTO) SetShelfLife(v TimePeriodDTO) {
	o.ShelfLife = &v
}

// GetLifeTime returns the LifeTime field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetLifeTime() TimePeriodDTO {
	if o == nil || IsNil(o.LifeTime) {
		var ret TimePeriodDTO
		return ret
	}
	return *o.LifeTime
}

// GetLifeTimeOk returns a tuple with the LifeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetLifeTimeOk() (*TimePeriodDTO, bool) {
	if o == nil || IsNil(o.LifeTime) {
		return nil, false
	}
	return o.LifeTime, true
}

// HasLifeTime returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasLifeTime() bool {
	if o != nil && !IsNil(o.LifeTime) {
		return true
	}

	return false
}

// SetLifeTime gets a reference to the given TimePeriodDTO and assigns it to the LifeTime field.
func (o *BaseOfferDTO) SetLifeTime(v TimePeriodDTO) {
	o.LifeTime = &v
}

// GetGuaranteePeriod returns the GuaranteePeriod field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetGuaranteePeriod() TimePeriodDTO {
	if o == nil || IsNil(o.GuaranteePeriod) {
		var ret TimePeriodDTO
		return ret
	}
	return *o.GuaranteePeriod
}

// GetGuaranteePeriodOk returns a tuple with the GuaranteePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetGuaranteePeriodOk() (*TimePeriodDTO, bool) {
	if o == nil || IsNil(o.GuaranteePeriod) {
		return nil, false
	}
	return o.GuaranteePeriod, true
}

// HasGuaranteePeriod returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasGuaranteePeriod() bool {
	if o != nil && !IsNil(o.GuaranteePeriod) {
		return true
	}

	return false
}

// SetGuaranteePeriod gets a reference to the given TimePeriodDTO and assigns it to the GuaranteePeriod field.
func (o *BaseOfferDTO) SetGuaranteePeriod(v TimePeriodDTO) {
	o.GuaranteePeriod = &v
}

// GetCustomsCommodityCode returns the CustomsCommodityCode field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetCustomsCommodityCode() string {
	if o == nil || IsNil(o.CustomsCommodityCode) {
		var ret string
		return ret
	}
	return *o.CustomsCommodityCode
}

// GetCustomsCommodityCodeOk returns a tuple with the CustomsCommodityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetCustomsCommodityCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CustomsCommodityCode) {
		return nil, false
	}
	return o.CustomsCommodityCode, true
}

// HasCustomsCommodityCode returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasCustomsCommodityCode() bool {
	if o != nil && !IsNil(o.CustomsCommodityCode) {
		return true
	}

	return false
}

// SetCustomsCommodityCode gets a reference to the given string and assigns it to the CustomsCommodityCode field.
func (o *BaseOfferDTO) SetCustomsCommodityCode(v string) {
	o.CustomsCommodityCode = &v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseOfferDTO) GetCertificates() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseOfferDTO) GetCertificatesOk() ([]string, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []string and assigns it to the Certificates field.
func (o *BaseOfferDTO) SetCertificates(v []string) {
	o.Certificates = v
}

// GetBoxCount returns the BoxCount field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetBoxCount() int32 {
	if o == nil || IsNil(o.BoxCount) {
		var ret int32
		return ret
	}
	return *o.BoxCount
}

// GetBoxCountOk returns a tuple with the BoxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetBoxCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BoxCount) {
		return nil, false
	}
	return o.BoxCount, true
}

// HasBoxCount returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasBoxCount() bool {
	if o != nil && !IsNil(o.BoxCount) {
		return true
	}

	return false
}

// SetBoxCount gets a reference to the given int32 and assigns it to the BoxCount field.
func (o *BaseOfferDTO) SetBoxCount(v int32) {
	o.BoxCount = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetCondition() OfferConditionDTO {
	if o == nil || IsNil(o.Condition) {
		var ret OfferConditionDTO
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetConditionOk() (*OfferConditionDTO, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given OfferConditionDTO and assigns it to the Condition field.
func (o *BaseOfferDTO) SetCondition(v OfferConditionDTO) {
	o.Condition = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetType() OfferType {
	if o == nil || IsNil(o.Type) {
		var ret OfferType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetTypeOk() (*OfferType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given OfferType and assigns it to the Type field.
func (o *BaseOfferDTO) SetType(v OfferType) {
	o.Type = &v
}

// GetDownloadable returns the Downloadable field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetDownloadable() bool {
	if o == nil || IsNil(o.Downloadable) {
		var ret bool
		return ret
	}
	return *o.Downloadable
}

// GetDownloadableOk returns a tuple with the Downloadable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetDownloadableOk() (*bool, bool) {
	if o == nil || IsNil(o.Downloadable) {
		return nil, false
	}
	return o.Downloadable, true
}

// HasDownloadable returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasDownloadable() bool {
	if o != nil && !IsNil(o.Downloadable) {
		return true
	}

	return false
}

// SetDownloadable gets a reference to the given bool and assigns it to the Downloadable field.
func (o *BaseOfferDTO) SetDownloadable(v bool) {
	o.Downloadable = &v
}

// GetAdult returns the Adult field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetAdult() bool {
	if o == nil || IsNil(o.Adult) {
		var ret bool
		return ret
	}
	return *o.Adult
}

// GetAdultOk returns a tuple with the Adult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetAdultOk() (*bool, bool) {
	if o == nil || IsNil(o.Adult) {
		return nil, false
	}
	return o.Adult, true
}

// HasAdult returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasAdult() bool {
	if o != nil && !IsNil(o.Adult) {
		return true
	}

	return false
}

// SetAdult gets a reference to the given bool and assigns it to the Adult field.
func (o *BaseOfferDTO) SetAdult(v bool) {
	o.Adult = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *BaseOfferDTO) GetAge() AgeDTO {
	if o == nil || IsNil(o.Age) {
		var ret AgeDTO
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseOfferDTO) GetAgeOk() (*AgeDTO, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given AgeDTO and assigns it to the Age field.
func (o *BaseOfferDTO) SetAge(v AgeDTO) {
	o.Age = &v
}

// GetParams returns the Params field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *BaseOfferDTO) GetParams() []OfferParamDTO {
	if o == nil {
		var ret []OfferParamDTO
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *BaseOfferDTO) GetParamsOk() ([]OfferParamDTO, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *BaseOfferDTO) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given []OfferParamDTO and assigns it to the Params field.
// Deprecated
func (o *BaseOfferDTO) SetParams(v []OfferParamDTO) {
	o.Params = v
}

func (o BaseOfferDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseOfferDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["offerId"] = o.OfferId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.MarketCategoryId) {
		toSerialize["marketCategoryId"] = o.MarketCategoryId
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if o.Pictures != nil {
		toSerialize["pictures"] = o.Pictures
	}
	if o.Videos != nil {
		toSerialize["videos"] = o.Videos
	}
	if o.Manuals != nil {
		toSerialize["manuals"] = o.Manuals
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if o.Barcodes != nil {
		toSerialize["barcodes"] = o.Barcodes
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.ManufacturerCountries != nil {
		toSerialize["manufacturerCountries"] = o.ManufacturerCountries
	}
	if !IsNil(o.WeightDimensions) {
		toSerialize["weightDimensions"] = o.WeightDimensions
	}
	if !IsNil(o.VendorCode) {
		toSerialize["vendorCode"] = o.VendorCode
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.ShelfLife) {
		toSerialize["shelfLife"] = o.ShelfLife
	}
	if !IsNil(o.LifeTime) {
		toSerialize["lifeTime"] = o.LifeTime
	}
	if !IsNil(o.GuaranteePeriod) {
		toSerialize["guaranteePeriod"] = o.GuaranteePeriod
	}
	if !IsNil(o.CustomsCommodityCode) {
		toSerialize["customsCommodityCode"] = o.CustomsCommodityCode
	}
	if o.Certificates != nil {
		toSerialize["certificates"] = o.Certificates
	}
	if !IsNil(o.BoxCount) {
		toSerialize["boxCount"] = o.BoxCount
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Downloadable) {
		toSerialize["downloadable"] = o.Downloadable
	}
	if !IsNil(o.Adult) {
		toSerialize["adult"] = o.Adult
	}
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

func (o *BaseOfferDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"offerId",
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

	varBaseOfferDTO := _BaseOfferDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBaseOfferDTO)

	if err != nil {
		return err
	}

	*o = BaseOfferDTO(varBaseOfferDTO)

	return err
}

type NullableBaseOfferDTO struct {
	value *BaseOfferDTO
	isSet bool
}

func (v NullableBaseOfferDTO) Get() *BaseOfferDTO {
	return v.value
}

func (v *NullableBaseOfferDTO) Set(val *BaseOfferDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseOfferDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseOfferDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseOfferDTO(val *BaseOfferDTO) *NullableBaseOfferDTO {
	return &NullableBaseOfferDTO{value: val, isSet: true}
}

func (v NullableBaseOfferDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseOfferDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


