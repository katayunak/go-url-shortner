package service

import (
	"context"

	"urlShortner/config"
	"urlShortner/repository/cache"
	"urlShortner/repository/database"
)

type service struct {
	cacheRepo    cache.Cache
	databaseRepo database.Database
	config       *config.Config
}

func NewService(
	config *config.Config,
	databaseRepo database.Database,
	cacheRepo cache.Cache,
) Service {
	return &service{
		config:       config,
		databaseRepo: databaseRepo,
		cacheRepo:    cacheRepo,
	}
}

type Service interface {
	FindURL(ctx context.Context, request FindRequest) (*FindResponse, error)
	Create(ctx context.Context, request CreateRequest) (*CreateResponse, error)
}
