package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	UNIX_SECONDS      = time.Second
	UNIX_MILLISECONDS = time.Millisecond
)

const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
	DateTime   = "2006-01-02 15:04:05"
	DateOnly   = "2006-01-02"
	TimeOnly   = "15:04:05"
)

func ParseTimestampUnit(unixTimestamp int64) int64 {
	var unit time.Duration

	if len(strconv.Itoa(int(unixTimestamp))) == 10 {
		unit = UNIX_SECONDS
	}

	if len(strconv.Itoa(int(unixTimestamp))) == 13 {
		unit = UNIX_MILLISECONDS
	}

	return int64(unixTimestamp) * int64(unit)
}

func TimestampToLocation(unixTimestamp int64, location string) time.Time {
	date := time.Unix(0, ParseTimestampUnit(unixTimestamp))

	timezone, err := time.LoadLocation(location)
	if err != nil {
		panic(err)
	}
	date = date.In(timezone)
	return date
}

func TimestampToRFC3339() {
	_ = time.Unix(int64(16200063800), 0)
	t := time.Now()
	// t.Unix() // time.Time -> timestamp
	s := t.Format(time.RFC3339)
	_, _ = time.Parse(time.RFC3339, s)
}

func RandomlyTimeDecent() {
	now := time.Now().UnixNano()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomNumber := int64(rand.Intn(999999999999999))
		descendingTime := now - randomNumber
		fmt.Println(time.Unix(0, descendingTime).Format(time.RFC3339))
		now = descendingTime
		fmt.Println("-------------")
	}
}

func main() {

}
