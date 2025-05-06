package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"maps"
	"net/http"

	"github.com/stnc/sinch-sdk-go/sdk/model"

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
func NewRequest(method, path string, token string, opt any) (result []byte, err error) {
	// var result []byte

	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	var body []byte
	// var err error
	var postBodyJsonBuffer *bytes.Buffer
	switch {
	case method == http.MethodPatch || method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		reqHeaders.Set("Authorization", "Bearer "+token)

		if opt != nil {
			body, err = json.Marshal(opt)
			if err != nil {
				return nil, err
			} else {
				postBodyJsonBuffer = bytes.NewBuffer(body)
			}
		}

	}

	resp, err := http.NewRequest(method, path, postBodyJsonBuffer)
	if err != nil {
		fmt.Println("NewRequest res", resp)
		return result, err
	}
	fmt.Println("req", resp)
	// Set the request specific headers.
	maps.Copy(resp.Header, reqHeaders)

	response, err_do := http.DefaultClient.Do(resp)
	fmt.Println("response", response)

	defer response.Body.Close()

	err_do = CheckResponse(response)

	if err_do != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further.
		return result, err_do
	}

	result, _ = io.ReadAll(response.Body)

	return result, nil
}
