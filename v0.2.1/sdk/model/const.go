package model
const (
	DefaultBaseURL = "https://auth.sinch.com"
	TokenUrl       = "oauth2/token"
	UserAgent      = "golang-sinch"
)

type ResponseTypeValue string

const (
	SendTextBatchResponseTypeText   ResponseTypeValue = "mt_text"
	SendTextBatchResponseTypeBinary ResponseTypeValue = "mt_binary"
	SendTextBatchResponseTypeMedia  ResponseTypeValue = "mt_media"
)


type Client struct {
	UserAgent    string `json:"userAgent"`
	ProjectId    string `json:"projectId"` // or ServicePlanId
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Region       string `json:"region"`
}
