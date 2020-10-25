package character

import (
	"github.com/hikvineh/go-rest-game-character/internal/auth"
	"github.com/hikvineh/go-rest-game-character/internal/entity"
	"github.com/hikvineh/go-rest-game-character/internal/test"
	"github.com/hikvineh/go-rest-game-character/pkg/log"
	"net/http"
	"testing"
	"time"
)

func TestAPI(t *testing.T) {
	logger, _ := log.NewForTest()
	router := test.MockRouter(logger)
	repo := &mockRepository{items: []entity.Character{
		{"123", "Frodo", 3, 100, 300, time.Now(), time.Now()},
	}}
	RegisterHandlers(router.Group(""), NewService(repo, logger), auth.MockAuthHandler, logger)
	header := auth.MockAuthHeader()

	tests := []test.APITestCase{
		{"get all", "GET", "/characters", "", nil, http.StatusOK, `*"total_count":1*`},
		{"get 123", "GET", "/characters/123", "", nil, http.StatusOK, `*Frodo*`},
		{"get unknown", "GET", "/characters/1234", "", nil, http.StatusNotFound, ""},
		{"create ok", "POST", "/characters", `{"name":"test"}`, header, http.StatusCreated, "*test*"},
		{"create ok count", "GET", "/characters", "", nil, http.StatusOK, `*"total_count":2*`},
		{"create auth error", "POST", "/characters", `{"name":"test"}`, nil, http.StatusUnauthorized, ""},
		{"create input error", "POST", "/characters", `"name":"test"}`, header, http.StatusBadRequest, ""},
		{"update ok", "PUT", "/characters/123", `{"name":"Frodoxyz"}`, header, http.StatusOK, "*Frodoxyz*"},
		{"update ok", "PUT", "/characters/123", `{"name":"Frodoxyz", "character_power":19}`, header, http.StatusOK, "*38*"},
		{"update ok", "PUT", "/characters/123", `{"name":"Frodoxyz", "character_power":20}`, header, http.StatusOK, "*60*"},
		{"update verify", "GET", "/characters/123", "", nil, http.StatusOK, `*Frodoxyz*`},
		{"update auth error", "PUT", "/characters/123", `{"name":"Frodoxyz"}`, nil, http.StatusUnauthorized, ""},
		{"update input error", "PUT", "/characters/123", `"name":"Frodoxyz"}`, header, http.StatusBadRequest, ""},
		{"delete ok", "DELETE", "/characters/123", ``, header, http.StatusOK, "*Frodoxyz*"},
		{"delete verify", "DELETE", "/characters/123", ``, header, http.StatusNotFound, ""},
		{"delete auth error", "DELETE", "/characters/123", ``, nil, http.StatusUnauthorized, ""},
	}
	for _, tc := range tests {
		test.Endpoint(t, router, tc)
	}
}
