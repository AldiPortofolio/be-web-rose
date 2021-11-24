package services

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	//"gopkg.in/gomail.v2"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/redis/redis_single"

	"ottodigital.id/library/ottomail"
	"ottodigital.id/library/utils"

	//"strconv"
	"strings"
	"time"
)

// var (
// 	userName string
// )

// BpActivation ...
func BpActivation(reqData dto.BpActivationReq) (res models.Response, key string) {
	fmt.Println(">>> Portal Activation Group - Function <<<")
	resp := models.Response{}
	var req dto.ReqPortalActivation
	var merchantGroup dbmodels.MerchantGroup
	var emailPortalMerchant string
	var emailPortalGroupMerchant string

	switch reqData.Category {
	case 0:
		merchantGroup, _ = db.InitMerchantGroupRepository().Get(dbmodels.MerchantGroup{ID: reqData.Id})
		kelurahanID, _ := strconv.ParseInt(merchantGroup.Kelurahan, 10, 64)
		kelurahan, _ := db.InitKelurahanRepository().Get(dbmodels.Kelurahan{Id: kelurahanID})
		kecamatanID, _ := strconv.ParseInt(merchantGroup.Kecamatan, 10, 64)
		kecamatan, _ := db.InitKelurahanRepository().Get(dbmodels.Kelurahan{Id: kecamatanID})
		kabID, _ := strconv.ParseInt(merchantGroup.FkLookupKabupatenKota, 10, 64)
		kabupaten, _ := db.InitLookupRepository().Get(dbmodels.Lookup{Id: kabID})
		provID, _ := strconv.ParseInt(merchantGroup.FkLookupProvinsi, 10, 64)
		province, _ := db.InitLookupRepository().Get(dbmodels.Lookup{Id: provID})
		emailPortal, _ := db.InitPortalListActivationDataRepository().GetEmailGroup(reqData.Id)
		emailPortalGroupMerchant = emailPortal
		contactPerson, _ := db.InitMerchantGroupInternalContactPersonRepository().Get(dbmodels.MerchantGroupInternalContactPerson{Id: merchantGroup.InternalContactPersonID})

		req = dto.ReqPortalActivation{
			MerchantName:      merchantGroup.MerchantGroupName,
			OutletName:        "",
			Name:              merchantGroup.NamaPt,
			MerchantGroupName: merchantGroup.MerchantGroupName,
			MerchantGroupId:   strconv.FormatInt(merchantGroup.ID, 10),
			MID:               strconv.FormatInt(merchantGroup.ID, 10),
			MPAN:              "",
			TerminalId:        "",
			Alamat:            merchantGroup.Alamat,
			Kelurahan:         kelurahan.Name,
			Kecamatan:         kecamatan.Name,
			KabupatenKota:     kabupaten.Name,
			Provinsi:          province.Name,
			ProfilePict:       "",
			Email:             reqData.Email,
			OwnerID:           0,
			OwnerName:         contactPerson.BusinessPic,
			MerchantOutletId:  strconv.FormatInt(merchantGroup.ID, 10),
			TipeMerchant:      "",
			StorePhoneNumber:  merchantGroup.NoTelpPic,
			Password:          randomPassword(),
			Action:            reqData.Action,
			Category:          "GROUP",
		}
		
	case 1:
		merchant, _ := db.InitMerchantRepository().Get(dbmodels.Merchant{ID: reqData.Id})
		merchantGroup, _ := db.InitMerchantGroupRepository().Get(dbmodels.MerchantGroup{ID: merchant.MerchantGroupId})
		kelurahanID, _ := strconv.ParseInt(merchant.Kelurahan, 10, 64)
		kelurahan, _ := db.InitKelurahanRepository().Get(dbmodels.Kelurahan{Id: kelurahanID})
		kecamatanID, _ := strconv.ParseInt(merchant.Kecamatan, 10, 64)
		kecamatan, _ := db.InitKelurahanRepository().Get(dbmodels.Kelurahan{Id: kecamatanID})
		kabID, _ := strconv.ParseInt(merchant.KabupatenKota, 10, 64)
		kabupaten, _ := db.InitDati2Repository().Get(dbmodels.Dati2{Id: kabID})
		provID, _ := strconv.ParseInt(merchant.Provinsi, 10, 64)
		province, _ := db.InitProvinsiRepository().Get(dbmodels.Provinsi{Id: provID})
		owner, _ := db.InitOwnerRepository().Get(dbmodels.Owner{Id: merchant.OwnerId})
		emailPortal, _ := db.InitPortalListActivationDataRepository().GetEmailGroup(merchant.OwnerId)
		emailPortalMerchant = emailPortal
		req = dto.ReqPortalActivation{
			MerchantName:      merchant.StoreName,
			OutletName:        "",
			Name:              merchant.StoreName,
			MerchantGroupName: merchantGroup.MerchantGroupName,
			MerchantGroupId:   strconv.FormatInt(merchantGroup.ID, 10),
			MID:               strconv.FormatInt(merchant.ID, 10),
			MPAN:              merchant.MerchantPan,
			TerminalId:        "",
			Alamat:            merchant.Alamat,
			Kelurahan:         kelurahan.Name,
			Kecamatan:         kecamatan.Name,
			KabupatenKota:     kabupaten.Name,
			Provinsi:          province.Name,
			ProfilePict:       merchant.SelfiePath,
			Email:             reqData.Email,
			OwnerID:           owner.Id,
			OwnerName:         owner.OwnerFirstName + owner.OwnerLastName,
			MerchantOutletId:  merchant.MerchantOutletID,
			TipeMerchant:      merchant.MerchantType,
			StorePhoneNumber:  merchant.StorePhoneNumber,
			Password:          "passwordRahasia1!",
			Action:            reqData.Action,
			Category:          "MERCHANT",
		}
		
	}

	portalActivationData, err := json.Marshal(req)
	if err != nil {
		resp.ErrCode = "05"
		resp.ErrDesc = "Error Marshal"
		return resp,""
	}

	keyredis := "ROSE:PORTAL:" + req.Email

	if emailPortalGroupMerchant == req.Email {
		redis_single.DelRedisKey(keyredis)
	}

	if emailPortalMerchant == req.Email {
		redis_single.DelRedisKey(keyredis)
	}

	val, _ := redis_single.GetRedisKey(keyredis)
	var keyURL string
	if val == "" {
		/* key/data tidak ada di redis */

		keyURL = b64.StdEncoding.EncodeToString([]byte(keyredis))

		errEmail := composeEmail(req, keyURL)
		if errEmail != nil {
			resp.ErrCode = "05"
			resp.ErrDesc = errEmail.Error()
			return resp,""
		}

		err := redis_single.SaveRedis(keyredis, portalActivationData)
		if err != nil {
			log.Println(fmt.Sprintf("Error save redis : %v", err))
		}

	} else {
		/* key/data masih ada di redis */
		log.Println("Data already exist in redis")
		resp.ErrCode = "05"
		resp.ErrDesc = "Email already axist : " + req.Email
		return resp,""
	}

	switch reqData.Category {
	case 0:
		merchantGroup.EmailPortal = reqData.Email
		merchantGroup.PortalStatus = 1
		merchantGroup, err = db.InitMerchantGroupRepository().ActivationBp(merchantGroup)
		if err != nil {
			resp.ErrCode = "05"
			resp.ErrDesc = "Error Update Merchant Group"
			return resp,""
		}
		
	case 1:
		er := db.InitPortalListActivationDataRepository().UpdateEmailMerchant(req)
		if er != nil {
			resp.ErrCode = "05"
			resp.ErrDesc = "Error Update email owner merchant"
			return resp,""
		}

		if err := db.InitPortalListActivationDataRepository().UpdatePortalCategory(req); err != nil {
			resp.ErrCode = "05"
			resp.ErrDesc = "Error Update portalCategory merchant"
			return resp,""
		}
	}

	resp.ErrCode = "00"
	resp.ErrDesc = "Send email success"

	return resp, keyURL

}

// GetPortalActivation ...
func GetPortalActivation(req dto.ReqPortalActivation) models.Response {
	resp := models.Response{}

	req.Password = randomPassword()

	switch req.Category {
	case "OUTLET":
		req.Name = req.OutletName
		total := db.ActivationOutlet(req.MPAN, req.TerminalId, req.Email)
		if total < 1 {
			resp.ErrCode = constants.EC_FAIL_DATA_NOTFOUND
			resp.ErrDesc = "Error Update Email Merchant Outlet"
			return resp
		}
	default:
		req.Name = req.MerchantName
		er := db.InitPortalListActivationDataRepository().UpdateEmailMerchant(req)
		if er != nil {
			resp.ErrCode = "05"
			resp.ErrDesc = "Error Update email owner merchant"
			return resp
		}
	}

	if err := db.InitPortalListActivationDataRepository().UpdatePortalCategory(req); err != nil {
		resp.ErrCode = "05"
		resp.ErrDesc = "Error Update portalCategory merchant"
		return resp
	}

	portalActivationData, err := json.Marshal(req)
	if err != nil {
		resp.ErrCode = "05"
		resp.ErrDesc = "Error Marshal"
		return resp
	}

	keyredis := "ROSE:PORTAL:" + req.Email
	val, _ := redis_single.GetRedisKey(keyredis)
	if val == "" {
		/* key/data tidak ada di redis */

		keyURL := b64.StdEncoding.EncodeToString([]byte(keyredis))

		errEmail := composeEmail(req, keyURL)
		if errEmail != nil {
			resp.ErrCode = "05"
			resp.ErrDesc = errEmail.Error()
			return resp
		}

		go redis_single.SaveRedis(keyredis, portalActivationData)

	} else {
		/* key/data masih ada di redis */
		log.Println("Data already exist in redis")
		resp.ErrCode = "05"
		resp.ErrDesc = "Email already axist : " + req.Email
		return resp
	}

	resp.ErrCode = "00"
	resp.ErrDesc = "Send email success"

	return resp
}

// GetPortalCallback ...
func GetPortalCallback(req dto.ReqPortalCallback) models.Response {
	resp := models.Response{}
	var err error

	if req.Type != "group" {
		err = db.InitPortalListActivationDataRepository().UpdateDbPortalStatus(req)
	} else {
		err = db.InitPortalListActivationDataRepository().UpdateGroupPortalStatus(req)
	}

	if err != nil {
		resp.ErrCode = "05"
		resp.ErrDesc = "Error update db portal status:" + err.Error()
		return resp
	}

	resp.ErrCode = "00"
	resp.ErrDesc = "Callback status aktivasi success"

	return resp
}

// GetPortalReset ...
func GetPortalReset(req dto.ReqPortalActivation) models.Response {
	resp := models.Response{}

	req.Password = randomPassword()

	errEmail := composeEmailReset(req)
	if errEmail != nil {
		resp.ErrCode = "05"
		resp.ErrDesc = errEmail.Error()
		return resp
	}

	jsonData := map[string]string{"username": req.Email, "new_password": req.Password}
	jsonValue, _ := json.Marshal(jsonData)
	url := utils.GetEnv("PORTAL_BE_HOST_URL", "http://13.228.25.85:8000/") + "forgot-password"
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		resp.ErrCode = "05"
		resp.ErrDesc = "Reset password failed -> " + err.Error()

		return resp
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response forgot password bussiness portal :", string(data))

	resp.ErrCode = "00"
	resp.ErrDesc = "Reset password success"

	return resp
}

// ResetPassword ...
func ResetPassword(req dto.ReqPortalActivation) models.Response {
	resp := models.Response{}

	req.Action = "reset-password"

	req.Password = randomPassword()

	errEmail := composeEmailReset(req)
	if errEmail != nil {
		resp.ErrCode = "05"
		resp.ErrDesc = errEmail.Error()
		return resp
	}

	jsonData := map[string]string{"username": req.Email, "new_password": req.Password}
	jsonValue, _ := json.Marshal(jsonData)
	url := utils.GetEnv("PORTAL_BE_HOST_URL", "http://13.228.25.85:8000/") + "forgot-password"
	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		resp.ErrCode = "05"
		resp.ErrDesc = "Reset password failed -> " + err.Error()

		return resp
	}

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response forgot password bussiness portal :", string(data))

	resp.ErrCode = "00"
	resp.ErrDesc = "Reset password success"

	return resp
}

// composeEmailReset ..
func composeEmailReset(req dto.ReqPortalActivation) error {
	urlLogin := utils.GetEnv("PORTAL_FE_URL", "http://18.139.224.183:8080/login")

	to := req.Email
	subject := "Reset Password"
	templateEmail := `
		<p>Hi, Kami telah menerima permintaanmu untuk melakukan reset password akun Business Portal.</p>
		<p>Silahkan login ulang menggunakan username dan password dibawah: </p>
		<p>Username: ` + req.Email + `</p>
		<p>Password: ` + req.Password + `</p>
		<p>Klik link ` + urlLogin + ` untuk menuju halaman login.</p>
		<p>Abaikan email ini jika kamu tidak pernah meminta untuk melakukan reset password.</p>`

	err := sendMail(to, subject, templateEmail, req)
	if err != nil {
		// log.Fatal(err.Error())
		return err
	}

	return nil
}

// composeEmail ..
func composeEmail(req dto.ReqPortalActivation, keyURL string) error {
	url := utils.GetEnv("EMAIL_URL", "http://18.139.224.183:8080/activate/")
	urlkey := url + keyURL

	to := req.Email
	subject := "Merchant Activation"
	templateEmail := `
	<h3>Dear Customer,</h3>
	<p>Selamat! Akun Anda berhasil di daftarkan.</p>
	<p>Silahkan Aktifkan Akun Merchant Anda menggunakan username dan password berikut,</p>
	<h4>Username: ` + req.Email + ` </h4>
	<h4>Password: ` + req.Password + ` </h4>
	<p>di link ` + urlkey + ` </p>
	<p>Terima Kasih,<br>Ottopay</p>
	`

	err := sendMail(to, subject, templateEmail, req)
	if err != nil {
		// log.Fatal(err.Error())
		return err
	}

	return nil
}

// sendMail ..
//func sendMail(to string, subject, message string, req dto.ReqPortalActivation) error {
//	configSMTP := utils.GetEnv("EMAIL_SMTP_ADDRESS", "smtp.office365.com")
//	configSMTPPort, _ := strconv.Atoi(utils.GetEnv("EMAIL_SMTP_PORT", "587"))
//	configEmailSender := utils.GetEnv("EMAIL_SENDER", "ottopay@ottopay.id")
//	configPassSender := utils.GetEnv("EMAIL_PASSWORD", "Mutiara2019")
//
//	mailer := gomail.NewMessage()
//	mailer.SetHeader("From", configEmailSender)
//	mailer.SetHeader("To", to)
//	mailer.SetHeader("Subject", subject)
//	mailer.SetBody("text/html", message)
//
//	dialer := gomail.NewDialer(configSMTP, configSMTPPort, configEmailSender, configPassSender)
//	fmt.Println("Dialer >>> ", dialer)
//	fmt.Println("Mailer >>> ", mailer)
//
//	err := dialer.DialAndSend(mailer)
//	if err != nil {
//		errs := saveLogPortal(req, configEmailSender, to, "Failed to send email")
//		if errs != nil {
//			fmt.Println("error save log portal ==> ", errs.Error())
//			return errs
//		}
//		fmt.Println("error while sending mail ==> ", err.Error())
//		return err
//	}
//
//	log.Println("Send email to:", to, "|| Error:", err)
//
//	error := saveLogPortal(req, configEmailSender, to, "Send email success")
//	if error != nil {
//		fmt.Println("error save log portal ==> ", error.Error())
//		return error
//	}
//
//	return nil
//}
func sendMail(to string, subject, message string, req dto.ReqPortalActivation) error {
	configEmailSender := utils.GetEnv("EMAIL_SENDER", "ottopay@ottopay.id")

	recipient := to
	subjectMail := subject
	body := message
	mail := ottomail.NewMail(configEmailSender)
	err := mail.Send(recipient, subjectMail, body)
	if err != nil {
		fmt.Println("There's error")
		return err
	}

	log.Println("Send email to:", to, "|| Error:", err)

	error := saveLogPortal(req, configEmailSender, to, "Send email success")
	if error != nil {
		fmt.Println("error save log portal ==> ", error.Error())
		return error
	}

	return nil
}

func saveLogPortal(req dto.ReqPortalActivation, from string, to string, message string) error {
	userName := auth.UserLogin.Name

	log := models.LogPortal{
		User:    userName,
		Mid:     req.MID,
		From:    from,
		To:      to,
		Action:  req.Action,
		Message: message,
	}

	errs := db.InitPortalListActivationDataRepository().SaveDbLogPortal(log)
	if errs != nil {
		fmt.Println("Save log portal ==> ", errs.Error())
		return errs
	}

	return nil
}

func randomPassword() string {

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")
	length := 10
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	generatePassword := b.String()

	return generatePassword
}
