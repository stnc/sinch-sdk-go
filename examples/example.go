package main

import (
	"fmt"
	"os"

	model "github.com/stnc/sinch-sdk-go/sdk/model"
	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"
	sdk "github.com/stnc/sinch-sdk-go/sdk/sms"

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
		Body:    "Hello from Sinch! via golang sdk ",
		From:    os.Getenv("FROM"),
		SmsType: model.SendTextBatchResponseTypeText, //or "mt_text",
		To:      []string{os.Getenv("TO")}}

		sms, err := sdk.Clients(os.Getenv("PROJECT_ID"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), os.Getenv("REGION"))

		if err != nil {
			panic(err)
		}
	

	response, err_send := sms.Batches.Send(data)

	if err_send != nil {
		fmt.Println(err_send)
	} else {
		fmt.Println(response)
	}

	// response, err_send := sms.Batches.List()

	// if err_send != nil {
	// 	fmt.Println(err_send)
	// } else {
	// 	fmt.Println(response)
	// }


	// a := sms.Batches.DryRun()

	// fmt.Println(a)

	// b := sms.Groups.Create()
	// fmt.Println(b)

}
