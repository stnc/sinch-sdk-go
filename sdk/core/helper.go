package core

import (
	"strings"
)

func ReplaceUrl(url string, projectId string, region string) string {
	urltemp := strings.Replace(url, "{Region}", region, 1)
	return strings.Replace(urltemp, "{ProjectId}", projectId, 1)
}
