package main

import (
	"net/http"
	"sort"

	"github.com/enkaysi/Chirpy/internal/database"
	"github.com/google/uuid"
)

func authorID(r *http.Request) (uuid.UUID, error) {
	authorIDString := r.URL.Query().Get("author_id")
	if authorIDString == "" {
		return uuid.Nil, nil
	}
	authorID, err := uuid.Parse(authorIDString)
	if err != nil {
		return uuid.Nil, err
	}
	return authorID, nil
}

func (cfg *apiConfig) handlerReadChirps(w http.ResponseWriter, r *http.Request) {
	authorID, err := authorID(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
		return
	}

	var dbChirps []database.Chirp

	if authorID != uuid.Nil {
		dbChirps, err = cfg.db.ReadUsersChirps(r.Context(), authorID)
	} else {
		dbChirps, err = cfg.db.ReadChirps(r.Context())
	}
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to get chirps", err)
		return
	}

	respChirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		respChirps = append(respChirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body:      dbChirp.Body,
			UserID:    dbChirp.UserID,
		})
	}

	sortOrder := r.URL.Query().Get("sort")
	sort.Slice(respChirps, func(i, j int) bool {
		if sortOrder == "desc" {
			return respChirps[i].CreatedAt.After(respChirps[j].CreatedAt)
		}
		return respChirps[i].CreatedAt.Before(respChirps[j].CreatedAt)
	})

	respondWithJSON(w, http.StatusOK, respChirps)

}

func (cfg *apiConfig) handlerReadChirp(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.ReadChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to get chirps", err)
		return
	}

	path := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(path)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	for _, chirp := range chirps {
		if chirp.ID == chirpID {
			respondWithJSON(w, http.StatusOK, Chirp{
				ID:        chirp.ID,
				CreatedAt: chirp.CreatedAt,
				UpdatedAt: chirp.UpdatedAt,
				Body:      chirp.Body,
				UserID:    chirp.UserID,
			})
			return
		}
	}
	respondWithError(w, http.StatusNotFound, "File not found", err)
}
