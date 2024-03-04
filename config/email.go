package config

type Email struct {
	Sender   string `mapstructure:"sender" yaml:"sender"`
	Title    string `mapstructure:"title" yaml:"title"`
	Smtp     string `mapstructure:"smtp" yaml:"smtp"`
	Port     int    `mapstructure:"port" yaml:"port"`
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
}
