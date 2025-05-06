package batches

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stnc/sinch-sdk-go/sdk/core"
	"github.com/stnc/sinch-sdk-go/sdk/model"
	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"
)

// Send  : Depending on the length of the body, one message might be split into multiple parts and charged accordingly.
// Any groups targeted in a scheduled batch will be evaluated at the time of sending.
// If a group is deleted between batch creation and scheduled date, it will be considered empty.
func (s *BranchRepo) SendOLD(req *batchesModel.SendBatchRequest) (batchesModel.SendBatchResponse, error) {
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
	fmt.Println("err_do", err_do)
	//Handle Error
	// TODO:  is it path right ?
	if err_do != nil {
		fmt.Println("result", result)
		return result, err_do
	}

	body2, _ := io.ReadAll(response.Body)

	if err := json.Unmarshal(body2, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("result2", result)
		return result, err
	}
	fmt.Println(" last result2", result)
	return result, err

}

// Send  : Depending on the length of the body, one message might be split into multiple parts and charged accordingly.
// Any groups targeted in a scheduled batch will be evaluated at the time of sending.
// If a group is deleted between batch creation and scheduled date, it will be considered empty.
func (s *BranchRepo) Send(req *batchesModel.SendBatchRequest) (batchesModel.SendBatchResponse, error) {
	var result batchesModel.SendBatchResponse
	token := core.GetToken(s.C)

	url := core.ReplaceUrl(model.APIUrl + "/batches", s.C.ProjectId, s.C.Region)
	
	response1, err := core.NewRequest( http.MethodPost, url, token.AccessToken, req)

	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(response1, &result); err != nil { // Parse []byte to go struct pointer
		return result, err
	}
	return result, err

}
