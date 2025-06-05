package batches

type SendDryRunRequest struct {
	// PerRecipient            bool   `json:"per_recipient"`
	// NumberOfRecipients      int    `json:"number_of_recipients"`
	To      []string `json:"to"` //required
	From                    string `json:"from"`
	Body                    string `json:"body"`
	SmsType                    string `json:"type"`
	// Udh                     string `json:"udh"`
	// DeliveryReport          string `json:"delivery_report"`

	SendAt                  string `json:"send_at"`
	ExpireAt                string `json:"expire_at"`
	// CallbackURL             string `json:"callback_url"`
	ClientReference         string `json:"client_reference"`
	FlashMessage            bool   `json:"flash_message"`
	// MaxNumberOfMessageParts int    `json:"max_number_of_message_parts"`
	// Parameters struct {
	// 	Name struct {
	// 		Phone   string `json:"phone"`
	// 		Default string `json:"default"`
	// 	} `json:"name"`
	// } `json:"parameters"`
}

type SendDryRunResponse struct {
	NumberOfRecipients int `json:"number_of_recipients"`
	NumberOfMessages   int `json:"number_of_messages"`
	PerRecipient []PerRecipient `json:"per_recipient"`
}

type PerRecipient struct {
	Recipient     string `json:"recipient"`
	NumberOfParts int    `json:"number_of_parts"`
	Body          string `json:"body"`
	Encoding      string `json:"encoding"`
}
