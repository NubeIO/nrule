package apirules

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NubeIO/nrule/pprint"
	"github.com/go-resty/resty/v2"
	"regexp"
)

type CurrentWeatherResponse struct {
	Result *CurrentWeather
	Error  string
}

type CurrentWeather struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzId           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

type HTTPGet struct {
	Result any
	Error  string
}

type HTTPBody struct {
	Url          string `json:"url"`
	Method       string
	ResponseType string `json:"response_type"` //json, string
	Headers      map[string]string
}

var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)

type conventionalMarshaller struct {
	Value interface{}
}

func (c conventionalMarshaller) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
				match,
				[]byte(`${1}_${2}`),
			))
		},
	)
	return converted, err
}

func httpBody(body any) (*HTTPBody, error) {
	result := &HTTPBody{}
	//dbByte, err := json.Marshal(body)
	//dbByte, err := json.MarshalIndent(body, "", "  ")
	//if err != nil {
	//	return result, err
	//}
	//err = json.Unmarshal(dbByte, &result)

	encoded, _ := json.MarshalIndent(conventionalMarshaller{body}, "", "  ")
	err := json.Unmarshal(encoded, &result)

	return result, err
}

func (inst *Client) HTTPGet(body *HTTPBody) *HTTPGet {
	b, _ := httpBody(body)
	fmt.Println(11111)
	pprint.PrintJSON(b)
	fmt.Println(11111)
	client := resty.New()
	pprint.PrintJSON(body)
	var resp *resty.Response
	var err error
	if body.Method == "GET" {
		resp, err = client.R().
			SetHeaders(body.Headers).
			Get(body.Url)
	}

	var response interface{}
	err = json.Unmarshal(resp.Body(), &response)
	return &HTTPGet{
		Result: response,
		Error:  errorString(err),
	}
}

func (inst *Client) GetCurrentWeather(apiKey, city string) *CurrentWeatherResponse {
	client := resty.New()
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, city)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetResult(&CurrentWeather{}).
		Get(url)
	if resp.StatusCode() >= 400 {
		err = errors.New(resp.String())
	}
	res := resp.Result().(*CurrentWeather)
	r := &CurrentWeatherResponse{
		Result: res,
		Error:  errorString(err),
	}
	return r
}
