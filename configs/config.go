package configs

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

type Conf struct {
	DBAutoMigrate bool   `mapstructure:"DB_AUTO_MIGRATE"`
	DBUser        string `mapstructure:"DB_USER"`
	DBSSLmode     string `mapstructure:"DB_SSL_MODE"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBhost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort int    `mapstructure:"DB_WEB_SERVER_PORT"`
	DBtimezone    string `mapstructure:"DB_TIMEZONE"`
	// Jwt related firlds
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExperesIn int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth    *jwtauth.JWTAuth
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv() // give OS EvVars preference instead of .env or any env var file
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}
