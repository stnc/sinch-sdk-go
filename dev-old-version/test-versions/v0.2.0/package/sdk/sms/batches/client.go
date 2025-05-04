package batches

import (
	"sinch/sdk/sms/model"
)


type Client struct {
	UserAgent    string `json:"userAgent"`
	ProjectId    string `json:"projectId"` // or ServicePlanId
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`

	Region   string `json:"region"`
	PostBody model.SendBatchRequest
	Batches  BatchesInterface
	// Inbounds inbounds.InboundsInterface
}


func NewClient() (*Client, error) {
	c := &Client{UserAgent:model.UserAgent}

	c.Batches = &BatchesService{Client: c}

	return c, nil
}
