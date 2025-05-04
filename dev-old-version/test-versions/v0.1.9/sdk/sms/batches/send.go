package batches

import (
	"fmt"
	svc "sinch/sdk/services/sms/batches"

)

// D:\go\sinch-sdk-go\SandBOX\0stnc\tutorials\v0.5-interface-perfect1\cmd\main.go
//usttekinden yapildi

type Init struct {
	ProjectId string `json:"projectId"` // or ServicePlanId
}



type Client struct {
	Batches svc.BatchesInterface
}

func (s *Init) Index() {
	fmt.Println(" HELLO")
	fmt.Println(s.ProjectId)
}

func (s *Init) Send() string {
	return " SEND SMS " + s.ProjectId
}
