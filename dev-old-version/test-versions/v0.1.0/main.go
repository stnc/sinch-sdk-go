package main

import (
	sms "sinch/sdk/sms/batches"
	// repository "avia/app/domain/repository"

	"fmt"

)





func main() {

	reminder := sms.Batches()
	fmt.Println(reminder.Send())
	reminder2 := sms.Batches().Send() // sms.Batches.Send(data)

	fmt.Println(reminder2)

	// http.HandleFunc("/list", reminder.Index)

	//Sms.Inbounds.Get
	//Sms.Inbounds.List(request);
	//sms.batches.send

}
