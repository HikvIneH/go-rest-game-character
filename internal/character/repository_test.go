package character

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/hikvineh/go-rest-game-character/internal/entity"
	"github.com/hikvineh/go-rest-game-character/internal/test"
	"github.com/hikvineh/go-rest-game-character/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	logger, _ := log.NewForTest()
	db := test.DB(t)
	test.ResetTables(t, db, "character")
	repo := NewRepository(db, logger)

	ctx := context.Background()

	// initial count
	count, err := repo.Count(ctx)
	assert.Nil(t, err)

	// create
	err = repo.Create(ctx, entity.Character{
		ID:             "test1",
		Name:           "character1",
		CharacterCode:  1,
		CharacterPower: 100,
		CharacterValue: 150,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	assert.Nil(t, err)
	count2, _ := repo.Count(ctx)
	assert.Equal(t, 1, count2-count)

	// get
	character, err := repo.Get(ctx, "test1")
	assert.Nil(t, err)
	assert.Equal(t, "character1", character.Name)
	assert.Equal(t, int64(100), character.CharacterPower)
	assert.Equal(t, int64(150), character.CharacterValue)

	_, err = repo.Get(ctx, "test0")
	assert.Equal(t, sql.ErrNoRows, err)

	// update
	err = repo.Update(ctx, entity.Character{
		ID:             "test1",
		Name:           "character1 updated",
		CharacterPower: 10,
		CharacterValue: 15,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	})
	assert.Nil(t, err)
	character, _ = repo.Get(ctx, "test1")
	assert.Equal(t, "character1 updated", character.Name)
	assert.Equal(t, int64(10), character.CharacterPower)
	assert.Equal(t, int64(15), character.CharacterValue)

	// query
	characters, err := repo.Query(ctx, 0, count2)
	assert.Nil(t, err)
	assert.Equal(t, count2, len(characters))

	// delete
	err = repo.Delete(ctx, "test1")
	assert.Nil(t, err)
	_, err = repo.Get(ctx, "test1")
	assert.Equal(t, sql.ErrNoRows, err)
	err = repo.Delete(ctx, "test1")
	assert.Equal(t, sql.ErrNoRows, err)
}
