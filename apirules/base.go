package apirules

import (
	"context"
	"github.com/NubeIO/rubix-os-client/rubixoscli"
	"github.com/NubeIO/rubix-os/installer"
)

type Client struct {
	Return         interface{} `json:"return"`
	Err            string      `json:"err"`
	TimeTaken      string      `json:"time_taken"`
	PdfApplication *PDFApplication
	CTX            context.Context
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
