package batches

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stnc/sinch-sdk-go/sdk/core"
	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"
	 "github.com/stnc/sinch-sdk-go/sdk/model"
)


// Send  : Depending on the length of the body, one message might be split into multiple parts and charged accordingly.
// Any groups targeted in a scheduled batch will be evaluated at the time of sending.
// If a group is deleted between batch creation and scheduled date, it will be considered empty.
func (s *BranchRepo) Send(req *batchesModel.SendBatchRequest) (batchesModel.SendBatchResponse, error) {
	var result batchesModel.SendBatchResponse

	token := core.GetToken(s.C)

	var url string = model.APIUrl + "/batches" 

	url = core.ReplaceUrl(url, s.C.ProjectId, s.C.Region)

	postBodyJson, err := json.Marshal(req)

	if err != nil {
		fmt.Println(err)
	}

	postBodyJsonBuffer := bytes.NewBuffer(postBodyJson)

	resp, err := http.NewRequest(http.MethodPost, url, postBodyJsonBuffer)

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
