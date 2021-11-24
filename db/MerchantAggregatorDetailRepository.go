package db

import (
	"fmt"
	"log"
	"rose-be-go/constants/status_approval"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"strings"
)

type MerchantAggregatorDetailRepository struct {

}

func InitMerchantAggregatorDetailRepository() *MerchantAggregatorDetailRepository {
	return &MerchantAggregatorDetailRepository{}
}

func (repo *MerchantAggregatorDetailRepository) FilterListAggregatorDetail(req dto.ReqFilterDto) ([]dbmodels.MerchantAggregatorDetail, error) {
	db := GetDbCon()

	var data []dbmodels.MerchantAggregatorDetail

	if req.MidFilter != "" {
		db = db.Where("a.mid_merchant = ?", req.MidFilter)
	}

	if req.MidAggregator != "" {
		db = db.Where("a.mid_aggregator = ?", req.MidAggregator)
	}


	err := db.Table("merchant_aggregator_detail a").Select(" a.*, b.name partner_name").Joins("left join merchant_aggregator b on a.mid_aggregator = b.mid").Order("a.id asc").Find(&data).Error // query

	if err != nil {
		log.Println("err -> ", err)
		fmt.Println("<<< Error get data Aggregator Detail >>>")
		return data, err
	}

	//fmt.Println("<<< Success get data Aggregator Detail >>> {}", data)

	return data, nil
}

func (repo *MerchantAggregatorDetailRepository) FilterDataAggregatorDetail(req dto.ReqFilterDto) ([]dbmodels.MerchantAggregatorDetail, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit

	var data []dbmodels.MerchantAggregatorDetail
	var total int

	if req.MidFilter != "" {
		db = db.Where("a.mid_merchant = ?", req.MidFilter)
	}

	fmt.Println("midFilter-->", req.MidFilter)

	err := db.Table("merchant_aggregator_detail a").Select(" a.*, b.store_name merchant_name, c.name partner_name").Joins("left join merchant b on a.mid_merchant = b.merchant_outlet_id").Joins("left join merchant_aggregator c on a.mid_aggregator = c.mid").Limit(limit).Offset((page - 1) * limit).Order("a.id asc").Find(&data).Limit(-1).Offset(0).Count(&total).Error // query

	if err != nil {
		log.Println("err -> ", err)
		fmt.Println("<<< Error get data Aggregator Detail >>>")
		return data, 0, err
	}

	return data, total, nil
}

func (repo *MerchantAggregatorDetailRepository) SaveMerchantAggregatorDetailTemp(req *dbmodels.MerchantAggregatorDetailTemp) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *MerchantAggregatorDetailRepository) FindMerchantAggregatorDetailTempByIdDetail(id int64) (dbmodels.MerchantAggregatorDetailTemp, error) {
	db := GetDbCon()

	var res dbmodels.MerchantAggregatorDetailTemp

	err := db.Where("merchant_aggregator_detail_id = ? and action_type in (?,?,?)", id, status_approval.CREATE, status_approval.EDIT, status_approval.DELETE).First(&res).Error

	return res, err
}

func (repo *MerchantAggregatorDetailRepository) FindListMerchantAggregatorApproval(req dto.ReqFilterDto) ([]dbmodels.MerchantAggregator, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit
	var total int

	name := fmt.Sprintf("%v%v%v", "%",req.Name,"%")
	log.Println("name ==> ", name)

	var res []dbmodels.MerchantAggregator

	err := db.Table("merchant_aggregator a").Select("a.*").
				Joins("JOIN merchant_aggregator_detail_temp	b on b.mid_aggregator = a.mid").
					Where("lower(a.name) like ?", strings.ToLower(name)).
					Where("b.action_type in (?,?,?)", status_approval.CREATE, status_approval.EDIT, status_approval.DELETE).
				Group("a.id").

			Limit(limit).Offset((page - 1) * limit).Order("id asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	return res,total, err
}

func (repo *MerchantAggregatorDetailRepository) FindMerchantAggregatorDetailTempByMidAggregator(req dto.ReqFilterDto) ([]dbmodels.MerchantAggregatorDetailTemp, int, error) {
	db := GetDbCon()

	page := req.Page
	limit := req.Limit
	var total int
	if page == 0 {
		limit = 9999999
	}

	var res []dbmodels.MerchantAggregatorDetailTemp

	err := db.Where("mid_aggregator = ? and action_type in (?,?,?)", req.MidAggregator, status_approval.CREATE, status_approval.EDIT, status_approval.DELETE).Limit(limit).Offset((page - 1) * limit).Order("id asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	return res,total, err
}

func (repo *MerchantAggregatorDetailRepository) FindMerchantApprovalByMidAggregator(req dto.ReqFilterDto) ([]dbmodels.Merchant, int, error) {
	db := GetDbCon()
	page := req.Page
	limit := req.Limit
	var total int

	var res []dbmodels.Merchant

	err := db.Table("merchant_aggregator_detail_temp b").Select("a.id, a.store_name, a.merchant_pan, a.merchant_outlet_id, a.n_mid, b.action_type").
		Joins("LEFT JOIN merchant a on b.mid_merchant = a.merchant_outlet_id").
		Where("b.action_type in (?,?,?)", status_approval.CREATE, status_approval.EDIT, status_approval.DELETE).
		Where("b.mid_aggregator = ?", req.MidAggregator).


		Limit(limit).Offset((page - 1) * limit).Order("id asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	//err := db.Where("mid_aggregator = ? and action_type in (?,?)", req.MidAggregator, status_approval.CREATE, status_approval.EDIT).Limit(limit).Offset((page - 1) * limit).Order("id asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	return res,total, err
}

func (repo *MerchantAggregatorDetailRepository) FindMerchantAggregatorDetailTempById(id int64) (dbmodels.MerchantAggregatorDetailTemp, error) {
	db := GetDbCon()
	var res dbmodels.MerchantAggregatorDetailTemp
	err := db.Where("id = ?", id).First(&res).Error

	return res, err
}

func (repo *MerchantAggregatorDetailRepository) FindMerchantAggregatorDetailById(id int64) (dbmodels.MerchantAggregatorDetail, error) {
	db := GetDbCon()
	var res dbmodels.MerchantAggregatorDetail
	err := db.Where("id = ?", id).First(&res).Error

	return res, err
}

func (repo *MerchantAggregatorDetailRepository) SaveMerchantAggregatorDetail(req *dbmodels.MerchantAggregatorDetail) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *MerchantAggregatorDetailRepository) UpdateMerchantAggregatorDetail(midAgg string, midMerchant string) error {
	db := GetDbCon()

	err := db.Table("merchant_aggregator_detail").Where("mid_aggregator = ? AND mid_merchant = ?", midAgg, midMerchant).Update("status, 0").Error

	return err
}

func (repo *MerchantAggregatorDetailRepository) FilterMerchantAggList(req dto.ReqFilterDto) ([]dbmodels.MerchantAggregatorDetail, int, error) {
	db := GetDbCon()

	var data []dbmodels.MerchantAggregatorDetail
	var total int

	page := req.Page
	limit := req.Limit

	if page == 0 {
		limit = 9999999999
	}

	db = db.Where("mad.mid_aggregator = ?", req.MidAggregator)

	if req.MidFilter != "" {
		db = db.Where("mad.mid_merchant ilike ?", "%" + req.MidFilter + "%")
	}

	if req.Name != "" {
		db = db.Where("m.store_name ilike ?","%"+req.Name+"%")
	}

	if req.MpanMerchant != "" {
		db = db.Where("m.merchant_pan = ?", "%" + req.MpanMerchant + "%")
	}

	if req.NmidMerchant != "" {
		db = db.Where("m.n_mid = ?", req.NmidMerchant)
	}

	err := db.Table("merchant_aggregator_detail mad").Select("mad.*, m.store_name as merchant_name, m.merchant_pan, m.n_mid as merchant_nmid").Joins("left join merchant m on mad.mid_merchant = m.merchant_outlet_id").Where("mad.status = 1").Limit(limit).Offset((page - 1) * limit).Order("mad.id desc").Find(&data).Offset(0).Limit(-1).Count(&total).Error

	if err != nil {
		return data, 0, err
	}

	return data, total, nil

}

func (repo *MerchantAggregatorDetailRepository) FilterMerchantAggTempList(req dto.ReqFilterDto) ([]dbmodels.MerchantAggregatorDetailTemp, int, error) {
	db := GetDbCon()

	var data []dbmodels.MerchantAggregatorDetailTemp
	var total int

	db = db.Where("mad.mid_aggregator = ? AND mad.merchant_detail_aggregator_id = 0 AND mad.action_type = 0", req.MidAggregator)

	if req.MidFilter != "" {
		db = db.Where("mad.mid_merchant = ?", req.MidFilter)
	}

	if req.Name != "" {
		db = db.Where("m.store_name ilike ?","%"+req.Name+"%")
	}

	if req.MpanMerchant != "" {
		db = db.Where("m.merchant_pan = ?", req.MpanMerchant)
	}

	if req.NmidMerchant != "" {
		db = db.Where("m.n_mid = ?", req.NmidMerchant)
	}

	err := db.Table("merchant_aggregator_detail_temp mad").Select("mad.*, m.store_name as merchant_name, m.merchant_pan, m.n_mid as merchant_nmid").Joins("left join merchant m on mad.mid_merchant = m.merchant_outlet_id").Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Order("mad.id desc").Find(&data).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return data, 0, err
	}

	return data, total, nil

}

