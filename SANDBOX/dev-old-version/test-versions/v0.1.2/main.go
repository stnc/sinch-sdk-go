package main

import (
	sms "sinch/sdk/sms/batches"
	servce "sinch/sdk/services/sms/batches"

	"fmt"
)

func main() {



	var sms servce.BatchesInterface = &sms.Client{ProjectId: "xxxx"}


	fmt.Println(sms.Send()) // sms.Batches.Send(data)

	//Sms.Inbounds.Get
	//Sms.Inbounds.List(request);
	//sms.batches.send

}


