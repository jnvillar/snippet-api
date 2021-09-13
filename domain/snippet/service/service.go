package snippetservice

import (
	"context"
	"fmt"

	"snippetapp/customerrors"
	snippetmodel "snippetapp/domain/snippet/model"
	snippetrepository "snippetapp/domain/snippet/repository"
	"snippetapp/utils"
)

const (
	defaultExpirationInSeconds = 30
	snippetsURLDomain          = "http://localhost:8080/snippets"
)

type Service struct {
	repository snippetrepository.Repository
}

func NewSnippetService(repository snippetrepository.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, name, content string, expiration int) (*snippetmodel.Snippet, error) {
	expiresAt := utils.GetCurrentTimePlusSeconds(expiration)
	URL := fmt.Sprintf("%s/%s", snippetsURLDomain, name)
	snippet := snippetmodel.NewSnippet(name, content, URL, expiresAt)
	return s.repository.Create(ctx, snippet)
}

func (s *Service) Like(ctx context.Context, name string) (*snippetmodel.Snippet, error) {
	snippet, err := s.getByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return snippet, nil
}

func (s *Service) GetByName(ctx context.Context, name string) (*snippetmodel.Snippet, error) {
	snippet, err := s.getByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return snippet, nil
}

func (s *Service) getByName(ctx context.Context, name string) (*snippetmodel.Snippet, error) {
	snippet, err := s.repository.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if utils.TimeHasPassed(snippet.ExpiresAt) {
		err = s.repository.DeleteSnippet(ctx, snippet.Name)
		if err != nil {
			return nil, err
		}
		return nil, customerrors.NewNotFoundError(fmt.Errorf("snippet %s not found", name))
	}
	go func() {
		_ = s.repository.UpdateExpiration(ctx, snippet.Name, utils.GetCurrentTimePlusSeconds(defaultExpirationInSeconds))
	}()
	return snippet, err
}
