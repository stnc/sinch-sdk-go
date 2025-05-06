package batches

import (
	"context"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	// "sinch/sdk/sms/services"
	"sinch/sdk/sms/model"

)

func GetToken(s *BatchesService) *oauth2.Token {

	config := clientcredentials.Config{
		ClientID:     s.Client.ClientId,
		ClientSecret: s.Client.ClientSecret,
		TokenURL:     model.DefaultBaseURL + "/" + model.TokenUrl,
		Scopes:       []string{""},
	}

	token, err := config.Token(context.Background())

	if err != nil {
		log.Fatal(err.Error())
	}

	return token
}

/*
func GetHttp(token string, url string, postBody *SendTextBatchRequest) (SendTextBatchResponse, error) {

	postBodyJson, err := json.Marshal(postBody)

	if err != nil {
		fmt.Println(err)
	}

	postBodyJsonBuffer := bytes.NewBuffer(postBodyJson)

	resp, err := http.NewRequest(http.MethodPost, url, postBodyJsonBuffer)

	resp.Header.Add("Content-Type", "application/json")

	resp.Header.Add("Authorization", "Bearer "+token)

	response, err_do := http.DefaultClient.Do(resp)

	defer response.Body.Close()

	err_do = CheckResponse(response)

	var result SendTextBatchResponse
	//Handle Error
	// TODO:  is it path right ?
	if err_do != nil {
		return result, err_do
	}

	body2, _ := io.ReadAll(response.Body)

	if err := json.Unmarshal(body2, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(PrettyPrint(result))

	return result, err

}
*/

// if resp.StatusCode == http.StatusUnauthorized && c.authType == BasicAuth { //  http.StatusUnauthorized - 401
