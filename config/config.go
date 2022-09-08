package config

import (
	"fmt"
	"itdp-group3-backend/utils"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type Config struct {
	DBConfig
	APIConfig
	TokenConfig
	RedisClient
	MediaPath
}

type DBConfig struct {
	DBHost      string `mapstructure:"DB_HOST"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
	DBPort      string `mapstructure:"DB_PORT"`
	SSLMode     string `mapstructure:"SSL_MODE"`
	TimeZone    string `mapstructure:"TIME_ZONE"`
	Environment string `mapstructure:"ENV"`
}

type APIConfig struct {
	APIUrl string `mapstructure:"API_URL"`
}

type TokenConfig struct {
	ApplicationName     string `mapstructure:"APP_NAME"`
	JwtSignatureKey     string `mapstructure:"SECRET_KEY"`
	TokenDuration       string `mapstructure:"TOKEN_DURATION"`
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
	Redis               *redis.Client
}

type RedisClient struct {
	RedisAddress string `mapstructure:"REDIS_ADDRESS"`
}

type MediaPath struct {
	Path           string `mapstructure:"FILEPATH"`
	PathProduct    string `mapstructure:"FILEPATH_PRODUCT"`
	PathClientFeed string `mapstructure:"FILEPATH_CLIENT_FEED"`
	PathFeed       string `mapstructure:"FILEPATH_FEED"`
}

// loadConfig : get configuration from .env
func (c *Config) loadConfig() (config Config, err error) {
	var tokenDur int

	v := viper.New()
	v.AutomaticEnv()
	v.BindEnv("DB_HOST")
	v.BindEnv("DB_USER")
	v.BindEnv("DB_PASSWORD")
	v.BindEnv("DB_NAME")
	v.BindEnv("DB_PORT")
	v.BindEnv("SSL_MODE")
	v.BindEnv("TIME_ZONE")
	v.BindEnv("ENV")
	v.BindEnv("APP_NAME")
	v.BindEnv("SECRET_KEY")
	v.BindEnv("TOKEN_DURATION")
	v.BindEnv("REDIS_ADDRESS")
	v.BindEnv("SECRET")

	if err = v.Unmarshal(&config.APIConfig); err != nil {
		return
	}

	if err = v.Unmarshal(&config.DBConfig); err != nil {
		return
	}

	if err = v.Unmarshal(&config.TokenConfig); err != nil {
		return
	}
	fmt.Println("cek", err)
	if tokenDur, err = utils.StringToInt64(config.TokenDuration); err != nil {
		return
	}
	config.TokenConfig.JwtSigningMethod = jwt.SigningMethodHS256
	config.TokenConfig.AccessTokenLifeTime = time.Duration(tokenDur) * time.Minute

	if err = v.Unmarshal(&config.RedisClient); err != nil {
		return
	}
	newRedisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
		DB:   0,
	})
	config.TokenConfig.Redis = newRedisClient

	if err = v.Unmarshal(&config.MediaPath); err != nil {
		return
	}
	return
}

// NewConfig : export config to be used
func NewConfig() Config {
	cfg := Config{}
	cfg, err := cfg.loadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	return cfg
}
