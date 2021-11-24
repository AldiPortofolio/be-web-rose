package utils

import (
	"time"
)

// ConvertTime ...
func ConvertTime(t time.Time) string {
	return t.Format("2006-01-02T15:04:05")

}

// ConvertDateFormat ...
func ConvertDateFormat(reqDate string) string {
	reqDateLayout := "02-01-2006"
	dateLayout := "2006-01-02"

	date, _ := time.Parse(reqDateLayout, reqDate)

	return date.Format(dateLayout)
}

func ConverDateStringToTime(reqDate string) time.Time {
	date := ConvertDateFormat(reqDate)
	layout := "2006-01-02"
	//str := "2014-11-12T11:45:26.371Z"
	t, _ := time.Parse(layout, date)
	return t
}

func  ConverDateStringYYYYMMDDToTime(reqDate string) time.Time {
	layout := "2006-01-02"
	//str := "2014-11-12T11:45:26.371Z"
	t, _ := time.Parse(layout, reqDate)
	return t
}

func ShortDateFromString(ds string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", ds)
	if err != nil {
		return t, err
	}
	return t, nil
}