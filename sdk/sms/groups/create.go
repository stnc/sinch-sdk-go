package groups

import (
	"sinch/sdk/model"
)

func GroupsRepositoryInit(c *model.Client) *GroupsRepo {
	return &GroupsRepo{c}
}

func (s *GroupsRepo) Create() string {
	return " Create data " + s.C.ClientId
}

// BranchRepo struct
type GroupsRepo struct {
	C *model.Client
}
