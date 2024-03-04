package initialize

import (
	"github.com/axliupore/axapi/axapi-backend/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitMysql 初始化 mysql 数据库
func InitMysql() {
	m := &global.Config.Mysql
	if m.DbName == "" {
		panic("config.yaml未正确配置数据库")
	}
	// 使用 gorm 进行数据库的配置
	mysqlConfig := m.InitConfig()
	// 进行连接的操作
	mysqlConnection := m.Connection()
	// 进行连接
	if db, err := gorm.Open(mysql.New(mysqlConfig), mysqlConnection); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		global.Db = db
	}
}
