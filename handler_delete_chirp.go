package main

import (
	"net/http"

	"github.com/enkaysi/Chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unable to retrieve token", err)
		return
	}
	userID, err := auth.ValidateJWT(token, cfg.secret)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Invalid token", err)
		return
	}

	path := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(path)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}
	chirp, err := cfg.db.GetUserByChirpID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect chirpID for user", err)
		return
	}
	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "User is not owner of chirp", err)
		return
	}
	cfg.db.DeleteChirp(r.Context(), chirpID)
	w.WriteHeader(http.StatusNoContent)
}
