package batches

import (
	"fmt"
)

// D:\go\sinch-sdk-go\SandBOX\0stnc\tutorials\v0.5-interface-perfect1\cmd\main.go
//usttekinden yapildi

type Calc struct {

}

type RequestHandler interface {
	Index()
	Send() string
}

type Client struct {
	ProjectId      string
	Batches RequestHandler
}

func (s *Calc) Index() {
	fmt.Println( " HELLO")
}

func (s *Calc) Send() string {
	return " SEND SMS "
}
