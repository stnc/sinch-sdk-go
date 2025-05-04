package main

import (
	"fmt"
	"os"
	sdk "sinch/sdk/sms"
)

func main() {

	// data := &model.SendBatchRequest{
	// 	Body: "Hello from Sinch! via golang sdk ",
	// 	From: os.Getenv("FROM"),
	// 	To:   []string{os.Getenv("TO")}}

	db := sdk.Clients(os.Getenv("ProjectId"), os.Getenv("ClientId"), os.Getenv("ClientSecret"), "us")
	sms, err := sdk.Init(db)
	if err != nil {
		panic(err)
	}
	a := sms.Batches.DryRun()
	fmt.Println(a)

	b := sms.Groups.Create()
	fmt.Println(b)



}
