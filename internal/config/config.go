package config

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type (
	Configuration struct {
		Application		Application		`mapstructure:"app"`
		PostgreSQL		PostgreSQL		`mapstructure:"psql"`
		Google			Google			`mapstructure:"google"`
	}
	Application struct {
		Name			string			`mapstructure:"name"`
		Version			string			`mapstructure:"version"`
		Port			int				`mapstructure:"port"`
		Environment		string			`mapstructure:"environment"`
		Host			string			`mapstructure:"host"`
		Timeout			time.Duration	`mapstructure:"timeout"`
		LogOption		string			`mapstructure:"log_option"`
		LogLevel		string			`mapstructure:"log_level"`
	}
	PostgreSQL struct {
		User			string			`mapstructure:"user"`
		Password		string			`mapstructure:"password"`
		Host			string			`mapstructure:"host"`
		Name			string			`mapstructure:"name"`
		Port			int				`mapstructure:"port"`
		SSLMode			string			`mapstructure:"ssl_mode"`
	}
	Google struct {
		ClientID		string			`mapstructure:"client_id"`
		ClientSecret	string			`mapstructure:"client_secret"`
		Redirect		string			`mapstructure:"redirect"`
		State			string			`mapstructure:"state"`
	}
)

func New(ctx context.Context) (*Configuration, error) {
	var config Configuration

	viper.AutomaticEnv()
	environment := strings.ToLower(viper.GetString("env"))
	configName := fmt.Sprintf("config.%s", environment)

	viper.AddConfigPath("./config")
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}