package db

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"

	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// EnvOfin ..
type EnvOfin struct {
	DbUser  string `envconfig:"OTTOFIN_POSTGRES_USER" default:"ottoagcfg"`
	DbPass  string `envconfig:"OTTOFIN_POSTGRES_PASS" default:"dTj*&56$es"`
	DbName  string `envconfig:"OTTOFIN_POSTGRES_NAME" default:"ottofindb"`
	DbHost  string `envconfig:"OTTOFIN_POSTGRES_HOST" default:"13.228.23.160"`
	DbPort  string `envconfig:"OTTOFIN_POSTGRES_PORT" default:"8432"`
	DbDebug bool   `envconfig:"OTTOFIN_POSTGRES_DEBUG" default:"true"`
	DbType  string `envconfig:"OTTOFIN_TYPE" default:"POSTGRES"`
	SslMode string `envconfig:"OTTOFIN_POSTGRES_SSL_MODE" default:"disable"`
}

var (
	DbOfinCon *gorm.DB
	DbOfinErr error
	dbOfinEnv EnvOfin
)

func init() {

	err := envconfig.Process("", &dbOfinEnv)
	if err != nil {
		fmt.Println("Failed to get DB env:", err)
	}

	if DbOfinOpen() != nil {
		//panic("DB Can't Open")
		fmt.Println("Can Open db Postgres")
	}
	DbOfinCon = GetDbOfinCon()
	DbOfinCon = DbOfinCon.LogMode(true)

}

// DbOfinOpen ..
func DbOfinOpen() error {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbOfinEnv.DbHost, dbOfinEnv.DbPort, dbOfinEnv.DbUser, dbOfinEnv.DbPass, dbOfinEnv.DbName, dbOfinEnv.SslMode)
	DbOfinCon, DbOfinErr = gorm.Open("postgres", args)

	if DbOfinErr != nil {
		logs.Error("open db Err ", DbOfinErr)
		return DbOfinErr
	}

	if errping := DbOfinCon.DB().Ping(); errping != nil {
		return errping
	}
	return nil
}

// GetDbOfinCon ..
func GetDbOfinCon() *gorm.DB {
	//TODO looping try connection until timeout
	// using channel timeout
	if errping := DbOfinCon.DB().Ping(); errping != nil {
		logs.Error("Db Not Connect test Ping :", errping)
		errping = nil
		if errping = DbOpen(); errping != nil {
			logs.Error("try to connect again but error :", errping)
		}
	}
	DbOfinCon.LogMode(dbEnv.DbDebug)
	return DbOfinCon
}
