package batches

import (
	"github.com/stnc/sinch-sdk-go/sdk/model"
)

func BatchesRepositoryInit(c *model.Client) *BranchRepo {
	return &BranchRepo{c}
}

// BranchRepo struct
type BranchRepo struct {
	C *model.Client
}
