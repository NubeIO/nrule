package js

import (
	"fmt"
	"github.com/NubeIO/nrule"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Rule1 struct {
	Name  int
	Name2 int
	Title string
}

type User struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (p *Rule1) GetUser() *User {
	client := resty.New()
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/1")
	resp, err := client.R().
		SetResult(&User{}).
		Get(url)
	if err != nil {
		return nil
	}
	return resp.Result().(*User)
}

func (p *Rule1) Add100() int {
	return p.Name + 100
}

func TestCycleCallRule2(t *testing.T) {

	script := `
	R1.Name = R1.Add100()
	R1.Name2 = 99
	R1.Title = R1.GetUser().Title // gets the title from an API call
`

	eng := NewRuleEngine()
	err := eng.Start()
	assert.Nil(t, err)

	name := "R1"

	r := &Rule1{Name: 10}

	props := make(nrule.PropertiesMap)
	props[name] = r

	err = eng.AddRule(name, script, props)
	if err != nil {
		fmt.Println(1111, err)
	}

	err = eng.Execute(name)
	if err != nil {
		fmt.Println(3333, err)
	}

	fmt.Println(r.Name)
	fmt.Println(r.Name2)
	fmt.Println(r.Title)

}
