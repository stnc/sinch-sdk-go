package core

import (
	"context"
	"log"
	"net/http"
	"net/url"
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
func  NewRequest(method, path string, token string) (*http.Request, error) {




	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")


	var body interface{}
	switch {
	case method == http.MethodPatch || method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		reqHeaders.Add("Authorization", "Bearer "+token)
	}

	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}



	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}