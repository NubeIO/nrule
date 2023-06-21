package apirules

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// docs
// https://date.nager.at/swagger/index.html

type PublicHolidays struct {
	Date        string   `json:"date"`
	LocalName   string   `json:"localName"`
	Name        string   `json:"name"`
	CountryCode string   `json:"countryCode"`
	Fixed       bool     `json:"fixed"`
	Global      bool     `json:"global"`
	Counties    []string `json:"counties"`
	LaunchYear  int      `json:"launchYear"`
	Types       []string `json:"types"`
}

type PublicHolidaysResponse struct {
	Result []PublicHolidays
	Error  string
}

func (inst *Client) GetPublicHolidays(year, countryCode string) *PublicHolidaysResponse {
	client := resty.New()
	url := fmt.Sprintf("https://date.nager.at/api/v3/PublicHolidays/%s/%s", year, countryCode)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]PublicHolidays{}).
		Get(url)
	if resp.StatusCode() >= 400 {
		err = errors.New(resp.String())
	}
	var out []PublicHolidays
	out = *resp.Result().(*[]PublicHolidays)
	r := &PublicHolidaysResponse{
		Result: out,
		Error:  errorString(err),
	}

	return r
}

func (inst *Client) GetPublicHolidaysByState(year, countryCode, state string) *PublicHolidaysResponse {
	client := resty.New()
	url := fmt.Sprintf("https://date.nager.at/api/v3/PublicHolidays/%s/%s", year, countryCode)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]PublicHolidays{}).
		Get(url)
	if resp.StatusCode() >= 400 {
		err = errors.New(resp.String())
	}
	var out []PublicHolidays
	var outFiltered []PublicHolidays
	out = *resp.Result().(*[]PublicHolidays)

	for _, holiday := range out {
		for _, county := range holiday.Counties {
			if county == fmt.Sprintf("%s-%s", countryCode, state) {
				outFiltered = append(outFiltered, holiday)
			}
		}
	}
	r := &PublicHolidaysResponse{
		Result: outFiltered,
		Error:  errorString(err),
	}
	return r
}

type IsPublicHoliday struct {
	IsPublicHoliday bool
	Name            string
	Locations       []string `json:"locations"`
}

type IsPublicHolidayResponse struct {
	Result *IsPublicHoliday
	Error  string
}

func (inst *Client) IsPublicHoliday(year, countryCode, date string) *IsPublicHolidayResponse {
	client := resty.New()
	url := fmt.Sprintf("https://date.nager.at/api/v3/PublicHolidays/%s/%s", year, countryCode)
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetResult(&[]PublicHolidays{}).
		Get(url)
	if resp.StatusCode() >= 400 {
		err = errors.New(resp.String())
	}
	var out []PublicHolidays
	out = *resp.Result().(*[]PublicHolidays)
	isHoliday := &IsPublicHoliday{}
	for _, holiday := range out {
		if holiday.Date == date {
			isHoliday.IsPublicHoliday = true
			isHoliday.Locations = holiday.Counties
			isHoliday.Name = holiday.Name

		}
	}
	r := &IsPublicHolidayResponse{
		Result: isHoliday,
		Error:  errorString(err),
	}
	return r
}
