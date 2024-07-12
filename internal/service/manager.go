package service

import "example.com/tracker/internal/repository"

type Manager struct {
	Repository  repository.Repository
	ExternalApi repository.ExternalAPI
}

func New(repository repository.Repository, externalAPI repository.ExternalAPI) *Manager {
	return &Manager{
		Repository:  repository,
		ExternalApi: externalAPI,
	}
}
