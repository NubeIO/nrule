package apirules

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/pkg/v1/model"
	"github.com/NubeIO/rubix-os-client/rubixoscli"
	"github.com/NubeIO/rubix-os/installer"
	pprint "github.com/NubeIO/rubix-ui/backend/helpers/print"
	"github.com/go-gota/gota/dataframe"
	"strings"
)

type Client struct {
	Result interface{} `json:"result"`
	Err    interface{} `json:"err"`
}

func (p *Client) GetPoints(hostIDName string) *Points {
	resp, err := cli.GetPoints(hostIDName)
	return &Points{
		Points: resp,
		Err:    errorString(err),
	}
}

func (p *Client) GetPoint(hostIDName, uuid string) *Point {
	resp, err := cli.GetPoint(hostIDName, uuid)
	return &Point{
		Point: resp,
		Err:   errorString(err),
	}
}

func (p *Client) GetPointHistories(hostIDName string, pointUUIDs []string) *Histories {
	resp, err := cli.GetPointHistories(hostIDName, pointUUIDs)
	return &Histories{
		Histories: resp,
		Err:       errorString(err),
	}
}

func (p *Client) PrintJson(x interface{}) {
	fmt.Println("CALL PRINT")
	pprint.PrintJOSN(x)
}

func (p *Client) Print(x interface{}) {
	fmt.Println(x)
}

func (p *Client) PrintMany(x ...interface{}) {
	fmt.Printf("%v\n", x)
}

func (p *Client) JsonToDF(data any) dataframe.DataFrame {
	b, err := json.Marshal(data)
	if err != nil {
		return dataframe.DataFrame{}
	}
	df := dataframe.ReadJSON(strings.NewReader(string(b)))
	return df
}

func (p *Client) Tags(tag ...string) {
	var includeList []string
	var excludeList []string
	for _, s := range tag {
		exclude := strings.Contains(s, "!")
		if exclude {
			t := strings.Trim(s, "!")
			excludeList = append(excludeList, t)
		} else {
			includeList = append(includeList, s)
		}
	}

	for i, s := range includeList {
		fmt.Println("includeList", i, s)
	}
	for i, s := range excludeList {
		fmt.Println("excludeList", i, s)
	}

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

type Points struct {
	Points []model.Point `json:"points"`
	Err    string        `json:"err"`
}

type Point struct {
	Point *model.Point `json:"point"`
	Err   string       `json:"err"`
}

type Histories struct {
	Histories []model.PointHistory `json:"histories"`
	Err       string               `json:"err"`
}
