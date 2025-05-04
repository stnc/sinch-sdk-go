package sdk

import (
	"sinch/sdk/model"
	svc "sinch/sdk/services/sms/batches"
	"sinch/sdk/sms/batches"
	group "sinch/sdk/sms/groups"
)

func Clients(projectId string, clientId string, clientSecret string, region string) (client *model.Client) {
	return &model.Client{
		ClientId:     clientId,
		ProjectId:    projectId,
		ClientSecret: clientSecret,
		Region:       region,
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
