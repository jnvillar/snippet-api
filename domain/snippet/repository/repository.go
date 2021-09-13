package snippetrepository

import (
	"context"
	"time"

	snippetmodel "snippetapp/domain/snippet/model"
)

type Repository interface {
	Create(ctx context.Context, snippet *snippetmodel.Snippet) (*snippetmodel.Snippet, error)
	Like(ctx context.Context, name string) (*snippetmodel.Snippet, error)
	GetByName(ctx context.Context, name string) (*snippetmodel.Snippet, error)
	UpdateExpiration(ctx context.Context, name string, time time.Time) error
	DeleteSnippet(ctx context.Context, name string) error
}
