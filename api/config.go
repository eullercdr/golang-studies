package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var config *conf

type conf struct {
	DB_Driver     string `mapstructure:"DB_DRIVER"`
	DB_Host       string `mapstructure:"DB_HOST"`
	DB_Port       string `mapstructure:"DB_PORT"`
	DB_User       string `mapstructure:"DB_USER"`
	DB_Password   string `mapstructure:"DB_PASSWORD"`
	DB_Name       string `mapstructure:"DB_NAME"`
	WebserverPort string `mapstructure:"WEBSERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigFile("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	config.TokenAuth = jwtauth.New("HS256", []byte(config.JWTSecret), nil)
	return config, nil
}
