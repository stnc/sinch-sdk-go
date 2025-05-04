package batches

import (
	"fmt"
)

type Init struct{}

func Batches() *Init {
	return &Init{}
}

func (index *Init) Index() {
	fmt.Println("HELLO")
}

func (index *Init) Send() string {
	return "SEND SMS "
}
