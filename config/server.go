package config

// Server 服务配置
type Server struct {
	Port          int    `mapstructure:"port" yaml:"port"`
	Mode          string `mapstructure:"mode" yaml:"mode"`
	RouterPrefix  string `mapstructure:"router-prefix" yaml:"router-prefix"`
	UseMultipoint bool   `mapstructure:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	UseOSS        bool   `mapstructure:"use-oss" yaml:"use-oss"`
	FilePath      string `mapstructure:"file-path" yaml:"file-path"`
}
