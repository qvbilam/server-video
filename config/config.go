package config

type ServerConfig struct {
	Name             string           `mapstructure:"name" json:"name"`
	Tags             []string         `mapstructure:"tags" json:"tags"`
	DBConfig         DBConfig         `mapstructure:"db" json:"db"`
	ESConfig         ESConfig         `mapstructure:"es" json:"es"`
	UserServerConfig UserServerConfig `mapstructure:"user-server" json:"user-server"`
}

type DBConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
}

type ESConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type UserServerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int64  `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
