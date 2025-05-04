package batches

import (
	"encoding/json"

	"io"
	"net/http"

	"sinch/sdk/core"

	batchesModel "sinch/sdk/model/sms/batches"
	 "sinch/sdk/model"
)

// List Batches  : With the list operation you can list batch messages created in the last 14 days that you have created. This operation supports pagination
func (s *BranchRepo) List() (batchesModel.ListBatchResponse, error) {

	var result batchesModel.ListBatchResponse

	token := core.GetToken(s.C)

	var url string = model.APIUrl + "/batches" 

	url = core.ReplaceUrl(url, s.C.ProjectId, s.C.Region)

	resp, err := http.NewRequest(http.MethodGet, url, nil)

	resp.Header.Add("Content-Type", "application/json")

	resp.Header.Add("Authorization", "Bearer "+token.AccessToken)

	response, err_do := http.DefaultClient.Do(resp)

	defer response.Body.Close()

	err_do = core.CheckResponse(response)

	//Handle Error
	// TODO:  is it path right ?
	if err_do != nil {
		return result, err_do
	}

	body2, _ := io.ReadAll(response.Body)

	if err := json.Unmarshal(body2, &result); err != nil { // Parse []byte to go struct pointer
		return result, err
	}

	return result, err
}
