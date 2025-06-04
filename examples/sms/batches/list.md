# SINCH sms.Batches.Send

Send a message or a batch of messages.

Depending on the length of the body, one message might be split into multiple parts and charged accordingly.

Any groups targeted in a scheduled batch will be evaluated at the time of sending. If a group is deleted between batch creation and scheduled date, it will be considered empty.

Be sure to use the correct region in the server URL.


| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**sms.Batches.List**] | **GET**  | Send a message or a batch of messages |


## `sms.Batches.List`

```golang

sms.Batches.List
```


### Example

```golang

package main

import (
	"encoding/json"
	"net/http"
	"github.com/stnc/sinch-sdk-go/sdk/core"
	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"
	 "github.com/stnc/sinch-sdk-go/sdk/model"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		fmt.Println("no env gotten")
	}
}

func main() {

	var result batchesModel.ListBatchResponse

	token := core.GetToken(s.C)

	url := core.ReplaceUrl(model.APIUrl+"/batches", s.C.ProjectId, s.C.Region)

	response, err1 := core.NewRequest(http.MethodGet, url, token.AccessToken, nil)

	if err1 != nil {
		return result, err1
	}

	if err1 := json.Unmarshal(response, &result); err1 != nil { // Parse []byte to go struct pointer
		return result, err1
	}

	return result, err1

}

```





[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)

[[Back to README]](../../README.md)




### Example with api [Golang Native ]


```golang


package main

import (
  "fmt"
  "net/http"
  "io"
)

func main() {

  url := "https://${region}.sms.api.sinch.com/xms/v1/${servicePlanId}/batches"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Authorization", "••••••")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}


```