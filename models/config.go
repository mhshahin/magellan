package models

type Config struct {
	Tile38 Tile38Cfg `mapstructure:"tile38"`
	Redis  RedisCfg  `mapstructure:"redis"`
}

type Tile38Cfg struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type RedisCfg struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
