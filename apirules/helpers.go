package apirules

import (
	"encoding/json"
	"fmt"
	pprint "github.com/NubeIO/nrule/helpers/print"
	"github.com/go-gota/gota/dataframe"
	"strings"
	"time"
)

func (p *Client) TimeUTC() time.Time {
	return time.Now().UTC()
}

func (p *Client) TimeDate() string {
	return time.Now().Format("2006.01.02 15:04:05")
}

func (p *Client) TimeDateDay() string {
	return time.Now().Format("2006-01-02 15:04:05 Monday")
}

func (p *Client) Time() time.Time {
	return time.Now()
}

func (p *Client) PrintJson(x interface{}) {
	pprint.PrintJOSN(x)
}

func (p *Client) PrintString(x string) {
	fmt.Println(x)
}

func (p *Client) Print(x interface{}) {
	fmt.Println(x)
}

func (p *Client) ToString(x interface{}) string {
	return fmt.Sprint(x)
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
