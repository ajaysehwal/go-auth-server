package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ajaysehwal/go-rest-server/internal/api/handlers"
	"github.com/ajaysehwal/go-rest-server/internal/config"
	"github.com/ajaysehwal/go-rest-server/internal/services"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddlware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		authHeader:=r.Header.Get("Authorization")
		if authHeader == ""{
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(handlers.APIError{Message: "Mission token"})
			return;

		}
		fmt.Print([]byte(config.LoadConfig().JWTSecret))
		tokenStr:=strings.TrimPrefix(authHeader,"Bearer ")
		claims:=&services.Claims{}
	    token, err:=jwt.ParseWithClaims(tokenStr,claims,func(token *jwt.Token)(any, error){
			return []byte(config.LoadConfig().JWTSecret),nil
		})
		if err != nil || !token.Valid {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(handlers.APIError{Message: "Invalid token"})
			return
		}
		log.Printf("Authenticated user ID: %d", claims.UserID)
		next.ServeHTTP(w, r)
	})
}