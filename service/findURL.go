package service

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

func (s *service) FindURL(ctx context.Context, request FindRequest) (*FindResponse, error) {
	long, err := s.cacheRepo.Get(ctx, request.Short)
	if err == nil {
		return &FindResponse{Long: long}, nil
	}

	long, err = s.databaseRepo.GetLongByShort(ctx, request.Short)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid short")
		}
		return nil, err
	}
	s.cacheRepo.Set(ctx, request.Short, long, s.config.MaxURLCache)

	return &FindResponse{Long: long}, nil
}
