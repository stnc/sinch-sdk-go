package batches

import (
	"github.com/stnc/sinch-sdk-go/sdk/model"
)

<<<<<<< HEAD
func BatchesRepositoryInit(c *model.Client) *BranchRepo {
	return &BranchRepo{c}
}

// BranchRepo struct
type BranchRepo struct {
=======
func BatchesInit(c *model.Client) *Batches {
	return &Batches{c}
}


type Batches struct {
>>>>>>> dev
	C *model.Client
}
