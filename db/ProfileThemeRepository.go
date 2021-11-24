package db

import (
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
)

type ProfileThemeRepository struct {

}

func InitProfileThemeRepository()  *ProfileThemeRepository{
	return &ProfileThemeRepository{}
}

func (repo *ProfileThemeRepository) FindByID(id int64) (dbmodels.ProfileTheme, error) {
	db := GetDbCon()

	var data dbmodels.ProfileTheme

	err := db.Where(dbmodels.ProfileTheme{
		ID: id,
	}).First(&data).Error

	return data, err
}

func (repo *ProfileThemeRepository) Save(req *dbmodels.ProfileTheme) error {
	db := GetDbCon()

	err := db.Save(&req).Error

	return err
}

func (repo *ProfileThemeRepository) Filter(req dto.ReqProfileThemeDto) ([]dto.ResProfileThemeDto, int, error) {
	db := GetDbCon()
	var res []dto.ResProfileThemeDto
	var total int
	page := req.Page
	limit := req.Limit

	db = db.Table("profile_theme PT").
		Select("PT.*," +
			" UC.name as user_category_name," +
			" LM.name as level_merchant_name").
		Joins("LEFT JOIN user_category UC on UC.id = PT.user_category_id").
		Joins("LEFT JOIN level_merchant LM on LM.id = PT.level_merchant_id")

	if req.ID  != 0 {
		db = db.Where("PT.id = ?",  req.ID )
	}

	if req.UserCategoryId  != 0 {
		db = db.Where("PT.user_category_id = ?",  req.UserCategoryId )
	}

	if req.LevelMerchantId  != 0 {
		db = db.Where("PT.level_merchant_id = ?",  req.LevelMerchantId )
	}

	if req.DashboardTopBackground != "" {
		db = db.Where("PT.dashboard_top_background like ?", "%" + req.DashboardTopBackground + "%")
	}

	if req.ThemeColor != "" {
		db = db.Where("PT.theme_color like ?", "%" + req.ThemeColor + "%")
	}

	if req.DashboardLogo != "" {
		db = db.Where("PT.dashboard_logo like ?", "%" + req.DashboardLogo + "%")
	}

	if req.DashboardText != "" {
		db = db.Where("PT.dashboard_text like ?", "%" + req.DashboardText + "%")
	}

	if req.ProfileBackgroundImage != "" {
		db = db.Where("PT.profile_background_image like ?", "%" + req.ProfileBackgroundImage + "%")
	}

	if req.Status != "" {
		db = db.Where("PT.status like ?", "%" + req.Status + "%")
	}

	err := db.Limit(limit).Offset((page-1)*limit).Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}