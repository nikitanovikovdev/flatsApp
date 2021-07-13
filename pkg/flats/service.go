package flats

import (
	"context"
	"encoding/json"
	"flatApp/pkg/platform/flat"
	"github.com/pkg/errors"
)

type Service struct {
	repo Repository
}

func NewService(r *RepositorySQL) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, f []byte) (flat.Flat, error) {
	var fl flat.Flat

	if err := json.Unmarshal(f, &fl); err != nil {
		return fl, errors.Wrap(err, "failed to unmarshal data in Service.Create")
	}

	return s.repo.Create(ctx, fl)
}

func (s *Service) Read(ctx context.Context, id string) (flat.Flat, error) {
	fl, err := s.repo.Read(ctx, id)
	if err != nil {
		return fl, err
	}
	return fl, nil
}

func (s *Service) ReadAll(ctx context.Context) ([]flat.Flat, error) {
	fl, err := s.repo.ReadAll(ctx)
	if err != nil {
		return fl, errors.Wrap(err, "failed to read all data in Service.ReadAll")
	}

	return fl, err
}

func (s *Service) Update(ctx context.Context, id string, f []byte) error {
	var fl flat.Flat

	if err := json.Unmarshal(f, &fl); err != nil {
		return errors.Wrap(err, "failed to unmarshal data in Service.Update")
	}

	return s.repo.Update(ctx, id, fl)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "failed to delete data in Service.Update")
	}
	return s.repo.Delete(ctx, id)
}
