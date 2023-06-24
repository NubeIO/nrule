package apirules

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/nrule/storage"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
)

type VarsResponse struct {
	Result []storage.RQLVariables
	Error  string
}

type VarResponse struct {
	Result *storage.RQLVariables
	Error  string
}

func (inst *Client) GetVariables() *VarsResponse {
	out, err := inst.Storage.SelectAllVariables()
	return &VarsResponse{
		Result: out,
		Error:  errorString(err),
	}
}

/*
VarParseArray
[1, 2, "hello"]

let data = JSON.parse(RQL.VarParseArray("array"));
RQL.Return = data;
*/
func (inst *Client) VarParseArray(uuidName string) interface{} {
	r := inst.GetVariable(uuidName)
	if r == nil {
		return nil
	}
	b, err := json.Marshal(r.Result.Variable)
	if err != nil {
		return 0
	}
	jsonStr := string(b)
	a := gjson.Parse(jsonStr).Array()
	if len(a) > 0 {
		return a[0].String()
	}
	return a
}

/*
VarParseObject
`{"sp":22.3,"db":99.9}`

let data = RQL.VarParseObject("obj");
let sp = JSON.parse(data);
RQL.Return = sp["sp"];
*/
func (inst *Client) VarParseObject(uuidName string) interface{} {
	r := inst.GetVariable(uuidName)
	if r == nil {
		return nil
	}
	b, err := json.Marshal(r.Result.Variable)
	if err != nil {
		return nil
	}
	jsonStr := string(b)
	a := gjson.Parse(jsonStr).String()
	t := strings.ReplaceAll(a, "'", "")
	t = strings.ReplaceAll(t, "`", "")
	return t
}

func (inst *Client) VarParseString(uuidName string) string {
	r := inst.GetVariable(uuidName)
	if r == nil {
		return ""
	}
	f := r.Result.Variable
	return fmt.Sprint(f)
}

func (inst *Client) VarParseNumber(uuidName string) float64 {
	r := inst.GetVariable(uuidName)
	if r == nil {
		return 0
	}
	f := r.Result.Variable
	if s, err := strconv.ParseFloat(fmt.Sprint(f), 64); err == nil {
		return s
	}
	return 0
}

func (inst *Client) GetVariable(uuidName string) *VarResponse {
	out, err := inst.Storage.SelectAllVariables()
	for _, variables := range out {
		if variables.Name == uuidName {
			return &VarResponse{
				Result: &variables,
				Error:  errorString(err),
			}
		}
		if variables.UUID == uuidName {
			return &VarResponse{
				Result: &variables,
				Error:  errorString(err),
			}
		}
	}
	return &VarResponse{
		Result: nil,
		Error:  errorString(err),
	}
}
