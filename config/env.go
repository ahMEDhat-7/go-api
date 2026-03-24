package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type Config struct{
	PublicHost	string
	Port 		string
	DBUser		string
	DBPassword	string
	DBAddress	string
	DBName		string
}

func initConfig() Config{
	godotenv.Load()
	return Config{
		PublicHost:		getEnv("PUBLIC_HOST","http://localhost"),
		Port:			getEnv("PORT","8000"),
		DBUser:			getEnv("DB_USER","root"),
		DBPassword:		getEnv("DB_PASSWORD","mysql"),
		DBAddress:		fmt.Sprintf("%s:%s",getEnv("DB_HOST","localhost"),getEnv("DB_PORT","3306")),
		DBName:			getEnv("DB_NAME","mysql"),
	}

}

func getEnv(key string,fallback string) string {
	if value ,ok := os.LookupEnv(key);ok {
		return value
	} 

	return fallback

}


var Envs = initConfig()