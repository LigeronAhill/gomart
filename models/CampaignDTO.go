package models

// CampaignDTO Информация о магазине.
type CampaignDTO struct {
	// URL магазина.
	Domain string `json:"domain,omitempty"`
	// Идентификатор кампании.
	Id int `json:"id,omitempty"`
	// Идентификатор плательщика в Яндекс Балансе.
	ClientId      int          `json:"clientId,omitempty"`
	Business      *BusinessDTO `json:"business,omitempty"`
	PlacementType string       `json:"placementType,omitempty"`
}
