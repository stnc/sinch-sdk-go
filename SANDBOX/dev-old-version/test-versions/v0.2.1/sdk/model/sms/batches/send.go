package batches



type SendBatchRequest struct {
	From    string   `json:"from"` // required
	Body    string   `json:"body"`
	To      []string `json:"to"` //required
	SmsType string   `json:"type"`
	/*
		CallbackUrl             string `json:"callback_url "`
		ClientReference         string `json:"client_reference "`
		FeedbackEnabled         bool   `json:"feedback_enabled "`
		FlashMessage            bool   `json:"flash_message "`
		TruncateConcat          bool   `json:"truncate_concat  "`
		MaxNumberOfMessageParts int32  `json:"max_number_of_message_Parts"`
		FromTon                 int32  `json:"from_ton"`
		FromNpi                 int32  `json:"from_npi"`
		ExpireAt                string `json:"expire_at"`
		SendAt                  string `json:"send_at"`
		Delivery_report // //TODO https://developers.sinch.com/docs/sms/api-reference/sms/tag/Batches/#tag/Batches/operation/SendSMS
	*/
}

type SendBatchResponse struct {
	ID             string   `json:"id"`
	To             []string `json:"to"`
	Body           string   `json:"body"`
	From           string   `json:"from"`
	Type           string   `json:"type"` //required
	ResponseType   string   `json:"ResponseType"`
	CreatedAt      string   `json:"created_at"`
	ModifiedAt     string   `json:"modified_at"`
	DeliveryReport string   `json:"delivery_report"`
	ExpireAt       string   `json:"expire_at"`
	FlashMessage   bool     `json:"flash_message"`
	Canceled       bool     `json:"canceled"`
	/*

		CallbackUrl             string `json:"callback_url "`
		ClientReference         string `json:"client_reference "`
		SendAt                  string `json:"send_at"`
		FeedbackEnabled         bool   `json:"feedback_enabled "`
		StrictValidation         bool   `json:"strict_validation "`
	*/
}
