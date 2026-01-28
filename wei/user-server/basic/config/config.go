package config

type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}
type AppConfig struct {
	Mysql
}
