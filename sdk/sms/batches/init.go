package batches

import (
	"github.com/stnc/sinch-sdk-go/sdk/model"
)

func BatchesInit(c *model.Client) *Batches {
	return &Batches{c}
}


type Batches struct {
	C *model.Client
}
