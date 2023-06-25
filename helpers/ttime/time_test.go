package ttime

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/andanhm/go-prettytime"
)

const (
	layout = "2006-01-02T15:04:05Z"
)

func TestRealTime_Now(t *testing.T) {
	firstDate := time.Date(2023, 6, 25, 9, 7, 0, 0, time.Local)

	log.Printf("%s \n", prettytime.Format(firstDate))

}

func TestRealDate(t *testing.T) {
	firstDate := time.Date(2022, 4, 13, 6, 0, 0, 0, time.UTC)
	secondDate := time.Date(2022, 4, 13, 6, 1, 0, 0, time.UTC)
	difference := firstDate.Sub(secondDate)

	fmt.Printf("Years: %d\n", int64(difference.Hours()/24/365))
	fmt.Printf("Hours: %.f\n", difference.Hours())
	fmt.Printf("Minutes: %.f\n", difference.Minutes())
	tt, err := AdjustTime(firstDate, "3 hour 15 min")
	fmt.Println(err)
	fmt.Println("was", firstDate, "now", tt)

}
