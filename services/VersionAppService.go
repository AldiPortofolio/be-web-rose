package services

import (
	"errors"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"log"
	ottoutils "ottodigital.id/library/utils"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/constants/app_name"
	"rose-be-go/db"
	"rose-be-go/models"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"rose-be-go/redis/redis_cluster"
	"time"
)

type VersionAppService struct {
	General models.GeneralModel
	GetRedisKey  func(string) (string, error)
	SaveRedisKey  func(string, interface{}) (error)
	HistoryAppVersionRepository *db.HistoryAppVersionRepository

}

func InitVersionAppService(gen models.GeneralModel) *VersionAppService  {
	return &VersionAppService{
		General:gen,
		GetRedisKey: redis_cluster.GetRedisKey,
		SaveRedisKey: redis_cluster.SaveRedis,
		HistoryAppVersionRepository: db.InitHistoryAppVersionRepository(),
	}
}

var (
	ottomartKey string
	nfcKey string
	indomarcoKey string
	sfaKey string
)

func init()  {
	ottomartKey = ottoutils.GetEnv("OTTOMART_KEY_VERSION", "OTTOMART:ANDROID-VERSION")
	nfcKey = ottoutils.GetEnv("NFC_KEY_VERSION", "NFC:ANDROID-VERSION")
	indomarcoKey = ottoutils.GetEnv("INDOMARCO_KEY_VERSION", "INDOMARCO:ANDROID-VERSION")
	sfaKey = ottoutils.GetEnv("SFA_KEY_VERSION", "SFA:ANDROID-VERSION")


}

func (service *VersionAppService) GetVersion() models.Response {
	fmt.Println(">>> VersionAppService - GetVersion <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("VersionAppService: GetVersion",
		zap.Any("req", ""))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "MerchantQrisStatusService: Filter")
	defer span.Finish()

	var res models.Response

	versionOttomartChan := make(chan string)
	versionNfcChan := make(chan string)
	versionIndomarcoChan := make(chan string)
	versionSfaChan := make(chan string)

	go service.GetVersionOttomart(versionOttomartChan)
	go service.GetVersionNfc(versionNfcChan)
	go service.GetVersionIndomarco(versionIndomarcoChan)
	go service.GetVersionSfa(versionSfaChan)


	versionOttomart := <-versionOttomartChan
	versionNfc := <- versionNfcChan
	versionIndomarco := <-versionIndomarcoChan
	versionSfa := <-versionSfaChan

	data := dto.ResVersionAppDto{
		Indomarco: versionIndomarco,
		Nfc: versionNfc,
		Ottomart: versionOttomart,
		Sfa: versionSfa,
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG
	res.Contents = data


	return res
}


func (service *VersionAppService) Update(req dto.ReqUpdateVersionAppDto) models.Response {

	fmt.Println(">>> VersionAppService - Update <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("VersionAppService: Update",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "VersionAppService: Update")
	defer span.Finish()

	var res models.Response

	err := service.UpdateVersion(req.AppName, req.Version)


	if err != nil {
		res.ErrCode = constants.EC_FAIL_SAVE
		res.ErrDesc = constants.EC_FAIL_SAVE_DESC
		return res
	}

	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}
func (service *VersionAppService) UpdateVersion(appName string, version string) error{
	var key string
	switch appName {
	case app_name.SFA:
		key = sfaKey
		
	case app_name.NFC:
		key = nfcKey
		
	case app_name.INDOMARCO:
		key = indomarcoKey
		
	case app_name.OTTOMART:
		key = ottomartKey
		
	default:
		key = ""
		return errors.New("key null")
		
	}

	err := service.SaveRedisKey(key, version);
	if err!=nil {
		log.Println("err -> ",key, err)
		return err
	}
	service.InsertLog(appName, key, version)

	return err
}

func (service *VersionAppService) InsertLog(appName string, key string, version string) {
	fmt.Println("<< VersionAppService - Insert log to DB >> ")
	req:=dbmodels.HistoryAppVersion{
		AppName:appName,
		Key: key,
		Version:version,
		CreatedBy: auth.UserLogin.Name,
		CreatedAt:time.Now(),
	}
	log.Println(req)

	service.HistoryAppVersionRepository.Save(req)

}


func (service *VersionAppService) GetVersionOttomart(version chan string) {
	fmt.Println("<< GetVersionOttomart >>")
	fmt.Println("key : ", ottomartKey)

	
	res,_:=service.GetRedisKey(ottomartKey)

	version <- res
	close(version)
}

func (service *VersionAppService) GetVersionNfc(version chan string) {
	fmt.Println("<< GetVersionNfc >>")
	fmt.Println("key : ", nfcKey)

	res, _:=service.GetRedisKey(nfcKey)

	version <- res
	close(version)
}

func (service *VersionAppService) GetVersionIndomarco(version chan string) {
	fmt.Println("<< GetVersionIndomarco >>")
	fmt.Println("key : ", indomarcoKey)

	res,_:=service.GetRedisKey(indomarcoKey)

	version <- res
	close(version)
}

func (service *VersionAppService) GetVersionSfa(version chan string) {
	fmt.Println("<< GetVersionSfa >>")
	fmt.Println("key : ", sfaKey)

	res,_:=service.GetRedisKey(sfaKey)

	version <- res
	close(version)
}

