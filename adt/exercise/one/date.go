package one

import (
	"strconv"
	"strings"
)

func toUint8(n string) uint8 {
	integer, _ := strconv.Atoi(n)

	return uint8(integer)
}

type Date struct {
	day   uint8
	month uint8
	year  uint8
}

func (d *Date) validate() bool {

	if ((d.month == 4 || d.month == 6 || d.month == 9 || d.month == 11) && d.day > 30) ||
		(d.month == 2 && d.day > 28) || d.day > 31 {
		return false
	}

	return true
}

func Insert(day uint8, month uint8, year uint8, date *Date) bool {
	*date = Date{day, month, year}

	return date.validate()
}

func InsertRaw(rawDate string, date *Date) bool {
	dateSlice := strings.Split(rawDate, "/")

	*date = Date{toUint8(dateSlice[0]), toUint8(dateSlice[1]), toUint8(dateSlice[2])}

	return date.validate()
}
