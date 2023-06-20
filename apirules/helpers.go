package apirules

import (
	"encoding/json"
	"fmt"
	pprint "github.com/NubeIO/nrule/helpers/print"
	"github.com/go-gota/gota/dataframe"
	"strings"
	"time"
)

func (inst *Client) TimeUTC() time.Time {
	return time.Now().UTC()
}

func (inst *Client) TimeDate() string {
	return time.Now().Format("2006.01.02 15:04:05")
}

func (inst *Client) TimeDateDay() string {
	return time.Now().Format("2006-01-02 15:04:05 Monday")
}

func (inst *Client) Time() time.Time {
	return time.Now()
}

func (inst *Client) PrintJson(x interface{}) {
	pprint.PrintJOSN(x)
}

func (inst *Client) PrintString(x string) {
	fmt.Println(x)
}

func (inst *Client) Print(x interface{}) {
	fmt.Println(x)
}

func (inst *Client) ToString(x interface{}) string {
	return fmt.Sprint(x)
}

func (inst *Client) PrintMany(x ...interface{}) {
	fmt.Printf("%v\n", x)
}

func (inst *Client) JsonToDF(data any) dataframe.DataFrame {
	b, err := json.Marshal(data)
	if err != nil {
		return dataframe.DataFrame{}
	}
	df := dataframe.ReadJSON(strings.NewReader(string(b)))
	return df
}

func (inst *Client) Tags(tag ...string) {
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
