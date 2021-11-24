package services

import (
	"fmt"
	"log"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"strconv"
	"time"

	"ottodigital.id/library/logger/v2"
)

type UpdatedDataMerchantService struct {
	Ottolog logger.OttologInterface
	UpdatedDataMerchantRepository *db.UpdatedDataMerchantRepository
	MerchantRepository *db.MerchantRepository
	OwnerRepository *db.OwnerRepository
	MerchantBankLoanRepository *db.MerchantBankLoanRepository
	PartnerLinkRepository *db.PartnerLinkRepository
	LookupRepository *db.LookupRepository
}

func InitUpdatedDataMerchantService(logs logger.OttologInterface) *UpdatedDataMerchantService {
	return &UpdatedDataMerchantService{
		Ottolog:logs,
		UpdatedDataMerchantRepository: db.InitUpdatedDataMerchantRepository(),
		MerchantRepository: db.InitMerchantRepository(),
		MerchantBankLoanRepository: db.InitMerchantBankLoanRepository(),
		PartnerLinkRepository: db.InitPartnerLinkRepository(logs),
		LookupRepository: db.InitLookupRepository(),
	}
}

func (svc *UpdatedDataMerchantService) Approve(req dto.ReqUpdateDataMerchantDto, res *models.Response)  {
	svc.Ottolog.Info("MerchantBankAccountService - Approve")
	user := auth.UserLogin

	data, err := svc.UpdatedDataMerchantRepository.FindByID(req.ID)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return
	}

	merchant, err :=svc.MerchantRepository.FindByMid(data.Mid)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return
	}

	var tipeBisnis dbmodels.Lookup
	tipeBisnis.Code = data.TipeBisnis
	tipeBisnis.LookupGroup = "JENIS_USAHA"
	lookupTipeBisnis, err := svc.LookupRepository.Get(tipeBisnis)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_TIPE_BISNIS_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_TIPE_BISNIS_NOTFOUND_DESC
		return
	}

	merchant.MerchantGroupId = data.MerchantGroupId
	merchant.Category = data.UserCategoryCode
	merchant.StoreName = data.StoreName
	merchant.StorePhoneNumber = data.StorePhoneNumber
	merchant.Alamat = data.Address
	merchant.Provinsi = data.Province
	merchant.KabupatenKota = data.City
	merchant.Kecamatan = data.District
	merchant.Kelurahan = data.Village
	merchant.TipeBisnis = strconv.Itoa(int(lookupTipeBisnis.Id))
	merchant.SrId = data.SrId
	merchant.Longitude = data.Longitude
	merchant.Latitude = data.Latitude
	merchant.LokasiBisnis = data.LokasiBisnis
	merchant.JenisLokasiBisnis = data.JenisLokasiBisnis
	merchant.BestVisit = data.BestVisit
	merchant.CategoryBisnis = data.CategoryBisnis
	merchant.Patokan = data.Patokan

	if err :=svc.MerchantRepository.Save(&merchant); err!=nil{
		log.Println(err)
		res.ErrCode = constants.EC_FAILED_SAVE_MERCHANT
		res.ErrDesc = constants.EC_FAILED_SAVE_MERCHANT_DESC
		return
	}
	
	if data.PartnerCustomerId != data.PartnerCustomerIdBefore {
		partnerLink, err := svc.PartnerLinkRepository.FindByIdAndCode(int(merchant.ID), data.PartnerCode)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			partnerLink.Code = data.PartnerCode
			partnerLink.PartnerId = data.PartnerCustomerId
			partnerLink.MerchantId = strconv.Itoa(int(merchant.ID))
			partnerLink.CreatedBy = auth.UserLogin.Name
			partnerLink.UpdatedAt = time.Now()
		} else {
			partnerLink.UpdatedAt = time.Now()
			partnerLink.PartnerId = data.PartnerCustomerId
		}

		if err = svc.PartnerLinkRepository.Save(&partnerLink); err != nil {
			log.Println(err)
			res.ErrCode = constants.EC_FAILED_SAVE_PARTNER_LINK
			res.ErrDesc = constants.EC_FAILED_SAVE_PARTNER_LINK_DESC
			return
		}
	}
	
	if data.LoanBankAccount != data.LoanBankAccountBefore {
		merchantBankLoan, err := svc.MerchantBankLoanRepository.FindByMidAndCode(data.Mid, data.LoanBankCode)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return
		}

		merchantBankLoan.AccountNumber = data.LoanBankAccount
		if err:=svc.MerchantBankLoanRepository.Save(&merchantBankLoan); err!=nil{
			log.Println(err)
			res.ErrCode = constants.EC_FAILED_SAVE_MERCHANT_BANK_LOAN
			res.ErrDesc = constants.EC_FAILED_SAVE_MERCHANT_BANK_LOAN_DESC
			return
		}
	}

	if data.PartnerCustomerId != data.PartnerCustomerIdBefore {
		partnerLink, err := svc.PartnerLinkRepository.FindByIdAndCode(int(merchant.ID), data.PartnerCode)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return
		}
		partnerLink.PartnerId = data.PartnerCustomerId

		if err = svc.PartnerLinkRepository.Save(&partnerLink); err != nil {
			log.Println(err)
			res.ErrCode = constants.EC_FAILED_SAVE_PARTNER_LINK
			res.ErrDesc = constants.EC_FAILED_SAVE_PARTNER_LINK_DESC
			return
		}
	}

	if data.LoanBankAccount != data.LoanBankAccountBefore {
		merchantBankLoan, err := svc.MerchantBankLoanRepository.FindByMidAndCode(data.Mid, data.LoanBankCode)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return
		}

		merchantBankLoan.AccountNumber = data.LoanBankAccount
		if err:=svc.MerchantBankLoanRepository.Save(&merchantBankLoan); err!=nil{
			log.Println(err)
			res.ErrCode = constants.EC_FAILED_SAVE_MERCHANT_BANK_LOAN
			res.ErrDesc = constants.EC_FAILED_SAVE_MERCHANT_BANK_LOAN_DESC
			return
		}
	}

	owner, err := svc.OwnerRepository.FindByID(merchant.OwnerId)
	if err != nil {
		fmt.Println("service error : ", err)
	}
	owner.OwnerTanggalExpiredId = data.ExpireDate
	owner.OwnerFirstName = data.OwnerName
	owner.OwnerAddress = data.OwnerAddress
	owner.OwnerTanggalLahir = data.OwnerTanggalLahir
	owner.OwnerNoTelp = data.OwnerPhoneNumber
	owner.OwnerNoId = data.OwnerNoId
	owner.OwnerKodePos = data.OwnerPostalCode
	owner.OwnerProvinsi = data.OwnerProvinsi
	owner.OwnerKabupaten = data.OwnerCity
	owner.OwnerKecamatan = data.OwnerKecamatan
	owner.OwnerKelurahan = data.OwnerKelurahan
	owner.OwnerRt = data.Rt
	owner.OwnerRw = data.Rw

	if _, err:=svc.OwnerRepository.Update(&owner); err!=nil{
		log.Println(err)
		res.ErrCode = constants.EC_FAILED_SAVE_OWNER
		res.ErrDesc = constants.EC_FAILED_SAVE_OWNER_DESC
		return
	}

	

	data.Status = constants.APPROVED
	data.UpdatedBy = user.Name

	if err := svc.UpdatedDataMerchantRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS


}

func (svc *UpdatedDataMerchantService) Reject(req dto.ReqUpdateDataMerchantDto, res *models.Response)  {
	svc.Ottolog.Info("MerchantBankAccountService - Reject")

	data, err := svc.UpdatedDataMerchantRepository.FindByID(req.ID)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return
	}
	user := auth.UserLogin


	data.Status = constants.REJECTED
	data.UpdatedBy = user.Name
	if err := svc.UpdatedDataMerchantRepository.Save(&data); err!=nil{
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS


}

func (svc *UpdatedDataMerchantService) Filter(req dto.ReqUpdatedDataMerchantDto, res *models.Response)  {
	svc.Ottolog.Info("UpdatedDataMerchantService - Filter")

	data, total, err := svc.UpdatedDataMerchantRepository.Filter(req)
	if err != nil {

		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND

		svc.Ottolog.Error(fmt.Sprintf("Failed to get data from database: %s", fmt.Sprintf("ERR:%s", err.Error())))
		return
	}
	log.Println(total)

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS
	res.Contents = data
	res.TotalData = total


}




