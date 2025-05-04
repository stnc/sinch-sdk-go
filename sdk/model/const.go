package model
const (
	DefaultTokenUrl = "https://auth.sinch.com"
	APIUrl = "https://zt.{Region}.sms.api.sinch.com/xms/v1/{ProjectId}"
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
