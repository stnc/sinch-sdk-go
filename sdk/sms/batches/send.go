package batches

import (

	"encoding/json"

	"net/http"

	"github.com/stnc/sinch-sdk-go/sdk/core"
	"github.com/stnc/sinch-sdk-go/sdk/model"
	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"
)



// Send  : Depending on the length of the body, one message might be split into multiple parts and charged accordingly.
// Any groups targeted in a scheduled batch will be evaluated at the time of sending.
// If a group is deleted between batch creation and scheduled date, it will be considered empty.
func (s *Batches) Send(req *batchesModel.SendBatchRequest) (batchesModel.SendBatchResponse, error) {
	var result batchesModel.SendBatchResponse
	token := core.GetToken(s.C)

	url := core.ReplaceUrl(model.APIUrl+"/batches", s.C.ProjectId, s.C.Region)

	response1, err := core.NewRequest(http.MethodPost, url, token.AccessToken, req)

	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(response1, &result); err != nil { // Parse []byte to go struct pointer
		return result, err
	}
	return result, err
}
