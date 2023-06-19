package apirules

import "github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"

type Points struct {
	Result []model.Point `json:"result"`
	Error  string        `json:"error"`
}

type Point struct {
	Result *model.Point `json:"result"`
	Error  string       `json:"error"`
}

func (p *Client) GetPoints(hostIDName string) *Points {
	resp, err := cli.GetPoints(hostIDName)
	return &Points{
		Result: resp,
		Error:  errorString(err),
	}
}

func (p *Client) GetPoint(hostIDName, uuid string) *Point {
	resp, err := cli.GetPoint(hostIDName, uuid)
	return &Point{
		Result: resp,
		Error:  errorString(err),
	}
}
