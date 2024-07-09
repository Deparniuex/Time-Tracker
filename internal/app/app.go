package app

import (
	storage "example.com/tracker/internal/storage/postgres"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	//
}

func NewApp() *App {
	return &App{
		//
	}
}

func (a *App) SetupConfig(path string) error {
	viper.SetConfigFile(path)

	return viper.ReadInConfig()
}

func (a *App) Run() error {
	db, err := storage.ConnectDB(&storage.PostgresConfig{
		Host:     viper.GetString("DB_ADDRESS"),
		Port:     viper.GetInt("DB_PORT"),
		User:     viper.GetString("DB_USERNAME"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
	})
	if err != nil {
		return err
	}
	logrus.Info("Connection to DB success")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("../../migrations", "postgres", driver)
	return nil
}
