package db

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Env ..
type Env struct {
	DbUser  string `envconfig:"ROSE_POSTGRES_USER" default:"ottoagcfg"`
	DbPass  string `envconfig:"ROSE_POSTGRES_PASS" default:"dTj*&56$es"`
	DbName  string `envconfig:"ROSE_POSTGRES_NAME" default:"rosedb"`
	DbHost  string `envconfig:"ROSE_POSTGRES_HOST" default:"13.228.23.160"`
	DbPort  string `envconfig:"ROSE_POSTGRES_PORT" default:"8432"`
	DbDebug bool   `envconfig:"ROSE_POSTGRES_DEBUG" default:"true"`
	DbType  string `envconfig:"ROSE_TYPE" default:"POSTGRES"`
	SslMode string `envconfig:"ROSE_POSTGRES_SSL_MODE" default:"disable"`
}

var (
	DbCon *gorm.DB
	DbErr error
	dbEnv Env
)

func init() {

	err := envconfig.Process("", &dbEnv)
	if err != nil {
		fmt.Println("Failed to get DB env:", err)
	}

	if DbOpen() != nil {
		//panic("DB Can't Open")
		fmt.Println("Can Open db Postgres")
	}
	DbCon = GetDbCon()
	DbCon = DbCon.LogMode(true)

}

// DbOpen ..
func DbOpen() error {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbEnv.DbHost, dbEnv.DbPort, dbEnv.DbUser, dbEnv.DbPass, dbEnv.DbName, dbEnv.SslMode)
	DbCon, DbErr = gorm.Open("postgres", args)

	if DbErr != nil {
		logs.Error("open db Err ", DbErr)
		return DbErr
	}

	if errping := DbCon.DB().Ping(); errping != nil {
		return errping
	}
	return nil
}

// GetDbCon ..
func GetDbCon() *gorm.DB {
	//TODO looping try connection until timeout
	// using channel timeout
	if errping := DbCon.DB().Ping(); errping != nil {
		logs.Error("Db Not Connect test Ping :", errping)
		errping = nil
		if errping = DbOpen(); errping != nil {
			logs.Error("try to connect again but error :", errping)
		}
	}
	DbCon.LogMode(dbEnv.DbDebug)
	return DbCon
}
