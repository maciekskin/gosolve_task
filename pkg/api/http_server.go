package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type IndexService interface {
	GetIndex(int) (int, error)
}

type ApiSevices struct {
	IndexService IndexService
}

func StartHttpServer(services ApiSevices) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := GetIndexResponse{Index: -1}

		valueRaw := r.URL.Query().Get("value")
		value, err := strconv.Atoi(valueRaw)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.ErrorMessage = "value is missing or has invalid type"
			json.NewEncoder(w).Encode(response)
			return
		}

		index, err := services.IndexService.GetIndex(value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.ErrorMessage = err.Error()
		}
		response.Index = index

		json.NewEncoder(w).Encode(response)
	})
	return http.ListenAndServe(":8888", r)
}
