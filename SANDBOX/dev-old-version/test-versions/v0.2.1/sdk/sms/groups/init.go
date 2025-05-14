package groups

import (
	"github.com/stnc/sinch-sdk-go/sdk/model"
)

func GroupsInit(c *model.Client) *Groups {
	return &Groups{c}
}


// BranchRepo struct
type Groups struct {
	C *model.Client
}
