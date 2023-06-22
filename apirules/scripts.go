package apirules

import (
	"github.com/NubeIO/nrule/storage"
	"github.com/go-resty/resty/v2"
)

type ScriptsResponse struct {
	Result []storage.RQLRule
	Error  string
}

func (inst *Client) GetScripts() *ScriptsResponse {
	client := resty.New()
	url := "http://0.0.0.0:1666/api/rules"
	resp, err := client.R().
		SetResult(&[]storage.RQLRule{}).
		Get(url)
	var out []storage.RQLRule
	out = *resp.Result().(*[]storage.RQLRule)
	return &ScriptsResponse{
		Result: out,
		Error:  errorString(err),
	}
}
