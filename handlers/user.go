package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mariasilva795/go-api-rest/models"
	"github.com/mariasilva795/go-api-rest/repository"
	"github.com/mariasilva795/go-api-rest/server"
	"github.com/segmentio/ksuid"
)

type SignUpResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var user = models.User{
			Email:    request.Email,
			Password: request.Password,
			Id:       id.String(),
		}
		//(*models.User, error)
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			ID:    user.Id,
			Email: user.Email,
		})
	}
}
