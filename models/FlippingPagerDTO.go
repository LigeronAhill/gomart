package models

// FlippingPagerDTO Модель для пагинации.
type FlippingPagerDTO struct {
	// Сколько всего найдено элементов.
	Total int `json:"total,omitempty"`
	// Начальный номер найденного элемента на странице.
	From int `json:"from,omitempty"`
	// Конечный номер найденного элемента на странице.
	To int `json:"to,omitempty"`
	// Текущая страница.
	CurrentPage int `json:"currentPage,omitempty"`
	// Общее количество страниц.
	PagesCount int `json:"pagesCount,omitempty"`
	// Размер страницы.
	PageSize int `json:"pageSize,omitempty"`
}
