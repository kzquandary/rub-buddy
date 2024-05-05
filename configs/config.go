package configs

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ProgrammingConfig struct {
	ServerPort int
	DBPort     uint16
	DBHost     string
	DBUser     string
	DBPass     string
	DBName     string
	Secret     string
	OpenAI     string
	ProjectID  string
	BucketName string
}

func InitConfig() *ProgrammingConfig {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("Config: Cannot start program, failed to read configuration:", err)
		return nil
	}

	var res = new(ProgrammingConfig)
	res.ServerPort = viper.GetInt("SERVER")
	res.DBPort = uint16(viper.GetInt("DBPort"))
	res.DBHost = viper.GetString("DBHost")
	res.DBUser = viper.GetString("DBUser")
	res.DBPass = viper.GetString("DBPass")
	res.DBName = viper.GetString("DBName")
	res.Secret = viper.GetString("Secret")
	res.OpenAI = viper.GetString("KEY_OPEN_AI")
	res.ProjectID = viper.GetString("GOOGLE_PROJECT_ID")
	res.BucketName = viper.GetString("GOOGLE_BUCKET_NAME")

	return res
}
