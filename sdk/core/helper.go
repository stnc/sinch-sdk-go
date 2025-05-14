package core

import (
	"strings"
)

// ReplaceUrl replaces the placeholders in the URL with actual values.
/*
url:= "https://zt.{Region}.sms.api.sinch.com/xms/v1/{ProjectId}"
url = core.ReplaceUrl(url, s.C.ProjectId, s.C.Region)
*/
func ReplaceUrl(url string, projectId string, region string) string {
	urltemp := strings.Replace(url, "{Region}", region, 1)
	return strings.Replace(urltemp, "{ProjectId}", projectId, 1)
	//	u := fmt.Sprintf("users/%s/projects", projectId) fmt.println(u)
}

