package config

type ServerConfig struct {
	Name       string      `mapstructure:"name"`
	Port       int         `mapstructure:"port"`
	LogAddress string      `mapstructure:"logAddress"`
	MysqlInfo  MysqlConfig `mapstructure:"mysql"`
	RedisInfo  RedisConfig `mapstructure:"redis"`
}

type MysqlConfig struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
