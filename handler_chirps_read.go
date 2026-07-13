package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerReadChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.db.ReadChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to get chirps", err)
		return
	}

	respChirps := []Chirp{}
	for _, chirp := range chirps {
		respChirps = append(respChirps, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

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
