package configs

import (
	"fmt"
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type (
	conf struct {
		DBDriver      string `mapstructure:"DB_DRIVER"`
		DBHost        string `mapstructure:"DB_HOST"`
		DBPort        string `mapstructure:"DB_PORT"`
		DBUser        string `mapstructure:"DB_USER"`
		DBPassword    string `mapstructure:"DB_PASSWORD"`
		DBName        string `mapstructure:"DB_NAME"`
		WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
		JWTSecret     string `mapstructure:"JWT_SECRET"`
		JWtExperiesIn int    `mapstructure:"JWT_EXPERIESIN"`
		TokenAuth     *jwtauth.JWTAuth
	}
)

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.SetDefault("JWT_EXPERIESIN", 300)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	if cfg.JWtExperiesIn <= 0 {
		return nil, fmt.Errorf("JWT_EXPERIESIN deve ser maior que 0")
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}
