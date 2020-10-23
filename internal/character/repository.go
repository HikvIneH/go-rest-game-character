package character

import (
	"context"
	"github.com/hikvineh/go-rest-game-character/internal/entity"
	"github.com/hikvineh/go-rest-game-character/pkg/dbcontext"
	"github.com/hikvineh/go-rest-game-character/pkg/log"
)

// Repository encapsulates the logic to access albums from the data source.
type Repository interface {
	// Get returns the album with the specified album ID.
	Get(ctx context.Context, id string) (entity.Character, error)
	// Count returns the number of albums.
	Count(ctx context.Context) (int, error)
	// Query returns the list of albums with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]entity.Character, error)
	// Create saves a new album in the storage.
	Create(ctx context.Context, album entity.Character) error
	// Update updates the album with given ID in the storage.
	Update(ctx context.Context, album entity.Character) error
	// Delete removes the album with given ID from the storage.
	Delete(ctx context.Context, id string) error
}

// repository persists albums in database
type repository struct {
	db     *dbcontext.DB
	logger log.Logger
}

// NewRepository creates a new album repository
func NewRepository(db *dbcontext.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

// Get reads the album with the specified ID from the database.
func (r repository) Get(ctx context.Context, id string) (entity.Character, error) {
	var character entity.Character
	err := r.db.With(ctx).Select().Model(id, &character)
	return character, err
}

// Create saves a new album record in the database.
// It returns the ID of the newly inserted album record.
func (r repository) Create(ctx context.Context, character entity.Character) error {
	return r.db.With(ctx).Model(&character).Insert()
}

// Update saves the changes to an album in the database.
func (r repository) Update(ctx context.Context, character entity.Character) error {
	return r.db.With(ctx).Model(&character).Update()
}

// Delete deletes an album with the specified ID from the database.
func (r repository) Delete(ctx context.Context, id string) error {
	character, err := r.Get(ctx, id)
	if err != nil {
		return err
	}
	return r.db.With(ctx).Model(&character).Delete()
}

// Count returns the number of the album records in the database.
func (r repository) Count(ctx context.Context) (int, error) {
	var count int
	err := r.db.With(ctx).Select("COUNT(*)").From("character").Row(&count)
	return count, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r repository) Query(ctx context.Context, offset, limit int) ([]entity.Character, error) {
	var characters []entity.Character
	err := r.db.With(ctx).
		Select().
		OrderBy("id").
		Offset(int64(offset)).
		Limit(int64(limit)).
		All(&characters)
	return characters, err
}
