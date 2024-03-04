package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// Mysql 配置
type Mysql struct {
	Prefix       string `mapstructure:"prefix"  yaml:"prefix"`
	Port         string `mapstructure:"port" yaml:"port"`
	Config       string `mapstructure:"config" yaml:"config"`     // 高级配置
	DbName       string `mapstructure:"db-name" yaml:"db-name"`   // 数据库名
	Username     string `mapstructure:"username" yaml:"username"` // 数据库密码
	Password     string `mapstructure:"password" yaml:"password"` // 数据库密码
	Path         string `mapstructure:"path" yaml:"path"`
	Engine       string `mapstructure:"engine" yaml:"engine" default:"InnoDB"` //数据库引擎，默认InnoDB
	LogMode      string `mapstructure:"log-mode" yaml:"log-mode"`              // 是否开启Gorm全局日志
	MaxIdleConns int    `mapstructure:"max-idle-conns" yaml:"max-idle-conns"`  // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" yaml:"max-open-conns"`  // 打开到数据库的最大连接数
	Singular     bool   `mapstructure:"singular" yaml:"singular"`              //是否开启全局禁用复数，true表示开启
	LogZap       bool   `mapstructure:"log-zap" yaml:"log-zap"`
}

// Dsn 连接配置
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DbName + "?" + m.Config
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}

func (m *Mysql) InitConfig() mysql.Config {
	return mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
}

func (m *Mysql) Connection() *gorm.Config {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // （日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 使用彩色打印
		},
	)
	return &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,  // 单数表名
			NoLowerCase:   false, // 关闭小写转换
		},
		Logger: newLogger,
	}
}
