package service

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"urlShortner/pkg/encryption"
	"urlShortner/repository/database"

	"github.com/lib/pq"
)

var r = &rand.Rand{}

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func (s *service) Create(ctx context.Context, request CreateRequest) (*CreateResponse, error) {
	short := encryption.Encode(s.generateRandShort(s.config.MaxURLLength))

	err := s.databaseRepo.CreateNewURL(ctx, &database.URL{
		Long:  request.Long,
		Short: short,
	})
	if err == nil {
		s.cacheRepo.Set(ctx, short, request.Long, s.config.MaxURLCache)
		return &CreateResponse{Short: short}, nil
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		short = encryption.Encode(s.generateRandShort(s.config.MaxURLLength))
		err = s.databaseRepo.CreateNewURL(ctx, &database.URL{
			Long:  request.Long,
			Short: short,
		})
		if err == nil {
			s.cacheRepo.Set(ctx, short, request.Long, s.config.MaxURLCache)
			return &CreateResponse{Short: short}, nil
		}
	}

	return nil, err
}

func (s *service) generateRandShort(len int) int64 {
	maxInt := int64(1)
	for i := 0; i < len; i++ {
		maxInt *= 62
	}
	return r.Int63n(maxInt)
}
