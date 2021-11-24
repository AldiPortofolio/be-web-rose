package utils

import (
	"log"
	"testing"
)

func TestStringWithCharset(t *testing.T) {

	a := StringWithCharset(10, "OP1A000321312312312")
	log.Println("a --> ", a)
}

func TestGenerateRandom(t *testing.T) {
	a:= GenerateRandom(10)
	log.Println("---> ", a)
}