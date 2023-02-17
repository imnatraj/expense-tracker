package models

type Config struct {
	App struct {
		Server struct {
			Port string `mapstructure:"port"`
			Host string `mapstructure:"host"`
		} `mapstructure:"server"`
		Redis struct {
			Password string `mapstructure:"password"`
			Host     string `mapstructure:"host"`
		} `mapstructure:"redis"`
		Postgres struct {
			Password string `mapstructure:"password"`
			User     string `mapstructure:"user"`
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			DBName   string `mapstructure:"name"`
		} `mapstructure:"postgres"`
		Postgres2 struct {
			Password string `mapstructure:"password"`
			User     string `mapstructure:"user"`
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			DBName   string `mapstructure:"name"`
		} `mapstructure:"postgres2"`
	} `mapstructure:"app"`
}
