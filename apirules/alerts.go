package apirules

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

type Alert struct {
	Result *model.Alert `json:"result"`
	Error  string       `json:"error"`
}

func alertBody(body any) (*model.Alert, error) {
	result := &model.Alert{}
	dbByte, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(dbByte, &result)
	return result, err
}

func (p *Client) AddAlert(hostIDName string, body any) *Alert {
	b, err := alertBody(body)
	if err != nil {
		return &Alert{
			Result: nil,
			Error:  fmt.Sprintf("failed to parse body err:%s", err.Error()),
		}
	}
	resp, err := cli.AddAlert(hostIDName, b)
	return &Alert{
		Result: resp,
		Error:  errorString(err),
	}
}
