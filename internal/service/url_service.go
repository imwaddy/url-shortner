package service

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/imwaddy/url-shortner/internal/logger"
	"github.com/imwaddy/url-shortner/internal/repository"
	"github.com/imwaddy/url-shortner/pkg/generator"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(r *repository.URLRepository) *URLService {
	return &URLService{repo: r}
}

func (s *URLService) Create(original string) (string, error) {
	for i := 0; i < 5; i++ { // retry 5 times
		short := generator.Generate(6)

		err := s.repo.Save(short, original)
		if err == nil {
			return short, nil
		}

		// if duplicate, retry
		if isDuplicateError(err) {
			continue
		}

		logger.Errorf("Error while saving code %+v", err)
		return "", err
	}

	logger.Errorf("could not generate unique short code")
	return "", fmt.Errorf("could not generate unique short code")
}

func (s *URLService) Resolve(ctx context.Context, short string) (string, error) {
	return s.repo.Get(ctx, short)
}

func isDuplicateError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		return mysqlErr.Number == 1062
	}
	return false
}
