package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/services"
	ottoutils "ottodigital.id/library/utils"

)

type MerchantMasterTagController struct {

}

// @Summary MerchantMasterTag - Filter
// @Description MerchantMasterTag Filter Paging
// @ID MerchantMasterTagFilter
// @Param Authorization header string true "Bearer"
// @Param body body dto.ReqMerchantMasterTagDto true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response{contents=[]dbmodels.MerchantMasterTag}
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/merchant-master-tag/filter [post]
func (controller *MerchantMasterTagController) Filter(ctx *gin.Context)  {
	fmt.Println(">>> MerchantMasterTagController - Filter <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	var res models.Response

	var req dto.ReqMerchantMasterTagDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		logs.Error("err unmarshal "+err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	reqByte,_ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))


	services.InitMerchantMasterTagService().Filter(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MerchantMasterTagController  - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

func (controller *MerchantMasterTagController) FindByMid(ctx *gin.Context)  {
	fmt.Println(">>> MerchantMasterTagController - FindByMid <<<")
	// initiate logs
	logs := logger.InitLogs(ctx.Request)

	var res models.Response

	var req dto.ReqMerchantMasterTagByMidDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("Request body error:", err)
		res.ErrCode = constants.ERR_UNMARSHAL
		res.ErrDesc = constants.ERR_UNMARSHAL_MSG
		logs.Error("err unmarshal "+err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	reqByte,_ := json.Marshal(req)
	log.Println("req --> ", string(reqByte))


	services.InitMerchantMasterTagService().FindByMid(req, &res)

	bodyRes, _ := json.Marshal(res)

	logs.Info("Response MerchantMasterTagController  - Filter",
		logs.AddField("ResponseBody: ", string(bodyRes)))


	ctx.JSON(http.StatusOK, res)

}

// @Summary MerchantMasterTag - Download ALl
// @Description MerchantMasterTag DownloadAll
// @ID MerchantMasterTagDownloadAll
// @Param Authorization header string true "Bearer"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /rosego/v.0.1/merchant-master-tag/download-all [get]
func (controller *MerchantMasterTagController) DownloadAll(ctx *gin.Context)  {
	pathDir := ottoutils.GetEnv("PATH_DOWNLOAD_EXAMPLE_MERCHANT_MASTER_TAG","/opt/app-rose/merchant-master-tag/")
	path := pathDir + "alldata.xlsx"

	fmt.Println(">>> MerchantMasterTagController - DownloadAll <<<")
	// initiate logs

	var res models.Response

	w := ctx.Writer

	data, err:= db.InitMerchantMasterTagRepository().GetAll()
	if err != nil {
		http.Error(w,"File Not Found", http.StatusInternalServerError)
		fmt.Println("Failed to open file", err)
		return
	}

	ExcelExport(data, path)
	f, err := os.Open(path)

	if f != nil {
		defer f.Close()
	}
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w,"File Not Found", http.StatusInternalServerError)
		fmt.Println("Failed to open file", err)
		return
	}
	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	w.Header().Set("Content-Disposition", contentDisposition)

	if _, err := io.Copy(w, f); err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		http.Error(w,"File Not Found", http.StatusInternalServerError)
		fmt.Println("Failed to copy file", err)
		return
	}

	ctx.JSON(http.StatusOK, res)

}

func ExcelExport(data []dbmodels.MerchantMasterTag, path string) error {
	sheetName := "Sheet1"
	f := excelize.NewFile()

	f.SetCellValue(sheetName, fmt.Sprintf("A%d", 1), "NO")
	f.SetCellValue(sheetName, fmt.Sprintf("B%d", 1), "MID")
	f.SetCellValue(sheetName, fmt.Sprintf("C%d", 1), "TAG CODE")
	for i := 0; i < len(data); i++  {
		j:= i+2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", j), i+1)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", j), data[i].Mid)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", j), data[i].MasterTagCode)

	}

	err := f.SaveAs(path)
	if err != nil {
		return err
	}

	//f.set
	return nil
}
