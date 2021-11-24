package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type BannerRepository struct {
}

func InitBannerRepository() *BannerRepository {
	return &BannerRepository{}
}

func (repo *BannerRepository) FindByID(id int64) (dbmodels.Banner, error) {
	db := GetDbCon()

	var data dbmodels.Banner

	err := db.Where(dbmodels.Banner{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *BannerRepository) Save(req *dbmodels.Banner) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (r *BannerRepository) Delete(id int) error {
	db := GetDbCon()
	var banner dbmodels.Banner
	err := db.Where("id = ?", id).Delete(&banner).Error

	if err != nil {
		return err
	}

	return nil
}
func (repo *BannerRepository) Filter(req dto.ReqBannerDto) ([]dto.ResBannerDto, int, error) {
	db := GetDbCon()
	var res []dto.ResBannerDto
	var total int
	page := req.Page
	limit := req.Limit

	db = db.Table("banner B").
		Select("B.*," +
			" UC.name as user_category_name," +
			" LM.name as level_merchant_name").
		Joins("LEFT JOIN user_category UC on UC.id = B.user_category_id").
		Joins("LEFT JOIN level_merchant LM on LM.id = B.level_merchant_id")

	if req.ID != 0 {
		db = db.Where("B.id = ?", req.ID)
	}

	if req.UserCategoryId != 0 {
		db = db.Where("B.user_category_id = ?", req.UserCategoryId)
	}

	if req.LevelMerchantId != 0 {
		db = db.Where("B.level_merchant_id = ?", req.LevelMerchantId)
	}

	if req.AdsImage != "" {
		db = db.Where("B.ads_link like ?", "%"+req.AdsLink+"%")
	}

	if req.AdsLink != "" {
		db = db.Where("B.ads_image like ?", "%"+req.AdsImage+"%")
	}

	if req.Seq != "" {
		db = db.Where("B.seq like ?", "%"+req.Seq+"%")
	}

	if req.Status != "" {
		db = db.Where("B.status like ?", "%"+req.Status+"%")
	}

	if req.BannerName != "" {
		db = db.Where("B.banner_name ilike ?", "%"+req.BannerName+"%")
	}

	err := db.Limit(limit).Offset((page - 1) * limit).Order("lower(UC.name) DESC").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}
