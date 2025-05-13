package core

import (
	"encoding/json"
	"strings"
)

// ReplaceUrl replaces the placeholders in the URL with actual values.
/*
url:= "https://zt.{Region}.sms.api.sinch.com/xms/v1/{ProjectId}"
url = core.ReplaceUrl(url, s.C.ProjectId, s.C.Region)
*/
func ReplaceUrl(url string, projectId string, region string) string {
	temp := strings.Replace(url, "{Region}", region, 1)
	return strings.Replace(temp, "{ProjectId}", projectId, 1)
	//	u := fmt.Sprintf("test/%s/batches", projectId) fmt.println(u)
}

// if err := core.Byte2Json(response1, &result); err != nil {
// 	return result, err
// }

func Byte2Json(response []byte, data *any) error {
	if err := json.Unmarshal(response, &data); err != nil { // Parse []byte to go struct pointer
		return err
	}
	return nil
}
