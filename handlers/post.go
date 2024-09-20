package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mariasilva795/go-api-rest/helpers/auth"
	"github.com/mariasilva795/go-api-rest/models"
	"github.com/mariasilva795/go-api-rest/repository"
	"github.com/mariasilva795/go-api-rest/server"
	"github.com/segmentio/ksuid"
)

type UpsertPostRequest struct {
	PostContent string `json:"postContent"`
}

type PostResponse struct {
	Id          string `json:"id"`
	PostContent string `json:"postContent"`
}

type PostDeletedResponse struct {
	Message string `json:"message"`
}

func GetPostByIDHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		post, err := repository.GetPostById(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

func UpdatePostByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		claimsUserId, err := auth.ValidateToken(s, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var postRequest = UpsertPostRequest{}
		err = json.NewDecoder(r.Body).Decode(&postRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		post := models.Post{
			PostContent: postRequest.PostContent,
			Id:          params["id"],
		}

		err = repository.UpdatePost(r.Context(), &post, claimsUserId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PostDeletedResponse{
			Message: "Post Update",
		})
	}
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		claimsUserId, err := auth.ValidateToken(s, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var postRequest = UpsertPostRequest{}
		err = json.NewDecoder(r.Body).Decode(&postRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		post := models.Post{
			Id:          id.String(),
			PostContent: postRequest.PostContent,
			UserId:      claimsUserId,
		}

		err = repository.InsertPost(r.Context(), &post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(PostResponse{
			Id:          post.Id,
			PostContent: post.PostContent,
		})
	}
}
