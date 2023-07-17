package configurations

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	RapidAPIKey      string
	FootballHost     string
	HockeyHost       string
	BasketballHost   string
	TennisHost       string
	MMAHost          string
	CacheMinutes     int
	TimeZone         int
	PlusDay          int
	FootballLeague   string
	FootballTeam     string
	HockeyLeague     string
	BasketballLeague string
	TennisLeague     string
	Port             string
}

func LoadConfig() *Config {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("error loading .env file ", err)
	}
	return &Config{
		RapidAPIKey:      getStringValue("rapidApiKey"),
		FootballHost:     getStringValue("footballHost"),
		HockeyHost:       getStringValue("hockeyHost"),
		BasketballHost:   getStringValue("basketballHost"),
		TennisHost:       getStringValue("tennisHost"),
		MMAHost:          getStringValue("mmaHost"),
		CacheMinutes:     getIntValue("cacheMinutes"),
		TimeZone:         getIntValue("timeZone"),
		PlusDay:          getIntValue("plusDay"),
		FootballLeague:   getStringValue("footballLeague"),
		FootballTeam:     getStringValue("footballTeam"),
		HockeyLeague:     getStringValue("hockeyLeague"),
		BasketballLeague: getStringValue("basketballLeague"),
		TennisLeague:     getStringValue("tennisLeague"),
		Port:             getStringValue("port"),
	}
}

// проверка и получение строкогово значения
func getStringValue(e string) string {
	value, exists := os.LookupEnv(e)
	if !exists || value == "" {
		log.Fatalf("environment variable %s is not set", e)
	}
	return value
}

// проверка и получение числового значения
func getIntValue(e string) int {
	value, exists := os.LookupEnv(e)
	if !exists || value == "" {
		log.Fatalf("environment variable %s is not set", e)
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("error converting value for string %s to int: %s", e, err)
	}
	return intValue
}
