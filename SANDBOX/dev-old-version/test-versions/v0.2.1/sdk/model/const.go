package model

const (
	DefaultTokenUrl = "https://auth.sinch.com"
	APIUrl          = "https://zt.{Region}.sms.api.sinch.com/xms/v1/{ProjectId}"
	TokenUrl        = "oauth2/token"
	UserAgent       = "golang-sinch"
)

const (
	SendTextBatchResponseTypeText   string = "mt_text"
	SendTextBatchResponseTypeBinary string = "mt_binary"
	SendTextBatchResponseTypeMedia  string = "mt_media"
)

type Client struct {
	UserAgent    string `json:"userAgent"`
	ProjectId    string `json:"projectId"` // or ServicePlanId
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Region       string `json:"region"`
}
