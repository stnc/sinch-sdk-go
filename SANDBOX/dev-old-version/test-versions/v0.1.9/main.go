package main

import (
	sms "sinch/sdk/sms/batches"
)

func main() {



	var sms = &sms.Client{Batches: &sms.Init{ProjectId: "xxxx"}}

	// fmt.Println(sms.Batches.Send()) // sms.Batches.Send(data)
	sms.Batches.Index()

	//Sms.Inbounds.Get
	//Sms.Inbounds.List(request);
	//sms.batches.send

}
