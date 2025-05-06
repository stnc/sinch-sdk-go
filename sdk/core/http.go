package core

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"maps"
	"net/http"

	"github.com/stnc/sinch-sdk-go/sdk/model"
	"github.com/stnc/sinch-sdk-go/sdk/core"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func GetToken(s *model.Client) *oauth2.Token {
	config := clientcredentials.Config{
		ClientID:     s.ClientId,
		ClientSecret: s.ClientSecret,
		TokenURL:     model.DefaultTokenUrl + "/" + model.TokenUrl,
		Scopes:       []string{""},
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	return token
}

// NewRequest creates a new API request. The method expects a relative URL
// path that will be resolved relative to the base URL of the Client.
// Relative URL paths should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included
// as the request body.
func NewRequest(method, path string, token string, opt any) (*http.Request, error) {

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	var body []byte
	var err error
	var postBodyJsonBuffer *bytes.Buffer
	switch {
	case method == http.MethodPatch || method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		reqHeaders.Add("Authorization", "Bearer "+token)

		if opt != nil {
			body, err = json.Marshal(opt)
			if err != nil {
				return nil, err
			} else {
				postBodyJsonBuffer = bytes.NewBuffer(body)
			}
		}

	}

	req, err := http.NewRequest(method, path, postBodyJsonBuffer)
	if err != nil {
		return nil, err
	}

	// Set the request specific headers.
	maps.Copy(req.Header, reqHeaders)

	return req, nil
}



func  Do(req *http.Request, v any) (  any, error) {
response, err_do := http.DefaultClient.Do(req)

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
