package batches

import (
	"sinch/sdk/sms/model"
)

type BatchesInterface interface {

	//Depending on the length of the body, one message might be split into multiple parts and charged accordingly.
	//Any groups targeted in a scheduled batch will be evaluated at the time of sending.
	// If a group is deleted between batch creation and scheduled date, it will be considered empty.
	Send(*model.SendBatchRequest) (model.SendBatchResponse, error)

	//This operation will perform a dry run of a batch which calculates the bodies
	//and number of parts for all messages in the batch without actually sending any messages.
	// DryRun(*SendTextBatchRequest) (SendTextBatchResponse, error)
	//	DryRun(*model.SendBatchRequest)

	// List(string)
}

type BatchesService struct {
	Client *Client
	// request BatchesInterface
}

var _ BatchesInterface = &BatchesService{}

// func (f *BatchesService) Send(d *model.SendBatchRequest) (model.SendBatchResponse, error) {
// 	return f.request.Send(d)
// }

//  var _ BatchesInterface = (*BatchesService)(nil)
