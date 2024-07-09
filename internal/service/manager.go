package service

import "example.com/tracker/internal/repository"

type Manager struct {
	Repository repository.Repository
}
