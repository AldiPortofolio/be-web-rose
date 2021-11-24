package routers

import (
	"fmt"
	"io"
	"os"
	"rose-be-go/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/opentracing/opentracing-go"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"

	"rose-be-go/auth"
	"rose-be-go/controllers"

	"ottodigital.id/library/httpserver/ginserver"
	ottologer "ottodigital.id/library/logger"
	"ottodigital.id/library/ottotracing"
	"ottodigital.id/library/utils"
)

// RoseBeGoEnv struct
type RoseBeGoEnv struct {
	ReadTo  int `envconfig:"READ_TIMEOUT" default:"120"`
	WriteTo int `envconfig:"WRITE_TIMEOUT" default:"120"`
}

var (
	accessPointHealthCheck string

	accessPointest           string
	accessPointGroupMerchant string

	accessPointUpload             string
	accessPointGetDataUpload      string
	accessPointDownloadFile       string
	accessPointDownloadResultFile string

	redisPrefixToken string

	header    string
	version   string
	prefixUrl string

	nameService    string
	openTracingSvr string

	debugMode string
	//readTo    int
	//writeTo   int
	roseBeGoEnv RoseBeGoEnv
)

func init() {

	header = "/rosego"
	version = "/v.0.1"
	prefixUrl = header + version

	accessPointHealthCheck = "/health-check"

	accessPointest = "/test"
	accessPointGroupMerchant = "/merchant"

	//accessPointGroupMerchant+acessPointNmid+accessPointUpload
	accessPointUpload = "/merchant/nmid/upload"
	accessPointGetDataUpload = "/merchant/nmid/all"
	accessPointDownloadFile = "/merchant/nmid/downloadFile"
	accessPointDownloadResultFile = "/merchant/nmid/download/resultFile"

	debugMode = utils.GetEnv("APPS_DEBUG", "debug")

	err := envconfig.Process("ROSE_BE_GO", &roseBeGoEnv)
	if err != nil {
		fmt.Println("Failed to get ROSE_BE_GO env:", err)
	}

	nameService = utils.GetEnv("ROSE_BE_GO", "ROSE-BE-GO")
	openTracingSvr = utils.GetEnv("JAEGER_HOSTURL", "13.250.21.165:5775")
	redisPrefixToken = utils.GetEnv("ROSE_REDIS_PREFIX_TOKEN", "ROSE:TOKEN:")
}

// Server ...
func Server(listenAddress string) error {
	sugarLogger := ottologer.GetLogger()

	ottoRouter := OttoRouter{}
	ottoRouter.InitTracing()
	ottoRouter.Routers()
	defer ottoRouter.Close()

	err := ginserver.GinServerUp(listenAddress, ottoRouter.Router)

	if err != nil {
		fmt.Println("Error:", err)
		sugarLogger.Error("Error ", zap.Error(err))
		return err
	}

	fmt.Println("Server UP")
	sugarLogger.Info("Server UP ", zap.String("listenAddress", listenAddress))

	return err
}

// OttoRouter ..
type OttoRouter struct {
	Tracer   opentracing.Tracer
	Reporter jaeger.Reporter
	Closer   io.Closer
	Err      error
	GinFunc  gin.HandlerFunc
	Router   *gin.Engine
}

//Routers
func (ottoRouter *OttoRouter) Routers() {

	gin.SetMode(debugMode)

	router := gin.New()
	router.Use(ottoRouter.GinFunc)
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "authorization", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		//AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 86400,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.Title = "ROSE BE GO"
	docs.SwaggerInfo.Description = "Be for web rose"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	switch utils.GetEnv("APPS_ENV", "local") {
	case "local":
		docs.SwaggerInfo.Host = "localhost:8098"
	case "dev":
		docs.SwaggerInfo.Host = "13.228.25.85:8098"
	}

	router.GET(prefixUrl+accessPointHealthCheck, controllers.HealthCheck)
	router.GET("/rose-be-go/version", controllers.Version)
	router.POST(prefixUrl+accessPointest, auth.CekToken, controllers.TestController)

	var api *gin.RouterGroup

	UploadNmidController := new(controllers.UploadNmidController)
	router.POST(prefixUrl+accessPointUpload, auth.CekToken, UploadNmidController.Upload)
	router.POST(prefixUrl+accessPointGetDataUpload, auth.CekToken, UploadNmidController.GetFilterPaging)
	router.POST(prefixUrl+accessPointDownloadFile, auth.CekToken, UploadNmidController.HandleDownload)
	router.POST(prefixUrl+accessPointDownloadResultFile, auth.CekToken, UploadNmidController.HandleResultDownload)

	ReportFinishController := new(controllers.ReportFinishController)
	api = router.Group(prefixUrl + "/report-finish")
	api.POST("/send", auth.CekToken, ReportFinishController.Send)
	api.POST("/all", auth.CekToken, ReportFinishController.GetReportFinished)
	api.POST("/download", auth.CekToken, ReportFinishController.DownloadPath)

	MerchantController := new(controllers.MerchantController)
	api = router.Group(prefixUrl + "/merchant")
	api.POST("/upgrade", auth.CekToken, MerchantController.Upgrade)
	api.POST("/filter", auth.CekToken, MerchantController.GetFilterPaging)
	api.POST("/by-group", auth.CekToken, MerchantController.GetListMerchantByGroup)
	api.POST("/find-tgl-akuisisi", auth.CekToken, MerchantController.FindTanggalAkuisisi)
	api.POST("/dashboard", auth.CekToken, MerchantController.GetFilterDashboardPaging)
	// api.POST("/merchant-detail", auth.CekToken, MerchantController.FindMerchantDetail)
	//api = &nil

	MerchantGroupController := new(controllers.MerchantGroupController)
	api = router.Group(prefixUrl + "/merchantGroup")
	api.GET("/id/:id", auth.CekToken, MerchantGroupController.GetById)
	api.POST("/merchant-group-activation", auth.CekToken, MerchantGroupController.MerchantGroupActivationController)

	ReportQrController := new(controllers.ReportQrController)
	api = router.Group(prefixUrl + "/report-qr")
	api.POST("/all", auth.CekToken, ReportQrController.GetReportQr)
	api.POST("/download", auth.CekToken, ReportQrController.DownloadPath)
	api.POST("/download/result", auth.CekToken, ReportQrController.DownloadResultPath)

	PortalActivationController := new(controllers.PortalActivationController)
	api = router.Group(prefixUrl + "/portal")
	api.POST("/activation", auth.CekToken, PortalActivationController.Activation)
	api.POST("/bpactivation", auth.CekToken, PortalActivationController.BpActivation)
	api.POST("/list-account/filter", auth.CekToken, PortalActivationController.Filter)
	api.POST("/callback", PortalActivationController.CallBack)
	api.POST("/reset-password", auth.CekToken, PortalActivationController.Reset)
	api.POST("/reset-password-bp", auth.CekToken, PortalActivationController.ResetPassword)
	api.POST("/list-outlet/filter", auth.CekToken, PortalActivationController.FilterOutlet)

	MerchantWipController := new(controllers.MerchantWipController)
	api = router.Group(prefixUrl + "/merchantwip")
	api.POST("action/:action", auth.CekToken, MerchantWipController.Approve)

	MerchantAggregatorDetailController := new(controllers.MerchantAggregatorDetailController)
	api = router.Group(prefixUrl + "/merchant-aggregator-detail")
	api.POST("", auth.CekToken, MerchantAggregatorDetailController.Save)
	api.POST("/list-data-approval", auth.CekToken, MerchantAggregatorDetailController.ListDataApproval)
	api.POST("/filter-data-approval", auth.CekToken, MerchantAggregatorDetailController.FilterDetailDataApproval)
	api.POST("/filter", auth.CekToken, MerchantAggregatorDetailController.GetFilterPaging)
	api.POST("/filter-temp", auth.CekToken, MerchantAggregatorDetailController.GetFilterTempPaging)
	api.POST("/approve", auth.CekToken, MerchantAggregatorDetailController.Approve)
	api.POST("/list-aggregator", auth.CekToken, MerchantAggregatorDetailController.GetListMerchantAggregator)

	MerchantAggUploadController := new(controllers.MerchantAggUploadController)
	api = router.Group(prefixUrl + "/merchant/agg")
	api.POST("/upload", auth.CekToken, MerchantAggUploadController.UploadFile)
	api.POST("/upload-data", auth.CekToken, MerchantAggUploadController.GetDataMerchantAggUpload)
	api.POST("/download", auth.CekToken, MerchantAggUploadController.Download)
	api.POST("/approve", auth.CekToken, MerchantAggUploadController.ApproveMerchantAggregatorUpload)
	api.POST("/download-template", auth.CekToken, MerchantAggUploadController.HandleTemplateDownload)

	ReportRejectedController := new(controllers.ReportRejectedController)
	api = router.Group(prefixUrl + "/report-rejected")
	api.POST("/all", auth.CekToken, ReportRejectedController.GetReportRejected)
	api.POST("/send", auth.CekToken, ReportRejectedController.Send)
	api.POST("/download", auth.CekToken, ReportRejectedController.DownloadPath)
	api.GET("/reason/:wip_id", auth.CekToken, ReportRejectedController.GetReason)

	MasterLimitationTempController := new(controllers.MasterLimitationTempController)
	api = router.Group(prefixUrl + "/masterlimitationtemp")
	api.POST("/filter", auth.CekToken, MasterLimitationTempController.GetFilterPaging)
	api.GET("/id/:id", auth.CekToken, MasterLimitationTempController.GetById)
	api.POST("/approve", auth.CekToken, MasterLimitationTempController.Approve)
	api.POST("/reject", auth.CekToken, MasterLimitationTempController.Reject)

	MasterLimitationController := new(controllers.MasterLimitationController)
	api = router.Group(prefixUrl + "/masterlimitation")
	api.POST("/filter", auth.CekToken, MasterLimitationController.GetFilterPaging)
	api.POST("", auth.CekToken, MasterLimitationController.Save)
	api.GET("/id/:id", auth.CekToken, MasterLimitationController.GetById)

	MdrAggregatorController := new(controllers.MdrAggregatorController)
	api = router.Group(prefixUrl + "/mdr-aggregator")
	api.POST("", auth.CekToken, MdrAggregatorController.Save)
	api.POST("/filter", auth.CekToken, MdrAggregatorController.GetFilterPaging)
	api.GET("/id/:id", auth.CekToken, MdrAggregatorController.GetById)

	MdrAggregatorTempController := new(controllers.MdrAggregatorTempController)
	api = router.Group(prefixUrl + "/mdr-aggregator-temp")
	api.POST("/filter", auth.CekToken, MdrAggregatorTempController.GetFilterPaging)
	api.GET("/id/:id", auth.CekToken, MdrAggregatorTempController.GetById)
	api.POST("/approve", auth.CekToken, MdrAggregatorTempController.Approve)
	api.POST("/reject", auth.CekToken, MdrAggregatorTempController.Reject)

	DepositBankController := new(controllers.DepositBankController)
	api = router.Group(prefixUrl + "/deposit-bank")
	api.POST("/all", auth.CekToken, DepositBankController.GetAllInfo)

	MerchantQrisStatusController := new(controllers.MerchantQrisStatusController)
	api = router.Group(prefixUrl + "/qris-status")
	api.POST("/filter", auth.CekToken, MerchantQrisStatusController.Filter)

	VersionAppController := new(controllers.VersionAppController)
	api = router.Group(prefixUrl + "/version-app")
	api.GET("", auth.CekToken, VersionAppController.GetVersion)
	api.POST("/update", auth.CekToken, VersionAppController.Update)

	ClearSessionController := new(controllers.ClearSessionController)
	api = router.Group(prefixUrl + "/clear-session")
	api.GET("", auth.CekToken, ClearSessionController.Clear)
	api.GET("/last-updated", auth.CekToken, ClearSessionController.GetLastUpdated)

	ReportMagTransactionController := new(controllers.ReportMagTransactionsController)
	api = router.Group(prefixUrl + "/mag-transactions")
	api.POST("/all", auth.CekToken, ReportMagTransactionController.GetAll)
	api.POST("/export", auth.CekToken, ReportMagTransactionController.Export)

	QrisConfigController := new(controllers.QrisConfigController)
	api = router.Group(prefixUrl + "/qris-config")
	api.POST("", auth.CekToken, QrisConfigController.Save)
	api.POST("/filter", auth.CekToken, QrisConfigController.GetFilterPaging)

	UploadMerchantController := new(controllers.UploadMerchantController)
	api = router.Group(prefixUrl + "/upload-merchant")
	api.POST("/upload", auth.CekToken, UploadMerchantController.Upload)
	api.POST("/all", auth.CekToken, UploadMerchantController.GetFilterPaging)
	api.POST("/download", auth.CekToken, UploadMerchantController.HandleDownload)
	api.POST("/result-download", auth.CekToken, UploadMerchantController.HandleResultDownload)

	UploadMerchantWipController := new(controllers.UploadMerchantWipController)
	api = router.Group(prefixUrl + "/upload-merchant-wip")
	api.POST("/upload", auth.CekToken, UploadMerchantWipController.Upload)
	api.POST("/all", auth.CekToken, UploadMerchantWipController.GetFilterPaging)
	api.POST("/download", auth.CekToken, UploadMerchantWipController.HandleDownload)
	api.POST("/result-download", auth.CekToken, UploadMerchantWipController.HandleResultDownload)

	UploadMerchantNonWipController := new(controllers.UploadMerchantNonWipController)
	api = router.Group(prefixUrl + "/upload-merchant-non-wip")
	api.POST("/upload", auth.CekToken, UploadMerchantNonWipController.Upload)
	api.POST("/all", auth.CekToken, UploadMerchantNonWipController.GetFilterPaging)
	api.POST("/download", auth.CekToken, UploadMerchantNonWipController.HandleDownload)
	api.POST("/download-example", auth.CekToken, UploadMerchantNonWipController.HandleDownloadExample)
	api.POST("/result-download", auth.CekToken, UploadMerchantNonWipController.HandleResultDownload)

	FeatureProductController := new(controllers.FeatureProductController)
	api = router.Group(prefixUrl + "/feature-product")
	api.POST("", auth.CekToken, FeatureProductController.Save)
	api.POST("/filter", auth.CekToken, FeatureProductController.GetFilterPaging)

	LimitTransactionController := new(controllers.LimitTransactionController)
	api = router.Group(prefixUrl + "/limit-transaction")
	api.POST("", auth.CekToken, LimitTransactionController.Save)
	api.POST("/filter", auth.CekToken, LimitTransactionController.GetFilterPaging)

	ProductController := new(controllers.ProductController)
	api = router.Group(prefixUrl + "/product")
	api.POST("", auth.CekToken, ProductController.Save)
	api.POST("/filter", auth.CekToken, ProductController.GetFilterPaging)

	UserCategoryController := new(controllers.UserCategoryController)
	api = router.Group(prefixUrl + "/user-category")
	api.POST("", auth.CekToken, UserCategoryController.Save)
	api.POST("/filter", auth.CekToken, UserCategoryController.GetFilterPaging)
	api.GET("/dropdown", auth.CekToken, UserCategoryController.DropdownList)

	CategoryLevelFeatureController := new(controllers.CategoryLevelFeatureController)
	api = router.Group(prefixUrl + "/category-level-feature")
	api.POST("", auth.CekToken, CategoryLevelFeatureController.Save)
	api.POST("/filter", auth.CekToken, CategoryLevelFeatureController.GetFilterPaging)

	BannerController := new(controllers.BannerController)
	api = router.Group(prefixUrl + "/banner")
	api.POST("", auth.CekToken, BannerController.Save)
	api.POST("/filter", auth.CekToken, BannerController.GetFilterPaging)
	api.DELETE("/delete/:id", auth.CekToken, BannerController.Delete)

	ImageManagementController := new(controllers.ImageManagementController)
	api = router.Group(prefixUrl + "/image-management")
	api.POST("", auth.CekToken, ImageManagementController.Save)
	api.POST("/filter", auth.CekToken, ImageManagementController.GetFilterPaging)

	ProfileThemeController := new(controllers.ProfileThemeController)
	api = router.Group(prefixUrl + "/profile-theme")
	api.POST("", auth.CekToken, ProfileThemeController.Save)
	api.POST("/filter", auth.CekToken, ProfileThemeController.GetFilterPaging)

	LevelMerchantController := new(controllers.LevelMerchantController)
	api = router.Group(prefixUrl + "/level-merchant")
	api.GET("/:limit", auth.CekToken, LevelMerchantController.GetAll)

	MdrBankController := new(controllers.MdrBankController)
	api = router.Group(prefixUrl + "/mdr-bank")
	api.POST("", auth.CekToken, MdrBankController.Save)
	api.POST("/filter", auth.CekToken, MdrBankController.GetFilterPaging)
	api.GET("/all", auth.CekToken, MdrBankController.FindAll)

	LimitTransactionDepositController := new(controllers.LimitTransactionDepositController)
	api = router.Group(prefixUrl + "/limit-transaction-deposit")
	api.POST("", auth.CekToken, LimitTransactionDepositController.Save)
	api.POST("/filter", auth.CekToken, LimitTransactionDepositController.GetFilter)

	FeeMdrSettingController := new(controllers.FeeMdrSettingController)
	api = router.Group(prefixUrl + "/fee-mdr-setting")
	api.POST("", auth.CekToken, FeeMdrSettingController.Save)
	api.POST("/filter", auth.CekToken, FeeMdrSettingController.GetFilterPaging)
	api.GET("/all", auth.CekToken, FeeMdrSettingController.FindAll)

	FeeMdrSettingMerchantGroupController := new(controllers.FeeMdrSettingMerchantGroupController)
	api = router.Group(prefixUrl + "/fee-mdr-setting-merchant-group")
	api.POST("", auth.CekToken, FeeMdrSettingMerchantGroupController.Save)
	api.POST("/filter", auth.CekToken, FeeMdrSettingMerchantGroupController.GetFilterPaging)
	api.GET("/all", auth.CekToken, FeeMdrSettingMerchantGroupController.FindAll)

	MdrTenorController := new(controllers.MdrTenorController)
	api = router.Group(prefixUrl + "/mdr-tenor")
	api.POST("", auth.CekToken, MdrTenorController.Save)
	api.POST("/filter", auth.CekToken, MdrTenorController.GetFilterPaging)
	api.GET("/all", auth.CekToken, MdrTenorController.FindAll)

	MasterServiceController := new(controllers.MasterServiceController)
	api = router.Group(prefixUrl + "/master-service")
	api.POST("", auth.CekToken, MasterServiceController.Save)
	api.GET("/all", auth.CekToken, MasterServiceController.FindAll)

	MasterTypeController := new(controllers.MasterTypeController)
	api = router.Group(prefixUrl + "/master-type")
	api.POST("", auth.CekToken, MasterTypeController.Save)
	api.GET("/all", auth.CekToken, MasterTypeController.FindAll)

	MerchantCustomerController := new(controllers.MerchantCustomerController)
	api = router.Group(prefixUrl + "/merchant-customer")
	api.POST("/filter", auth.CekToken, MerchantCustomerController.GetFilterPaging)
	api.GET("/all", auth.CekToken, MerchantCustomerController.FindAll)
	api.POST("/export", auth.CekToken, MerchantCustomerController.Export)

	QrPrePrintedController := new(controllers.QrPrePrintedController)
	api = router.Group(prefixUrl + "/qr-preprinted")
	api.POST("/send", auth.CekToken, QrPrePrintedController.Send)
	api.POST("/all", auth.CekToken, QrPrePrintedController.GetFilterPaging)
	api.POST("/download", auth.CekToken, QrPrePrintedController.HandleDownload)

	FeeCicilanController := new(controllers.FeeCicilanController)
	api = router.Group(prefixUrl + "/fee-cicilan")
	api.POST("", auth.CekToken, FeeCicilanController.Save)
	api.GET("/find", auth.CekToken, FeeCicilanController.Find)

	UploadMerchantOkController := new(controllers.UploadMerchantOkController)
	api = router.Group(prefixUrl + "/upload-merchant-ok")
	api.POST("/upload", auth.CekToken, UploadMerchantOkController.Upload)
	api.POST("/all", auth.CekToken, UploadMerchantOkController.GetFilterPaging)
	api.POST("/download", auth.CekToken, UploadMerchantOkController.HandleDownload)
	api.POST("/result-download", auth.CekToken, UploadMerchantOkController.HandleResultDownload)
	api.POST("/template-download", auth.CekToken, UploadMerchantOkController.HandleTemplateDownload)

	BankListController := new(controllers.BankListController)
	api = router.Group(prefixUrl + "/bank-list")
	api.POST("", auth.CekToken, BankListController.Save)
	api.POST("/filter", auth.CekToken, BankListController.Filter)

	MerchantBankAccountController := new(controllers.MerchantBankAccountController)
	api = router.Group(prefixUrl + "/merchant-bank-account")
	api.POST("", auth.CekToken, MerchantBankAccountController.Save)
	api.GET("/mid/:mid", auth.CekToken, MerchantBankAccountController.FindAllByMid)
	api.POST("/approval", auth.CekToken, MerchantBankAccountController.FilterApproval)
	api.POST("/approval/approve", auth.CekToken, MerchantBankAccountController.Approve)
	api.POST("/approval/reject", auth.CekToken, MerchantBankAccountController.Reject)
	api.POST("/approval/resendpushnotif", auth.CekToken, MerchantBankAccountController.ResendPushNotif)
	api.POST("/validation-bank-account", auth.CekToken, MerchantBankAccountController.ValidationBankAccount)

	AkuisisiSfaController := new(controllers.AkuisisiSfaController)
	api = router.Group(prefixUrl + "/akuisisi-sfa")
	api.POST("/filter", auth.CekToken, AkuisisiSfaController.GetFilterPaging)

	UploadFeeMdrSettingController := new(controllers.UploadFeeMdrSettingController)
	api = router.Group(prefixUrl + "/upload-fee-mdr-setting")
	api.POST("/upload", auth.CekToken, UploadFeeMdrSettingController.Upload)
	api.POST("/all", auth.CekToken, UploadFeeMdrSettingController.GetFilterPaging)
	api.POST("/download", auth.CekToken, UploadFeeMdrSettingController.HandleDownload)
	api.POST("/result-download", auth.CekToken, UploadFeeMdrSettingController.HandleResultDownload)
	api.POST("/template-download", auth.CekToken, UploadFeeMdrSettingController.HandleTemplateDownload)

	ReportQrPreprintedController := new(controllers.ReportQrPreprintedController)
	api = router.Group(prefixUrl + "/report-qr-preprinted")
	api.POST("/send", auth.CekToken, ReportQrPreprintedController.Send)
	api.POST("/all", auth.CekToken, ReportQrPreprintedController.GetReport)
	api.POST("/download", auth.CekToken, ReportQrPreprintedController.DownloadPath)

	ReportCrmPenugasanController := new(controllers.ReportCrmPenugasanController)
	api = router.Group(prefixUrl + "/report-crm-penugasan")
	api.POST("/send", auth.CekToken, ReportCrmPenugasanController.Send)
	api.POST("/all", auth.CekToken, ReportCrmPenugasanController.GetReport)
	api.POST("/download", auth.CekToken, ReportCrmPenugasanController.DownloadPath)

	ValidationCodeController := new(controllers.ValidationCodeController)
	api = router.Group(prefixUrl + "/validation-code")
	api.POST("", auth.CekToken, ValidationCodeController.Save)
	api.POST("/filter", auth.CekToken, ValidationCodeController.Filter)

	LogUpgradeFdsController := new(controllers.LogUpgradeFdsController)
	api = router.Group(prefixUrl + "/upgrade-fds")
	api.POST("", auth.CekToken, LogUpgradeFdsController.Save)
	api.POST("/filter", auth.CekToken, LogUpgradeFdsController.Filter)

	LookupGroupController := new(controllers.LookupGroupController)
	api = router.Group(prefixUrl + "/lookup-group")
	api.POST("", auth.CekToken, LookupGroupController.Save)
	api.POST("/filter", auth.CekToken, LookupGroupController.Filter)
	api.GET("/all", auth.CekToken, LookupGroupController.All)

	LookupController := new(controllers.LookupController)
	api = router.Group(prefixUrl + "/lookup")
	api.POST("", auth.CekToken, LookupController.Save)
	api.POST("/filter", auth.CekToken, LookupController.Filter)

	MasterTagController := new(controllers.MasterTagController)
	api = router.Group(prefixUrl + "/master-tag")
	api.POST("", auth.CekToken, MasterTagController.Save)
	api.POST("/filter", auth.CekToken, MasterTagController.Filter)
	api.GET("/all", auth.CekToken, MasterTagController.All)

	UploadMerchantMasterTagController := new(controllers.UploadMerchantMasterTagController)
	api = router.Group(prefixUrl + "/upload-merchant-master-tag")
	api.POST("/upload", auth.CekToken, UploadMerchantMasterTagController.Upload)
	api.POST("/all", auth.CekToken, UploadMerchantMasterTagController.GetFilterPaging)
	api.POST("/download", auth.CekToken, UploadMerchantMasterTagController.HandleDownload)
	api.POST("/result-download", auth.CekToken, UploadMerchantMasterTagController.HandleResultDownload)
	api.GET("/download-example", auth.CekToken, UploadMerchantMasterTagController.HandleDownloadExample)

	MerchantMasterTagController := new(controllers.MerchantMasterTagController)
	api = router.Group(prefixUrl + "/merchant-master-tag")
	//api.POST("", auth.CekToken, MasterTagController.Save)
	api.POST("/filter", auth.CekToken, MerchantMasterTagController.Filter)
	api.GET("/download-all", auth.CekToken, MerchantMasterTagController.DownloadAll)
	api.POST("/find-by-mid", auth.CekToken, MerchantMasterTagController.FindByMid)
	//api.GET("/all", auth.CekToken, MasterTagController.All)

	ValidationCodeMasterTagController := new(controllers.ValidationCodeMasterTagController)
	api = router.Group(prefixUrl + "/validation-code-master-tag")
	api.POST("", auth.CekToken, ValidationCodeMasterTagController.Save)
	api.POST("/filter", auth.CekToken, ValidationCodeMasterTagController.Filter)

	UploadMerchantBankLoanController := new(controllers.UploadMerchantBankLoanController)
	api = router.Group(prefixUrl + "/upload-merchant-bank-loan")
	api.POST("/upload", auth.CekToken, UploadMerchantBankLoanController.Upload)
	api.POST("/all", auth.CekToken, UploadMerchantBankLoanController.GetFilterPaging)
	api.POST("/download", auth.CekToken, UploadMerchantBankLoanController.HandleDownload)
	api.POST("/result-download", auth.CekToken, UploadMerchantBankLoanController.HandleResultDownload)
	api.GET("/download-example", auth.CekToken, UploadMerchantBankLoanController.HandleDownloadExample)

	MerchantBankLoanController := new(controllers.MerchantBankLoanController)
	api = router.Group(prefixUrl + "/merchant-bank-loan")
	api.POST("/filter", auth.CekToken, MerchantBankLoanController.Filter)
	api.POST("/find-sub", auth.CekToken, MerchantBankLoanController.FindSubMerchantBankLoan)

	UpdatedDataMerchantController := new(controllers.UpdatedDataMerchantController)
	api = router.Group(prefixUrl + "/updated-data")
	api.POST("/filter", auth.CekToken, UpdatedDataMerchantController.Filter)
	api.POST("/approve", auth.CekToken, UpdatedDataMerchantController.Approve)
	api.POST("/reject", auth.CekToken, UpdatedDataMerchantController.Reject)

	ReportUpdatedDataMerchantController := new(controllers.ReportUpdatedDataMerchantController)
	api = router.Group(prefixUrl + "/report-updated-data-merchant")
	api.POST("/send", auth.CekToken, ReportUpdatedDataMerchantController.Send)
	api.POST("/all", auth.CekToken, ReportUpdatedDataMerchantController.GetReport)
	api.POST("/download", auth.CekToken, ReportUpdatedDataMerchantController.DownloadPath)

	LoanProductMaintenanceController := new(controllers.LoanProductMaintenanceController)
	api = router.Group(prefixUrl + "/loan-product-maintenance")
	api.POST("", auth.CekToken, LoanProductMaintenanceController.Save)
	api.POST("/filter", auth.CekToken, LoanProductMaintenanceController.Filter)

	GenerateQrController := new(controllers.GenerateQrController)
	api = router.Group(prefixUrl + "/generate-qr")
	api.POST("/generate", auth.CekToken, GenerateQrController.Generate)

	BlastNotifController := new(controllers.BlastNotifController)
	api = router.Group(prefixUrl + "/blast-notif")
	api.POST("/send-all", auth.CekToken, BlastNotifController.SendAll)

	KategoriBisnisController := new(controllers.KategoriBisnisController)
	api = router.Group(prefixUrl + "/kategori-bisnis")
	api.GET("/all", auth.CekToken, KategoriBisnisController.FindAll)

	AppUserController := new(controllers.AppUserController)
	api = router.Group(prefixUrl + "/app-user")
	api.POST("/update", AppUserController.Update)

	ReportExportMerchantController := new(controllers.ReportExportMerchantController)
	api = router.Group(prefixUrl + "/report-export-merchant")
	api.POST("/filter", auth.CekToken, ReportExportMerchantController.GetReportExportMerchant)
	api.POST("/send", auth.CekToken, ReportExportMerchantController.Send)
	api.POST("/download", auth.CekToken, ReportExportMerchantController.DownloadPath)

	UploadMissingDataController := new(controllers.UploadMissingDataController)
	api = router.Group(prefixUrl + "/upload-missing")
	api.POST("/upload", auth.CekToken, UploadMissingDataController.Upload)
	api.POST("/all", auth.CekToken, UploadMissingDataController.GetFilterPaging)
	api.POST("/download", auth.CekToken, UploadMissingDataController.HandleDownload)
	api.POST("/result-download", auth.CekToken, UploadMissingDataController.HandleResultDownload)
	api.GET("/download-example", auth.CekToken, UploadMissingDataController.HandleDownloadExample)

	UploadMerchantActivatedController := new(controllers.UploadMerchantActivatedController)
	api = router.Group(prefixUrl + "/upload-merchant-activated")
	api.POST("/upload", auth.CekToken, UploadMerchantActivatedController.Upload)
	api.POST("/all", auth.CekToken, UploadMerchantActivatedController.GetFilterPaging)
	api.POST("/download", auth.CekToken, UploadMerchantActivatedController.HandleDownload)
	api.POST("/result-download", auth.CekToken, UploadMerchantActivatedController.HandleResultDownload)
	api.POST("/template-download", auth.CekToken, UploadMerchantActivatedController.HandleTemplateDownload)

	PartnerLinkController := new(controllers.PartnerLinkController)
	api = router.Group(prefixUrl + "/partner-link")
	api.POST("", auth.CekToken, PartnerLinkController.Save)
	api.POST("/filter", auth.CekToken, PartnerLinkController.Filter)
	api.DELETE("/delete/:id", auth.CekToken, PartnerLinkController.Delete)

	SettlementConfigController := new(controllers.SettlementConfigController)
	api = router.Group(prefixUrl + "/settlement-config")
	api.POST("", auth.CekToken, SettlementConfigController.Save)
	api.POST("/filter", auth.CekToken, SettlementConfigController.Filter)

	MerchantSettlementConfigController := new(controllers.MerchantSettlementConfigController)
	api = router.Group(prefixUrl + "/merchant-settlement-config")
	api.POST("", auth.CekToken, MerchantSettlementConfigController.Save)

	ReportExportAkuisisiSfaController := new(controllers.ReportExportAkuisisiSfaController)
	api = router.Group(prefixUrl + "/report-export-akuisisi-sfa")
	api.POST("/filter", auth.CekToken, ReportExportAkuisisiSfaController.GetReportExportAkuisisiSfa)
	api.POST("/send", auth.CekToken, ReportExportAkuisisiSfaController.Send)
	api.POST("/download", auth.CekToken, ReportExportAkuisisiSfaController.DownloadPath)

	MonitoringActivationFDSController := new(controllers.MonitoringActivationFDSController)
	api = router.Group(prefixUrl + "/monitoring-activation-fds")
	api.POST("/filter", auth.CekToken, MonitoringActivationFDSController.GetFilterPaging)

	AkuisisiSfaFailedController := new(controllers.AkuisisiSfaFailedController)
	api = router.Group(prefixUrl + "/akuisisi-sfa-failed")
	api.POST("/filter", auth.CekToken, AkuisisiSfaFailedController.GetFilterPaging)

	AcquititionsController := new(controllers.AcquititionsController)
	api = router.Group(prefixUrl + "/acquititions")
	api.POST("", auth.CekToken, AcquititionsController.Save)
	api.POST("/filter", auth.CekToken, AcquititionsController.GetFilterPaging)
	api.DELETE("/delete/:id", auth.CekToken, AcquititionsController.Delete)

	InstructionListController := new(controllers.InstructionListController)
	api = router.Group(prefixUrl + "/instruction-list")
	api.POST("", auth.CekToken, InstructionListController.Save)
	api.POST("/filter", auth.CekToken, InstructionListController.GetFilterPaging)
	api.DELETE("/delete/:id", auth.CekToken, InstructionListController.Delete)

	UploadImageController := new(controllers.UploadImageController)
	api = router.Group(prefixUrl + "/upload-image")
	api.POST("",auth.CekToken, UploadImageController.Upload )

	ConfigVaMerchantGroupController := new(controllers.ConfigVaMerchantGroupController)
	api = router.Group(prefixUrl + "/config-va-merchant-group")
	api.POST("/save", auth.CekToken , ConfigVaMerchantGroupController.Save )
	api.GET("/find-by-group-id/:id", auth.CekToken , ConfigVaMerchantGroupController.FindByGroupId )
	api.POST("/save-to-redis", auth.CekToken , ConfigVaMerchantGroupController.SaveToRedis )
	api.GET("/find-by-group-id-from-redis/:id" , auth.CekToken , ConfigVaMerchantGroupController.FindByGroupIdFromRedis )

	ottoRouter.Router = router

}

// InitTracing ..
func (ottoRouter *OttoRouter) InitTracing() {
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "PROD"
	}

	tracer, reporter, closer, err := ottotracing.InitTracing(fmt.Sprintf("%s::%s", nameService, hostName), openTracingSvr, ottotracing.WithEnableInfoLog(true))
	if err != nil {
		fmt.Println("Error :", err)
	}
	opentracing.SetGlobalTracer(tracer)

	ottoRouter.Closer = closer
	ottoRouter.Reporter = reporter
	ottoRouter.Tracer = tracer
	ottoRouter.Err = err
	ottoRouter.GinFunc = ottotracing.OpenTracer([]byte("api-request-"))
}

// Close ..
func (ottoRouter *OttoRouter) Close() {
	ottoRouter.Closer.Close()
	ottoRouter.Reporter.Close()
}
