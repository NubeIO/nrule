package apirules

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/nrule/pprint"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
)

type Points struct {
	Result []model.Point `json:"result"`
	Error  string        `json:"error"`
}

type Point struct {
	Result *model.Point `json:"result"`
	Error  string       `json:"error"`
}

func (inst *Client) GetPoints(hostIDName string) *Points {
	resp, err := cli.GetPoints(hostIDName)
	return &Points{
		Result: resp,
		Error:  errorString(err),
	}
}

func (inst *Client) GetPoint(hostIDName, uuid string) *Point {
	resp, err := cli.GetPoint(hostIDName, uuid)
	return &Point{
		Result: resp,
		Error:  errorString(err),
	}
}

func pointWriteBody(body any) (*model.Priority, error) {
	result := &model.Priority{}
	dbByte, err := json.Marshal(body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(dbByte, &result)
	return result, err
}

func (inst *Client) WritePointValue(hostIDName, uuid string, value *model.Priority) *Point {
	body, err := pointWriteBody(value)
	fmt.Println(111, err)
	pprint.PrintJSON(body)
	resp, err := cli.WritePointValue(hostIDName, uuid, body)
	return &Point{
		Result: resp,
		Error:  errorString(err),
	}
}
