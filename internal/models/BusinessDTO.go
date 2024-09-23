package models

// BusinessDTO Информация о кабинете.
type BusinessDTO struct {
	// Идентификатор кабинета.
	Id int `json:"id,omitempty"`
	// Название бизнеса.
	Name string `json:"name,omitempty"`
}
