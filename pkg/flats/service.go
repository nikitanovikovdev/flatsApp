package flats

import (
	"context"
	"encoding/json"
	"flatApp/pkg/platform/flat"
	"fmt"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, f []byte) (string, error) {
	var flats flat.Flat

	if err := json.Unmarshal(f, &flats); err != nil {
		fmt.Println(err.Error())
	}

	return s.repo.Create(ctx, flats)
}

func (s *Service) Read(ctx context.Context, id string) (flat.Flat, error) {
	f, err := s.repo.Read(ctx, id)
	if err != nil {
		return flat.Flat{}, err
	}
	return f, nil
}

func (s *Service) Update(ctx context.Context, id string, f []byte) error {
	var fl flat.Flat

	if err := json.Unmarshal(f, &fl); err != nil {
		return err
	}

	return s.repo.Update(ctx, id, &fl)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
