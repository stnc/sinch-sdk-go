package main

import (
	"fmt"
	"os"

<<<<<<< HEAD
	model "github.com/stnc/sinch-sdk-go/sdk/model"
	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"
	sdk "github.com/stnc/sinch-sdk-go/sdk/sms"
=======
	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"
	"github.com/stnc/sinch-sdk-go/sdk"
>>>>>>> dev

	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		fmt.Println("no env gotten")
	}
}

<<<<<<< HEAD

func main() {

	data := &batchesModel.SendBatchRequest{
		Body:    "Hello from Sinch! via golang sdk ",
		From:    os.Getenv("FROM"),
		SmsType: model.SendTextBatchResponseTypeText, //or "mt_text",
		To:      []string{os.Getenv("TO")}}

	client := sdk.Clients(os.Getenv("PROJECT_ID"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REGION"))
	sms, err := sdk.Init(client)
=======
func main() {

	data := &batchesModel.SendBatchRequest{
		Body: "Hello from Sinch! via golang sdk ",
		From: os.Getenv("FROM"),
		// SmsType: model.SendTextBatchResponseTypeText, //or "mt_text",
		SmsType: "mt_eetext",
		To:      []string{os.Getenv("TO")},
	}

	sms, err := sdk.Clients(os.Getenv("PROJECT_ID"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REGION"))

>>>>>>> dev
	if err != nil {
		panic(err)
	}

	response, err_send := sms.Batches.Send(data)

	if err_send != nil {
<<<<<<< HEAD
		fmt.Println(err_send)
	} else {
=======
		fmt.Println("err_send")
		fmt.Println(err_send)
	} else {
		fmt.Println("response main")
>>>>>>> dev
		fmt.Println(response)
	}

	// response, err_send := sms.Batches.List()

	// if err_send != nil {
	// 	fmt.Println(err_send)
	// } else {
	// 	fmt.Println(response)
	// }

<<<<<<< HEAD

=======
>>>>>>> dev
	// a := sms.Batches.DryRun()

	// fmt.Println(a)

	// b := sms.Groups.Create()
	// fmt.Println(b)

}
