package apirules

import (
	"encoding/json"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

type Alert struct {
	Result *model.Alert `json:"result"`
	Error  string       `json:"error"`
}

func bindAlert(body any) (*model.Alert, error) {
	result := &model.Alert{}
	dbByte, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(dbByte, &result)
	return result, err
}

func (p *Client) AddAlert(hostIDName string, body any) *Alert {
	b, err := bindAlert(body)
	if err != nil {
		return &Alert{
			Result: nil,
			Error:  errorString(err),
		}
	}
	resp, err := cli.AddAlert(hostIDName, b)
	return &Alert{
		Result: resp,
		Error:  errorString(err),
	}
}
