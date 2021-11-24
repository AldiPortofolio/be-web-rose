swag init
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
export ROSE_REDIS_HOST="13.228.23.160"
export ROSE_REDIS_PORT="8377"
export ROSE_REDIS_DB_TYPE="3"


#db/postgres.go
export ROSE_POSTGRES_USER=ottoagcfg
#ROSE_POSTGRES_PASS="dTj*&56$es", ditambah backslash (\) untuk escape karakter '$'
export ROSE_POSTGRES_PASS="dTj*&56\$es"
export ROSE_POSTGRES_NAME="rosedb"
export ROSE_POSTGRES_HOST="34.101.150.6"
export ROSE_POSTGRES_PORT="5432"
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

#redis/redis_cluster_redis.go
export REDIS_HOST_CLUSTER1="13.228.23.160:8079"
export REDIS_HOST_CLUSTER2="13.228.23.160:8078"
export REDIS_HOST_CLUSTER3="13.228.23.160:8077"
export REDIS_HOST="13.228.23.160"
export REDIS_PORT="8377"

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


export OTTOPAY_HOST="http://34.101.87.199:8989"
export OTTOPAY_ENDPOINT_PUSH_NOTIF="/ottokonek/v0.1.0/sendnotif"

go run main.go