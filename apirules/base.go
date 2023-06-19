package apirules

import (
	"github.com/NubeIO/rubix-os-client/rubixoscli"
	"github.com/NubeIO/rubix-os/installer"
)

type Client struct {
	Result interface{} `json:"result"`
	Err    interface{} `json:"err"`
}

var cli = rubixoscli.New(&rubixoscli.Client{
	Rest:          nil,
	Installer:     nil,
	Ip:            "0.0.0.0",
	Port:          1659,
	HTTPS:         false,
	ExternalToken: "",
}, &installer.Installer{})

func errorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
