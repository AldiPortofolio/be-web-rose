package db

import (
	"fmt"
	"rose-be-go/models"
	"rose-be-go/models/dto"
	"strconv"
	"sync"

	"github.com/jinzhu/gorm"
)

// PortalListActivationDataRepository ..
type PortalListActivationDataRepository struct {
}

// InitPortalListActivationDataRepository ..
func InitPortalListActivationDataRepository() *PortalListActivationDataRepository {
	return &PortalListActivationDataRepository{}
}

var (
	query      string
	queryTotal string
)

// SaveDbLogPortal ...
func (repo *PortalListActivationDataRepository) SaveDbLogPortal(data models.LogPortal) error {
	db := GetDbCon()

	//logPortal := models.LogPortal{
	//	User:    data.User,
	//	Mid:     data.Mid,
	//	From:    data.From,
	//	To:      data.To,
	//	Action:  data.Action,
	//	Message: data.Message,
	//}

	err := db.Create(&data).Error
	if err != nil {
		fmt.Println("Error insert Log Portal...")
		return err
	}

	return nil
}

// UpdateEmailMerchant ..
func (repo *PortalListActivationDataRepository) UpdateEmailMerchant(req dto.ReqPortalActivation) error {
	db := GetDbCon()

	err := db.Table("owner").Where("id = ?", req.OwnerID).Update("owner_email", req.Email).Error

	if err != nil {
		fmt.Println("Error update email owner merchant ...")
		return err
	}

	return nil
}

// UpdatePortalCategory ..
func (repo *PortalListActivationDataRepository) UpdatePortalCategory(req dto.ReqPortalActivation) error {
	db := GetDbCon()

	err := db.Table("merchant").Where("id = ?", req.MID).Update("portal_category", req.Category).Error

	if err != nil {
		fmt.Println("Error update portalCategory merchant ...")
		return err
	}

	return nil
}

// UpdateDbPortalStatus ..
func (repo *PortalListActivationDataRepository) UpdateDbPortalStatus(req dto.ReqPortalCallback) error {
	db := GetDbCon()

	err := db.Table("merchant").Where("id = ?", req.Mid).Update("portal_status", req.PortalStatus).Error

	if err != nil {
		fmt.Println("Error update status merchant aktivasi ...")
		return err
	}

	return nil
}

// UpdateGroupPortalStatus ..
func (repo *PortalListActivationDataRepository) UpdateGroupPortalStatus(req dto.ReqPortalCallback) error {
	db := GetDbCon()

	err := db.Table("merchant_group").Where("id = ?", req.Mid).Update("portal_status", req.PortalStatus).Error

	if err != nil {
		fmt.Println("Error update status merchant aktivasi ...")
		return err
	}

	return nil
}

// GetDbPortalListActivation ...
func (repo *PortalListActivationDataRepository) GetDbPortalListActivation(req dto.ReqPortalListAccountFilter) ([]models.PortalListMerchantAccount, int, error) {
	db := GetDbCon()
	condition := " from merchant as m left join owner as o on o.id = m.owner_id left join merchant_group as mg on mg.id = m.merchant_group_id left join provinsi p on m.provinsi::INT8 = p.id left join kecamatan kec on m.kecamatan::INT8 = kec.id left join kelurahan kel on m.kelurahan::INT8 = kel.id left join dati2 dat on m.kabupaten_kota::INT8 = dat.id"
	querySelect := " select m.id, m.store_name, m.merchant_pan, m.alamat, kel.name as kelurahan, kec.name as kecamatan, dat.name as kabupaten, p.name as provinsi, m.selfie_path as profile_pict, m.merchant_type, m.store_phone_number, m.merchant_outlet_id, m.portal_status, m.owner_id, o.owner_email, m.merchant_group_id, mg.merchant_group_name, m.portal_category" + condition
	queryCount := " select count(*) " + condition

	var PortalListAcccount []models.PortalListMerchantAccount
	var total int

	if req.Search != "" {
		search := "%" + req.Search + "%"

		query = querySelect + fmt.Sprintf(" where m.store_name ilike '%s' or mg.merchant_group_name ilike '%s' or m.store_phone_number ilike '%s'", search, search, search)
		queryTotal = queryCount + fmt.Sprintf(" where m.store_name ilike '%s' or mg.merchant_group_name ilike '%s' or m.store_phone_number ilike '%s'", search, search, search)

	} else if req.Filter_by != "" {
		search := "%" + req.Keyword + "%"
		if req.Filter_by == "1" { // filter by Merchant Outlet ID
			query = querySelect + fmt.Sprintf(" where m.merchant_outlet_id ilike '%s'", search)
			queryTotal = queryCount + fmt.Sprintf(" where m.merchant_outlet_id ilike '%s'", search)
		} else if req.Filter_by == "2" { // filter by MPAN
			query = querySelect + fmt.Sprintf(" where m.merchant_pan ilike '%s'", search)
			queryTotal = queryCount + fmt.Sprintf(" where m.merchant_pan ilike '%s'", search)
		} else if req.Filter_by == "3" { // filter by Email
			query = querySelect + fmt.Sprintf(" where o.owner_email ilike '%s'", search)
			queryTotal = queryCount + fmt.Sprintf(" where o.owner_email ilike '%s'", search)
		} else if req.Filter_by == "4" { // filter by Nomor Hp
			query = querySelect + fmt.Sprintf(" where m.store_phone_number ilike '%s'", search)
			queryTotal = queryCount + fmt.Sprintf(" where m.store_phone_number ilike '%s'", search)
		}

	} else if req.FilterAction != "" { // filter by Action
		key, _ := strconv.Atoi(req.FilterAction)
		if key == 1 {
			query = querySelect + fmt.Sprintf(" where m.portal_status = %b", key)
			queryTotal = queryCount + fmt.Sprintf(" where m.portal_status = %b", key)
		} else {
			query = querySelect + fmt.Sprintf(" where m.portal_status = %b", key)
			queryTotal = queryCount + fmt.Sprintf(" where m.portal_status = %b", key)
		}

	} else if req.FilterMtype == "8079" || req.FilterMtype == "8080" || req.FilterMtype == "8081" { // filter by Merchant Type
		query = querySelect + fmt.Sprintf(" where m.merchant_type = '%s'", req.FilterMtype)
		queryTotal = queryCount + fmt.Sprintf(" where m.merchant_type = '%s'", req.FilterMtype)
	}

	if req.Filter_by == "" && req.Search == "" && req.FilterMtype == "" && req.FilterAction == "" {
		query = querySelect
		queryTotal = queryCount
	}

	wg := sync.WaitGroup{}

	wg.Add(2)
	errQuery := make(chan error)
	errCount := make(chan error)

	go repo.queryData(db, query, req.Limit, req.Page, &PortalListAcccount, errQuery)
	go repo.queryCountData(db, queryTotal, &total, errCount)

	resErrQuery := <-errQuery
	resErrCount := <-errCount

	wg.Done()

	if resErrQuery != nil {
		fmt.Println("error get data")
	}
	if resErrCount != nil {
		fmt.Println("error get count data")
	}

	return PortalListAcccount, total, nil
}

func (repo *PortalListActivationDataRepository) queryData(db *gorm.DB, query string, limit int, page int, PortalListAcccount *[]models.PortalListMerchantAccount, resChan chan error) {
	err := db.Raw(query).Limit(limit).Offset((page - 1) * limit).Scan(&PortalListAcccount).Limit(-1).Offset(-1).Error
	if err != nil {
		fmt.Println("Error get List Account Portal...")
		resChan <- err
	}
	resChan <- nil
}

func (repo *PortalListActivationDataRepository) queryCountData(db *gorm.DB, queryTotal string, total *int, resChan chan error) {
	err := db.Raw(queryTotal).Count(&*total).Error
	if err != nil {
		fmt.Println("Error SQL...", err)
		resChan <- err
	}
	resChan <- nil
}

func (repo *PortalListActivationDataRepository) setLimitOffset(db **gorm.DB, requestPage int, requestLimit int) {
	count := 10
	if requestLimit > 0 {
		count = requestLimit
	}

	offset := 0
	if requestPage > 0 {
		offset = (requestPage - 1) * count
	}

	*db = (*db).Limit(count).Offset(offset)
}

// FilterOutlet ..
func FilterOutlet(req dto.ReqFilterOutlet) ([]models.PortalOutletAccount, int, error) {
	db := GetDbCon()
	var res []models.PortalOutletAccount
	var total int

	if req.GroupName != "" {
		db = db.Where(" mg.merchant_group_name ~* ?", req.GroupName)
	}

	if req.OutletName != "" {
		db = db.Where("mo.merchant_name ~* ?", req.OutletName)
	}

	err := db.Table("merchant_outlet mo").Select("mo.id, mo.merchant_id, mo.terminal_id, m.merchant_group_id, m.merchant_outlet_id, m.merchant_pan, m.merchant_type, mo.merchant_name as outlet_name, m.store_name as merchant_name, mg.merchant_group_name, m.alamat, kel.name as kelurahan, kec.name as kecamatan, dat.name as kabupaten, p.name as provinsi, m.selfie_path as profile_pict, m.store_phone_number, o.owner_first_name as owner_name, mo.email_outlet, mo.portal_status, m.id as mid").Joins("left join merchant m on mo.merchant_id = m.id").Joins("left join merchant_group mg on m.merchant_group_id = mg.id").Joins("left join owner o on m.owner_id = o.id").Joins("left join provinsi p on m.provinsi::INT8 = p.id").Joins("left join kecamatan kec on m.kecamatan::INT8 = kec.id").Joins("left join kelurahan kel on m.kelurahan::INT8 = kel.id").Joins("left join dati2 dat on m.kabupaten_kota::INT8 = dat.id").Where("m.merchant_pan is not null").Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Order("mo.merchant_name asc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		return res, 0, err
	}

	return res, total, nil
}

// ActivationOutlet ..
func ActivationOutlet(mpan string, terminalid string, email string) int64 {
	db := GetDbCon()
	var resp models.PortalOutletAccount

	err := db.Table("merchant_outlet mo").Select("mo.id, mo.merchant_id, mo.terminal_id, m.merchant_group_id, m.merchant_outlet_id, m.merchant_pan, m.merchant_type, mo.merchant_name as outlet_name, m.store_name as merchant_name, mg.merchant_group_name, m.alamat, m.store_phone_number, o.owner_first_name as owner_name").Joins("left join merchant m on mo.merchant_id = m.id").Joins("left join merchant_group mg on m.merchant_group_id = mg.id").Joins("left join owner o on m.owner_id = o.id").Where("m.merchant_pan is not null AND m.merchant_pan = ? and mo.terminal_id = ?", mpan, terminalid).Find(&resp).Limit(-1).Offset(0).Error

	if err != nil {
		return 0
	}

	total := db.Table("merchant_outlet").Where("id = ?", resp.Id).Update("email_outlet", email).RowsAffected

	return total
}

// GetEmailMerchant ..
func (repo *PortalListActivationDataRepository) GetEmailMerchant(req dto.ReqPortalActivation) ( string, error) {
	db := GetDbCon()

	var email dto.ResEmailDto
	ownerId := strconv.Itoa(int(req.OwnerID))

	err := db.Raw("select owner_email email from owner where id = " + ownerId).Scan(&email).Error

	if err != nil {
		fmt.Println("Error get email owner merchant ...")
		return "",err
	}

	return email.Email, nil
}

// GetEmailMerchant ..
func (repo *PortalListActivationDataRepository) GetEmailGroup(id int64) ( string, error) {
	db := GetDbCon()

	var email dto.ResEmailDto
	groupId := strconv.Itoa(int(id))

	err := db.Raw("select email_portal email from merchant_group where id = " + groupId).Scan(&email).Error

	if err != nil {
		fmt.Println("Error get email owner merchant ...")
		return "",err
	}

	return email.Email, nil
}
