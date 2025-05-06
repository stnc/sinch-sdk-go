package batches

//This operation will perform a dry run of a batch which calculates the bodies
//and number of parts for all messages in the batch without actually sending any messages.
/*
func (s *services.BatchesService) DryRun(opt *model.SendBatchRequest) {
	// func (s *BatchesService) DryRun(opt *SendTextBatchRequest) (SendTextBatchResponse, error) {
	//https://developers.sinch.com/docs/sms/api-reference/sms/tag/Batches/#tag/Batches/operation/Dry_Run
	// token := GetToken(s)
	//token := `eyJhbGciOiJSUzI1NiIsImtpZCI6InB1YmxpYzpkMzc5NzhhMC1hODNlLTRjMTAtYTVmMC02NGVhMzYxMzM2YjUiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOltdLCJjbGllbnRfaWQiOiJkZjc0MTU3Yi1mNGNjLTQ4NmItYmY1OS1jN2M5ZDVmNzRkMjUiLCJleHAiOjE3NDUzNTA5MDIsImV4dCI6e30sImlhdCI6MTc0NTM0NzMwMiwiaXNzIjoiaHR0cHM6Ly9hdXRoLnNpbmNoLmNvbS8iLCJqdGkiOiIxNjI5MjE0NS1jMjc1LTQ0MjgtYWJhNC02N2M3ODhiN2FiYWUiLCJuYmYiOjE3NDUzNDczMDIsInNjcCI6W10sInN1YiI6ImRmNzQxNTdiLWY0Y2MtNDg2Yi1iZjU5LWM3YzlkNWY3NGQyNSJ9.b2E_1Gfov9VPjf4plSxFzxi--NpYWq2BcrrKiyS39Oj5Lk1jLGwwRXvmvXuQW2J0mvBGPM0NRUd204qWh7e7LFYObSB9N6boPdP4m7Dfo7P5To77TeBpiFG_FK5KBQAUZOrCMFYzH-VwvUFYv2UbIkSra9ajhFvyQton99wk0iz2zq_ic-8ZL-NJ5wyWizMQt-91UGVh1ap8gxV1NaOY45Cnuh9GzcgGffWUPpm1jQBgnB5vDBvKH00zVlu82lZafjiG436L70e9Om7c_uGgXMxpCOdcNLfiUA62T2v4hqn7J_iolc9mz7RlP6xqQZwUYCyORri1sYhr-e6L1MG9HUE0KS5lVcXj6OVuUrq2zPcy8BHaLANhpUw1SGl908YyMGbzaWUC_sqbtgGbhzScftj8s7j4kVfMPnWnaMgEDB_eLv1Eriyh_64Odgspz42aTYR4y0B2F-GvizyaMykGWK_wgU97cvS3QTfrtKhy04iFJxsN9x97SR4gLkHR05Px4cF8-BVMqYPXvzK6hA0oP7Ae3kas-bngqgG-IaG5eNMrZeVLvcTlp0q-twrbPUdGI9t5ug81NYEjGIZSmYcW4wY82aYKzK-fVqYrVDE0re0sriDhDZk2i0-IgniX7HHDuXJhFd-OKlBr4GtIhbIA1bGxlSj2OgO1ysmCFult32c`
	//	url := "https://zt.us.sms.api.sinch.com/xms/v1/" + s.client.ProjectId + "/batches/dry_run"

	//return GetHttp(token, url, opt)
	// return s.GetHttp(token.AccessToken, opt)
}
*/