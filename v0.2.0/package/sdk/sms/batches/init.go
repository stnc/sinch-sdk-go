package batches



func SinchClient(projectId string, clientId string, clientSecret string, region string) (*Client, error) {
	client, err :=NewClient()
	client.ClientId = clientId
	client.ProjectId = projectId
	client.ClientSecret = clientSecret
	client.Region = region
	if err != nil {
		return nil, err
	}
	return client, nil
}

