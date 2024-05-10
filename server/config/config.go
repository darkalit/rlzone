package config

import "github.com/spf13/viper"

type Config struct {
	AppEnv  string `mapstructure:"APP_ENV"`
	AppPort string `mapstructure:"APP_PORT"`

	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`

	DBMaxOpenConns    int `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBConnMaxLifetime int `mapstructure:"DB_CONN_MAX_LIFETIME"`
	DBMaxIdleConn     int `mapstructure:"DB_MAX_IDLE_CONN"`
	DBConnMaxIdleTime int `mapstructure:"DB_CONN_MAX_IDLE_TIME"`

	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`

	AuthCookieName     string `mapstructure:"AUTH_COOKIE_NAME"`
	AuthCookieMaxAge   int    `mapstructure:"AUTH_COOKIE_MAX_AGE"`
	AuthCookieSecure   bool   `mapstructure:"AUTH_COOKIE_SECURE"`
	AuthCookieHttpOnly bool   `mapstructure:"AUTH_COOKIE_HTTP_ONLY"`
}

func GetConfig() (*Config, error) {
	config := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
