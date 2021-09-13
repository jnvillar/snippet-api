package snippetrepository

import (
	"context"
	"fmt"
	"time"

	"snippetapp/customerrors"
	snippetmodel "snippetapp/domain/snippet/model"
)

type memoryRepository struct {
	snippets map[string]*snippetmodel.Snippet
}

func (m *memoryRepository) DeleteSnippet(ctx context.Context, name string) error {
	snippet, err := m.GetByName(ctx, name)
	if err != nil {
		return err
	}
	delete(m.snippets, snippet.Name)
	return nil
}

func (m *memoryRepository) UpdateExpiration(ctx context.Context, name string, time time.Time) error {
	snippet, err := m.GetByName(ctx, name)
	if err != nil {
		return err
	}
	snippet.ExpiresAt = time
	return nil
}

func (m *memoryRepository) Like(ctx context.Context, name string) (*snippetmodel.Snippet, error) {
	snippet, err := m.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	snippet.Likes++
	return snippet, nil
}

func (m *memoryRepository) Create(ctx context.Context, snippet *snippetmodel.Snippet) (*snippetmodel.Snippet, error) {
	_, err := m.GetByName(ctx, snippet.Name)
	if err == nil {
		return nil, customerrors.NewBadRequestError(fmt.Errorf("recipe already exists"))
	}
	m.snippets[snippet.Name] = snippet
	return snippet, nil
}

func (m *memoryRepository) GetByName(ctx context.Context, name string) (*snippetmodel.Snippet, error) {
	snippet, found := m.snippets[name]
	if !found {
		return nil, customerrors.NewNotFoundError(fmt.Errorf("snippet %s not found", name))
	}
	return snippet, nil
}

func NewRepository() Repository {
	return &memoryRepository{
		snippets: map[string]*snippetmodel.Snippet{},
	}
}
