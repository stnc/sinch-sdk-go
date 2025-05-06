package batches

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"sinch/sdk/sms/model"
	"sinch/sdk/sms/core"



)

// Depending on the length of the body, one message might be split into multiple parts and charged accordingly.
// Any groups targeted in a scheduled batch will be evaluated at the time of sending.
// If a group is deleted between batch creation and scheduled date, it will be considered empty.
func (s *BatchesService) Send(opt *model.SendBatchRequest) (model.SendBatchResponse, error) {
	// token := GetToken(s)
	token := `xxxxx`
	url := "https://zt.us.sms.api.sinch.com/xms/v1/" + s.Client.ProjectId + "/batches"

	postBodyJson, err := json.Marshal(opt)

	if err != nil {
		fmt.Println(err)
	}

	postBodyJsonBuffer := bytes.NewBuffer(postBodyJson)

	resp, err := http.NewRequest(http.MethodPost, url, postBodyJsonBuffer)

	resp.Header.Add("Content-Type", "application/json")

	resp.Header.Add("Authorization", "Bearer "+token)

	response, err_do := http.DefaultClient.Do(resp)

	defer response.Body.Close()

	err_do = core.CheckResponse(response)

	var result model.SendBatchResponse
	//Handle Error
	// TODO:  is it path right ?
	if err_do != nil {
		return result, err_do
	}

	body2, _ := io.ReadAll(response.Body)

	if err := json.Unmarshal(body2, &result); err != nil { // Parse []byte to go struct pointer
		return result, err_do
	}
	// fmt.Println(PrettyPrint(result))

	return result, err

	// return s.GetHttp(token.AccessToken, opt)
}

// 	fmt.Println("HTTP Response Status:", response.Status) HTTP Response Status: 201 Created
