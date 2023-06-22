package apirules

import (
	"github.com/NubeDev/flow-eng/helpers"
)

type pingResult struct {
	Ip string `json:"ip"`
	Ok bool   `json:"ok"`
}

type PingResponse struct {
	Result []pingResult
	Error  string
}

func (inst *Client) Ping(ipList []string) *PingResponse {
	var r pingResult
	var out []pingResult
	for _, ip := range ipList {
		ok := helpers.CommandPing(ip)
		r.Ip = ip
		r.Ok = ok
		out = append(out, r)
	}
	return &PingResponse{
		Result: out,
		Error:  "",
	}
}
