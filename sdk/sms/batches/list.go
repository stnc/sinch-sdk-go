package batches

import (
	"encoding/json"


	"net/http"

	"github.com/stnc/sinch-sdk-go/sdk/core"

	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"
	 "github.com/stnc/sinch-sdk-go/sdk/model"
)

// List Batches  : With the list operation you can list batch messages created in the last 14 days that you have created. This operation supports pagination
func (s *Batches) List() (batchesModel.ListBatchResponse, error) {
<<<<<<< HEAD


	var result1 batchesModel.ListBatchResponse

	token1 := core.GetToken(s.C)

	url1 := core.ReplaceUrl(model.APIUrl+"/batches", s.C.ProjectId, s.C.Region)

	response1, err1 := core.NewRequest(http.MethodGet, url1, token1.AccessToken, nil)
=======

	var result batchesModel.ListBatchResponse

	token := core.GetToken(s.C)

	url := core.ReplaceUrl(model.APIUrl+"/batches", s.C.ProjectId, s.C.Region)

	response, err1 := core.NewRequest(http.MethodGet, url, token.AccessToken, nil)
>>>>>>> dev

	if err1 != nil {
		return result, err1
	}

	if err1 := json.Unmarshal(response, &result); err1 != nil { // Parse []byte to go struct pointer
		return result, err1
	}

	return result, err1
}
