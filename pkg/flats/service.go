package flats

import (
	"context"
	"encoding/json"
	flat2 "github.com/nikitanovikovdev/flatsApp-flats/pkg/platform/flat"
)

type Service struct {
	repo Repository
}

func NewService(r *RepositorySQL) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(ctx context.Context, f []byte, username string) (flat2.Flat, error) {
	var flUser flat2.FlatWithUsername
	var fl flat2.Flat
	var u flat2.Username

	u.Username = username
	flUser.Username = u

	if err := json.Unmarshal(f, &flUser); err != nil {
		return fl, err
	}

	return s.repo.Create(ctx, flUser)
}

func (s *Service) Read(ctx context.Context, username string) ([]flat2.Flat, error) {
	var u flat2.Username
	u.Username = username

	fl, err := s.repo.Read(ctx, u)
	if err != nil {
		return fl, err
	}
	return fl, nil
}

func (s *Service) ReadAll(ctx context.Context) ([]flat2.Flat, error) {
	fl, err := s.repo.ReadAll(ctx)
	if err != nil {
		return fl, err
	}

	return fl, err
}

func (s *Service) Update(ctx context.Context, id string, f []byte, username string) error {
	var flUser flat2.FlatWithUsername
	var u flat2.Username

	u.Username = username
	flUser.Username = u

	if err := json.Unmarshal(f, &flUser); err != nil {
		return err
	}

	return s.repo.Update(ctx, id, flUser)
}

func (s *Service) Delete(ctx context.Context, id string, username string) error {
	var usr flat2.Username
	usr.Username = username

	if err := s.repo.Delete(ctx, id, usr); err != nil {
		return err
	}
	return s.repo.Delete(ctx, id, usr)
}

