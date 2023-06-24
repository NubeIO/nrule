package apirules

import "time"

func (inst *Client) SleepMs(duration int) {
	d := time.Duration(duration)
	time.Sleep(d * time.Millisecond)
}

func (inst *Client) Sleep(duration int) {
	d := time.Duration(duration)
	time.Sleep(d * time.Second)
}

func (inst *Client) TimeUTC() time.Time {
	return time.Now().UTC()
}

func (inst *Client) TimeDate() string {
	return time.Now().Format("2006.01.02 15:04:05")
}

func (inst *Client) Time() string {
	return time.Now().Format("15:04:05") // 2006.01.02
}

func (inst *Client) Date() string {
	return time.Now().Format("2006.01.02") // 2006.01.02
}

func (inst *Client) Year() string {
	return time.Now().Format("2006") // 2006.01.02
}

func (inst *Client) Day() string {
	return time.Now().Format("Monday") // 2006.01.02
}

func (inst *Client) TimeDateFormat(format string) string {
	return time.Now().Format(format) // 2006.01.02
}

func (inst *Client) TimeDateDay() string {
	return time.Now().Format("2006-01-02 15:04:05 Monday")
}
