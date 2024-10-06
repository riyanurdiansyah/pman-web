package constanta

import (
	"os"

	"kalbenutritionals.com/pman/app/helper/model"
)

func Get() model.Config {
	return model.Config{
		AppConfig: model.AppConfig{
			Name:       os.Getenv("APP_NAME"),
			Port:       os.Getenv("APP_PORT"),
			Version:    os.Getenv("APP_VERSION"),
			VersionAPI: os.Getenv("APP_VERSION_API"),
		},
		ConnectionDB: model.ConnectionDB{
			PostgreConnection:   os.Getenv("PostgreConnection"),
			OracleConnection:    os.Getenv("OracleConnection"),
			SqlServerConnection: os.Getenv("SqlServerConnection"),
		},
		UserSession: model.UserSession{
			SessionID:        os.Getenv("SESSION_ID"),
			SessionKey:       os.Getenv("SESSION_KEY"),
			SessionRedisName: os.Getenv("SESSION_REDIS_NAME"),
		},
		Redis: model.Redis{
			Host:     os.Getenv("ReddisConnection"),
			Password: "",
		},
		Rijndael: model.Rijndael{
			Key: os.Getenv("RijndaelKey"),
		},
	}
}
