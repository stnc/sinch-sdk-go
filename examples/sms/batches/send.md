# SINCH sms.Batches.Send

Send a message or a batch of messages.

Depending on the length of the body, one message might be split into multiple parts and charged accordingly.

Any groups targeted in a scheduled batch will be evaluated at the time of sending. If a group is deleted between batch creation and scheduled date, it will be considered empty.

Be sure to use the correct region in the server URL.


| Method | HTTP request | Description |
| ------------- | ------------- | ------------- |
| [**sms.Batches.Send**] | **POST**  | Send a message or a batch of messages |


## `sms.Batches.Send`

```golang
sms.Batches.Send
```


### Example

```golang
package main

import (
	"fmt"
	"os"

	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"
	"github.com/stnc/sinch-sdk-go/sdk"
	"github.com/stnc/sinch-sdk-go/sdk/model"
	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		fmt.Println("no env gotten")
	}
}

func main() {

	data := &batchesModel.SendBatchRequest{
		Body: "Hello from Sinch! via golang sdk ",
		From: os.Getenv("FROM"),
		SmsType: model.SendTextBatchResponseTypeText, //or "mt_text",
		To:      []string{os.Getenv("TO")},
	}

	sms, err := sdk.Clients(os.Getenv("PROJECT_ID"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REGION"))

	if err != nil {
		panic(err)
	}

	response, err_send := sms.Batches.Send(data)

	if err_send != nil {
		fmt.Println("err_send")
		fmt.Println(err_send)
	} else {
		fmt.Println("response main")
		fmt.Println(response)
	}

}

```

### Parameters

```golang


	data := &batchesModel.SendBatchRequest{
		Body: "Hello from Sinch! via golang sdk ",
		From: os.Getenv("FROM"),
	    SmsType: model.SendTextBatchResponseTypeText, //or "mt_text",

		To:      []string{os.Getenv("TO")},
	}


```



[[Back to top]](#) [[Back to API list]](../../README.md#endpoints)

[[Back to README]](../../README.md)


### Example with api [Golang Native]

```golang

package main

import (
  "fmt"
  "strings"
  "net/http"
  "io"
)

func main() {

  url := "https://${region}.sms.api.sinch.com/xms/v1/${servicePlanId}/batches"
  method := "POST"

  payload := strings.NewReader(`{`+"
"+`
    "from": "YOUR_Sinch_virtual_number",`+"
"+`
    "body": "Hello from Sinch! via go native api ",`+"
"+`
    "to": [`+"
"+`
        "YOUR_recipient_number"`+"
"+`
    ]`+"
"+`
}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")
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

