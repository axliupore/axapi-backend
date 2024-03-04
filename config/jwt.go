package config

// JWT 配置
type JWT struct {
	SigningKey  string `mapstructure:"signing-key" yaml:"signing-key"`   // jwt签名
	ExpiresTime string `mapstructure:"expires-time" yaml:"expires-time"` // 过期时间
	BufferTime  string `mapstructure:"buffer-time" yaml:"buffer-time"`   // 缓冲时间
	Issuer      string `mapstructure:"issuer" yaml:"issuer"`             // 发行者
}
