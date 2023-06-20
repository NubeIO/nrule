package apirules

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

type Host struct {
	Result *model.Host `json:"result"`
	Error  string      `json:"error"`
}

type Hosts struct {
	Result []model.Host `json:"result"`
	Error  string       `json:"error"`
}

func (inst *Client) GetHosts() *Hosts {
	resp, msg := cli.GetHosts()
	var err string
	if msg != nil {
		if msg.StatusCode > 300 {
			err = fmt.Sprint(msg.Message)
		}
	}
	return &Hosts{
		Result: resp,
		Error:  err,
	}
}
