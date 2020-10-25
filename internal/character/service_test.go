package character

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/hikvineh/go-rest-game-character/internal/entity"
	"github.com/hikvineh/go-rest-game-character/pkg/log"
	"github.com/stretchr/testify/assert"
)

var errCRUD = errors.New("error crud")

func TestCreateCharacterRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		model     CreateCharacterRequest
		wantError bool
	}{
		{"success", CreateCharacterRequest{Name: "test", CharacterCode: 1, CharacterPower: 100}, false},
		{"required", CreateCharacterRequest{Name: "", CharacterCode: 1, CharacterPower: 100}, true},
		{"too long", CreateCharacterRequest{Name: "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func TestUpdateCharacterRequest_Validate(t *testing.T) {
	tests := []struct {
		name      string
		model     UpdateCharacterRequest
		wantError bool
	}{
		{"success", UpdateCharacterRequest{Name: "test"}, false},
		{"required", UpdateCharacterRequest{Name: ""}, true},
		{"too long", UpdateCharacterRequest{Name: "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.model.Validate()
			assert.Equal(t, tt.wantError, err != nil)
		})
	}
}

func Test_service_CRUD(t *testing.T) {
	logger, _ := log.NewForTest()
	s := NewService(&mockRepository{}, logger)

	ctx := context.Background()

	// initial count
	count, _ := s.Count(ctx)
	assert.Equal(t, 0, count)

	// successful creation
	character, err := s.Create(ctx, CreateCharacterRequest{Name: "test", CharacterCode: 1, CharacterPower: 100})

	characterWizard, err := s.Create(ctx, CreateCharacterRequest{Name: "test", CharacterCode: 1, CharacterPower: 100})
	characterElf, err := s.Create(ctx, CreateCharacterRequest{Name: "test", CharacterCode: 2, CharacterPower: 60})
	characterHobbit, err := s.Create(ctx, CreateCharacterRequest{Name: "test", CharacterCode: 3, CharacterPower: 100})
	characterHobbit2, err := s.Create(ctx, CreateCharacterRequest{Name: "test", CharacterCode: 3, CharacterPower: 10})
	assert.Nil(t, err)
	assert.NotEmpty(t, character.ID)
	id := character.ID
	assert.Equal(t, "test", character.Name)
	assert.Equal(t, int64(100), characterWizard.CharacterPower)
	assert.Equal(t, int64(150), characterWizard.CharacterValue)
	assert.Equal(t, int64(60), characterElf.CharacterPower)
	assert.Equal(t, int64(68), characterElf.CharacterValue)
	assert.Equal(t, int64(100), characterHobbit.CharacterPower)
	assert.Equal(t, int64(300), characterHobbit.CharacterValue)
	assert.Equal(t, int64(10), characterHobbit2.CharacterPower)
	assert.Equal(t, int64(20), characterHobbit2.CharacterValue)
	assert.NotEmpty(t, characterWizard.CreatedAt)
	assert.NotEmpty(t, characterWizard.UpdatedAt)
	count, _ = s.Count(ctx)
	assert.Equal(t, 5, count)

	// validation error in creation
	_, err = s.Create(ctx, CreateCharacterRequest{Name: ""})
	assert.NotNil(t, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 5, count)

	// unexpected error in creation
	_, err = s.Create(ctx, CreateCharacterRequest{Name: "error"})
	assert.Equal(t, errCRUD, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 5, count)

	_, _ = s.Create(ctx, CreateCharacterRequest{Name: "test2"})

	// update
	character, err = s.Update(ctx, id, UpdateCharacterRequest{Name: "test updated"})

	assert.Nil(t, err)
	assert.Equal(t, "test updated", character.Name)
	_, err = s.Update(ctx, "none", UpdateCharacterRequest{Name: "test updated"})
	assert.NotNil(t, err)

	// validation error in update
	_, err = s.Update(ctx, id, UpdateCharacterRequest{Name: ""})

	assert.NotNil(t, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 6, count)

	// unexpected error in update
	_, err = s.Update(ctx, id, UpdateCharacterRequest{Name: "error"})
	assert.Equal(t, errCRUD, err)
	count, _ = s.Count(ctx)
	assert.Equal(t, 6, count)

	// get
	_, err = s.Get(ctx, "none")
	assert.NotNil(t, err)
	character, err = s.Get(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, "test updated", character.Name)
	assert.Equal(t, id, character.ID)

	// query
	characters, _ := s.Query(ctx, 0, 0)
	assert.Equal(t, 6, len(characters))

	// delete
	_, err = s.Delete(ctx, "none")
	assert.NotNil(t, err)
	character, err = s.Delete(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, id, character.ID)
	count, _ = s.Count(ctx)
	assert.Equal(t, 5, count)
}

type mockRepository struct {
	items []entity.Character
}

func (m mockRepository) Get(ctx context.Context, id string) (entity.Character, error) {
	for _, item := range m.items {
		if item.ID == id {
			return item, nil
		}
	}
	return entity.Character{}, sql.ErrNoRows
}

func (m mockRepository) Count(ctx context.Context) (int, error) {
	return len(m.items), nil
}

func (m mockRepository) Query(ctx context.Context, offset, limit int) ([]entity.Character, error) {
	return m.items, nil
}

func (m *mockRepository) Create(ctx context.Context, character entity.Character) error {
	if character.Name == "error" {
		return errCRUD
	}
	m.items = append(m.items, character)
	return nil
}

func (m *mockRepository) Update(ctx context.Context, character entity.Character) error {
	if character.Name == "error" {
		return errCRUD
	}
	for i, item := range m.items {
		if item.ID == character.ID {
			m.items[i] = character
			break
		}
	}
	return nil
}

func (m *mockRepository) Delete(ctx context.Context, id string) error {
	for i, item := range m.items {
		if item.ID == id {
			m.items[i] = m.items[len(m.items)-1]
			m.items = m.items[:len(m.items)-1]
			break
		}
	}
	return nil
}
