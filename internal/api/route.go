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
	userSvc:=services.NewUserService(db)
    
	r.HandleFunc("/register",handlers.RegisterHandler(authSvc)).Methods("POST")
	r.HandleFunc("/login",handlers.LoginHandler(authSvc)).Methods("POST");
    protected:=r.PathPrefix("/api").Subrouter()
    protected.Use(middleware.AuthMiddlware)
	protected.HandleFunc("/user/{id}",handlers.GetUserHandler(userSvc)).Methods("GET")
	return middleware.LoggingMiddleware(r)
}