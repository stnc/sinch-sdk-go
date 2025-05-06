package main

import (
	"fmt"
	"log"
	"os"
	sdk "sinch/sdk/sms/batches"
	"sinch/sdk/sms/model"
)

func main() {

	sms, err := sdk.SinchClient(os.Getenv("ProjectId"), os.Getenv("ClientId"), os.Getenv("ClientSecret"), "us")

	if err != nil {
		log.Fatal(err)
	}

	data := &model.SendBatchRequest{
		Body: "Hello from Sinch! via golang sdk ",
		From: os.Getenv("FROM"),
		To:   []string{os.Getenv("TO")}}

	response, err_send := sms.Batches.Send(data)

	if err_send != nil {
		fmt.Println(err_send)
	} else {
		fmt.Println(response)
	}

	// sms.Batches.DryRun(data)
	// fmt.Println(sms.KeyId)

}
