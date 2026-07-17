# Chirpy
Chirpy is a guided project that simulates a social network's server. Its purpose is to practice combining Go and SQL to create a functional database that will store user and activity information in a database. 

Various authentication methods to prove ownership. Authentication of the user and authorisation of their actions. There is no usable front-end of Chirpy.

Chirpy has the following CRUD endpoints:
	"GET /api/healthz", handlerReadiness
	"GET /api/chirps", cfg.handlerReadChirps
	"GET /api/chirps/{chirpID}", cfg.handlerReadChirp

	"POST /api/chirps", cfg.handlerCreateChirp
	"POST /api/users", cfg.handlerCreateUser
	"POST /api/login", cfg.handlerLogin
	"POST /api/refresh", cfg.handlerRefresh
	"POST /api/revoke", cfg.handlerRevoke
	"POST /api/polka/webhooks", cfg.handlerUpgrade

	"PUT /api/users", cfg.handlerUpdateUser

	"DELETE /api/chirps/{chirpID}", cfg.handlerDeleteChirp
	
    //admin
	"GET /admin/metrics", cfg.handlerMetrics
	"POST /admin/reset", cfg.handlerReset

Chirpy is a CLI tool with minimal functionality. 