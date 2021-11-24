package services

import (
	"encoding/json"
	"fmt"
	"ottodigital.id/library/logger/v2"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/constants/status_push_notif"
	"rose-be-go/db"
	"rose-be-go/host/op_bank"
	"rose-be-go/host/ottopay"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/models/opbankmodels"
	"rose-be-go/models/ottopaymodels"
	"time"
)

type MerchantBankAccountService struct {
	Ottolog logger.OttologInterface
	MerchantBankAccountRepository *db.MerchantBankAccountRepository
	MerchantRepository *db.MerchantRepository
	BankListRepository *db.BankListRepository
	OttopayHost *ottopay.OttopayHost
	OpBankHost *op_bank.OpBankHost
}

func InitMerchantBankAccountService(logs logger.OttologInterface) *MerchantBankAccountService {
	return &MerchantBankAccountService{
		Ottolog:logs,
		MerchantBankAccountRepository: db.InitMerchantBankAccountRepository(logs),
		MerchantRepository: db.InitMerchantRepository(),
		BankListRepository: db.InitBankListRepository(logs),
		OttopayHost: ottopay.InitOttopayHost(logs),
		OpBankHost: op_bank.InitOpBankHost(),
	}
}

func (svc *MerchantBankAccountService) Save(req dto.ReqMerchantBankAccountDto, res *models.Response)  {
	svc.Ottolog.Info("MerchantBankAccountService - Save")
	var data dbmodels.MerchantBankAccount
	var err error

	data.CreatedAt = time.Now()
	if req.ID >0  {
		data, err = svc.MerchantBankAccountRepository.FindByID(req.ID)
		if err != nil {
			res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
			return
		}
	}
	data.BankCode = req.BankCode
	data.AccountName = req.AccountName
	data.AccountNumber = req.AccountNumber
	data.Notes = req.Notes
	data.Mid = req.Mid
	data.Status = constants.PENDING
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := svc.MerchantBankAccountRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS

}

func (svc *MerchantBankAccountService) FindAllAccount(mid string, res *models.Response)  {
	svc.Ottolog.Info("MerchantBankAccountService - FindAllAccount")


	data, err := svc.MerchantBankAccountRepository.GetDataByMid(mid)
	if err != nil {

		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND

		svc.Ottolog.Error(fmt.Sprintf("Failed to get data from database: %s", fmt.Sprintf("ERR:%s", err.Error())))
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS
	res.Contents = data

}

func (svc *MerchantBankAccountService) FilterApproval(req dto.ReqMerchantBankAccountDto, res *models.Response)  {
	svc.Ottolog.Info("MerchantBankAccountService - FilterApproval")

	data, total, err := svc.MerchantBankAccountRepository.FilterApproval(req)
	if err != nil {

		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND

		svc.Ottolog.Error(fmt.Sprintf("Failed to get data from database: %s", fmt.Sprintf("ERR:%s", err.Error())))
		return
	}

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS
	res.Contents = data
	res.TotalData = total

}

func (svc *MerchantBankAccountService) Approve(req dto.ReqApprovalMerchantBankAccountDto, res *models.Response)  {
	svc.Ottolog.Info("MerchantBankAccountService - Approve")

	//var data dbmodels.MerchantBankAccount
	//var err error

	data, err := svc.MerchantBankAccountRepository.FindByID(req.ID)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return
	}

	merchant,bank, err := svc.GetMerchantAndBank(data.Mid, data.BankCode)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC

		return
	}

	data.Notes = req.Notes
	data.Status = constants.APPROVED
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := svc.MerchantBankAccountRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	reqPushNotif := models.PushNotifBankAccount{
		StorePhoneNumber:merchant.StorePhoneNumber,
		BankName: bank.FullName,
		AccountNumber: data.AccountNumber,
		StatusApprove: constants.APPROVED,
		Notes: req.Notes,
	}

	if err := svc.PushNotif(reqPushNotif); err != nil{
		svc.UpdateStatusPushNotif(data,  status_push_notif.FAILED, reqPushNotif)
		res.ErrDesc = constants.EC_FAIL_SEND_PUSH_NOTIF_DESC
		res.ErrCode = constants.EC_FAIL_SEND_PUSH_NOTIF
		return
	}
	svc.UpdateStatusPushNotif(data,  status_push_notif.SUCCESS, reqPushNotif)

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS

}

func (svc *MerchantBankAccountService)GetMerchantAndBank(mid, bankCode string) (dbmodels.Merchant, dbmodels.BankList, error) {
	merchant, err := svc.MerchantRepository.FindByMid(mid)
	if err != nil {
		return dbmodels.Merchant{}, dbmodels.BankList{}, err
	}

	bank, err := svc.BankListRepository.FindByCode(bankCode)
	if err != nil {
		return dbmodels.Merchant{}, dbmodels.BankList{}, err
	}

	return merchant, bank, nil
}

func (scv *MerchantBankAccountService) UpdateStatusPushNotif(data dbmodels.MerchantBankAccount, status string, req models.PushNotifBankAccount) error {

	dataByte,_ := json.Marshal(req)
	data.PushNotifStatus = status
	data.PushNotifData = string(dataByte)

	return scv.MerchantBankAccountRepository.Save(&data)

}

func (svc *MerchantBankAccountService) PushNotif(req models.PushNotifBankAccount) (error) {

	reqPushNotif := ottopaymodels.ReqPushNotif{
		Title: "Approval Bank Account",
		Target: "inbox",
		CustAccount: req.StorePhoneNumber,
		Desc: fmt.Sprintf("Your bank account %s %s has been %s. %s",req.BankName, req.AccountNumber, req.StatusApprove, req.Notes),
	}

	_, err := svc.OttopayHost.Send(reqPushNotif, constants.SEND_NOTIF)
	if err != nil {
		fmt.Println("failed send push notif")
		return err

	}
	return nil
}

func (svc *MerchantBankAccountService) Reject(req dto.ReqApprovalMerchantBankAccountDto, res *models.Response)  {
	svc.Ottolog.Info("MerchantBankAccountService - Reject")

	//var data dbmodels.MerchantBankAccount
	//var err error

	data, err := svc.MerchantBankAccountRepository.FindByID(req.ID)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return
	}

	merchant,bank, err := svc.GetMerchantAndBank(data.Mid, data.BankCode)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC

		return
	}

	data.Notes = req.Notes
	data.Status = constants.REJECTED
	data.UpdatedAt = time.Now()
	data.UpdatedBy = auth.UserLogin.Name

	if err := svc.MerchantBankAccountRepository.Save(&data); err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return
	}

	reqPushNotif := models.PushNotifBankAccount{
		StorePhoneNumber:merchant.StorePhoneNumber,
		BankName: bank.FullName,
		AccountNumber: data.AccountNumber,
		StatusApprove: constants.REJECTED,
		Notes: req.Notes,
	}

	if err := svc.PushNotif(reqPushNotif); err != nil{
		svc.UpdateStatusPushNotif(data,  status_push_notif.FAILED, reqPushNotif)
		res.ErrDesc = constants.EC_FAIL_SEND_PUSH_NOTIF_DESC
		res.ErrCode = constants.EC_FAIL_SEND_PUSH_NOTIF
		return
	}

	svc.UpdateStatusPushNotif(data,  status_push_notif.SUCCESS, reqPushNotif)

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS

}

func (svc *MerchantBankAccountService) ResendPushNotif(req dto.ReqApprovalMerchantBankAccountDto, res *models.Response)  {
	svc.Ottolog.Info("MerchantBankAccountService - ResendPushNotif")
	
	data, err := svc.MerchantBankAccountRepository.FindByID(req.ID)
	if err != nil {
		res.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
		res.ErrDesc = constants.EC_FAIL_DATA_NOTFOUND_DESC
		return
	}

	var reqPushNotif models.PushNotifBankAccount

	json.Unmarshal([]byte(data.PushNotifData), &reqPushNotif)


	if err := svc.PushNotif(reqPushNotif); err != nil{
		svc.UpdateStatusPushNotif(data,  status_push_notif.FAILED, reqPushNotif)
		res.ErrDesc = constants.EC_FAIL_SEND_PUSH_NOTIF_DESC
		res.ErrCode = constants.EC_FAIL_SEND_PUSH_NOTIF
		return
	}

	svc.UpdateStatusPushNotif(data,  status_push_notif.SUCCESS, reqPushNotif)

	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.ErrCode = constants.ERR_SUCCESS

}

func (svc *MerchantBankAccountService) ValidationBankAccount(req dto.ReqValidationBankAccount, res *models.Response)  {
	fmt.Println("MerchantBankAccountService - ValidationBankAccount")


	reqInq := opbankmodels.ReqInquiry{
		AccountNo: req.AccountNo,
		BankCode: req.BankCode,
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	merchant, err := db.InitMerchantRepository().FindByMid(req.Mid)
	if err != nil {
		svc.Ottolog.Error("Failed get data merchant")
	}

	owner, err := db.InitOwnerRepository().FindByID(merchant.OwnerId)
	if err != nil {
		svc.Ottolog.Error("Failed get data owner")
	}

	if req.BankCode == constants.MANDIRI_BANK_CODE {
		resInq, _ := svc.OpBankHost.Send(reqInq, constants.INQUIRY_INTERNAL)
		var resData opbankmodels.ResInquiryInternal
		json.Unmarshal(resInq, &resData)

		data := dto.ResValidationBankAccountDto{
			AccountNo: resData.Data.AccountNo,
			AccountName: resData.Data.AccountName,
			AccountBankName: "PT. BANK MANDIRI, TBK",
			BankCode: reqInq.BankCode,
			OwnerFirstName: owner.OwnerFirstName,
			OwnerLastName: owner.OwnerLastName,
		}

		res.Contents = data
		return
	}

	resInq, _ := svc.OpBankHost.Send(reqInq, constants.INQUIRY_EXTERNAL)
	var resData opbankmodels.ResInquiryExternal
	json.Unmarshal(resInq, &resData)

	
	
	data := dto.ResValidationBankAccountDto{
		AccountNo: resData.Data.DestinationAccountNo,
		AccountName: resData.Data.DestinationAccountName,
		AccountBankName: resData.Data.DestinationBankName,
		BankCode: reqInq.BankCode,
		OwnerFirstName: owner.OwnerFirstName,
		OwnerLastName: owner.OwnerLastName,
	}

	res.Contents = data
	
}

