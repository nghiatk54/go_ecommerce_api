package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/nghiatk54/go_ecommerce_api/global"
	"github.com/nghiatk54/go_ecommerce_api/internal/model"
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
		checkErrorPanic(err, "Mysql connect error")
	}
	global.Logger.Info("Initializing Mysqlc successfully")
	global.Mdbc = db
	// Set pool
	SetPoolC()
	// Migrate tables
	migrateTablesC()
}

func SetPoolC() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Println("Mysql error: ", err)
	}
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTablesC() {
	err := global.Mdb.AutoMigrate(
		// &po.User{},
		// &po.Role{},
		&model.GoCrmUserV2{},
	)
	if err != nil {
		fmt.Println("Migrating tables error: ", err)
	}
}
