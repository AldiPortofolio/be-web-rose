
export JAEGER_HOSTURL="13.250.21.165:5775"
export CONSULREG_HOSTADDRES=13.250.21.165:8500


#main.go
export LOGGER_FILENAME="./otto-logger.log"
export MAXPROCS="1"
export ROSE_BE_GO_PORT="0.0.0.0:8098"

#routers/route.go
export ROSE_BE_GO="ROSE-BE-GO"
export APPS_DEBUG="debug"
export ROSE_BE_GO_READ_TIMEOUT="120"
export ROSE_BE_GO_WRITE_TIMEOUT="120"

#redis/redis_singleredis.go
export ROSE_REDIS_HOST="34.101.222.33"
export ROSE_REDIS_PORT="6077"
export ROSE_REDIS_DB_TYPE="3"


#db/postgres.go
export ROSE_POSTGRES_USER=ottoagcfg
#ROSE_POSTGRES_PASS="dTj*&56$es", ditambah backslash (\) untuk escape karakter '$'
export ROSE_POSTGRES_PASS="dTj*&56\$es"
export ROSE_POSTGRES_NAME="rosedb"
export ROSE_POSTGRES_HOST="13.228.23.160"
export ROSE_POSTGRES_PORT="8432"
export ROSE_POSTGRES_DEBUG="true"
export ROSE_TYPE="postgres"
export ROSE_POSTGRES_SSL_MODE="disable"

#db/postgresOfin.go
export OTTOFIN_POSTGRES_USER=ottoagcfg
export OTTOFIN_POSTGRES_PASS="dTj*&56\$es"
export OTTOFIN_POSTGRES_NAME="ottofindb"
export OTTOFIN_POSTGRES_HOST="13.228.23.160"
export OTTOFIN_POSTGRES_PORT="8432"
export OTTOFIN_POSTGRES_DEBUG="true"
export OTTOFIN_TYPE="postgres"
export OTTOFIN_POSTGRES_SSL_MODE="disable"


#kafka/publisher/kafka.go
export KAFKA_PUBLISH_REST="http://13.228.25.85:8922/v1/publish"
#prod 10.10.43.21:8922

#pathCsv
export PATH_UPLOAD_NMID="/opt/app-rose/nmid/upload/"
export PATH_DOWNLOAD_NMID="/opt/app-rose/nmid/process/"
export PATH_DOWNLOAD_REPORT_FINISHED="/opt/app-rose/report/finish/"
export PATH_DOWNLOAD_RESULT_NMID="/opt/app-rose/nmid/result/nmid/"
export PATH_DOWNLOAD_RESULT_QR="/opt/app-rose/nmid/result/qr/"
export PATH_MERCHANT_AGG_UPLOAD="/opt/app-rose/merchant-agg/upload/"
export PATH_DOWNLOAD_REPORT_REJECTED="/opt/app-rose/report/rejected/"

#utils/http/http.go
export HTTP_DEBUG_CLIENT="true"
export HTTP_TIMEOUT="60s"
export HTTP_RETRY_BAD="1"

#services/PortalActivationService.go
export EMAIL_SMTP_ADDRESS="smtp.office365.com"
export EMAIL_SENDER="ottopay@ottopay.id"
export EMAIL_PASSWORD="Mutiara2019"
export EMAIL_SMTP_PORT="587"
export EMAIL_URL="http://18.139.224.183:8080/activate/"

export PORTAL_BE_HOST_URL="http://13.228.25.85:8000/"
export PORTAL_FE_URL="http://18.139.224.183:8080/login"

#services/MerchantService.go
export ROSE_BE_GO_UPGRADE_MERCHANT_TOPIC="rose-worker-upgrade-fds-topic"
export ROSE_WORKER_APPROVE_AGG_UPLOAD_KAFKA_TOPICS="rose-worker-approve-agg-upload-topic"
export ROSE_BE_GO_REPORT_REJECTED_TOPIC="rose-report-rejected-topic"
export ROSE_BE_GO_REPORT_EXPORT_MERCHANT_TOPIC="rose-report-export-merchant-topic"

#redis/redis_cluster_redis.go
export REDIS_HOST_CLUSTER1="34.101.208.23:8177"
export REDIS_HOST_CLUSTER2="34.101.208.23:8178"
export REDIS_HOST_CLUSTER3="34.101.208.23:8179"
export REDIS_HOST_SLAVE1="34.101.208.23:8174"
export REDIS_HOST_SLAVE2="34.101.208.23:8175"
export REDIS_HOST_SLAVE3="34.101.208.23:8176"
export REDIS_HOST="34.101.222.33"
export REDIS_PORT="6077"

#services/VersionAppService.go
export OTTOMART_KEY_VERSION="OTTOMART:ANDROID-VERSION"
export NFC_KEY_VERSION="NFC:ANDROID-VERSION"
export INDOMARCO_KEY_VERSION="INDOMARCO:ANDROID-VERSION"
export SFA_KEY_VERSION="SFA:ANDROID-VERSION"

#service/UploadMerchantService.go
export ROSE_UPLOAD_MERCHANT_PATH="/opt/app-rose/merchant/upload/"

#controller/UploadMerchantController.go
export PATH_DOWNLOAD_MERCHANT="/opt/app-rose/merchant/process/"
export PATH_DOWNLOAD_RESULT_MERCHANT="/opt/app-rose/merchant/result/"

#host/ottomart/ottomart.go
export OTTOMART_HOST="http://13.228.25.85:8999/"


#service/UploadMerchantWipService.go
export ROSE_UPLOAD_MERCHANT_WIP_PATH="/opt/app-rose/merchant-wip/upload/"

#controller/UploadMerchantWipController.go
export PATH_DOWNLOAD_MERCHANT_WIP="/opt/app-rose/merchant-wip/process/"
export PATH_DOWNLOAD_RESULT_MERCHANT_WIP="/opt/app-rose/merchant-wip/result/"

#new
#service/QrPrePrintedService.go
export ROSE_BE_GO_QR_PREPRINTED_TOPIC="qr-preprinted-topic"
export PATH_DOWNLOAD_QR_PREPRINTED_RESULT="/opt/app-rose/qrpreprinted/result/"

export ROSE_UPLOAD_MERCHANT_OK_PATH="/opt/app-rose/merchant-ok/upload/"
export PATH_DOWNLOAD_MERCHANT_OK="/opt/app-rose/merchant-ok/process/"
export PATH_DOWNLOAD_RESULT_MERCHANT_OK="/opt/app-rose/merchant-ok/result/"

export OTTOPAY_HOST="http://13.228.25.85:8987"
export OTTOPAY_ENDPOINT_PUSH_NOTIF="/ottopay/v0.1.0/sendnotif"

export ROSE_UPLOAD_FEE_MDR_SETTING_PATH="/opt/app-rose/fee-mdr-setting/upload/"
export PATH_DOWNLOAD_FEE_MDR="/opt/app-rose/fee-mdr-setting/process/"
export PATH_DOWNLOAD_RESULT_FEE_MDR="/opt/app-rose/fee-mdr-setting/result/"
export PATH_DOWNLOAD_TEMPLATE_FEE_MDR="/opt/app-rose/fee-mdr-setting/"

#new
export ROSE_BE_GO_REPORT_PREPRINTED_TOPIC="rose-report-preprinted-topic"
export PATH_DOWNLOAD_REPORT_PREPRINTED="/opt/app-rose/report/qr-preprinted/"

export ROSE_BE_GO_REPORT_CRM_PENUGASAN_TOPIC="rose-report-crm-penugasan-topic"
export PATH_DOWNLOAD_REPORT_CRM_PENUGASAN="/opt/app-rose/report/crm-penugasan/"

#controller/UploadMerchantNonWipController.go
export PATH_DOWNLOAD_MERCHANT_NON_WIP="/opt/app-rose/merchant-non-wip/process/"
export PATH_DOWNLOAD_RESULT_MERCHANT_NON_WIP="/opt/app-rose/merchant-non-wip/result/"

#service/UploadMerchantWipService.go
export ROSE_UPLOAD_MERCHANT_NON_WIP_PATH="/opt/app-rose/merchant-non-wip/upload/"


#service/UploadMerchantMasterTagService.go
# export ROSE_UPLOAD_MERCHANT_MASTER_TAG_PATH="/opt/app-rose/merchant-master-tag/upload/"
export ROSE_DOWNLOAD_MERCHANT_MASTER_TAG_PATH="/opt/app-rose/merchant-master-tag/process/"
export PATH_DOWNLOAD_RESULT_MERCHANT_MASTER_TAG="/opt/app-rose/merchant-master-tag/result/"
# export PATH_DOWNLOAD_EXAMPLE_MERCHANT_MASTER_TAG="/opt/app-rose/merchant-master-tag/"
# local
export ROSE_UPLOAD_MERCHANT_MASTER_TAG_PATH="/Users/abdulah/Documents/opt/app-rose/merchant-master-tag/upload/"
export PATH_DOWNLOAD_EXAMPLE_MERCHANT_MASTER_TAG="/Users/abdulah/Documents/opt/app-rose/merchant-master-tag/"

#ininew
export ROSE_UPLOAD_MERCHANT_BANK_LOAN_PATH="/opt/app-rose/merchant-bank-loan/upload/"
export ROSE_DOWNLOAD_MERCHANT_BANK_LOAN_PATH="/opt/app-rose/merchant-bank-loan/process/"
export PATH_DOWNLOAD_RESULT_MERCHANT_BANK_LOAN="/opt/app-rose/merchant-bank-loan/result/"
export PATH_DOWNLOAD_EXAMPLE_MERCHANT_BANK_LOAN="/opt/app-rose/merchant-bank-loan/"

export ROSE_BE_GO_REPORT_UPDATED_DATA_MERCHANT_TOPIC="rose-report-updated-data-merchant-topic"
export PATH_DOWNLOAD_REPORT_UPDATED_DATA_MERCHANT="/opt/app-rose/report/updated-data-merchant/"

#host/op_bank
export OP_BANK_HOST="http://13.228.25.85:8989"
export OP_BANK_INQUIRY_INTERNAL="/v1.0/mandiri/account/inquiry"
export OP_BANK_INQUIRY_EXTERNAL="/v1.0/mandiri/account/inquiry/external"

#host/blast_notif
export BLAST_NOTIF_HOST="http://13.228.25.85:8987"
export BLAST_NOTIF_SEND_ALL="/ottopay/v0.1.0/sendnotif/all"


export ROSE_UPLOAD_DATA_MISSING_PATH="/opt/app-rose/merchant-missing/upload/"
export ROSE_DOWNLOAD_MISSING_DATA_PATH="/opt/app-rose/merchant-missing/process/"
export PATH_DOWNLOAD_RESULT_DATA_MISSING="/opt/app-rose/merchant-missing/result/"
export PATH_DOWNLOAD_EXAMPLE_MERCHANT_MISSING="/opt/app-rose/merchant-missing/"

#AuthActivation

go run main.go