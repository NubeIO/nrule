package apirules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/NubeIO/nrule/rules"
	"github.com/go-gota/gota/dataframe"
	"github.com/stretchr/testify/assert"
	"github.com/yukithm/json2csv"
	"io"
	"log"
	"strings"
	"testing"
)

func TestCycleCallRule(t *testing.T) {

	script := `
	// JS script
	let cli = Client // go-lang code
	let hist = cli.GetPoints("rc") // core lib by nube with all basic functions
	cli.PrintMany("points count:", hist.Points.length)
	if (hist.Err != "") {
		cli.PrintMany("hist.Err:", hist.Err)
	} else {
		let df = cli.JsonToDF(hist.Points)
		let value = df.Col("present_value").Max() // get max values from the points
		cli.PrintMany("max point value",":",value*100)
	}

`

	eng := rules.NewRuleEngine()
	err := eng.Start()
	assert.Nil(t, err)

	name := "Core"
	props := make(rules.PropertiesMap)
	props[name] = eng

	client := "Client"
	newClient := &Client{Name: 10}
	props[client] = newClient

	err = eng.AddRule(name, script, props)
	if err != nil {
		fmt.Println("AddRule Err", err)
	}

	err = eng.Execute(name)
	if err != nil {
		fmt.Println("Execute Err", err)
	}
}

func TestCSV(t *testing.T) {

	script := `
	let hist = Client.GetPoints("rc") // core lib by nube with all basic functions


`

	eng := rules.NewRuleEngine()
	err := eng.Start()
	assert.Nil(t, err)

	name := "Core"
	props := make(rules.PropertiesMap)
	props[name] = eng

	client := "Client"
	newClient := &Client{Name: 10}
	props[client] = newClient

	err = eng.AddRule(name, script, props)
	if err != nil {
		fmt.Println(1111, err)
	}

	err = eng.Execute(name)
	if err != nil {
		fmt.Println(3333, err)
	}

	fmt.Println(eng.Result)
	fmt.Println(newClient.Name)

	pnts := newClient.GetPoints("rc")

	data, err := toInterface(pnts.Points)
	s, err := jsonToSri(pnts.Points)
	jsonDf := dataframe.ReadJSON(strings.NewReader(s))
	fmt.Println(44)
	fmt.Println(jsonDf)
	fmt.Println(jsonDf)

	csv, err := json2csv.JSON2CSV(data)

	b := &bytes.Buffer{}
	wr := json2csv.NewCSVWriter(b)
	wr.HeaderStyle = json2csv.DotNotationStyle
	wr.Transpose = false
	err = wr.WriteCSV(csv)
	if err != nil {
		log.Fatal(err)
	}
	wr.Flush()
	got := b.String()

	fmt.Println(11111)
	println(got)
	fmt.Println(11111)
	//err = printCSV(os.Stdout, csv, json2csv.DotNotationStyle, false)
	//if err != nil {
	//	log.Fatal(err)
	//}

	df := dataframe.ReadCSV(strings.NewReader(got))

	//aaa := df.Select([]string{"name", "uuid", "present_value"}).GroupBy("name")

	var buf bytes.Buffer
	//err = df.Select(
	//	[]string{"name", "uuid", "present_value"},
	//).WriteCSV(&buf)
	//if err != nil {
	//	log.Fatal(err)
	//}

	err = df.Select(
		[]string{"name", "present_value"},
	).GroupBy(
		"name", "present_value",
	).Aggregation(
		[]dataframe.AggregationType{dataframe.Aggregation_SUM},
		[]string{"present_value"},
	).Select(
		[]string{"name", "present_value_SUM"},
	).WriteCSV(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(err)
	fmt.Println(buf.String())
	fmt.Println(222)

}

func TestDf(t *testing.T) {

	newClient := &Client{Name: 10}

	pnts := newClient.GetPoints("rc")
	df, err := jsonToDF(pnts.Points)
	if err != nil {
		fmt.Println(err)
	}
	df.Col("present_value").Max()
	fmt.Printf("Summary of the dataframe\n%v+", df.Describe())

	fmt.Println(df.Select([]string{"name", "uuid", "present_value", "tags"}))

	//dfTeams := df.GroupBy("name")
	//teamsByGroup := dfTeams.GetGroups()

	//// print @soypete top 5 teams USMNT, Uruguay, Argentina, Brazil, Mexico
	//brazil := teamsByGroup["pnt1"].Filter(
	//	dataframe.F{
	//		Colname:    "player",
	//		Comparator: series.Neq,
	//		Comparando: "",
	//	},
	//)
	//fmt.Println("Brazil National Team")
	//fmt.Println(brazil.Select([]string{"player", "age", "position", "birth_year"}))

}

func toInterface(data any) (any, error) {
	b, err := json.Marshal(data)
	var a interface{}
	err = json.Unmarshal(b, &a)
	return a, err

}

func jsonToSri(data any) (string, error) {
	b, err := json.Marshal(data)
	return string(b), err

}

func jsonToDF(data any) (dataframe.DataFrame, error) {
	b, err := json.Marshal(data)
	df := dataframe.ReadJSON(strings.NewReader(string(b)))
	return df, err
}

func printCSV(w io.Writer, results []json2csv.KeyValue, headerStyle json2csv.KeyStyle, transpose bool) error {
	csv := json2csv.NewCSVWriter(w)
	csv.HeaderStyle = headerStyle
	csv.Transpose = transpose
	if err := csv.WriteCSV(results); err != nil {
		return err
	}

	output := fmt.Sprint(w)
	fmt.Println(111)
	fmt.Println(output)
	fmt.Println(111)
	return nil
}
