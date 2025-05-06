package model

// import "sdk/sms/inbounds"

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

