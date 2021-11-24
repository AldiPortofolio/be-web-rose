package utils

import (
	"log"
	"testing"
	"time"
)

func TestConverDateStringToTime(t *testing.T) {
	res:=ConverDateStringToTime("21-04-2020")

	log.Println("res->", res)

	log.Println("res->", res.Add(time.Hour * 24))
}

func TestConvertDateFormat(t *testing.T) {
	res:=ConvertDateFormat("21-04-2020")
	log.Println("res-->", res)
}

func TestConverDateStringYYYYMMDDToTime(t *testing.T) {
	res:=ConverDateStringYYYYMMDDToTime("2020-12-26")

	log.Println("res->", res)

	log.Println("res->", res.Add(time.Hour * 24))

}