package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/maciekskin/gosolve_task/pkg/numbers"
)

type IndexService interface {
	GetIndex(int) (numbers.Number, error)
}

type ApiSevices struct {
	IndexService IndexService
}

func StartHttpServer(services ApiSevices, port int) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/numbers/{value}", func(w http.ResponseWriter, r *http.Request) {
		response := GetIndexResponse{Index: -1}

		valueRaw := chi.URLParam(r, "value")
		value, err := strconv.Atoi(valueRaw)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.ErrorMessage = "value has invalid type"
			json.NewEncoder(w).Encode(response)
			return
		}

		number, err := services.IndexService.GetIndex(value)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			response.ErrorMessage = err.Error()
		}
		response.Index = number.Index
		response.Value = number.Value

		json.NewEncoder(w).Encode(response)
	})
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
