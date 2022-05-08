package configs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type appConfig struct {
	Name     string
	Stage    string
	Timezone string
	Cost     int
}

// App for service configuration
var App appConfig

func Load(name string) {
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	initConfig()
	initTimeZone()
}

func initConfig() {
	App = appConfig{
		Name:     viper.GetString("app.name"),
		Stage:    viper.GetString("app.stage"),
		Timezone: viper.GetString("app.timezone"),
		Cost:     viper.GetInt("app.cost"),
	}

	Database = databaseConfig{
		Driver:   viper.GetString("database.driver"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
		SSLMode:  viper.GetString("database.sslmode"),
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation(App.Timezone)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	time.Local = ict
}
