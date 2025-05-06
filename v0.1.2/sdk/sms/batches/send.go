package batches

import (
	"fmt"
)



type Client struct {
	ProjectId string

}



func (c Client) Index() {
	fmt.Println(c.ProjectId + " HELLO")
}

func (c Client) Send() string {
	return c.ProjectId + " SEND SMS "
}
