package batches

import (
	"encoding/json"

	"net/http"

	"github.com/stnc/sinch-sdk-go/sdk/core"
	"github.com/stnc/sinch-sdk-go/sdk/model"
	batchesModel "github.com/stnc/sinch-sdk-go/sdk/model/sms/batches"
)

//            "per_recipient": str(self.request_data.per_recipient),
//            "number_of_recipients": self.request_data.number_of_recipients

func (s *Batches) DryRun(req *batchesModel.SendDryRunRequest) (batchesModel.SendDryRunResponse, error) {
	var result batchesModel.SendDryRunResponse
	token := core.GetToken(s.C)

	url := core.ReplaceUrl(model.APIUrl+"/batches/dry_run", s.C.ProjectId, s.C.Region)

	response, err := core.NewRequest(http.MethodPost, url, token.AccessToken, req)

	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(response, &result); err != nil { 
		return result, err
	}
	return result, err
}
