package app

import (
	"os"
	"os/signal"

	"example.com/tracker/internal/handler"
	"example.com/tracker/internal/httpserver"
	"example.com/tracker/internal/repository/pgrepo"
	"example.com/tracker/internal/service"
	storage "example.com/tracker/internal/storage/postgres"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		logrus.Fatal(err)
		return err
	}
	//running from cmd
	migrator, err := migrate.NewWithDatabaseInstance("file://../migrations/", "postgres", driver)
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		logrus.Fatal(err)
		return err
	}
	repository := pgrepo.New(db)
	services := service.New(repository)
	handler := handler.New(services)
	server := httpserver.NewServer(handler.InitRouter(), &httpserver.ServerConfig{
		Port: viper.GetString("API_PORT"),
	})
	server.Start()
	logrus.Info("Server started")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		logrus.Infof("signal received: %s", s.String())
	case err = <-server.Notify():
		logrus.Infof("server notify: %s", err.Error())
	}
	return nil
}
