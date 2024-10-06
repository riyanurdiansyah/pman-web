package model

type Config struct {
	AppConfig    AppConfig
	ConnectionDB ConnectionDB
	UserSession  UserSession
	Rijndael     Rijndael
	Redis        Redis
}

type UserSession struct {
	SessionID        string
	SessionKey       string
	SessionRedisName string
}

type ConnectionDB struct {
	PostgreConnection   string
	OracleConnection    string
	SqlServerConnection string
}

type Rijndael struct {
	Key string
}

type AppConfig struct {
	Name       string
	Port       string
	Version    string
	VersionAPI string
}

type Redis struct {
	Host     string
	Password string
}
