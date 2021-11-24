package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"github.com/unknwon/com"
)

// MerchantCustomerRepository ..
type MerchantCustomerRepository struct {
	// struct attributes
}

// InitMerchantCustomerRepository ..
func InitMerchantCustomerRepository() *MerchantCustomerRepository {
	return &MerchantCustomerRepository{}
}

// Filter ..
func (repo *MerchantCustomerRepository) Filter(req dto.ReqMerchantCustomerDto) ([]dbmodels.CustomMerchantCustomer, int, error) {
	db := GetDbCon()
	var res []dbmodels.CustomMerchantCustomer
	var total int
	page := req.Page
	limit := req.Limit

	// select
	db = db.Select("mc.*, ms.name AS service_name, mt.name AS type_name," +
		"pr.name AS province_name, dt.name AS city_name, kc.name AS district_name, kl.name as village_name")

	// from
	db = db.Table("merchant_customer mc")
	db = db.Joins("JOIN master_service ms ON mc.service_id = ms.id")
	db = db.Joins("JOIN master_type mt ON mc.type_id = mt.id")
	db = db.Joins("LEFT JOIN kelurahan kl ON mc.village_id = kl.id")
	db = db.Joins("LEFT JOIN kecamatan kc ON mc.district_id = kc.id OR kl.kecamatan_id = kc.id")
	db = db.Joins("LEFT JOIN dati2 dt ON mc.city_id = dt.id OR kc.dati2_id = dt.id")
	db = db.Joins("LEFT JOIN provinsi pr ON mc.province_id = pr.id OR dt.provinsi_id = pr.id")

	// where
	if req.ServiceID > 0 {
		db = db.Where("mc.service_id = ?", req.ServiceID)
	}
	if req.Name != "" {
		db = db.Where("mc.name ilike ?", "%"+req.Name+"%")
	}
	if req.Phone != "" {
		db = db.Where("mc.phone = ?", req.Phone)
	}
	if req.TypeID > 0 {
		db = db.Where("mc.type_id = ?", req.TypeID)
	}
	if req.CitizenIdNo != "" {
		db = db.Where("mc.citizen_id_no = ?", req.CitizenIdNo)
	}
	if req.Merchant != "" {
		db = db.Where("mc.merchant = ?", req.Merchant)
	}
	if req.ProvinceID > 0 {
		db = db.Where("mc.province_id = ?", req.ProvinceID)
	}
	if req.CityID > 0 {
		db = db.Where("mc.city_id = ?", req.CityID)
	}
	if req.DistrictID > 0 {
		db = db.Where("mc.district_id = ?", req.DistrictID)
	}
	if req.VillageID > 0 {
		db = db.Where("mc.village_id = ?", req.VillageID)
	}

	// order by
	db = db.Order("mc.id ASC")

	err := db.Limit(limit).Offset((page - 1) * limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	return res, total, err
}

// FindAll ..
func (repo *MerchantCustomerRepository) FindAll() ([]dbmodels.CustomMerchantCustomer, int, error) {
	db := GetDbCon()
	var res []dbmodels.CustomMerchantCustomer
	var total int

	// select
	db = db.Select("mc.*, ms.name AS service_name, mt.name AS type_name," +
		"pr.name AS province_name, dt.name AS city_name, kc.name AS district_name, kl.name as village_name")

	// from
	db = db.Table("merchant_customer mc")
	db = db.Joins("JOIN master_service ms ON mc.service_id = ms.id")
	db = db.Joins("JOIN master_type mt ON mc.type_id = mt.id")
	db = db.Joins("LEFT JOIN kelurahan kl ON mc.village_id = kl.id")
	db = db.Joins("LEFT JOIN kecamatan kc ON mc.district_id = kc.id OR kl.kecamatan_id = kc.id")
	db = db.Joins("LEFT JOIN dati2 dt ON mc.city_id = dt.id OR kc.dati2_id = dt.id")
	db = db.Joins("LEFT JOIN provinsi pr ON mc.province_id = pr.id OR dt.provinsi_id = pr.id")

	// order by
	db = db.Order("mc.id ASC")

	err := db.Find(&res).Count(&total).Error

	return res, total, err
}

// Export ...
func (repo *MerchantCustomerRepository) Export(data []dbmodels.CustomMerchantCustomer) ([]dbmodels.ExportCustomMerchants) {
	var file []dbmodels.ExportCustomMerchants

	for _, value := range data {
		//no := strconv.Itoa(i + 1)

		row := dbmodels.ExportCustomMerchants{
			Service:      value.ServiceName,
			Name:         value.Name,
			PhoneNumber:  value.Phone,
			Type:         value.TypeName,
			CitizenIdNo:  value.CitizenIdNo,
			Merchant:     value.Merchant,
			Dob:          value.Pob + ", " + value.Dob.Format("2006-01-02"),
			Gender:       value.Gender,
			Address:      value.Address,
			Province:     value.ProvinceName,
			City:         value.CityName,
			District:     value.DistrictName,
			Village:      value.VillageName,
			VillageId:    com.ToStr(value.VillageID),
			RTRW:         value.RT + "/" + value.RW,
			Occupation:   value.Occupation,
			Longitude:    com.ToStr(value.Longitude),
			Latitude:     com.ToStr(value.Latitude),
		}

		file = append(file, row)

	}

	return file
}