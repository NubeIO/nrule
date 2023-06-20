package apirules

import (
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

type Histories struct {
	Result []model.PointHistory `json:"result"`
	Error  string               `json:"error"`
}

func (inst *Client) GetPointHistories(hostIDName string, pointUUIDs []string) *Histories {
	resp, err := cli.GetPointHistories(hostIDName, pointUUIDs)
	return &Histories{
		Result: resp,
		Error:  errorString(err),
	}
}
