package batches

import "time"

type ListBatchResponse struct {
	Count    int        `json:"count"`
	Page     int        `json:"page"`
	Batches  []LBatches `json:"batches"`
	PageSize int        `json:"page_size"`
}

type LBatches struct {
	ID             string    `json:"id"`
	To             []string  `json:"to"`
	From           string    `json:"from"`
	Canceled       bool      `json:"canceled"`
	Body           string    `json:"body"`
	Type           string    `json:"type"`
	CreatedAt      time.Time `json:"created_at"`
	ModifiedAt     time.Time `json:"modified_at"`
	DeliveryReport string    `json:"delivery_report"`
	ExpireAt       time.Time `json:"expire_at"`
	FlashMessage   bool      `json:"flash_message"`
}

// test - don't use this 
type ListBatchResponse2 struct {
	Count   int `json:"count"`
	Page    int `json:"page"`
	Batches []struct {
		ID             string    `json:"id"`
		To             []string  `json:"to"`
		From           string    `json:"from"`
		Canceled       bool      `json:"canceled"`
		Body           string    `json:"body"`
		Type           string    `json:"type"`
		CreatedAt      time.Time `json:"created_at"`
		ModifiedAt     time.Time `json:"modified_at"`
		DeliveryReport string    `json:"delivery_report"`
		ExpireAt       time.Time `json:"expire_at"`
		FlashMessage   bool      `json:"flash_message"`
	} `json:"batches"`
	PageSize int `json:"page_size"`
}
