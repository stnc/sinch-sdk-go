package main

import (
	"fmt"
	"os"

	// model "sinch/sdk/model"
	// batchesModel "sinch/sdk/model/sms/batches"
	sdk "sinch/sdk/sms"

	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		fmt.Println("no env gotten")
	}
}

type User struct {
	Name        string
	Age         int
	Active      bool
	lastLoginAt string
}

func main() {




	// data := &batchesModel.SendBatchRequest{
	// 	Body:    "Hello from Sinch! via golang sdk ",
	// 	From:    os.Getenv("FROM"),
	// 	SmsType: model.SendTextBatchResponseTypeText, //or "mt_text",
	// 	To:      []string{os.Getenv("TO")}}

	client := sdk.Clients(os.Getenv("PROJECT_ID"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REGION"))
	sms, err := sdk.Init(client)
	if err != nil {
		panic(err)
	}

	// response, err_send := sms.Batches.Send(data)

	// if err_send != nil {
	// 	fmt.Println(err_send)
	// } else {
	// 	fmt.Println(response)
	// }

	response, err_send := sms.Batches.List()

	if err_send != nil {
		fmt.Println(err_send)
	} else {
		fmt.Println(response)
	}


	a := sms.Batches.DryRun()

	fmt.Println(a)

	b := sms.Groups.Create()
	fmt.Println(b)

}
