package sdk

import (
	"sinch/sdk/model"
	svc "sinch/sdk/services/sms/batches"
	group "sinch/sdk/sms/groups"
    "sinch/sdk/sms/batches"
)

func Clients(projectId string, clientId string, clientSecret string, region string) (client *model.Client ){

	return &model.Client{
		ClientId:     "Anikaa",
		ProjectId:    "Ballia",
		ClientSecret: "277001",
		Region:       "ddd",
	}

	// var client *model.Client
	// client.ClientId = clientId
	// client.ProjectId = projectId
	// client.ClientSecret = clientSecret
	// client.Region = region
	// return client
}

type Repositories struct {
	Batches svc.BatchesInterface
	Groups  svc.GroupsInterface
	Client  *model.Client
}


func Init(c *model.Client) (*Repositories, error) {

	return &Repositories{

		Batches: batches.BatchesRepositoryInit(c),
		Groups:  group.GroupsRepositoryInit(c),

		Client: c,
	}, nil
}
