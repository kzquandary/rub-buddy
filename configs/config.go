package configs

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type ProgrammingConfig struct {
	ServerPort              int
	DBPort                  uint16
	DBHost                  string
	DBUser                  string
	DBPass                  string
	DBName                  string
	Secret                  string
	OpenAI                  string
	BaseURLFE               string
}

func InitConfig() *ProgrammingConfig {
	var res = new(ProgrammingConfig)
	res, errorRes := loadConfig()

	logrus.Error(errorRes)
	if res == nil {
		logrus.Error("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func readData() *ProgrammingConfig {
	var data = new(ProgrammingConfig)
	data, _ = loadConfig()

	if data == nil {
		err := godotenv.Load(".env")
		data, errorData := loadConfig()

		fmt.Println(errorData)

		if err != nil || data == nil {
			return nil
		}
	}
	return data
}

func loadConfig() (*ProgrammingConfig, error) {
	var error error
	var res = new(ProgrammingConfig)
	var permit = true

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid port value,", err.Error())
			permit = false
		}
		res.ServerPort = port
	} else {
		permit = false
		error = errors.New("Port undefined")
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid port value,", err.Error())
			permit = false
		}

		res.DBPort = uint16(port)
	} else {
		permit = false
		error = errors.New("DB Port undefined")
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = val
	} else {
		permit = false
		error = errors.New("DB Host undefined")
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = val
	} else {
		permit = false
		error = errors.New("DB User undefined")
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		res.DBPass = val
	} else {
		permit = false
		error = errors.New("DB Pass undefined")
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBName = val
	} else {
		permit = false
		error = errors.New("DB Name undefined")
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	} else {
		permit = false
		error = errors.New("Secret undefined")
	}

	if val, found := os.LookupEnv("KEY_OPEN_AI"); found {
		res.OpenAI = val
	} else {
		permit = false
		error = errors.New("KEY OPEN AI undefined")
	}

	if val, found := os.LookupEnv("BASE_URL_FE"); found {
		res.BaseURLFE = val
	} else {
		permit = false
		error = errors.New("Base URL FE undefined")
	}

	if !permit {
		return nil, error
		
		//DEV MODE
		// return res
	}

	return res, nil
}
