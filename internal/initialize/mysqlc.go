package initialize

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/nghiatk54/go_ecommerce_api/global"
	"go.uber.org/zap"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlC() {
	m := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.UserName, m.Password, m.Host, m.Port, m.DbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		checkErrorPanicC(err, "Mysql connect error")
	}
	global.Logger.Info("Initializing Mysqlc successfully")
	global.Mdbc = db
	// Set pool
	// SetPoolC()
}

// func SetPoolC() {
// 	m := global.Config.Mysql
// 	sqlDb, err := global.Mdbc.DB()
// 	if err != nil {
// 		fmt.Println("Mysql error: ", err)
// 	}
// 	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
// 	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
// 	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
// }
