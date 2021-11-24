package db

import (
	"fmt"
	"log"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type MerchantRepository struct {

}

func InitMerchantRepository() *MerchantRepository {
	return &MerchantRepository{}
}

func (r *MerchantRepository) Save(merchant *dbmodels.Merchant) (err error) {
	db:= GetDbCon()

	return  db.Save(&merchant).Error

}

func (repo *MerchantRepository) Filter(req dto.ReqMerchantDto) ([]dbmodels.Merchant, int, error) {

	db := GetDbCon()
	var res []dbmodels.Merchant
	var total int
	page := req.Page
	limit := req.Limit

	if req.MerchantGroupID != 0 {
		db = db.Where("merchant_group_id = ? ", req.MerchantGroupID)
	}
	if req.StoreName != "" {
		db = db.Where("store_name ilike ?", "%" + req.StoreName + "%")
	}

	err := db.Limit(limit).Offset((page-1)*limit).Order("store_name asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error


	if err != nil {
		return res, 0, err
	}

	return res, total, nil

}

func (repo *MerchantRepository) Get(data dbmodels.Merchant) (dbmodels.Merchant, error) {
	db := GetDbCon()
	var res dbmodels.Merchant

	err := db.Where(&data).First(&res).Error

	return res, err
}

func (repo *MerchantRepository) FindByMid(mid string) (dbmodels.Merchant, error) {
	db:= GetDbCon()
	var res dbmodels.Merchant

	err := db.Where("merchant_outlet_id = ?", mid).First(&res).Error

	return res, err
}

func (repo *MerchantRepository) FindById(id int) (dbmodels.Merchant, error) {
	db:= GetDbCon()
	var res dbmodels.Merchant

	err := db.Where("id = ?", id).First(&res).Error

	return res, err
}

//GetListByGroupId...
func (repo *MerchantRepository) GetListByGroupId(req dto.ReqMerchantGroupDto) ([]dbmodels.Merchant, int, error){
	fmt.Println(">>> Get List Merchant By Group Id - DB <<<")
	var merchants []dbmodels.Merchant
	var total int
	var portalStatus int

	db:= GetDbCon()

	if req.StoreName != "" {
		db = db.Where("LOWER(store_name) LIKE ?", "%" + strings.ToLower(req.StoreName) + "%")
	}

	if req.PortalStatus != "" {
		switch req.PortalStatus {
		case "0":
			portalStatus = 0
		case "1":
			portalStatus = 1
		}

		db = db.Where("portal_status = ?", portalStatus)
	}

	if err := db.Where("merchant_group_id = ?", req.MerchantGroupID).Offset((req.Page-1)*req.Limit).Limit(req.Limit).Find(&merchants).Offset(0).Limit(-1).Count(&total).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		fmt.Println("===============> Error Get Merchant : ", err)
		return merchants, total, err
	}

	return merchants, total, nil
}

func (repo *MerchantRepository) FindTanggalAkuisisi(phone string)(tglAkuisisi time.Time, err error)  {

	db := GetDbCon()

	var merchant struct{
		TanggalSalesAkuisisi time.Time `json:"tanggalSalesAkuisisi"`
	}
	//var tglAkuisisi time.Time

	err =db.Table("merchant").Select("tanggal_sales_akuisisi").Where("store_phone_number = ?", phone).First(&merchant).Error
	tglAkuisisi = merchant.TanggalSalesAkuisisi

	return tglAkuisisi, err
}

func (repo *MerchantRepository) FilterDashboard(req dto.ReqDashboardMerchantDto) ([]dbmodels.DashboardMerchant, int, error) {

	db := GetDbCon()
	var data []dbmodels.DashboardMerchant
	
	page := req.Page
	limit := req.Limit

	filterName := req.Name
	filterStatus := req.PortalStatus
	filterType := req.TipeMerchant
	var queryFilter string

	if filterName != "" && filterStatus != "" && filterType != "" {
		queryFilter = "Where a.portal_status =  '" + filterStatus + "' and e.name = '" + filterType + "' and ( lower(a.merchant_outlet_id) ilike lower('%"+ filterName +"%')" + " or lower(a.merchant_pan) ilike lower('%"+ filterName +"%')" + " or a.store_phone_number ilike '%"+ filterName + "%' or lower(a.store_name) ilike lower('%"+ filterName +"%'))"
	} else if filterName == "" && filterStatus != "" && filterType != "" {
		queryFilter = "Where a.portal_status =  '" + filterStatus + "' and e.name = '" + filterType +"'"
	} else if filterName == "" && filterStatus == "" && filterType != "" {
		queryFilter = "Where e.name = '" + filterType + "'"
	} else if filterName != "" && filterStatus != "" && filterType == "" {
		queryFilter = "Where a.portal_status =  '" + filterStatus + "' and ( lower(a.merchant_outlet_id) ilike lower('%"+ filterName +"%')" + " or lower(a.merchant_pan) ilike lower('%"+ filterName +"%')" + " or a.store_phone_number ilike '%"+ filterName + "%' or lower(a.store_name) ilike lower('%"+ filterName +"%'))"
	} else if filterName != "" && filterStatus == "" && filterType == "" {
		queryFilter = "Where lower(a.merchant_outlet_id) ilike lower('%"+ filterName +"%')" + " or lower(a.merchant_pan) ilike lower('%"+ filterName +"%')" + " or a.store_phone_number ilike '%"+ filterName + "%' or lower(a.store_name) ilike lower('%"+ filterName +"%')"
	} else if filterName == "" && filterStatus != "" && filterType == "" {
		queryFilter = "Where a.portal_status =  '" + filterStatus + "'"
	} else if filterName != "" && filterStatus == "" && filterType != "" {
		queryFilter = "Where e.name =  '" + filterType + "' and ( lower(a.merchant_outlet_id) ilike lower('%"+ filterName +"%')" + " or lower(a.merchant_pan) ilike lower('%"+ filterName +"%')" + " or a.store_phone_number ilike '%"+ filterName + "%' or lower(a.store_name) ilike lower('%"+ filterName +"%'))"
	} 


	// err := db.Raw("select a.id, a.store_name, d.merchant_group_name merchant_group, e.name merchant_type, f.name jenis_usaha, a.merchant_outlet_id mid, a.merchant_pan mpan, a.status_suspense, c.status_registration approval_status, a.portal_status " +
	// 		"from Merchant a " +
	// 		"LEFT JOIN merchant_wip b ON b.id_merchant = A.ID and b.wip_status = 1 " +
	// 		"LEFT JOIN wip_queue C ON b.id = C.wip_id and c.status = 0 " +
	// 		"LEFT JOIN merchant_group d ON A.merchant_group_id = d.ID " +
	// 		"LEFT JOIN lookup e on a.merchant_type = CAST(e.id as VARCHAR) " +
	// 		"LEFT JOIN lookup f on a.jenis_usaha =  CAST(f.id as VARCHAR) " +
	// 		queryFilter).Limit(limit).Offset((page-1)*limit).Order("store_name asc").
	// 		Scan(&data).Limit(-1).Offset(0).Error

	err := db.Raw( "select a.id, a.store_name, d.merchant_group_name merchant_group, e.name merchant_type, f.name jenis_usaha, a.merchant_outlet_id mid, a.merchant_pan mpan, a.status_suspense, c.status_registration approval_status, a.portal_status " +
			"from Merchant a " +
			"LEFT JOIN merchant_wip b ON b.id_merchant = A.ID and b.wip_status = 1 " +
			"LEFT JOIN wip_queue C ON b.id = C.wip_id and c.status = 0 "+
			"LEFT JOIN merchant_group d ON A.merchant_group_id = d.ID "+
			"LEFT JOIN (select id,name from lookup where lookup_group = 'TIPE_MERCHANT') e on a.merchant_type = CAST(e.id as VARCHAR) "+
			"LEFT JOIN (select id,name from lookup where lookup_group = 'JENIS_USAHA') f on a.jenis_usaha = CAST(f.id as VARCHAR) "+
			queryFilter).Limit(limit).Offset((page-1)*limit).Order("store_name asc").
			Scan(&data).Limit(-1).Offset(0).Error
	// err := db.Limit(limit).Offset((page-1)*limit).Order("store_name asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error
	if err != nil {
		log.Println("err get db :", err, data)
		return data, 0, err
	}
	
	type Count struct {
		Count int `json:"count"`
	}
	var total Count
	
	// err = db.Raw("select count (a.id)" +
	// 	"from Merchant a " +
	// 	"LEFT JOIN merchant_wip b ON b.id_merchant = A.ID and b.wip_status = 1 " +
	// 	"LEFT JOIN wip_queue C ON b.id = C.wip_id and c.status = 0 " +
	// 	"LEFT JOIN merchant_group d ON A.merchant_group_id = d.ID " +
	// 	"LEFT JOIN lookup e on a.merchant_type = CAST(e.id as VARCHAR) " +
	// 	"LEFT JOIN lookup f on a.jenis_usaha =  CAST(f.id as VARCHAR) " +
	// 	queryFilter).Scan(&total).Error

	err = db.Raw("select count (a.id)" +
		"from Merchant a " +
			"LEFT JOIN merchant_wip b ON b.id_merchant = A.ID and b.wip_status = 1 " +
			"LEFT JOIN wip_queue C ON b.id = C.wip_id and c.status = 0 "+
			"LEFT JOIN merchant_group d ON A.merchant_group_id = d.ID "+
			"LEFT JOIN (select id,name from lookup where lookup_group = 'TIPE_MERCHANT') e on a.merchant_type = CAST(e.id as VARCHAR) "+
			"LEFT JOIN (select id,name from lookup where lookup_group = 'JENIS_USAHA') f on a.jenis_usaha = CAST(f.id as VARCHAR) "+
			queryFilter).Scan(&total).Error
	
	log.Println("ini total data :", total)
	if err != nil {
		log.Println("err get db :", err, data, total)
		return data, 0, err
	}
	log.Println("success get db :", err, data, total)
	return data, total.Count, nil

}

func (repo *MerchantRepository) GetMerchantDetailByID(id int64)( dbmodels.MerchantDetail, error)  {

	db := GetDbCon()

	var res dbmodels.MerchantDetail
	err := db.Where("id = ?", id).First(&res).Error
	

	return res, err
}

