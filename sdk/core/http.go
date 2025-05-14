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

	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	var resp *http.Request
	var postBodyJsonBuffer *bytes.Buffer
	var err1 error

	reqHeaders.Set("Content-Type", "application/json")

	reqHeaders.Set("Authorization", "Bearer "+token)

	switch {
	case method == http.MethodPatch || method == http.MethodPost || method == http.MethodPut:

		if opt != nil {
			var body []byte

			body, err1 = json.Marshal(opt)
			fmt.Println(" json.Marshal", body)
			if err1 != nil {
				return nil, err1
			} else {
				fmt.Println(" postBodyJsonBuffer ", body)
				postBodyJsonBuffer = bytes.NewBuffer(body)
			}
			resp, err = http.NewRequest(method, path, postBodyJsonBuffer)
			if err != nil {
				fmt.Println("NewRequest res", resp)
				return result, err
			}
			fmt.Println("req", resp)

		}
	default:
		fmt.Println("fetch with get mothod ", resp)
		resp, err1 = http.NewRequest(method, path, nil)
		if err1 != nil {
			fmt.Println("get err block  ", resp)
			return result, err1
		}
	}

<<<<<<< HEAD
	resp, err := http.NewRequest(method, path, postBodyJsonBuffer)
	if err != nil {
		return result, err
	}

	// Set the request specific headers.
	maps.Copy(resp.Header, reqHeaders)

	response, err_do := http.DefaultClient.Do(resp)
=======
	// Set the request specific headers.
	maps.Copy(resp.Header, reqHeaders)

	response, err2 := http.DefaultClient.Do(resp)
	if err != nil {
		return result, err2
	}

	fmt.Println("response", response)
>>>>>>> dev

	defer response.Body.Close()

	err = CheckResponse(response)

	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further.
		return result, err
	}

	result, _ = io.ReadAll(response.Body)

	return result, nil
}
