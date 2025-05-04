package main

import (
	sms "sinch/sdk/sms/batches"

	"fmt"
)

func main() {



	var sms = &sms.Client{ProjectId: "xxxx", Batches: &sms.Calc{}}

	fmt.Println(sms.Batches.Send()) // sms.Batches.Send(data)

	//Sms.Inbounds.Get
	//Sms.Inbounds.List(request);
	//sms.batches.send

}
