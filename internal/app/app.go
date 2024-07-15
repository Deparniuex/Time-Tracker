package app

import (
	"os"
	"os/signal"

	"example.com/tracker/internal/handler"
	"example.com/tracker/internal/httpserver"
	"example.com/tracker/internal/repository/api"
	"example.com/tracker/internal/repository/pgrepo"
	"example.com/tracker/internal/service"
	storage "example.com/tracker/internal/storage/postgres"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetupConfig(path string) error {
	viper.SetConfigFile(path)
	return viper.ReadInConfig()
}

func SetupLogger() error {
	level, err := logrus.ParseLevel(viper.GetString("LOG_LEVEL"))
	if err != nil {
		return err
	}
	logrus.SetLevel(level)
	return nil
}

func Run() error {
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

	logrus.Info("connection to DB success")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	//running from cmd
	migrator, err := migrate.NewWithDatabaseInstance("file://../migrations/", "postgres", driver)
	if err != nil {
		return err
	}

	err = migrator.Up()
	logrus.Infof("migrations status: %s", err)
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	repository := pgrepo.New(db)
	externalAPI := api.New(api.ApiClientConfig{
		APIURL: viper.GetString("EXTERNAL_API_URL"),
	})

	services := service.New(repository, externalAPI)
	handler := handler.New(services)
	server := httpserver.NewServer(handler.InitRouter(), &httpserver.ServerConfig{
		Port: viper.GetString("API_PORT"),
	})

	server.Start()
	logrus.Info("server started")

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
