package api

import (
	"encoding/json"
	"farmyard/domain"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestApiRoutes(t *testing.T) {
	t.Run("creates correct Routes and occurrences", func(t *testing.T) {
		routes := make(map[string]int)

		walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			key := route + "_" + method
			if _, ok := routes[key]; !ok {
				routes[key] = 1
			} else {
				routes[key]++
			}
			return nil
		}

		err := chi.Walk(Routes(), walkFunc)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(routes))
		assert.Equal(t, 1, routes["/farm/{farmID}_GET"])
	})
}

func TestFarmRoutes(t *testing.T) {
	ts := httptest.NewServer(Routes())
	defer ts.Close()
	t.Run("decodes json for a farm", func(t *testing.T) {
		resp, err := ts.Client().Get(ts.URL + "/farm/1234")
		assert.Nil(t, err)

		var farm domain.Farm
		errRead := json.NewDecoder(resp.Body).Decode(&farm)
		assert.Nil(t, errRead)
		assert.Equal(t, "1234", farm.ID)
		assert.Equal(t, "Animal Farm", farm.Name)
	})
}
