package configs

import (
	"os"
	"strconv"
)

type ProgrammingConfig struct {
	ServerPort int
	DBPort     int
	DBHost     string
	DBUser     string
	DBPass     string
	DBName     string
	Secret     string
	OpenAI     string
	ProjectID  string
	BucketName string
	Midtrans   MidtransConfig
}

type MidtransConfig struct {
	ClientKey string
	ServerKey string
}

func InitConfig() *ProgrammingConfig {
	// err := godotenv.Load()
	// if err != nil {
	// 	logrus.Error("Config: Cannot start program, failed to load .env file:", err)
	// 	return nil
	// }
	var res = new(ProgrammingConfig)

	//Dev Mode
	// res.ServerPort = viper.GetInt("SERVER")
	// res.DBPort = uint16(viper.GetInt("DBPort"))
	// res.DBHost = viper.GetString("DBHost")
	// res.DBUser = viper.GetString("DBUser")
	// res.DBPass = viper.GetString("DBPass")
	// res.DBName = viper.GetString("DBName")
	// res.Secret = viper.GetString("Secret")
	// res.OpenAI = viper.GetString("KEY_OPEN_AI")
	// res.ProjectID = viper.GetString("GOOGLE_PROJECT_ID")
	// res.BucketName = viper.GetString("GOOGLE_BUCKET_NAME")
	// res.Midtrans.ClientKey = viper.GetString("MIDTRANS_CLIENT_KEY")
	// res.Midtrans.ServerKey = viper.GetString("MIDTRANS_SERVER_KEY")

	//Prod Mode
	res.ServerPort, _ = strconv.Atoi(os.Getenv("SERVER"))
	res.DBPort, _ = strconv.Atoi(os.Getenv("DBPort"))
	res.DBHost = os.Getenv("DBHost")
	res.DBUser = os.Getenv("DBUser")
	res.DBPass = os.Getenv("DBPass")
	res.DBName = os.Getenv("DBName")
	res.Secret = os.Getenv("Secret")
	res.OpenAI = os.Getenv("KEY_OPEN_AI")
	res.ProjectID = os.Getenv("GOOGLE_PROJECT_ID")
	res.BucketName = os.Getenv("GOOGLE_BUCKET_NAME")
	res.Midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	res.Midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")
	res.Midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	return res
}
