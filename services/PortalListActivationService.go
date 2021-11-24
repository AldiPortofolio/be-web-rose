package services

import (
	"fmt"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dto"
)

// GetPortalListAccountFilter ...
func GetPortalListAccountFilter(req dto.ReqPortalListAccountFilter) models.Response {
	resp := models.Response{}
	var total int

	fmt.Println(">>> PortalListActivationService - Filter <<<")

	list, total, err := db.InitPortalListActivationDataRepository().GetDbPortalListActivation(req)
	if err != nil {
		resp.ErrCode = "05"
		resp.ErrDesc = "Error Get data List Merchant"
		return resp
	}

	contents := make([]models.PortalListMerchantAccount, 0)
	for _, data := range list {
		contents = append(contents, models.PortalListMerchantAccount{
			Merchant_id: data.Merchant_id,
			Merchant_name: data.Merchant_name,
			Mpan: data.Mpan,
			Alamat: data.Alamat,
			Kelurahan: data.Kelurahan,
			Kecamatan: data.Kecamatan,
			Kabupaten_kota: data.Kabupaten_kota,
			Provinsi: data.Provinsi,
			Profile_pict: data.Profile_pict,
			Merchant_type: data.Merchant_type,
			Merchant_outlet_id: data.Merchant_outlet_id,
			Store_phone: data.Store_phone,
			Portal_status: data.Portal_status,
			Owner_id: data.Owner_id,
			Email_owner: data.Email_owner,
			Merchant_group_id: data.Merchant_group_id,
			Merchant_group_name: data.Merchant_group_name,
			PortalCategory: data.PortalCategory,
		})
	}

	resp.ErrCode = "00"
	resp.ErrDesc = "Success"
	resp.Contents = contents
	resp.TotalData = total

	return resp
}

func FilterOutletPortal(req dto.ReqFilterOutlet) models.Response {
	var res models.Response

	data, total, err := db.FilterOutlet(req)

	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = data
	res.TotalData = total

	return res
}