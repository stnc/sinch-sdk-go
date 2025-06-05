package batches

import (
	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"
)

type BatchesInterface interface {

	// Send  : Depending on the length of the body, one message might be split into multiple parts and charged accordingly.
	// Any groups targeted in a scheduled batch will be evaluated at the time of sending.
	// If a group is deleted between batch creation and scheduled date, it will be considered empty.
	Send(*batchesModel.SendBatchRequest) (batchesModel.SendBatchResponse, error)

	// List Batches  : With the list operation you can list batch messages created in the last 14 days that you have created. This operation supports pagination
	List() (batchesModel.ListBatchResponse, error)

	DryRun(*batchesModel.SendDryRunRequest) (batchesModel.SendDryRunResponse, error)
}

