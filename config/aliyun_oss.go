package config

type AliyunOSS struct {
	Endpoint        string `mapstructure:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `mapstructure:"access-key-id" yaml:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret" yaml:"access-key-secret"`
	BucketName      string `mapstructure:"bucket-name" yaml:"bucket-name"`
	BucketUrl       string `mapstructure:"bucket-url" yaml:"bucket-url"`
	BasePath        string `mapstructure:"base-path" yaml:"base-path"`
}
