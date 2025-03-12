package api

import (
	"database/sql"
	"net/http"

	"github.com/ajaysehwal/go-rest-server/internal/api/handlers"
	"github.com/ajaysehwal/go-rest-server/internal/api/middleware"
	"github.com/ajaysehwal/go-rest-server/internal/config"
	"github.com/ajaysehwal/go-rest-server/internal/services"
	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) http.Handler{
	r:=mux.NewRouter()
	authSvc:= services.NewAuthService(db,[]byte(config.LoadConfig().JWTSecret));
	r.HandleFunc("/register",handlers.RegisterHandler(authSvc)).Methods("POST")
    // protected:=r.PathPrefix("/api").Subrouter()

	return middleware.LoggingMiddleware(r)
}