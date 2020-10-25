package character

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hikvineh/go-rest-game-character/internal/entity"
	"github.com/hikvineh/go-rest-game-character/pkg/log"
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

// Character types
const (
	Wizard int64 = 1
	Elf    int64 = 2
	Hobbit int64 = 3
)

// CreateCharacterRequest represents an character creation request.
type CreateCharacterRequest struct {
	Name           string `json:"name"`
	CharacterCode  int64  `json:"character_code"`
	CharacterPower int64  `json:"character_power"`
}

// Validate validates the CreateAlbumRequest fields.
func (m CreateCharacterRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateCharacterRequest represents an album update request.
type UpdateCharacterRequest struct {
	Name           string `json:"name"`
	CharacterPower int64  `json:"character_power"`
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

// NewService creates a new album service.
func NewService(repo Repository, logger log.Logger) Service {
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

	power := req.CharacterPower
	var value int64

	if req.CharacterCode == Wizard {
		value = (power * 150 / 100)
	} else if req.CharacterCode == Elf {
		characterValue := power * 110 / 100
		value = 2 + characterValue
	} else if req.CharacterCode == Hobbit {
		if power < 20 {
			value = power * 200 / 100
		} else {
			value = power * 300 / 100
		}
	}

	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.Character{
		ID:             id,
		Name:           req.Name,
		CharacterCode:  req.CharacterCode,
		CharacterPower: power,
		CharacterValue: value,
		CreatedAt:      now,
		UpdatedAt:      now,
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
	power := req.CharacterPower
	var value int64

	if character.CharacterCode == Wizard {
		value = power * (150 / 100)
	} else if character.CharacterCode == Elf {
		value = 2 + (110 / 100)
	} else if character.CharacterCode == Hobbit {
		if power < 20 {
			value = power * (200 / 100)
		} else {
			value = power * (300 / 100)
		}
	}

	character.Name = req.Name
	character.CharacterPower = power
	character.CharacterValue = value

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
