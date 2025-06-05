package main

import (
	"fmt"
	"os"

	"github.com/stnc/sinch-sdk-go/sdk"
	"github.com/stnc/sinch-sdk-go/sdk/model"
	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"

	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		fmt.Println("no env gotten")
	}
}

func main() {
	// send()
	// list()
	dryrun()

}
func send() {
	data := &batchesModel.SendBatchRequest{
		Body:    "Hello from Sinch! via golang sdk ",
		From:    os.Getenv("FROM"),
		SmsType: model.SendTextBatchResponseTypeText, //or "mt_text",
		// SmsType: "mt_eetext",
		To: []string{os.Getenv("TO")},
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
func list() {
	sms, err := sdk.Clients(os.Getenv("PROJECT_ID"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REGION"))

	if err != nil {
		panic(err)
	}

	response, err_send := sms.Batches.List()

	if err_send != nil {
		fmt.Println(err_send)
	} else {
		fmt.Println(response)
	}

}

func dryrun() {
	sms, err := sdk.Clients(os.Getenv("PROJECT_ID"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REGION"))

	if err != nil {
		panic(err)
	}

	data := &batchesModel.SendDryRunRequest{
		Body:         "Hello from Sinch! via golang sdk ",
		From:         os.Getenv("FROM"),
		SmsType:      model.SendTextBatchResponseTypeText, //or "mt_text",
		To:           []string{os.Getenv("TO")},
		// PerRecipient: true,
	}

	response, err_send := sms.Batches.DryRun(data)

	if err_send != nil {
		fmt.Println(err_send)
	} else {
		fmt.Println(response)
	}

}
