package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"farmyard/domain"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	// Read operations
	router.Get("/{farmID}", GetFarm)

	// router.Get("/{farmID}/animals", GetAllAnimals)
	// router.Get("/{farmID}/animals/{animalID}", GetAnAnimal)
	return router
}

func GetFarm(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "farmID")
	farm := domain.Farm{
		ID:   id,
		Name: "Animal Farm",
	}
	render.JSON(w, r, farm)
}

/*
func GetAllAnimals(w http.ResponseWriter, r *http.Request) {
}

func GetAnAnimal(w http.ResponseWriter, r *http.Request) {
}
*/
