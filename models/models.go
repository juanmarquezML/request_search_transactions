package models

import "time"

type Transaction struct {
	ID                string    `json:"id"`
	ProfileID         string    `json:"profile_id"`
	Acquirer          string    `json:"acquirer"`
	InternalProfileID string    `json:"internal_profile_id"`
	PaymentID         string    `json:"payment_id"`
	Status            string    `json:"status"`
	Operation         string    `json:"operation"`
	ProcessingMode    string    `json:"processing_mode"`
	CreationDate      time.Time `json:"creation_date"`
	LastModifiedDate  time.Time `json:"last_modified_date"`
}

type SearchResult struct {
	Data   []Transaction `json:"data"`
	Paging struct {
		Total  uint64 `json:"total"`
		Limit  uint64 `json:"limit"`
		Offset string `json:"offset"`
	} `json:"paging"`
}
