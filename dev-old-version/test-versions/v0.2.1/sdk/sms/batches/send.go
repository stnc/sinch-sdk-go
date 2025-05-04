package batches


import (

	 "sinch/sdk/model"
)


func BatchesRepositoryInit(c *model.Client) *BranchRepo {
	return &BranchRepo{c}
}


func (s *BranchRepo) Send() string {
	return " SEND SMS " + s.C.ClientId
}
// BranchRepo struct
type BranchRepo struct {
	C *model.Client
}

