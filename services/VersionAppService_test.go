package services

import (
	"log"
	"rose-be-go/models"
	"testing"
)

func TestVersionAppService_GetVersionOttomart(t *testing.T) {
	gen := models.GeneralModel{}
	x :=make(chan string)
	go InitVersionAppService(gen).GetVersionOttomart(x)
	res := <-x
	log.Println(res)
}

func TestVersionAppService_GetVersionNfc(t *testing.T) {
	gen := models.GeneralModel{}
	x :=make(chan string)
	go InitVersionAppService(gen).GetVersionNfc(x)
	res := <-x
	log.Println(res)
}

func TestVersionAppService_GetVersionIndomarco(t *testing.T) {
	gen := models.GeneralModel{}
	x :=make(chan string)
	go InitVersionAppService(gen).GetVersionIndomarco(x)
	res := <-x
	log.Println(res)
}

func TestVersionAppService_GetVersionSfa(t *testing.T) {
	gen := models.GeneralModel{}
	x :=make(chan string)
	go InitVersionAppService(gen).GetVersionSfa(x)
	res := <-x
	log.Println(res)
}

func TestVersionAppService_UpdateVersion(t *testing.T) {
	gen := models.GeneralModel{}
	err := InitVersionAppService(gen).UpdateVersion("sfa", "1")
	log.Println(err)

}