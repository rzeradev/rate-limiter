package configs

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr    string
	RedisPass    string
	RedisDB      int
	IPMaxReq     int
	TokenMaxReq  int
	BlockTime    int
	RateLimitDur string
}

var AppConfig Config

func LoadConfig(path *string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	defaultPath := ".env"
	if path == nil {
		path = &defaultPath
	}
	println(pwd, *path)
	err = godotenv.Load(filepath.Join(pwd, *path))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	ipMaxReq, _ := strconv.Atoi(os.Getenv("IP_MAX_REQ"))
	tokenMaxReq, _ := strconv.Atoi(os.Getenv("TOKEN_MAX_REQ"))
	blockTime, _ := strconv.Atoi(os.Getenv("BLOCK_TIME"))

	AppConfig = Config{
		RedisAddr:    os.Getenv("REDIS_ADDR"),
		RedisPass:    os.Getenv("REDIS_PASS"),
		RedisDB:      redisDB,
		IPMaxReq:     ipMaxReq,
		TokenMaxReq:  tokenMaxReq,
		BlockTime:    blockTime,
		RateLimitDur: os.Getenv("RATE_LIMIT_DUR"),
	}
}
