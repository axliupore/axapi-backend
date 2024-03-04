package config

// Config 读取配置
type Config struct {
	Server    Server    `mapstructure:"server" yaml:"server"`
	Mysql     Mysql     `mapstructure:"mysql" yaml:"mysql"`
	Redis     Redis     `mapstructure:"redis" yaml:"redis"`
	JWT       JWT       `mapstructure:"jwt" yaml:"jwt"`
	Zap       Zap       `mapstructure:"zap" yaml:"zap"`
	AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" yaml:"aliyun-oss"`
	Email     Email     `mapstructure:"email" yaml:"email"`
}
