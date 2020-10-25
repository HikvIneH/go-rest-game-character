package character

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/hikvineh/go-rest-game-character/internal/errors"
	"github.com/hikvineh/go-rest-game-character/pkg/log"
	"github.com/hikvineh/go-rest-game-character/pkg/pagination"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHandlers(r *routing.RouteGroup, service Service, authHandler routing.Handler, logger log.Logger) {
	res := resource{service, logger}

	r.Get("/characters/<id>", res.get)
	r.Get("/characters", res.query)

	r.Use(authHandler)

	// the following endpoints require a valid JWT
	r.Post("/characters", res.create)
	r.Put("/characters/<id>", res.update)
	r.Delete("/characters/<id>", res.delete)
}

type resource struct {
	service Service
	logger  log.Logger
}

func (r resource) get(c *routing.Context) error {
	character, err := r.service.Get(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(character)
}

func (r resource) query(c *routing.Context) error {
	ctx := c.Request.Context()
	count, err := r.service.Count(ctx)
	if err != nil {
		return err
	}
	pages := pagination.NewFromRequest(c.Request, count)
	characters, err := r.service.Query(ctx, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.Items = characters
	return c.Write(pages)
}

func (r resource) create(c *routing.Context) error {
	var input CreateCharacterRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}
	character, err := r.service.Create(c.Request.Context(), input)
	if err != nil {
		return err
	}

	return c.WriteWithStatus(character, http.StatusCreated)
}

func (r resource) update(c *routing.Context) error {
	var input UpdateCharacterRequest
	if err := c.Read(&input); err != nil {
		r.logger.With(c.Request.Context()).Info(err)
		return errors.BadRequest("")
	}

	character, err := r.service.Update(c.Request.Context(), c.Param("id"), input)
	if err != nil {
		return err
	}

	return c.Write(character)
}

func (r resource) delete(c *routing.Context) error {
	character, err := r.service.Delete(c.Request.Context(), c.Param("id"))
	if err != nil {
		return err
	}

	return c.Write(character)
}
