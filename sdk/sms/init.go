package sdk

import (
	"github.com/stnc/sinch-sdk-go/sdk/model"
	svc "github.com/stnc/sinch-sdk-go/sdk/services/sms/batches"
	"github.com/stnc/sinch-sdk-go/sdk/sms/batches"
	group "github.com/stnc/sinch-sdk-go/sdk/sms/groups"
)

type Loader struct {
	Batches svc.BatchesInterface
	Groups  svc.GroupsInterface
	Client  *model.Client
}

func Clients(projectId string, clientId string, clientSecret string, region string) (*Loader, error) {
	c := &model.Client{
		ClientId:     clientId,
		ProjectId:    projectId,
		ClientSecret: clientSecret,
		Region:       region,
	}

	return &Loader{
		Batches: batches.BatchesRepositoryInit(c),
		Groups:  group.GroupsRepositoryInit(c),
		Client:  c,
	}, nil
}