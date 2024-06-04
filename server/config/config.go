package config

import (
	"os"
	"regexp"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string `env:"APP_ENV"`
	AppPort string `env:"APP_PORT"`

	DBHost string `env:"DB_HOST"`
	DBPort string `env:"DB_PORT"`
	DBUser string `env:"DB_USER"`
	DBPass string `env:"DB_PASS"`
	DBName string `env:"DB_NAME"`

	DBMaxOpenConns    int `env:"DB_MAX_OPEN_CONNS"`
	DBConnMaxLifetime int `env:"DB_CONN_MAX_LIFETIME"`
	DBMaxIdleConn     int `env:"DB_MAX_IDLE_CONN"`
	DBConnMaxIdleTime int `env:"DB_CONN_MAX_IDLE_TIME"`

	AccessTokenSecret      string `env:"ACCESS_TOKEN_SECRET"`
	AccessTokenExpiryHour  int    `env:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenSecret     string `env:"REFRESH_TOKEN_SECRET"`
	RefreshTokenExpiryHour int    `env:"REFRESH_TOKEN_EXPIRY_HOUR"`

	TokenIssuer string `env:"TOKEN_ISSUER"`

	AuthCookieName     string `env:"AUTH_COOKIE_NAME"`
	AuthCookieMaxAge   int    `env:"AUTH_COOKIE_MAX_AGE"`
	AuthCookieSecure   bool   `env:"AUTH_COOKIE_SECURE"`
	AuthCookieHttpOnly bool   `env:"AUTH_COOKIE_HTTP_ONLY"`
}

func GetConfig() (*Config, error) {
	config := Config{}
	projectName := regexp.MustCompile(`^(.*rlzone[/\\]server)`)
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	path = string(projectName.Find([]byte(path))) + "/.env"
	err = godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	err = env.Parse(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
