package sdk

import (
	"github.com/stnc/sinch-sdk-go/sdk/model"
	svc "github.com/stnc/sinch-sdk-go/sdk/services/sms/batches"
	"github.com/stnc/sinch-sdk-go/sdk/sms/batches"
	group "github.com/stnc/sinch-sdk-go/sdk/sms/groups"
)


type Repositories struct {
	Batches svc.BatchesInterface
	Groups  svc.GroupsInterface
	Client  *model.Client
}

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


func Init(c *model.Client) (*Repositories, error) {

	return &Repositories{

		Batches: batches.BatchesRepositoryInit(c),
		Groups:  group.GroupsRepositoryInit(c),

		Client: c,
	}, nil
}
