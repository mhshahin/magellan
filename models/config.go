package models

type Config struct {
	Redis RedisCfg `mapstructure:"redis"`
}

type RedisCfg struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
