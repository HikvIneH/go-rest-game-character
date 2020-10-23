package character

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hikvineh/go-rest-game-character/internal/entity"
	"github.com/hikvineh/go-rest-game-character/pkg/log"
	"time"
)

// Service encapsulates usecase logic for albums.
type Service interface {
	Get(ctx context.Context, id string) (Character, error)
	Query(ctx context.Context, offset, limit int) ([]Character, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateCharacterRequest) (Character, error)
	Update(ctx context.Context, id string, input UpdateCharacterRequest) (Character, error)
	Delete(ctx context.Context, id string) (Character, error)
}

// Character represents the data about an album.
type Character struct {
	entity.Character
}

// CreateCharacterRequest represents an character creation request.
type CreateCharacterRequest struct {
	Name string `json:"name"`
}

// Validate validates the CreateAlbumRequest fields.
func (m CreateCharacterRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateCharacterRequest represents an album update request.
type UpdateCharacterRequest struct {
	Name string `json:"name"`
}

// Validate validates the CreateAlbumRequest fields.
func (m UpdateCharacterRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewCharacterService creates a new album service.
func NewCharacterService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the album with the specified the album ID.
func (s service) Get(ctx context.Context, id string) (Character, error) {
	character, err := s.repo.Get(ctx, id)
	if err != nil {
		return Character{}, err
	}
	return Character{character}, nil
}

// Create creates a new album.
func (s service) Create(ctx context.Context, req CreateCharacterRequest) (Character, error) {
	if err := req.Validate(); err != nil {
		return Character{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.Character{
		ID:        id,
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return Character{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the album with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateCharacterRequest) (Character, error) {
	if err := req.Validate(); err != nil {
		return Character{}, err
	}

	character, err := s.Get(ctx, id)
	if err != nil {
		return character, err
	}
	character.Name = req.Name
	character.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, character.Character); err != nil {
		return character, err
	}
	return character, nil
}

// Delete deletes the album with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Character, error) {
	character, err := s.Get(ctx, id)
	if err != nil {
		return Character{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Character{}, err
	}
	return character, nil
}

// Count returns the number of albums.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the albums with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Character, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Character{}
	for _, item := range items {
		result = append(result, Character{item})
	}
	return result, nil
}
