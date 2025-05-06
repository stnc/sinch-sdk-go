package batches

import (
	"fmt"
)

type Client struct {
	ProjectId string
}

// func Batches() *Init {
// 	return &Init{}
// }

func (c *Client) Index() {
	fmt.Println(c.ProjectId + " HELLO")
}

func (c *Client) Send() string {
	return c.ProjectId + " SEND SMS "
}
