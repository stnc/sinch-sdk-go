package main

import (
	sms "sinch/sdk/sms/batches"
	// repository "avia/app/domain/repository"

	"fmt"
)

func main() {

	sms := sms.Client{ProjectId: "xxxx"}

	fmt.Println(sms.Send()) // sms.Batches.Send(data)


	//Sms.Inbounds.Get
	//Sms.Inbounds.List(request);
	//sms.batches.send

}
